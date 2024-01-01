package controllers

import (
	"net/http"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sellimenes/shadecom-backend/initializers"
	"github.com/sellimenes/shadecom-backend/models"
)

var trToEn = map[string]string{
    "ğ": "g",
    "ı": "i",
    "ö": "o",
    "ü": "u",
    "ç": "c",
    "ş": "s",
}

var reg = regexp.MustCompile("[^a-zA-Z0-9-]+")

func createSlug(name string) string {
    slug := strings.ToLower(name)
    slug = strings.ReplaceAll(slug, " ", "-")

    for old, new := range trToEn {
        slug = strings.ReplaceAll(slug, old, new)
    }

    slug = reg.ReplaceAllString(slug, "")

    return slug
}

func CategoryCreate(c *gin.Context){
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

	slug := createSlug(body.Name)

	// Create a category
	category := models.Category{Name: body.Name, Slug: slug}
	result := initializers.DB.Create(&category)

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
			"error": "An error occurred while creating the category",
		})
		return
	}

	// Return category
	c.JSON(200, gin.H{
		"category": category,
	})
}

func CategoryIndex(c *gin.Context){
	// Get all categories
	var categories []models.Category
	result := initializers.DB.Order("name asc").Find(&categories)

	// Handle potential database errors
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "An error occurred while fetching categories",
		})
		return
	}

	// Return all categories
	c.JSON(200, gin.H{
		"categories": categories,
	})
}

func CategoryShow(c *gin.Context){
    // Get category by id
    var category models.Category
    result := initializers.DB.First(&category, c.Param("id"))

    // Handle potential database errors
    if result.Error != nil {
        c.JSON(http.StatusNotFound, gin.H{
            "error": "Category not found",
        })
        return
    }

    // Return category
    c.JSON(http.StatusOK, gin.H{
        "category": category,
    })
}

func CategoryUpdate(c *gin.Context){
	// Get category by id
	var category models.Category
	result := initializers.DB.First(&category, c.Param("id"))

    // Handle potential database errors
    if result.Error != nil {
        c.JSON(http.StatusNotFound, gin.H{
            "error": "Category not found",
        })
        return
    }

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

	slug := createSlug(body.Name)


	// Update category
	category.Name = body.Name
	category.Slug = slug
	result = initializers.DB.Save(&category)

    // Handle potential database errors
    if result.Error != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "An error occurred while updating the category",
        })
        return
    }

	// Return category
	c.JSON(200, gin.H{
		"category": category,
	})
}

func CategoryDelete(c *gin.Context){
    // Get category by id
    var category models.Category
    result := initializers.DB.First(&category, c.Param("id"))

    // Handle potential database errors
    if result.Error != nil {
        c.JSON(http.StatusNotFound, gin.H{
            "error": "Category not found",
        })
        return
    }

    // Delete category
    result = initializers.DB.Delete(&category)

    if result.Error != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "An error occurred while deleting the category",
        })
        return
    }

    // Return category
    c.JSON(http.StatusOK, gin.H{
        "category": category,
    })
}