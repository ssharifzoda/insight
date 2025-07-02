package api

import (
	"encoding/json"
	"errors"
	"green/internal/models"
	"green/pkg/consts"
	"green/pkg/utils"
	"net/http"
)

// @Summary loginHandler
// @Tags Auth
// @Description sign-in in account
// @ID loginHandler
// @Accept json
// @Produce json
// @Param input body models.SingIn true "user info"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Failure default {object} models.ErrorResponse
// @Router /auth/login [post]
func (h *Handler) loginHandler(w http.ResponseWriter, r *http.Request) {
	var request *models.SingIn
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		h.logger.Error(err)
		utils.ErrorResponse(w, err, 400, 0)
		return
	}
	//get user from db
	user, err := h.service.GetUserByUsername(request.Username)
	if err != nil {
		if err.Error() == consts.UserNotFound {
			h.logger.Error(err)
			utils.ErrorResponse(w, errors.New(consts.UserNotFound), 400, 0)
			return
		} else {
			h.logger.Error(err)
			utils.ErrorResponse(w, errors.New(consts.InternalServerError), 500, 0)
			return
		}
	}
	//check password
	if user.Password != request.Password {
		h.logger.Error(consts.UsernameOrPasswordWrong)
		utils.ErrorResponse(w, errors.New(consts.UsernameOrPasswordWrong), 400, 0)
		return
	}
	//create token
	accessToken, refreshToken, err := utils.GenerateTokens(user.Username)
	if err != nil {
		h.logger.Error(err)
		utils.ErrorResponse(w, errors.New(consts.InternalServerError), 500, 0)
		return
	}
	//save refreshToken
	err = h.service.Authorization.UpdateRefreshToken(request.Username, accessToken, refreshToken)
	if err != nil {
		h.logger.Error(err)
		utils.ErrorResponse(w, errors.New(consts.InternalServerError), 500, 0)
		return
	}
	utils.Response(w, map[string]interface{}{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
		"role":          user.Role,
	})
}

// @Summary refreshToken
// @Security ApiKeyAuth
// @Tags Auth
// @Description Роут для обвноляения токена
// @ID refreshToken
// @Produce json
// @Success 200 {object} models.Response
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Failure default {object} models.ErrorResponse
// @Router /auth/refresh-token [post]
func (h *Handler) refreshToken(w http.ResponseWriter, r *http.Request) {
	refreshToken := r.Header.Get("Authorization")
	if refreshToken == "" {
		h.logger.Error(consts.TokenIsEmpty)
		utils.ErrorResponse(w, errors.New(consts.TokenIsEmpty), 401, 0)
		return
	}
	username, err := utils.ParseRefreshToken(refreshToken)
	if err != nil {
		h.logger.Error(err)
		utils.ErrorResponse(w, errors.New(consts.InternalServerError), 500, 0)
		return
	}
	_, refreshTokenDb, err := h.service.GetTokensByUsername(username)
	if err != nil {
		h.logger.Error(err)
		utils.ErrorResponse(w, errors.New(consts.InternalServerError), 500, 0)
		return
	}
	if refreshToken != refreshTokenDb {
		h.logger.Error(consts.UserNotFound)
		utils.ErrorResponse(w, errors.New(consts.UserNotFound), 404, 0)
		return
	}
	accessToken, refreshToken, err := utils.GenerateTokens(username)
	if err != nil {
		h.logger.Error(err)
		utils.ErrorResponse(w, errors.New(consts.InternalServerError), 500, 0)
		return
	}
	err = h.service.Authorization.UpdateRefreshToken(username, accessToken, refreshToken)
	if err != nil {
		h.logger.Error(err)
		utils.ErrorResponse(w, errors.New(consts.InternalServerError), 500, 0)
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
// @Param input body models.ChangePasswordSW true "Заполните поля"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Failure default {object} models.ErrorResponse
// @Router /auth/change-password [put]
func (h *Handler) changePassword(w http.ResponseWriter, r *http.Request) {
	var request *models.ChangePassword
	token := r.Header.Get("Authorization")
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		h.logger.Error(err)
		utils.ErrorResponse(w, err, 400, 0)
		return
	}
	//todo: validate password
	username, err := utils.ParseRefreshToken(token)
	if err != nil {
		h.logger.Error(err)
		utils.ErrorResponse(w, errors.New(consts.TokenIsEmpty), 400, 0)
		return
	}
	user, err := h.service.Authorization.GetUserByUsername(username)
	if err != nil {
		h.logger.Error(err)
		utils.ErrorResponse(w, errors.New(consts.InternalServerError), 500, 0)
		return
	}
	if user.Password != request.OldPassword {
		h.logger.Error(consts.UsernameOrPasswordWrong)
		utils.ErrorResponse(w, errors.New(consts.UsernameOrPasswordWrong), 400, 0)
		return
	}
	request.UserId = user.Id
	err = h.service.Authorization.ChangeUserPassword(request)
	if err != nil {
		h.logger.Error(err)
		utils.ErrorResponse(w, errors.New(consts.InternalServerError), 500, 0)
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
// @Success 200 {object} models.Response
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Failure default {object} models.ErrorResponse
// @Router /auth/log-out [put]
func (h *Handler) logoutHandler(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	if token == "" {
		h.logger.Error(consts.TokenIsEmpty)
		utils.ErrorResponse(w, errors.New(consts.TokenIsEmpty), 400, 0)
		return
	}
	username, err := utils.ParseRefreshToken(token)
	if err != nil {
		h.logger.Error(consts.TokenIsEmpty)
		utils.ErrorResponse(w, errors.New(consts.TokenIsEmpty), 400, 0)
		return
	}
	err = h.service.Authorization.UpdateRefreshToken(username, " ", " ")
	if err != nil {
		h.logger.Error(err)
		utils.ErrorResponse(w, errors.New(consts.InternalServerError), 500, 0)
		return
	}
	utils.Response(w, consts.Success)
}
