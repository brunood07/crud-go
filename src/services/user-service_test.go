package services

import (
	"bytes"
	"crud/src/db"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestGetUsers(t *testing.T) {
	// Mock the database
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer mockDB.Close()
	db.SetDB(mockDB)

	// Prepare mock rows
	rows := sqlmock.NewRows([]string{"id", "first_name", "last_name", "age", "email"}).
		AddRow(1, "John", "Doe", 30, "john@example.com").
		AddRow(2, "Jane", "Smith", 25, "jane@example.com")

	mock.ExpectQuery("SELECT id, first_name, last_name, age, email FROM users").WillReturnRows(rows)

	// Setup Gin router
	router := SetupRouter()
	router.GET("/users", GetUsers)

	// Create a request
	req, _ := http.NewRequest("GET", "/users", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "John")
	assert.Contains(t, w.Body.String(), "Jane")
}

func TestCreateUser(t *testing.T) {
	// Mock the database
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer mockDB.Close()
	db.SetDB(mockDB)

	// Prepare mock statement and result
	mock.ExpectPrepare("INSERT INTO users").
		ExpectQuery().
		WithArgs("John", "Doe", 30, "john@example.com").
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	// Setup Gin router
	router := SetupRouter()
	router.POST("/users", CreateUser)

	// Create a request
	userJSON := `{"first_name":"John","last_name":"Doe","age":30,"email":"john@example.com"}`
	req, _ := http.NewRequest("POST", "/users", bytes.NewBufferString(userJSON))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Contains(t, w.Body.String(), "John")
}

func TestUpdateUser(t *testing.T) {
	// Mock the database
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer mockDB.Close()
	db.SetDB(mockDB)

	// Prepare mock statement and result
	mock.ExpectPrepare("UPDATE users SET").
		ExpectExec().
		WithArgs("John", "Doe", 31, "john@example.com", 1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	// Setup Gin router
	router := SetupRouter()
	router.PUT("/users/:id", UpdateUser)

	// Create a request
	userJSON := `{"first_name":"John","last_name":"Doe","age":31,"email":"john@example.com"}`
	req, _ := http.NewRequest("PUT", "/users/1", bytes.NewBufferString(userJSON))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "user updated successfully")
}

func TestDeleteUser(t *testing.T) {
	// Mock the database
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer mockDB.Close()
	db.SetDB(mockDB)

	// Prepare mock statement and result
	mock.ExpectPrepare("DELETE FROM users WHERE id =").
		ExpectExec().
		WithArgs(1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	// Setup Gin router
	router := SetupRouter()
	router.DELETE("/users/:id", DeleteUser)

	// Create a request
	req, _ := http.NewRequest("DELETE", "/users/1", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "user deleted successfully")
}
