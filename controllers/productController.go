package controllers

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sellimenes/shadecom-backend/initializers"
	"github.com/sellimenes/shadecom-backend/models"
)

func ProductCreate(c *gin.Context) {
	var body struct {
		Name 			string
		Description 	string
		Price 			float64
		Stock 			int
		Images 			[]string
		CategoryID 		int
		IsActive 		bool
		IsSale 			bool
		IsFeatured 		bool
		SaleProcent 	int
	}

	c.Bind(&body)

	// Create slug
	slug := createSlug(body.Name)

	// Convert Images to json.RawMessage
	imagesJSON, err := json.Marshal(body.Images)
	if err != nil {
		fmt.Println(err)
	}
	
	// Create product
	product := models.Product{
		Name: body.Name,
		Slug: slug,
		Description: body.Description,
		Price: body.Price,
		Stock: body.Stock,
		Images: imagesJSON,
		CategoryID: body.CategoryID,
		IsActive: body.IsActive,
		IsSale: body.IsSale,
		IsFeatured: body.IsFeatured,
		SaleProcent: body.SaleProcent,
	}
	result := initializers.DB.Create(&product)

	if result.Error != nil {
		// If the error is due to a duplicate slug
		if strings.Contains(result.Error.Error(), "duplicate") && strings.Contains(result.Error.Error(), "slug") {
			c.JSON(400, gin.H{
				"error": "A category with this slug already exists",
			})
			return
		}

		// If the error is due to another reason
		c.JSON(500, gin.H{
			"error": "An error occurred while creating the product.",
		})
		return
	}

	// Return product
	c.JSON(200, gin.H{
		"product": product,
	})
}