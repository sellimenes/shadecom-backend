package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sellimenes/shadecom-backend/initializers"
	"github.com/sellimenes/shadecom-backend/models"
)

func AddBasket(c *gin.Context){
	var body struct {
		ProductID uint
		Quantity  int
	}
	if err := c.ShouldBind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}
    user, exists := c.Get("user")
    if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{
            "error": "User not authenticated",
        })
        return
    }

    userModel, ok := user.(models.User)
    if !ok {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": "Failed to process user information",
        })
        return
    }
	userEmail := userModel.Email
	
	basket := models.Basket{
		UserEmail: userEmail,
		ProductID: body.ProductID,
		Quantity:  body.Quantity,
	}
	initializers.DB.Create(&basket)
	c.JSON(http.StatusCreated, gin.H{
		"message": "Product added to basket",
	})
}

func GetBasket(c *gin.Context){
    user, exists := c.Get("user")
    if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{
            "error": "User not authenticated",
        })
        return
    }

    userModel, ok := user.(models.User)
    if !ok {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": "Failed to process user information",
        })
        return
    }
	userEmail := userModel.Email
	
    if !ok {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": "Failed to process user information",
        })
        return
    }
	var basket []models.Basket
	initializers.DB.Where("user_email = ?", userEmail).Find(&basket)
	c.JSON(http.StatusOK, gin.H{
		"basket": basket,
	})
}