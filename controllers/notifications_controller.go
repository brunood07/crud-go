package controller

import (
	"crud/models"
	usecase "crud/usecases"
	"fmt"
	"net/http"
	"strconv"

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
// @Param notification body models.CreateNotification true "Notification to create"
// @Success 201 {object} models.Notification
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /notifications [post]
func (nc *NotificationsController) CreateNotification(ctx *gin.Context) {

	var newNotification models.CreateNotification
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

// SetNotificationRead godoc
// @Summary Changes a notification to read
// @Description Update notification to read
// @Tags notifications
// @Accept  json
// @Produce  json
// @Param id path int true "Notification ID"
// @Success 200 {object} models.Notification
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /notifications/read/{id} [put]
func (nc *NotificationsController) SetNotificationRead(ctx *gin.Context) {
	id := ctx.Param("id")
	notificationId, err := strconv.Atoi(id)

	if err != nil {
		fmt.Println(err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid notification id to update", })
		return
	}

	readNotification, err := nc.notificationsUsecase.SetNotificationRead(notificationId)
	if err != nil {
		fmt.Println(err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "updating notification", })
		return
	}

	ctx.JSON(http.StatusOK, readNotification)
}

// GetUserNotifications godoc
// @Summary Get all user notifications
// @Description Get all user notifications
// @Tags notifications
// @Accept  json
// @Produce  json
// @Param id path int true "User ID"
// @Success 200 {object} []models.Notification
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /notifications/user/{id} [get]
func (nc *NotificationsController) GetUserNotifications(ctx *gin.Context) {
	id := ctx.Param("id")
	userId, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println(err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid user id", })
		return
	}

	notifications, err := nc.notificationsUsecase.GetUserNotifications(userId)
	if err != nil {
		fmt.Println(err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "while fetching user notifications", })
		return
	}

	ctx.JSON(http.StatusOK, notifications)
}

// GetUserUnreadNotifications godoc
// @Summary Get all user unread notifications
// @Description Get all user unread notifications
// @Tags notifications
// @Accept  json
// @Produce  json
// @Param id path int true "User ID"
// @Success 200 {object} []models.Notification
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /notifications/user/unread/{id} [get]
func (nc *NotificationsController) GetUserUnreadNotifications(ctx *gin.Context) {
	id := ctx.Param("id")
	userId, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println(err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid user id", })
		return
	}

	notifications, err := nc.notificationsUsecase.GetUserUnreadNotifications(userId)
	if err != nil {
		fmt.Println(err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "while fetching user notifications", })
		return
	}

	ctx.JSON(http.StatusOK, notifications)
}