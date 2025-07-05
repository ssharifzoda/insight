package api

import "net/http"

// @Summary addBrand
// @Tags Settings
// @Description Добавление бренда
// @ID addBrand
// @Accept json
// @Produce json
// @Param phone query string true "Введите номера телефона пользователя"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Failure default {object} models.ErrorResponse
// @Router /settings/brands [post]
func (h *Handler) addBrand(w http.ResponseWriter, r *http.Request) {

}

// @Summary getAllBrands
// @Tags Settings
// @Description Добавление бренда
// @ID getAllBrands
// @Accept json
// @Produce json
// @Param phone query string true "Введите номера телефона пользователя"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Failure default {object} models.ErrorResponse
// @Router /settings/brands [get]
func (h *Handler) getAllBrands(w http.ResponseWriter, r *http.Request) {

}
func (h *Handler) editBrand(w http.ResponseWriter, r *http.Request) {

}
func (h *Handler) deleteBrand(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) addNewCategory(w http.ResponseWriter, r *http.Request) {

}
func (h *Handler) getAllCategories(w http.ResponseWriter, r *http.Request) {

}
func (h *Handler) editCategory(w http.ResponseWriter, r *http.Request) {

}
func (h *Handler) deleteCategory(w http.ResponseWriter, r *http.Request) {

}
