package controllers

import (
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

	c.Bind(&body)

	// Create a role
	role := models.Role{Name: body.Name}
	result := initializers.DB.Create(&role)

	if result.Error != nil {
		// If the error is due to a duplicate slug
		if strings.Contains(result.Error.Error(), "duplicate") && strings.Contains(result.Error.Error(), "slug") {
			c.JSON(400, gin.H{
				"error": "A role with this name already exists",
			})
			return
		}

		c.JSON(400, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"role": role,
	})
}