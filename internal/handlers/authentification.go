package api

import (
	"encoding/json"
	"errors"
	"gorm.io/gorm"
	"insight/internal/models"
	"insight/pkg/consts"
	"insight/pkg/utils"
	"net/http"
)

// @Summary loginHandler
// @Tags Auth
// @Description sign-in in account
// @ID loginHandler
// @Accept json
// @Produce json
// @Param input body models.SingIn true "user info"
// @Success 200 {object} utils.DataResponse
// @Failure 400 {object} utils.DataResponse
// @Failure 500 {object} utils.DataResponse
// @Failure default {object} utils.DataResponse
// @Router /auth/login [post]
func (h *Handler) loginHandler(w http.ResponseWriter, r *http.Request) {
	var request *models.SingIn
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		h.logger.Error(err)
		utils.ErrorResponse(w, consts.InvalidRequestData, 400, 0)
		return
	}
	//get user from db
	user, err := h.service.Users.GetUserByPhone(request.Phone)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			h.logger.Error(err)
			utils.ErrorResponse(w, consts.UserNotFound, 404, 0)
			return
		} else {
			h.logger.Error(err)
			utils.ErrorResponse(w, consts.InvalidRequestData, 400, 0)
			return
		}
	}
	userAuth, err := h.service.Authorization.GetTokenByUserId(user.Id)
	if err != nil {
		h.logger.Error(err)
		utils.ErrorResponse(w, consts.InternalServerError, 500, 0)
		return
	}
	//check password
	if user.Password != request.Password {
		h.logger.Error(consts.UsernameOrPasswordWrong)
		utils.ErrorResponse(w, consts.UsernameOrPasswordWrong, 400, 0)
		return
	}
	if userAuth.TemporaryPass == 1 {
		h.logger.Error(consts.TemporaryPassResponse)
		utils.ErrorResponse(w, consts.TemporaryPassResponse, 451, 0)
		return
	}
	permissions, err := h.service.Authorization.GetUserPermission(user.Id)
	if err != nil {
		h.logger.Error(err)
		utils.ErrorResponse(w, consts.InternalServerError, 500, 0)
		return
	}
	//create token
	accessToken, refreshToken, err := utils.GenerateTokens(user.Id, permissions)
	if err != nil {
		h.logger.Error(err)
		utils.ErrorResponse(w, consts.InternalServerError, 500, 0)
		return
	}
	//save refreshToken
	err = h.service.Authorization.UpdateRefreshToken(user.Id, accessToken, refreshToken)
	if err != nil {
		h.logger.Error(err)
		utils.ErrorResponse(w, consts.InternalServerError, 500, 0)
		return
	}
	utils.Response(w, map[string]interface{}{
		"user_id":       user.Id,
		"access_token":  accessToken,
		"refresh_token": refreshToken,
		"role":          user.RoleId,
	})
}

// @Summary refreshToken
// @Security ApiKeyAuth
// @Tags Auth
// @Description Роут для обвноляения токена
// @ID refreshToken
// @Produce json
// @Success 200 {object} utils.DataResponse
// @Failure 400 {object} utils.DataResponse
// @Failure 500 {object} utils.DataResponse
// @Failure default {object} utils.DataResponse
// @Router /auth/refresh-token [post]
func (h *Handler) refreshToken(w http.ResponseWriter, r *http.Request) {
	refreshToken := r.Header.Get("Authorization")
	if refreshToken == "" {
		h.logger.Error(consts.TokenIsEmpty)
		utils.ErrorResponse(w, consts.TokenIsEmpty, 401, 0)
		return
	}
	userId, err := utils.ParseRefreshToken(refreshToken)
	if err != nil {
		h.logger.Error(err)
		utils.ErrorResponse(w, consts.InternalServerError, 500, 0)
		return
	}
	userAuth, err := h.service.GetTokenByUserId(userId)
	if err != nil {
		h.logger.Error(err)
		utils.ErrorResponse(w, consts.InternalServerError, 500, 0)
		return
	}

	if refreshToken != userAuth.RefreshToken {
		h.logger.Error(consts.UserNotFound)
		utils.ErrorResponse(w, consts.UserNotFound, 404, 0)
		return
	}
	permissions, err := h.service.Authorization.GetUserPermission(userId)
	if err != nil {
		h.logger.Error(err)
		utils.ErrorResponse(w, consts.InternalServerError, 500, 0)
		return
	}
	accessToken, refreshToken, err := utils.GenerateTokens(userId, permissions)
	if err != nil {
		h.logger.Error(err)
		utils.ErrorResponse(w, consts.InternalServerError, 500, 0)
		return
	}
	err = h.service.Authorization.UpdateRefreshToken(userId, accessToken, refreshToken)
	if err != nil {
		h.logger.Error(err)
		utils.ErrorResponse(w, consts.InternalServerError, 500, 0)
		return
	}
	utils.Response(w, map[string]interface{}{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}

// @Summary changePassword
// @Security ApiKeyAuth
// @Tags Auth
// @Description changePassword
// @ID changePassword
// @Accept json
// @Produce json
// @Param input body models.ChangePassword true "Заполните поля"
// @Success 200 {object} utils.DataResponse
// @Failure 400 {object} utils.DataResponse
// @Failure 500 {object} utils.DataResponse
// @Failure default {object} utils.DataResponse
// @Router /auth/change-password [put]
func (h *Handler) changePassword(w http.ResponseWriter, r *http.Request) {
	var request *models.ChangePassword
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		h.logger.Error(err)
		utils.ErrorResponse(w, consts.InvalidRequestData, 400, 0)
		return
	}
	user, err := h.service.Users.GetUserById(request.UserId)
	if err != nil {
		h.logger.Error(err)
		utils.ErrorResponse(w, consts.InternalServerError, 500, 0)
		return
	}
	if user.Password != request.OldPassword {
		h.logger.Error(consts.UsernameOrPasswordWrong)
		utils.ErrorResponse(w, consts.UsernameOrPasswordWrong, 400, 0)
		return
	}
	request.UserId = user.Id
	err = h.service.Authorization.ChangeUserPassword(request)
	if err != nil {
		h.logger.Error(err)
		utils.ErrorResponse(w, consts.InternalServerError, 500, 0)
		return
	}
	utils.Response(w, consts.Success)
}

// @Summary logoutHandler
// @Security ApiKeyAuth
// @Tags Auth
// @Description Роут для выхода из системы
// @ID logoutHandler
// @Accept json
// @Produce json
// @Success 200 {object} utils.DataResponse
// @Failure 400 {object} utils.DataResponse
// @Failure 500 {object} utils.DataResponse
// @Failure default {object} utils.DataResponse
// @Router /auth/log-out [put]
func (h *Handler) logoutHandler(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	if token == "" {
		h.logger.Error(consts.TokenIsEmpty)
		utils.ErrorResponse(w, consts.TokenIsEmpty, 400, 0)
		return
	}
	userId, err := utils.ParseRefreshToken(token)
	if err != nil {
		h.logger.Error(consts.TokenIsEmpty)
		utils.ErrorResponse(w, consts.TokenIsEmpty, 400, 0)
		return
	}
	err = h.service.Authorization.UpdateRefreshToken(userId, " ", " ")
	if err != nil {
		h.logger.Error(err)
		utils.ErrorResponse(w, consts.InternalServerError, 500, 0)
		return
	}
	utils.Response(w, consts.Success)
}
