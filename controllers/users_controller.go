package controller

import (
	"crud/models"
	usecase "crud/usecases"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UsersController struct {
	usersUsecase usecase.UsersUsecase
}

func NewUsersController(usecase usecase.UsersUsecase) UsersController {
	return UsersController{
		usersUsecase: usecase,
	}
}

func (uc *UsersController) CreateUser(ctx *gin.Context) {

	var newUser models.User
	err := ctx.BindJSON(&newUser)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	createdUser, err := uc.usersUsecase.CreateUser(newUser)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, createdUser)
}

func (uc *UsersController) GetUsers(ctx *gin.Context) {

	users, err := uc.usersUsecase.GetUsers()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, users)
}

func (uc *UsersController) UpdateUser(ctx *gin.Context) {
	id := ctx.Param("id")
	userID, err := strconv.Atoi(id)

	if err != nil {
		fmt.Println(err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid user id to update", })
		return
	}

	var updateUser models.User
	if err := ctx.BindJSON(&updateUser); err != nil {
		fmt.Println(err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	updatedUser, err := uc.usersUsecase.UpdateUser(userID, updateUser)
	if err != nil {
		fmt.Println(err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
	}

	ctx.JSON(http.StatusOK, updatedUser)
}

func (uc *UsersController) DeleteUser(ctx *gin.Context) {
	id := ctx.Param("id")
	userID, err := strconv.Atoi(id)

	if err != nil {
		fmt.Println(err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid user id to delete", })
		return
	}

	message, err := uc.usersUsecase.DeleteUser(userID)
	if err != nil {
		fmt.Println(err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "error while deleteing user"})
	}

	ctx.JSON(http.StatusOK, message)
}