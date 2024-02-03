package auth

import (
	"log"
	"net/http"
	"strings"

	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/sellimenes/shadecom-backend/initializers"
	"github.com/sellimenes/shadecom-backend/models"
	"golang.org/x/crypto/bcrypt"
)

type UserClaims struct {
    Name  string
    Email string
    Phone string
    Address string
    RoleID uint
}

type CustomClaims struct {
    User *UserClaims
    jwt.StandardClaims
}

var jwtKey = []byte("your_secret_key")

func CreateUser(c *gin.Context) {
    // Get data off req body
    var body struct {
        Name     string `json:"Name"`
        Email    string `json:"Email"`
        Password string `json:"Password"`
    }

    if err := c.ShouldBind(&body); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Invalid request body",
        })
        return
    }

    // Check if name, email and password are provided
    if body.Name == "" || body.Email == "" || body.Password == "" {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Name, email and password are required",
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

    // Create a user
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

func LoginUser(c *gin.Context) {
    // Get data from request body
    var body struct {
        Email    string
        Password string
    }

    if err := c.ShouldBind(&body); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Invalid request body",
        })
        return
    }

    // Find user by email
    var user models.User
    if err := initializers.DB.Where("email = ?", body.Email).First(&user).Error; err != nil {
        log.Println(err) // Log the actual error message
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Email or password is incorrect",
        })
        return
    }

    // Check password
    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Email or password is incorrect",
        })
        return
    }

    // Create a new UserClaims object from the user
    userClaims := &UserClaims{
    Name:  user.Name,
    Email: user.Email,
    Phone: user.Phone,
    Address: user.Address,
    RoleID: user.RoleID,
}

    // User is authenticated, create JWT token
    expirationTime := time.Now().Add(24 * time.Hour * 7)
    claims := &CustomClaims{
        User: userClaims,
        StandardClaims: jwt.StandardClaims{
            Subject:   user.Email,
            ExpiresAt: expirationTime.Unix(),
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    tokenString, err := token.SignedString(jwtKey)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": "Could not generate token",
        })
        return
    }

    // Return the token
    c.JSON(http.StatusOK, gin.H{
        "token": tokenString,
    })
}

func GetCurrentUser(c *gin.Context) {
    // Get token from Authorization header
    authHeader := c.GetHeader("Authorization")
    tokenString := strings.TrimPrefix(authHeader, "Bearer ")

    // Parse the token
    token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
        return jwtKey, nil
    })

    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{
            "error": "Invalid or expired token",
        })
        return
    }

    // Get the user's email from the token
    claims, ok := token.Claims.(*jwt.StandardClaims)
    if !ok || !token.Valid {
        c.JSON(http.StatusUnauthorized, gin.H{
            "error": "Invalid or expired token",
        })
        return
    }

    // Find the user by email
    var user models.User
    if err := initializers.DB.Where("email = ?", claims.Subject).First(&user).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{
            "error": "User not found",
        })
        return
    }

    // Return the user's information
    c.JSON(http.StatusOK, gin.H{
        "user": user,
    })
}