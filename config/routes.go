package config

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/khildamaililhaq/gorest-boilerplate/controllers"
	"log"
	"os"
)

var (
	authsController controllers.AuthsController = controllers.NewAuthsController()
	officesController controllers.OfficesController = controllers.NewOfficesController()
)

func Routes() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	routes := gin.Default()

	// Load Auths Route
	authRoutes := routes.Group("/api/v1/auths")
	{
		authRoutes.POST("/login", authsController.Login)
		authRoutes.POST("/register", authsController.Register)
		authRoutes.GET("/me", authsController.Me)
	}

	// Load Offices Route
	officeRoutes := routes.Group("/api/v1/offices")
	{
		officeRoutes.GET("/", officesController.Index)
		officeRoutes.POST("/", officesController.Create)
		officeRoutes.GET("/:id", officesController.Get)
		officeRoutes.PUT("/:id", officesController.Update)
		officeRoutes.PATCH("/:id", officesController.Update)
		officeRoutes.DELETE("/:id", officesController.Destroy)
	}

	routes.Run(os.Getenv("APP_PORT"))
}
