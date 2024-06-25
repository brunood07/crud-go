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

// CreateNotification godoc
// @Summary Create a notification
// @Description Create a new notification
// @Tags notifications
// @Accept  json
// @Produce  json
// @Param notification body models.Notification true "Notification to create"
// @Success 201 {object} models.Notification
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /notifications [post]
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