package controllers

import (
	"encoding/json"
	"net/http"
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

    if err := c.ShouldBind(&body); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Invalid request body",
        })
        return
    }

	// Create slug
	slug := createSlug(body.Name)

	// Convert Images to json.RawMessage
    imagesJSON, err := json.Marshal(body.Images)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": "An error occurred while processing the images",
        })
        return
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
            c.JSON(http.StatusBadRequest, gin.H{
                "error": "A product with this slug already exists",
            })
            return
        }

        // If the error is due to another reason
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": "An error occurred while creating the product",
        })
        return
    }

	// Return product
	c.JSON(200, gin.H{
		"product": product,
	})
}