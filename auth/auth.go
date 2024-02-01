package auth

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sellimenes/shadecom-backend/initializers"
	"github.com/sellimenes/shadecom-backend/models"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(c *gin.Context) {
    //   Get data off req body
    var body struct {
        Name     string
        Email    string
        Password string
    }

    if err := c.ShouldBind(&body); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Invalid request body",
        })
        return
    }

    // Hash the password
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": "Failed to hash password",
        })
        return
    }

    //   Create a user
    user := models.User{Name: body.Name, Email: body.Email, Password: string(hashedPassword)}
    result := initializers.DB.Create(&user)

    if result.Error != nil {
        // If the error is due to a duplicate email
        if strings.Contains(result.Error.Error(), "duplicate") && strings.Contains(result.Error.Error(), "email") {
            c.JSON(http.StatusBadRequest, gin.H{
                "error": "A user with this email already exists",
            })
            return
        }

        c.JSON(http.StatusBadRequest, gin.H{
            "error": result.Error.Error(),
        })
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "user": user,
    })
}