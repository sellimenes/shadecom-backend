package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
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
		CoverImage 		string
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
		CoverImage: body.CoverImage,
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

func ProductGetAll(c *gin.Context) {
    var products []models.Product
    db := initializers.DB

    // Get limit from query parameters, if provided
    if limit, ok := c.GetQuery("limit"); ok {
        limitValue, err := strconv.Atoi(limit)
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{
                "error": "Invalid limit value",
            })
            return
        }
        db = db.Limit(limitValue)
    }

    result := db.Preload("Category").Find(&products)

    if result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": "An error occurred while fetching the products",
        })
        return
    }

    // Return products
    c.JSON(200, gin.H{
        "products": products,
    })
}