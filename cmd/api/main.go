package main

import (
	"absence/internal"
	"absence/internal/middleware"
	"absence/pkg/database"
	"log"
	"os"

	_ "absence/docs" // This will be generated by swag

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Absence API
// @version         1.0
// @description     A attendance management system API.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Database configuration
	dbConfig := &database.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
	}

	// Initialize database
	db, err := internal.InitializeDB(dbConfig)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// Auto migrate database
	if err := database.AutoMigrate(db); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	// Initialize API
	api, err := internal.InitializeAPI(db)
	if err != nil {
		log.Fatalf("Failed to initialize API: %v", err)
	}

	// Setup Gin router
	router := gin.Default()

	// Swagger documentation
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Public routes
	router.POST("/api/register", api.UserHandler.Register)
	router.POST("/api/login", api.UserHandler.Login)

	// Protected routes
	apiGroup := router.Group("/api")
	apiGroup.Use(middleware.AuthMiddleware())
	{
		// User routes
		users := apiGroup.Group("/users")
		{
			users.GET("/:id", api.UserHandler.GetUser)
			users.PUT("/:id", api.UserHandler.UpdateUser)
			users.DELETE("/:id", api.UserHandler.DeleteUser)
			users.GET("/:id/attendance", api.AttendanceHandler.GetUserAttendances)
		}

		// Attendance routes
		attendance := apiGroup.Group("/attendance")
		{
			attendance.POST("/check-in", api.AttendanceHandler.CheckIn)
			attendance.POST("/check-out", api.AttendanceHandler.CheckOut)
			attendance.GET("/:id", api.AttendanceHandler.GetAttendance)
		}
	}

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
