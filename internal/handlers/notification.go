package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"insight/internal/models"
	"insight/pkg/consts"
	"insight/pkg/utils"
	"math"
	"net/http"
	"strconv"
)

// @Summary addNewNotification
// @Security ApiKeyAuth
// @Tags Notifications
// @Description Создание нового уведомления
// @ID addNewNotification
// @Accept json
// @Produce json
// @Param params body models.NotificationInput true "Введите данные"
// @Success 200 {object} utils.DataResponse
// @Failure 500 {object} utils.DataResponse
// @Failure 400 {object} utils.DataResponse
// @Failure default {object} utils.DataResponse
// @Router /notifications/new [post]
func (h *Handler) addNewNotification(w http.ResponseWriter, r *http.Request) {
	var params *models.NotificationInput
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		h.logger.Error(err)
		utils.ErrorResponse(w, consts.InvalidRequestData, 400, 0)
		return
	}
	err = h.service.Notifications.CreateNewNotification(params)
	if err != nil {
		h.logger.Error(err)
		utils.ErrorResponse(w, consts.InternalServerError, 500, 0)
		return
	}

	utils.Response(w, consts.Success)
}

// @Summary getAllNotifications
// @Security ApiKeyAuth
// @Tags Notifications
// @Description Просмотр списока всех уведомлений
// @ID getAllNotifications
// @Accept json
// @Produce json
// @Param page query string true "Введите данные"
// @Param limit query string true "Введите данные"
// @Success 200 {object} utils.DataResponse
// @Failure 500 {object} utils.DataResponse
// @Failure 400 {object} utils.DataResponse
// @Failure default {object} utils.DataResponse
// @Router /notifications/list [get]
func (h *Handler) getAllNotifications(w http.ResponseWriter, r *http.Request) {
	pageStr := mux.Vars(r)["page"]
	limitStr := mux.Vars(r)["limit"]
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		h.logger.Error(err)
		utils.ErrorResponse(w, consts.PageErrResponse, 400, 0)
		return
	}
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		h.logger.Error(err)
		utils.ErrorResponse(w, consts.CountErrResponse, 400, 0)
		return
	}
	notifications, count, err := h.service.Notifications.GetAllNotifications(page, limit)
	if err != nil {
		h.logger.Error(err)
		utils.ErrorResponse(w, consts.InternalServerError, 500, 0)
		return
	}
	utils.Response(w, map[string]interface{}{
		"reports":     notifications,
		"total_count": count,
		"count":       len(notifications),
		"total_page":  int(math.Ceil(float64(count) / float64(limit))),
	})
}

// @Summary getNotification
// @Security ApiKeyAuth
// @Tags Notifications
// @Description Просмотр уведомления
// @ID getNotification
// @Accept json
// @Produce json
// @Param notification_id query string true "Введите данные"
// @Success 200 {object} utils.DataResponse
// @Failure 500 {object} utils.DataResponse
// @Failure 400 {object} utils.DataResponse
// @Failure default {object} utils.DataResponse
// @Router /notifications/by-id [get]
func (h *Handler) getNotification(w http.ResponseWriter, r *http.Request) {
	notificationStr := mux.Vars(r)["notification_id"]
	notificationId, err := strconv.Atoi(notificationStr)
	if err != nil {
		h.logger.Error(err)
		utils.ErrorResponse(w, consts.InvalidRequestData, 400, 0)
		return
	}
	notification, err := h.service.Notifications.GetNotificationById(notificationId)
	if err != nil {
		h.logger.Error(err)
		utils.ErrorResponse(w, consts.InternalServerError, 500, 0)
		return
	}
	utils.Response(w, notification)
}

// @Summary deleteNotification
// @Security ApiKeyAuth
// @Tags Notifications
// @Description Удаление уведомления
// @ID deleteNotification
// @Accept json
// @Produce json
// @Param notification_id query string true "Введите данные"
// @Success 200 {object} utils.DataResponse
// @Failure 500 {object} utils.DataResponse
// @Failure 400 {object} utils.DataResponse
// @Failure default {object} utils.DataResponse
// @Router /notifications/rm [delete]
func (h *Handler) deleteNotification(w http.ResponseWriter, r *http.Request) {
	notificationStr := mux.Vars(r)["notification_id"]
	notificationId, err := strconv.Atoi(notificationStr)
	if err != nil {
		h.logger.Error(err)
		utils.ErrorResponse(w, consts.InvalidRequestData, 400, 0)
		return
	}
	err = h.service.Notifications.DeleteNotification(notificationId)
	if err != nil {
		h.logger.Error(err)
		utils.ErrorResponse(w, consts.InternalServerError, 500, 0)
		return
	}
	utils.Response(w, consts.Success)
}
