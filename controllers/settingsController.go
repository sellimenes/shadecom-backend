package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sellimenes/shadecom-backend/initializers"
	"github.com/sellimenes/shadecom-backend/models"
)

func SettingsUpdate(c *gin.Context) {
	var body models.Settings

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Find first or create settings
	var settings models.Settings
	if result := initializers.DB.FirstOrCreate(&settings, models.Settings{}); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	// Update settings
	if err := initializers.DB.Model(&settings).Updates(body).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return message
	c.JSON(http.StatusOK, gin.H{
		"message": "Settings updated successfully",
	})
}

func SettingsGet(c *gin.Context) {
	var settings models.Settings

	if result := initializers.DB.First(&settings, models.Settings{}); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, settings)
}
