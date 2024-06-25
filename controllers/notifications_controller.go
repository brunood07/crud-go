package controller

import (
	"crud/models"
	usecase "crud/usecases"
	"net/http"

	"github.com/gin-gonic/gin"
)

type NotificationsController struct {
	notificationsUsecase usecase.NotificationsUsecase
}

func NewNotificationsController(usecase usecase.NotificationsUsecase) NotificationsController {
	return NotificationsController{
		notificationsUsecase: usecase,
	}
}

func (nc *NotificationsController) CreateNotification(ctx *gin.Context) {

	var newNotification models.Notification
	err := ctx.BindJSON(&newNotification)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	createdNotification, err := nc.notificationsUsecase.CreateNotification(newNotification)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, createdNotification)
}