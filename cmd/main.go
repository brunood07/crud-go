package main

import (
	controller "crud/controllers"
	"crud/db"
	"crud/env"
	"crud/repositories"
	usecase "crud/usecases"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

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

	router.Run("localhost:" + cfg.Port);
}