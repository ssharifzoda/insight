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

// @Summary addBrand
// @Tags Settings
// @Description Добавление бренда
// @ID addBrand
// @Accept json
// @Produce json
// @Param params body models.Brand true "Введите данные"
// @Success 200 {object} utils.Response
// @Failure 500 {object} utils.ErrorResponse
// @Failure 400 {object} utils.ErrorResponse
// @Failure default {object} utils.ErrorResponse
// @Router /settings/brands [post]
func (h *Handler) addBrand(w http.ResponseWriter, r *http.Request) {
	var params *models.Brand
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		h.logger.Error(err)
		utils.ErrorResponse(w, consts.InvalidRequestData, 400, 0)
		return
	}
	err = h.service.AddBrand(params)
	if err != nil {
		h.logger.Error(err)
		utils.ErrorResponse(w, consts.InternalServerError, 500, 0)
		return
	}
	utils.Response(w, consts.Success)
}

// @Summary getAllBrands
// @Tags Settings
// @Description Просмотр всех брендов
// @ID getAllBrands
// @Accept json
// @Produce json
// @Success 200 {object} utils.Response
// @Failure 500 {object} utils.ErrorResponse
// @Failure 400 {object} utils.ErrorResponse
// @Failure default {object} utils.ErrorResponse
// @Router /settings/brands [get]
func (h *Handler) getAllBrands(w http.ResponseWriter, r *http.Request) {
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
	brands, err := h.service.GetAllBrands(page, limit)
	if err != nil {
		h.logger.Error(err)
		utils.ErrorResponse(w, consts.InternalServerError, 500, 0)
		return
	}
	utils.Response(w, brands)
}

// @Summary editBrand
// @Tags Settings
// @Description Изменение бренда
// @ID editBrand
// @Accept json
// @Produce json
// @Param params body models.Brand true "Введите данные"
// @Success 200 {object} utils.Response
// @Failure 500 {object} utils.ErrorResponse
// @Failure 400 {object} utils.ErrorResponse
// @Failure default {object} utils.ErrorResponse
// @Router /settings/brands [get]
func (h *Handler) editBrand(w http.ResponseWriter, r *http.Request) {
	var brand *models.Brand
	err := json.NewDecoder(r.Body).Decode(&brand)
	if err != nil {
		h.logger.Error(err)
		utils.ErrorResponse(w, consts.InvalidRequestData, 400, 0)
		return
	}
	err = h.service.EditBrand(brand)
	if err != nil {
		h.logger.Error(err)
		utils.ErrorResponse(w, consts.InternalServerError, 500, 0)
		return
	}
	utils.Response(w, consts.Success)
}

// @Summary deleteBrand
// @Tags Settings
// @Description Удаление бренда
// @ID deleteBrand
// @Accept json
// @Produce json
// @Param brand_id query string true "Введите данные"
// @Success 200 {object} utils.Response
// @Failure 500 {object} utils.ErrorResponse
// @Failure 400 {object} utils.ErrorResponse
// @Failure default {object} utils.ErrorResponse
// @Router /settings/brands [get]
func (h *Handler) deleteBrand(w http.ResponseWriter, r *http.Request) {
	brandIdStr := mux.Vars(r)["brand_id"]
	brandId, err := strconv.Atoi(brandIdStr)
	if err != nil {
		h.logger.Error(err)
		utils.ErrorResponse(w, consts.InvalidBrandId, 400, 0)
		return
	}
	err = h.service.DeleteBrand(brandId)
	if err != nil {
		h.logger.Error(err)
		utils.ErrorResponse(w, consts.InternalServerError, 500, 0)
		return
	}
	utils.Response(w, consts.Success)
}

func (h *Handler) addNewCategory(w http.ResponseWriter, r *http.Request) {

}
func (h *Handler) getAllCategories(w http.ResponseWriter, r *http.Request) {

}
func (h *Handler) editCategory(w http.ResponseWriter, r *http.Request) {

}
func (h *Handler) deleteCategory(w http.ResponseWriter, r *http.Request) {

}
