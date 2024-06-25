package main

import (
	controller "crud/controllers"
	"crud/db"
	"crud/env"
	"crud/repositories"
	usecase "crud/usecases"
	"log"

	_ "crud/docs"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @host						crud-go-4kur.onrender.com
func main() {
	cfg := env.LoadEnv()

	db.Init()

	if err := db.Migrate(); err != nil {
		log.Fatalf("Failed to apply migrations: %v", err)
}
	router := gin.Default()

	// USER ROUTES
	UsersRepository := repositories.NewUsersRepository(db.CON)
	UsersUsecase := usecase.NewUsersUsecase(UsersRepository)
	UsersController := controller.NewUsersController(UsersUsecase)

	router.GET("/users", UsersController.GetUsers)
	router.POST("/users", UsersController.CreateUser)
	router.PUT("/users/:id", UsersController.UpdateUser)
	router.DELETE("/users/:id", UsersController.DeleteUser)

	// NOTIFICATION ROUTES
	NotificationsRepository := repositories.NewNotificationsRepository(db.CON)
	NotificationsUsecase := usecase.NewNotificationsUsecase(NotificationsRepository)
	NotificationsController := controller.NewNotificationsController(NotificationsUsecase)
	router.POST("/notifications", NotificationsController.CreateNotification)
	router.PUT("/notifications/read/:id", NotificationsController.SetNotificationRead)
	router.GET("/notifications/user/:id", NotificationsController.GetUserNotifications)

	// SWAGGER
	swaggerHost := cfg.SwaggerHost
	if swaggerHost == "" {
		swaggerHost = "http://localhost:8080"
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.URL(swaggerHost+"/swagger/doc.json")))

	router.Run(cfg.AppHost + cfg.Port);
}