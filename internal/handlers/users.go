package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"insight/internal/models"
	"insight/pkg/consts"
	"insight/pkg/utils"
	"net/http"
	"strconv"
)

// @Summary addNewUser
// @Security ApiKeyAuth
// @Tags Users
// @Description Добавление нового пользователя
// @ID addNewUser
// @Accept json
// @Produce json
// @Param params body models.UserSW true "Введите данные"
// @Success 200 {object} utils.DataResponse
// @Failure 500 {object} utils.DataResponse
// @Failure 400 {object} utils.DataResponse
// @Failure default {object} utils.DataResponse
// @Router /users/new [post]
func (h *Handler) addNewUser(w http.ResponseWriter, r *http.Request) {
	var params *models.User
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		h.logger.Error(err)
		utils.ErrorResponse(w, consts.InvalidRequestData, 400, 0)
		return
	}
	userParams, err := h.service.AddNewUser(params)
	if err != nil {
		h.logger.Error(err)
		utils.ErrorResponse(w, consts.InternalServerError, 500, 0)
		return
	}

	utils.Response(w, userParams)
}

// @Summary editUser
// @Security ApiKeyAuth
// @Tags Users
// @Description Редактирование пользователя
// @ID editUser
// @Accept json
// @Produce json
// @Param params body models.UserSW true "Введите данные"
// @Success 200 {object} utils.DataResponse
// @Failure 500 {object} utils.DataResponse
// @Failure 400 {object} utils.DataResponse
// @Failure default {object} utils.DataResponse
// @Router /users/edit [put]
func (h *Handler) editUser(w http.ResponseWriter, r *http.Request) {
	var params *models.User
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		h.logger.Error(err)
		utils.ErrorResponse(w, consts.InvalidRequestData, 400, 0)
		return
	}
	err = h.service.UpdateUserParams(params)
	if err != nil {
		h.logger.Error(err)
		utils.ErrorResponse(w, consts.InternalServerError, 500, 0)
		return
	}
	utils.Response(w, consts.Success)
}

// @Summary getAllUsers
// @Security ApiKeyAuth
// @Tags Users
// @Description Просмотр списока пользователей
// @ID getAllUsers
// @Accept json
// @Produce json
// @Param page query string true "Введите данные"
// @Param limit query string true "Введите данные"
// @Success 200 {object} utils.DataResponse
// @Failure 500 {object} utils.DataResponse
// @Failure 400 {object} utils.DataResponse
// @Failure default {object} utils.DataResponse
// @Router /users/list [get]
func (h *Handler) getAllUsers(w http.ResponseWriter, r *http.Request) {
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
	users, err := h.service.GetAllUsers(page, limit)
	if err != nil {
		h.logger.Error(err)
		utils.ErrorResponse(w, consts.InternalServerError, 500, 0)
		return
	}
	utils.Response(w, users)
}

// @Summary getUserById
// @Security ApiKeyAuth
// @Tags Users
// @Description Просмотр пользователя
// @ID getUserById
// @Accept json
// @Produce json
// @Param id query string true "Введите данные"
// @Success 200 {object} utils.DataResponse
// @Failure 500 {object} utils.DataResponse
// @Failure 400 {object} utils.DataResponse
// @Failure default {object} utils.DataResponse
// @Router /users/by-id [get]
func (h *Handler) getUserById(w http.ResponseWriter, r *http.Request) {
	userIdStr := mux.Vars(r)["id"]
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		h.logger.Error(err)
		utils.ErrorResponse(w, consts.InvalidRequestData, 400, 0)
		return
	}
	user, err := h.service.GetUserById(userId)
	if err != nil {
		h.logger.Error(err)
		utils.ErrorResponse(w, consts.InternalServerError, 500, 0)
		return
	}
	utils.Response(w, user)
}

// @Summary getMe
// @Security ApiKeyAuth
// @Tags Users
// @Description Информация про меня
// @ID getMe
// @Accept json
// @Produce json
// @Success 200 {object} utils.DataResponse
// @Failure 500 {object} utils.DataResponse
// @Failure 400 {object} utils.DataResponse
// @Failure default {object} utils.DataResponse
// @Router /users/me [get]
func (h *Handler) getMe(w http.ResponseWriter, r *http.Request) {
	userId, _, _, err := utils.ParseToken(r.Header.Get("Authorization"))
	if err != nil {
		h.logger.Error(err)
		utils.ErrorResponse(w, consts.TokenIsEmpty, 400, 0)
		return
	}
	user, err := h.service.GetUserById(userId)
	if err != nil {
		h.logger.Error(err)
		utils.ErrorResponse(w, consts.InternalServerError, 500, 0)
		return
	}
	utils.Response(w, user)
}

// @Summary deleteUser
// @Security ApiKeyAuth
// @Tags Users
// @Description Удаление пользователя
// @ID deleteUser
// @Accept json
// @Produce json
// @Param id query string true "Введите данные"
// @Success 200 {object} utils.DataResponse
// @Failure 500 {object} utils.DataResponse
// @Failure 400 {object} utils.DataResponse
// @Failure default {object} utils.DataResponse
// @Router /users/rm [delete]
func (h *Handler) deleteUser(w http.ResponseWriter, r *http.Request) {
	userIdStr := mux.Vars(r)["id"]
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		h.logger.Error(err)
		utils.ErrorResponse(w, consts.InvalidRequestData, 400, 0)
		return
	}
	err = h.service.DeleteUser(userId)
	if err != nil {
		h.logger.Error(err)
		utils.ErrorResponse(w, consts.InternalServerError, 500, 0)
		return
	}
	utils.Response(w, consts.Success)
}
