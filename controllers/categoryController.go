package controllers

import (
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sellimenes/shadecom-backend/initializers"
	"github.com/sellimenes/shadecom-backend/models"
)

// Slug create function
func createSlug(name string) string {
    slug := strings.ToLower(name)
    slug = strings.ReplaceAll(slug, " ", "-")

    // Turkish characters to English
    trToEn := map[string]string{
        "ğ": "g",
        "ı": "i",
        "ö": "o",
        "ü": "u",
        "ç": "c",
        "ş": "s",
    }

    for old, new := range trToEn {
        slug = strings.ReplaceAll(slug, old, new)
    }

    // Only English characters, numbers and dashes
    reg, err := regexp.Compile("[^a-zA-Z0-9-]+")
    if err != nil {
        panic(err)
    }
    slug = reg.ReplaceAllString(slug, "")

    return slug
}

func CategoryCreate(c *gin.Context){
	// Get data off req body
	var body struct {
		Name string
	}

	c.Bind(&body)

	slug := createSlug(body.Name)

	// Create a category
	category := models.Category{Name: body.Name, Slug: slug}
	result := initializers.DB.Create(&category)

	if result.Error != nil {
		c.Status(400)
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
	initializers.DB.Find(&categories)

	// Return all categories
	c.JSON(200, gin.H{
		"categories": categories,
	})
}

func CategoryShow(c *gin.Context){
	// Get category by id
	var category models.Category
	initializers.DB.First(&category, c.Param("id"))

	// Return category
	c.JSON(200, gin.H{
		"category": category,
	})
}

func CategoryUpdate(c *gin.Context){
	// Get category by id
	var category models.Category
	initializers.DB.First(&category, c.Param("id"))

	// Get data off req body
	var body struct {
		Name string
	}

	c.Bind(&body)

	slug := createSlug(body.Name)


	// Update category
	category.Name = body.Name
	category.Slug = slug
	result := initializers.DB.Save(&category)

	if result.Error != nil {
		c.Status(400)
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
	initializers.DB.First(&category, c.Param("id"))

	// Delete category
	result := initializers.DB.Delete(&category)

	if result.Error != nil {
		c.Status(400)
		return
	}

	// Return category
	c.JSON(200, gin.H{
		"category": category,
	})
}