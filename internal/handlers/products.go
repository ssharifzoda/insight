package api

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"insight/internal/models"
	"insight/pkg/consts"
	"insight/pkg/utils"
	"net/http"
	"strconv"
)

// @Summary addNewProduct
// @Security ApiKeyAuth
// @Tags Products
// @Description Добавление нового продукта
// @ID addNewProduct
// @Accept json
// @Produce json
// @Param params body models.Product true "Введите данные"
// @Success 200 {object} utils.DataResponse
// @Failure 500 {object} utils.DataResponse
// @Failure 400 {object} utils.DataResponse
// @Failure default {object} utils.DataResponse
// @Router /products/new [post]
func (h *Handler) addNewProduct(w http.ResponseWriter, r *http.Request) {
	var params *models.Product
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		h.logger.Error(err)
		utils.ErrorResponse(w, consts.InvalidRequestData, 400, 0)
		return
	}
	err = h.service.Products.AddNewProduct(params)
	if err != nil {
		h.logger.Error(err)
		utils.ErrorResponse(w, consts.InternalServerError, 500, 0)
		return
	}

	utils.Response(w, consts.Success)
}

// @Summary editProduct
// @Security ApiKeyAuth
// @Tags Products
// @Description Редактирование продукта
// @ID editProduct
// @Accept json
// @Produce json
// @Param params body models.Product true "Введите данные"
// @Success 200 {object} utils.DataResponse
// @Failure 500 {object} utils.DataResponse
// @Failure 400 {object} utils.DataResponse
// @Failure default {object} utils.DataResponse
// @Router /products/edit [put]
func (h *Handler) editProduct(w http.ResponseWriter, r *http.Request) {
	var params *models.Product
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		h.logger.Error(err)
		utils.ErrorResponse(w, consts.InvalidRequestData, 400, 0)
		return
	}
	err = h.service.Products.EditProduct(params)
	if err != nil {
		h.logger.Error(err)
		utils.ErrorResponse(w, consts.InternalServerError, 500, 0)
		return
	}
	utils.Response(w, consts.Success)
}

// @Summary getAllProducts
// @Security ApiKeyAuth
// @Tags Products
// @Description Просмотр списока всех продуктов
// @ID getAllProducts
// @Accept json
// @Produce json
// @Param page query string true "Введите данные"
// @Param limit query string true "Введите данные"
// @Param search query string false "Введите данные"
// @Param brand_id query string false "Введите данные"
// @Param category_id query string false "Введите данные"
// @Param supplier_id query string false "Введите данные"
// @Success 200 {object} utils.DataResponse
// @Failure 500 {object} utils.DataResponse
// @Failure 400 {object} utils.DataResponse
// @Failure default {object} utils.DataResponse
// @Router /products/list [get]
func (h *Handler) getAllProducts(w http.ResponseWriter, r *http.Request) {
	pageStr := mux.Vars(r)["page"]
	limitStr := mux.Vars(r)["limit"]
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		h.logger.Error(err)
		utils.ErrorResponse(w, consts.PageErrResponse, 400, 0)
		return
	}
	brandStr := r.URL.Query().Get("brand_id")
	brandId, err := strconv.Atoi(brandStr)
	if err != nil && !errors.Is(err, strconv.ErrSyntax) {
		h.logger.Error(err)
		utils.ErrorResponse(w, consts.FilterError, 400, 0)
		return
	}
	categoryStr := r.URL.Query().Get("category_id")
	categoryId, err := strconv.Atoi(categoryStr)
	if err != nil && !errors.Is(err, strconv.ErrSyntax) {
		h.logger.Error(err)
		utils.ErrorResponse(w, consts.FilterError, 400, 0)
		return
	}
	supplierStr := r.URL.Query().Get("supplier_id")
	supplierId, err := strconv.Atoi(supplierStr)
	if err != nil && !errors.Is(err, strconv.ErrSyntax) {
		h.logger.Error(err)
		utils.ErrorResponse(w, consts.FilterError, 400, 0)
		return
	}
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		h.logger.Error(err)
		utils.ErrorResponse(w, consts.CountErrResponse, 400, 0)
		return
	}
	search := r.URL.Query().Get("search")
	filter := models.ProductFilter{
		BrandId:    brandId,
		SupplierId: supplierId,
		CategoryId: categoryId,
		Search:     search,
	}
	products, err := h.service.Products.GetAllProducts(page, limit, &filter)
	if err != nil {
		h.logger.Error(err)
		utils.ErrorResponse(w, consts.InternalServerError, 500, 0)
		return
	}
	utils.Response(w, products)
}

// @Summary getProduct
// @Security ApiKeyAuth
// @Tags Products
// @Description Просмотр магазина
// @ID getProduct
// @Accept json
// @Produce json
// @Param product_id query string true "Введите данные"
// @Success 200 {object} utils.DataResponse
// @Failure 500 {object} utils.DataResponse
// @Failure 400 {object} utils.DataResponse
// @Failure default {object} utils.DataResponse
// @Router /products/by-id [get]
func (h *Handler) getProduct(w http.ResponseWriter, r *http.Request) {
	productIdStr := mux.Vars(r)["product_id"]
	productId, err := strconv.Atoi(productIdStr)
	if err != nil {
		h.logger.Error(err)
		utils.ErrorResponse(w, consts.InvalidRequestData, 400, 0)
		return
	}
	product, err := h.service.Products.GetProductById(productId)
	if err != nil {
		h.logger.Error(err)
		utils.ErrorResponse(w, consts.InternalServerError, 500, 0)
		return
	}
	utils.Response(w, product)
}

// @Summary deleteProduct
// @Security ApiKeyAuth
// @Tags Products
// @Description Удаление продукта
// @ID deleteProduct
// @Accept json
// @Produce json
// @Param product_id query string true "Введите данные"
// @Success 200 {object} utils.DataResponse
// @Failure 500 {object} utils.DataResponse
// @Failure 400 {object} utils.DataResponse
// @Failure default {object} utils.DataResponse
// @Router /products/rm [delete]
func (h *Handler) deleteProduct(w http.ResponseWriter, r *http.Request) {
	productIdStr := mux.Vars(r)["product_id"]
	productId, err := strconv.Atoi(productIdStr)
	if err != nil {
		h.logger.Error(err)
		utils.ErrorResponse(w, consts.InvalidRequestData, 400, 0)
		return
	}
	err = h.service.Products.DeleteProduct(productId)
	if err != nil {
		h.logger.Error(err)
		utils.ErrorResponse(w, consts.InternalServerError, 500, 0)
		return
	}
	utils.Response(w, consts.Success)
}
