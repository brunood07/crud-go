package main

import (
	"crud/src/db"
	"crud/src/env"
	"crud/src/services"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	cfg := env.LoadEnv()

	db.Init()
	
	router := gin.Default()
	router.GET("/users", services.GetUsers)
	router.POST("/users", services.CreateUser)
	router.PUT("/users/:id", services.UpdateUser)
	router.DELETE("/users/:id", services.DeleteUser)
	router.Run("localhost:" + cfg.Port);
}