package controllers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sellimenes/shadecom-backend/initializers"
	"github.com/sellimenes/shadecom-backend/models"
)

func RoleCreate(c *gin.Context){
    // Get data off req body
    var body struct {
        Name string
    }

    if err := c.ShouldBind(&body); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Invalid request body",
        })
        return
    }

    // Create a role
    role := models.Role{Name: body.Name}
    result := initializers.DB.Create(&role)

    if result.Error != nil {
        // If the error is due to a duplicate slug
        if strings.Contains(result.Error.Error(), "duplicate") && strings.Contains(result.Error.Error(), "slug") {
            c.JSON(http.StatusBadRequest, gin.H{
                "error": "A role with this name already exists",
            })
            return
        }

        c.JSON(http.StatusBadRequest, gin.H{
            "error": result.Error.Error(),
        })
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "role": role,
    })
}