package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sellimenes/shadecom-backend/auth"
	"github.com/sellimenes/shadecom-backend/controllers"
	"github.com/sellimenes/shadecom-backend/initializers"
	"github.com/sellimenes/shadecom-backend/middleware"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main () {
	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000", "http://localhost:8080", "https://shadecom.vercel.app"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	config.AllowHeaders = []string{"Origin", "Content-Type"}

	r.Use(cors.New(config))

	r.POST("/api/category", controllers.CategoryCreate)
	r.GET("/api/category", controllers.CategoryIndex)
	r.GET("/api/category/:id", controllers.CategoryShow)
	r.PUT("/api/category/:id", controllers.CategoryUpdate)
	r.DELETE("/api/category/:id", controllers.CategoryDelete)

	r.POST("/api/role", controllers.RoleCreate)
	r.POST("/api/register", auth.CreateUser)
	r.POST("/api/login", auth.LoginUser)
	r.GET("/api/me", auth.GetCurrentUser)

	r.PUT("/api/settings", controllers.SettingsUpdate)
	r.GET("/api/settings", controllers.SettingsGet)


	r.POST("/api/upload", controllers.UploadImages)

	r.POST("/api/product", controllers.ProductCreate)
	r.GET("/api/products", controllers.ProductGetAll)
	r.GET("/api/product/:slug", controllers.ProductGetSingle)

	r.GET("/api/basket", middleware.AuthMiddleware(), controllers.GetBasket)
	r.POST("/api/basket", middleware.AuthMiddleware(), controllers.AddBasket)

	r.Run()
}