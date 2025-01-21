package main

import (
	"api1/logger"
	"api1/middleware"
	"api1/userfeature"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize logger
	log := logger.NewLogger()

	// Initialize service and handler
	userService := userfeature.NewUserService(log)
	userHandler := userfeature.NewUserHandler(userService, log)

	// Set up Gin router
	r := gin.New()
	r.Use(gin.Recovery())

	//set middlewares
	r.Use(middleware.LoggingMiddleware(log)) // Add the logging middleware

	// Routes
	v1 := r.Group("/api/v1")
	{
		users := v1.Group("/users")
		{
			users.POST("", userHandler.Create)
			users.GET("", userHandler.List)
			users.GET("/:id", userHandler.Get)
			users.PUT("/:id", userHandler.Update)
			users.DELETE("/:id", userHandler.Delete)
		}
	}

	// Start server
	log.Info("Starting server on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Server failed to start: ", err)
	}
}

/*
func main() {
	// Load config
	cfg := config.Load()

	// Initialize logger
	log := logger.NewLogger()

	// Enable file logging if configured
	if cfg.LogConfig.EnableFile {
		logger.EnableFileLogging(log, cfg.LogConfig)
	}

	// Use logger
	log.Info("Application starting...")

	// Initialize db
	//db := postgres.Connect(cfg.Database)

	// Initialize repositories
	userRepo := user.NewRepository(db)

	// Initialize services
	userService := user.NewService(userRepo, logger)

	// Initialize handlers
	userHandler := user.NewHandler(userService, logger)

	// Setup router
	router := gin.Default()

	// Register routes
	api := router.Group("/api/v1")
	{
		users := api.Group("/users")
		{
			users.POST("", userHandler.Create)
			users.GET("/:id", userHandler.Get)
			users.PUT("/:id", userHandler.Update)
			users.DELETE("/:id", userHandler.Delete)
		}
	}

	// Start server
	router.Run(":8080")
}
*/
