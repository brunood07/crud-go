package main

import (
	"crud/src/db"
	"crud/src/services"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)


func main() {
	err := godotenv.Load()
	if err != nil {
			log.Fatal("Error loading .env file")
	}
	
	db.Init()
	
	router := gin.Default()
	router.GET("/users", services.GetUsers)
	router.POST("/users", services.CreateUser)
	router.PUT("/users/:id", services.UpdateUser)
	router.DELETE("/users/:id", services.DeleteUser)
	router.Run("localhost:8080")
}