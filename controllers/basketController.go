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
    }
    if err := c.ShouldBind(&body); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Invalid request body",
        })
        return
    }

    userEmail, exists := c.Get("userEmail")
    if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{
            "error": "User not authenticated",
        })
        return
    }

    userEmailStr, ok := userEmail.(string)
    if !ok {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": "Failed to process user information",
        })
        return
    }
    
    var existingBasket models.Basket
    if err := initializers.DB.Where("user_email = ? AND product_id = ?", userEmailStr, body.ProductID).First(&existingBasket).Error; err == nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Product already in basket",
        })
        return
    }

    basket := models.Basket{
        UserEmail: userEmailStr,
        ProductID: body.ProductID,
    }
    initializers.DB.Create(&basket)
    c.JSON(http.StatusCreated, gin.H{
        "message": "Product added to basket",
    })
}

func GetBasket(c *gin.Context){
    userEmail, exists := c.Get("userEmail")
    if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{
            "error": "User not authenticated",
        })
        return
    }

    userEmailStr, ok := userEmail.(string)
    if !ok {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": "Failed to process user information",
        })
        return
    }

    var basket []models.Basket
    initializers.DB.Where("user_email = ?", userEmailStr).Preload("Product").Find(&basket)
    c.JSON(http.StatusOK, gin.H{
        "basket": basket,
    })
}