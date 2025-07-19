package api

import (
	"encoding/json"
	"insight/internal/models"
	"insight/pkg/consts"
	"insight/pkg/utils"
	"net/http"
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
