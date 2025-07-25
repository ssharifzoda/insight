package api

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"insight/internal/models"
	"insight/pkg/consts"
	"insight/pkg/utils"
	"io"
	"math"
	"net/http"
	"os"
	"strconv"
)

// @Summary addBrand
// @Security ApiKeyAuth
// @Tags Settings
// @Description Добавление бренда
// @ID addBrand
// @Accept json
// @Produce json
// @Param params body models.Brand true "Введите данные"
// @Success 200 {object} utils.DataResponse
// @Failure 500 {object} utils.DataResponse
// @Failure 400 {object} utils.DataResponse
// @Failure default {object} utils.DataResponse
// @Router /settings/brands [post]
func (h *Handler) addBrand(w http.ResponseWriter, r *http.Request) {
	var params *models.Brand
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		h.logger.Error(err)
		utils.ErrorResponse(w, consts.InvalidRequestData, 400, 0)
		return
	}
	brand, err := h.service.Settings.AddBrand(params)
	if err != nil {
		h.logger.Error(err)
		utils.ErrorResponse(w, consts.InternalServerError, 500, 0)
		return
	}
	utils.Response(w, brand)
}

// @Summary getAllBrands
// @Security ApiKeyAuth
// @Tags Settings
// @Description Просмотр всех брендов
// @ID getAllBrands
// @Accept json
// @Produce json
// @Param page query string true "Введите данные"
// @Param limit query string true "Введите данные"
// @Param search query string false "Введите данные"
// @Success 200 {object} utils.DataResponse
// @Failure 500 {object} utils.DataResponse
// @Failure 400 {object} utils.DataResponse
// @Failure default {object} utils.DataResponse
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
	search := r.URL.Query().Get("search")
	brands, count, err := h.service.GetAllBrands(page, limit, search)
	if err != nil {
		h.logger.Error(err)
		utils.ErrorResponse(w, consts.InternalServerError, 500, 0)
		return
	}
	utils.Response(w, map[string]interface{}{
		"reports":     brands,
		"total_count": count,
		"count":       len(brands),
		"total_page":  int(math.Ceil(float64(count) / float64(limit))),
	})
}

// @Summary editBrand
// @Security ApiKeyAuth
// @Tags Settings
// @Description Изменение бренда
// @ID editBrand
// @Accept json
// @Produce json
// @Param params body models.Brand true "Введите данные"
// @Success 200 {object} utils.DataResponse
// @Failure 500 {object} utils.DataResponse
// @Failure 400 {object} utils.DataResponse
// @Failure default {object} utils.DataResponse
// @Router /settings/brands [put]
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
// @Security ApiKeyAuth
// @Tags Settings
// @Description Удаление бренда
// @ID deleteBrand
// @Accept json
// @Produce json
// @Param brand_id query string true "Введите данные"
// @Success 200 {object} utils.DataResponse
// @Failure 500 {object} utils.DataResponse
// @Failure 400 {object} utils.DataResponse
// @Failure default {object} utils.DataResponse
// @Router /settings/brands [delete]
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

// @Summary addNewCategory
// @Security ApiKeyAuth
// @Tags Settings
// @Description Добавление категории
// @ID addNewCategory
// @Accept json
// @Produce json
// @Param params body models.Category true "Введите данные"
// @Success 200 {object} utils.DataResponse
// @Failure 500 {object} utils.DataResponse
// @Failure 400 {object} utils.DataResponse
// @Failure default {object} utils.DataResponse
// @Router /settings/categories [post]
func (h *Handler) addNewCategory(w http.ResponseWriter, r *http.Request) {
	var params *models.Category
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		h.logger.Error(err)
		utils.ErrorResponse(w, consts.InvalidRequestData, 400, 0)
		return
	}
	category, err := h.service.AddNewCategory(params)
	if err != nil {
		h.logger.Error(err)
		utils.ErrorResponse(w, consts.InternalServerError, 500, 0)
		return
	}
	utils.Response(w, category)
}

// @Summary getAllCategories
// @Security ApiKeyAuth
// @Tags Settings
// @Description Просмотр всех категорий
// @ID getAllCategories
// @Accept json
// @Produce json
// @Param page query string true "Введите данные"
// @Param limit query string true "Введите данные"
// @Param search query string false "Введите данные"
// @Success 200 {object} utils.DataResponse
// @Failure 500 {object} utils.DataResponse
// @Failure 400 {object} utils.DataResponse
// @Failure default {object} utils.DataResponse
// @Router /settings/categories [get]
func (h *Handler) getAllCategories(w http.ResponseWriter, r *http.Request) {
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
	search := r.URL.Query().Get("search")
	categories, count, err := h.service.GetAllCategories(page, limit, search)
	if err != nil {
		h.logger.Error(err)
		utils.ErrorResponse(w, consts.InternalServerError, 500, 0)
		return
	}
	utils.Response(w, map[string]interface{}{
		"reports":     categories,
		"total_count": count,
		"count":       len(categories),
		"total_page":  int(math.Ceil(float64(count) / float64(limit))),
	})
}

// @Summary editCategory
// @Security ApiKeyAuth
// @Tags Settings
// @Description Редактирование категории
// @ID editCategory
// @Accept json
// @Produce json
// @Param params body models.Category true "Введите данные"
// @Success 200 {object} utils.DataResponse
// @Failure 500 {object} utils.DataResponse
// @Failure 400 {object} utils.DataResponse
// @Failure default {object} utils.DataResponse
// @Router /settings/categories [put]
func (h *Handler) editCategory(w http.ResponseWriter, r *http.Request) {
	var category *models.Category
	err := json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		h.logger.Error(err)
		utils.ErrorResponse(w, consts.InvalidRequestData, 400, 0)
		return
	}
	err = h.service.EditCategory(category)
	if err != nil {
		h.logger.Error(err)
		utils.ErrorResponse(w, consts.InternalServerError, 500, 0)
		return
	}
	utils.Response(w, consts.Success)
}

// @Summary deleteCategory
// @Security ApiKeyAuth
// @Tags Settings
// @Description Удаление категории
// @ID deleteCategory
// @Accept json
// @Produce json
// @Param category_id query string true "Введите данные"
// @Success 200 {object} utils.DataResponse
// @Failure 500 {object} utils.DataResponse
// @Failure 400 {object} utils.DataResponse
// @Failure default {object} utils.DataResponse
// @Router /settings/categories [delete]
func (h *Handler) deleteCategory(w http.ResponseWriter, r *http.Request) {
	categoryIdStr := mux.Vars(r)["category_id"]
	categoryId, err := strconv.Atoi(categoryIdStr)
	if err != nil {
		h.logger.Error(err)
		utils.ErrorResponse(w, consts.InvalidBrandId, 400, 0)
		return
	}
	err = h.service.DeleteCategory(categoryId)
	if err != nil {
		h.logger.Error(err)
		utils.ErrorResponse(w, consts.InternalServerError, 500, 0)
		return
	}
	utils.Response(w, consts.Success)
}

// @Summary addNewCity
// @Security ApiKeyAuth
// @Tags Settings
// @Description Добавление города
// @ID addNewCity
// @Accept json
// @Produce json
// @Param params body models.City true "Введите данные"
// @Success 200 {object} utils.DataResponse
// @Failure 500 {object} utils.DataResponse
// @Failure 400 {object} utils.DataResponse
// @Failure default {object} utils.DataResponse
// @Router /settings/cities [post]
func (h *Handler) addNewCity(w http.ResponseWriter, r *http.Request) {
	var params *models.City
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		h.logger.Error(err)
		utils.ErrorResponse(w, consts.InvalidRequestData, 400, 0)
		return
	}
	city, err := h.service.AddNewCity(params)
	if err != nil {
		h.logger.Error(err)
		utils.ErrorResponse(w, consts.InternalServerError, 500, 0)
		return
	}
	utils.Response(w, city)
}

// @Summary getAllCities
// @Security ApiKeyAuth
// @Tags Settings
// @Description Просмотр всех городов
// @ID getAllCities
// @Accept json
// @Produce json
// @Param page query string true "Введите данные"
// @Param limit query string true "Введите данные"
// @Success 200 {object} utils.DataResponse
// @Failure 500 {object} utils.DataResponse
// @Failure 400 {object} utils.DataResponse
// @Failure default {object} utils.DataResponse
// @Router /settings/cities [get]
func (h *Handler) getAllCities(w http.ResponseWriter, r *http.Request) {
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
	cities, count, err := h.service.GetAllCities(page, limit)
	if err != nil {
		h.logger.Error(err)
		utils.ErrorResponse(w, consts.InternalServerError, 500, 0)
		return
	}
	utils.Response(w, map[string]interface{}{
		"reports":     cities,
		"total_count": count,
		"count":       len(cities),
		"total_page":  int(math.Ceil(float64(count) / float64(limit))),
	})
}

// @Summary editCity
// @Security ApiKeyAuth
// @Tags Settings
// @Description Редактирование города
// @ID editCity
// @Accept json
// @Produce json
// @Param params body models.City true "Введите данные"
// @Success 200 {object} utils.DataResponse
// @Failure 500 {object} utils.DataResponse
// @Failure 400 {object} utils.DataResponse
// @Failure default {object} utils.DataResponse
// @Router /settings/cities [put]
func (h *Handler) editCity(w http.ResponseWriter, r *http.Request) {
	var city *models.City
	err := json.NewDecoder(r.Body).Decode(&city)
	if err != nil {
		h.logger.Error(err)
		utils.ErrorResponse(w, consts.InvalidRequestData, 400, 0)
		return
	}
	err = h.service.EditCity(city)
	if err != nil {
		h.logger.Error(err)
		utils.ErrorResponse(w, consts.InternalServerError, 500, 0)
		return
	}
	utils.Response(w, consts.Success)
}

// @Summary deleteCity
// @Security ApiKeyAuth
// @Tags Settings
// @Description Удаление города
// @ID deleteCity
// @Accept json
// @Produce json
// @Param city_id query string true "Введите данные"
// @Success 200 {object} utils.DataResponse
// @Failure 500 {object} utils.DataResponse
// @Failure 400 {object} utils.DataResponse
// @Failure default {object} utils.DataResponse
// @Router /settings/cities [delete]
func (h *Handler) deleteCity(w http.ResponseWriter, r *http.Request) {
	cityIdStr := mux.Vars(r)["city_id"]
	cityId, err := strconv.Atoi(cityIdStr)
	if err != nil {
		h.logger.Error(err)
		utils.ErrorResponse(w, consts.InvalidBrandId, 400, 0)
		return
	}
	err = h.service.DeleteCity(cityId)
	if err != nil {
		h.logger.Error(err)
		utils.ErrorResponse(w, consts.InternalServerError, 500, 0)
		return
	}
	utils.Response(w, consts.Success)
}

// @Summary addNewPromotion
// @Security ApiKeyAuth
// @Tags Settings
// @Description Создание новой акции
// @ID addNewPromotion
// @Accept json
// @Produce json
// @Param params body models.Promotion true "Введите данные"
// @Success 200 {object} utils.DataResponse
// @Failure 500 {object} utils.DataResponse
// @Failure 400 {object} utils.DataResponse
// @Failure default {object} utils.DataResponse
// @Router /settings/promotions [post]
func (h *Handler) addNewPromotion(w http.ResponseWriter, r *http.Request) {
	var params *models.Promotion
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		h.logger.Error(err)
		utils.ErrorResponse(w, consts.InvalidRequestData, 400, 0)
		return
	}
	promotion, err := h.service.AddNewPromotion(params)
	if err != nil {
		h.logger.Error(err)
		utils.ErrorResponse(w, consts.InternalServerError, 500, 0)
		return
	}
	utils.Response(w, promotion)
}

// @Summary getAllPromotions
// @Security ApiKeyAuth
// @Tags Settings
// @Description Просмотр всех акций
// @ID getAllPromotions
// @Accept json
// @Produce json
// @Param page query string true "Введите данные"
// @Param limit query string true "Введите данные"
// @Success 200 {object} utils.DataResponse
// @Failure 500 {object} utils.DataResponse
// @Failure 400 {object} utils.DataResponse
// @Failure default {object} utils.DataResponse
// @Router /settings/promotions [get]
func (h *Handler) getAllPromotions(w http.ResponseWriter, r *http.Request) {
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
	promotions, count, err := h.service.GetAllPromotions(page, limit)
	if err != nil {
		h.logger.Error(err)
		utils.ErrorResponse(w, consts.InternalServerError, 500, 0)
		return
	}
	utils.Response(w, map[string]interface{}{
		"reports":     promotions,
		"total_count": count,
		"count":       len(promotions),
		"total_page":  int(math.Ceil(float64(count) / float64(limit))),
	})
}

// @Summary getPromotionById
// @Security ApiKeyAuth
// @Tags Settings
// @Description Просмотр подробной информации по акции
// @ID getPromotionById
// @Accept json
// @Produce json
// @Param promotion_id query string true "Введите данные"
// @Success 200 {object} utils.DataResponse
// @Failure 500 {object} utils.DataResponse
// @Failure 400 {object} utils.DataResponse
// @Failure default {object} utils.DataResponse
// @Router /settings/promotions [get]
func (h *Handler) getPromotionById(w http.ResponseWriter, r *http.Request) {
	promotionIdStr := mux.Vars(r)["promotion_id"]
	promotionId, err := strconv.Atoi(promotionIdStr)
	if err != nil {
		h.logger.Error(err)
		utils.ErrorResponse(w, consts.CountErrResponse, 400, 0)
		return
	}
	promotion, err := h.service.GetPromotionById(promotionId)
	if err != nil {
		h.logger.Error(err)
		utils.ErrorResponse(w, consts.InternalServerError, 500, 0)
		return
	}
	utils.Response(w, promotion)
}

// @Summary editPromotion
// @Security ApiKeyAuth
// @Tags Settings
// @Description Редактирование акции
// @ID editPromotion
// @Accept json
// @Produce json
// @Param params body models.Promotion true "Введите данные"
// @Success 200 {object} utils.DataResponse
// @Failure 500 {object} utils.DataResponse
// @Failure 400 {object} utils.DataResponse
// @Failure default {object} utils.DataResponse
// @Router /settings/promotions [put]
func (h *Handler) editPromotion(w http.ResponseWriter, r *http.Request) {
	var promotion *models.Promotion
	err := json.NewDecoder(r.Body).Decode(&promotion)
	if err != nil {
		h.logger.Error(err)
		utils.ErrorResponse(w, consts.InvalidRequestData, 400, 0)
		return
	}
	err = h.service.EditPromotion(promotion)
	if err != nil {
		h.logger.Error(err)
		utils.ErrorResponse(w, consts.InternalServerError, 500, 0)
		return
	}
	utils.Response(w, consts.Success)
}

// @Summary deletePromotion
// @Security ApiKeyAuth
// @Tags Settings
// @Description Удаление акции
// @ID deletePromotion
// @Accept json
// @Produce json
// @Param promotion_id query string true "Введите данные"
// @Success 200 {object} utils.DataResponse
// @Failure 500 {object} utils.DataResponse
// @Failure 400 {object} utils.DataResponse
// @Failure default {object} utils.DataResponse
// @Router /settings/promotions [delete]
func (h *Handler) deletePromotion(w http.ResponseWriter, r *http.Request) {
	promotionIdStr := mux.Vars(r)["promotion_id"]
	promotionId, err := strconv.Atoi(promotionIdStr)
	if err != nil {
		h.logger.Error(err)
		utils.ErrorResponse(w, consts.InvalidBrandId, 400, 0)
		return
	}
	err = h.service.DeletePromotion(promotionId)
	if err != nil {
		h.logger.Error(err)
		utils.ErrorResponse(w, consts.InternalServerError, 500, 0)
		return
	}
	utils.Response(w, consts.Success)
}

// @Summary addNewRole
// @Security ApiKeyAuth
// @Tags Settings
// @Description Создание новой роли
// @ID addNewRole
// @Accept json
// @Produce json
// @Param params body models.RoleInput true "Введите данные"
// @Success 200 {object} utils.DataResponse
// @Failure 500 {object} utils.DataResponse
// @Failure 400 {object} utils.DataResponse
// @Failure default {object} utils.DataResponse
// @Router /settings/roles [post]
func (h *Handler) addNewRole(w http.ResponseWriter, r *http.Request) {
	var params *models.RoleInput
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		h.logger.Error(err)
		utils.ErrorResponse(w, consts.InvalidRequestData, 400, 0)
		return
	}
	role, err := h.service.AddNewRole(params)
	if err != nil {
		h.logger.Error(err)
		utils.ErrorResponse(w, consts.InternalServerError, 500, 0)
		return
	}
	utils.Response(w, role)
}

// @Summary getAllRoles
// @Security ApiKeyAuth
// @Tags Settings
// @Description Просмотр всех ролей
// @ID getAllRoles
// @Accept json
// @Produce json
// @Param page query string true "Введите данные"
// @Param limit query string true "Введите данные"
// @Success 200 {object} utils.DataResponse
// @Failure 500 {object} utils.DataResponse
// @Failure 400 {object} utils.DataResponse
// @Failure default {object} utils.DataResponse
// @Router /settings/roles [get]
func (h *Handler) getAllRoles(w http.ResponseWriter, r *http.Request) {
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
	promotions, count, err := h.service.GetAllRoles(page, limit)
	if err != nil {
		h.logger.Error(err)
		utils.ErrorResponse(w, consts.InternalServerError, 500, 0)
		return
	}
	utils.Response(w, map[string]interface{}{
		"reports":     promotions,
		"total_count": count,
		"count":       len(promotions),
		"total_page":  int(math.Ceil(float64(count) / float64(limit))),
	})
}

// @Summary getRoleById
// @Security ApiKeyAuth
// @Tags Settings
// @Description Просмотр подробной информации по роли
// @ID getRoleById
// @Accept json
// @Produce json
// @Param role_id query string true "Введите данные"
// @Success 200 {object} utils.DataResponse
// @Failure 500 {object} utils.DataResponse
// @Failure 400 {object} utils.DataResponse
// @Failure default {object} utils.DataResponse
// @Router /settings/role [get]
func (h *Handler) getRoleById(w http.ResponseWriter, r *http.Request) {
	roleIdStr := mux.Vars(r)["role_id"]
	roleId, err := strconv.Atoi(roleIdStr)
	if err != nil {
		h.logger.Error(err)
		utils.ErrorResponse(w, consts.CountErrResponse, 400, 0)
		return
	}
	role, err := h.service.GetRoleById(roleId)
	if err != nil {
		h.logger.Error(err)
		utils.ErrorResponse(w, consts.InternalServerError, 500, 0)
		return
	}
	utils.Response(w, role)
}

// @Summary editRoleById
// @Security ApiKeyAuth
// @Tags Settings
// @Description Редактирование роли
// @ID editRoleById
// @Accept json
// @Produce json
// @Param params body models.RoleInput true "Введите данные"
// @Success 200 {object} utils.DataResponse
// @Failure 500 {object} utils.DataResponse
// @Failure 400 {object} utils.DataResponse
// @Failure default {object} utils.DataResponse
// @Router /settings/roles [put]
func (h *Handler) editRoleById(w http.ResponseWriter, r *http.Request) {
	var role *models.RoleInput
	err := json.NewDecoder(r.Body).Decode(&role)
	if err != nil {
		h.logger.Error(err)
		utils.ErrorResponse(w, consts.InvalidRequestData, 400, 0)
		return
	}
	err = h.service.EditRole(role)
	if err != nil {
		h.logger.Error(err)
		utils.ErrorResponse(w, consts.InternalServerError, 500, 0)
		return
	}
	utils.Response(w, consts.Success)
}

// @Summary deleteRole
// @Security ApiKeyAuth
// @Tags Settings
// @Description Удаление роли
// @ID deleteRole
// @Accept json
// @Produce json
// @Param role_id query string true "Введите данные"
// @Success 200 {object} utils.DataResponse
// @Failure 500 {object} utils.DataResponse
// @Failure 400 {object} utils.DataResponse
// @Failure default {object} utils.DataResponse
// @Router /settings/roles [delete]
func (h *Handler) deleteRole(w http.ResponseWriter, r *http.Request) {
	roleIdStr := mux.Vars(r)["role_id"]
	roleId, err := strconv.Atoi(roleIdStr)
	if err != nil {
		h.logger.Error(err)
		utils.ErrorResponse(w, consts.InvalidBrandId, 400, 0)
		return
	}
	err = h.service.DeletePromotion(roleId)
	if err != nil {
		h.logger.Error(err)
		utils.ErrorResponse(w, consts.InternalServerError, 500, 0)
		return
	}
	utils.Response(w, consts.Success)
}

// @Summary getAllPermissions
// @Security ApiKeyAuth
// @Tags Settings
// @Description Просмотр всех доступов
// @ID getAllPermissions
// @Accept json
// @Produce json
// @Success 200 {object} utils.DataResponse
// @Failure 500 {object} utils.DataResponse
// @Failure 400 {object} utils.DataResponse
// @Failure default {object} utils.DataResponse
// @Router /settings/permissions [get]
func (h *Handler) getAllPermissions(w http.ResponseWriter, r *http.Request) {
	permissions, err := h.service.Settings.GetAllPermissions()
	if err != nil {
		h.logger.Error(err)
		utils.ErrorResponse(w, consts.InternalServerError, 500, 0)
		return
	}
	utils.Response(w, permissions)
}

// @Summary Скачивание логов
// @Security ApiKeyAuth
// @Tags Settings
// @Description Просмотр логов системы
// @ID downloadSystemLogs
// @Produce json
// @Success 200 {object} utils.DataResponse
// @Failure 400 {object} utils.DataResponse
// @Failure 500 {object} utils.DataResponse
// @Failure default {object} utils.DataResponse
// @Router /settings/download-logs [get]
func (h *Handler) downloadSystemLogs(w http.ResponseWriter, r *http.Request) {
	filePath := consts.LogPath
	file, err := os.Open(filePath)
	if err != nil {
		h.logger.Error(err)
		utils.ErrorResponse(w, consts.InternalServerError, 500, 0)
		return
	}
	info, err := os.Stat(filePath)
	if err != nil {
		h.logger.Error(err)
		utils.ErrorResponse(w, consts.InternalServerError, 500, 0)
		return
	}
	size := info.Size()
	defer file.Close()
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", info.Name()))
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Transfer-Encoding", "binary")
	w.Header().Set("Content-Length", fmt.Sprintf("%d", size))
	_, err = io.Copy(w, file)
	if err != nil {
		h.logger.Error(err)
		utils.ErrorResponse(w, consts.InternalServerError, 500, 0)
		return
	}
}
