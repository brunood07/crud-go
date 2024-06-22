package services

import (
	"crud/src/db"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	rows, err := db.CON.Query("SELECT id, first_name, last_name, age, email FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var users []user
	for rows.Next() {
		var u user
		err := rows.Scan(&u.ID, &u.FirstName, &u.LastName, &u.Age, &u.Email)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, u)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	c.IndentedJSON(http.StatusOK, users)
}

func CreateUser(c *gin.Context) {
	var newUser user
	if err := c.BindJSON(&newUser); err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	stmt, err := db.CON.Prepare("INSERT INTO users (first_name, last_name, age, email) VALUES ($1, $2, $3, $4) RETURNING id")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	err = stmt.QueryRow(newUser.FirstName, newUser.LastName, newUser.Age, newUser.Email).Scan(&newUser.ID)
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusCreated, newUser)
}

func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	userID, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid user id to delete", })
		return
	}

	var updateUser user
	if err := c.BindJSON(&updateUser); err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	stmt, err := db.CON.Prepare("UPDATE users SET first_name=$1, last_name=$2, age=$3, email=$4 WHERE id=$5")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(updateUser.FirstName, updateUser.LastName, updateUser.Age, updateUser.Email, userID)
	if err != nil {
		log.Fatal(err)
	}

	// Check how many rows were affected
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}

	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "user not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user updated successfully"})
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	userID, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid user id to delete", })
		return
	}

	stmt, err := db.CON.Prepare("DELETE FROM users WHERE id = $1")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(userID)
	if err != nil {
		log.Fatal(err)
	}

	// Check how many rows were affected
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}

	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "user not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user deleted successfully"})
} 