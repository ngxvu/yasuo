package handler

import (
	"gitlab.com/merakilab9/meracore/ginext"
	"gitlab.com/merakilab9/yasuo/pkg/model"
	"gitlab.com/merakilab9/yasuo/pkg/service"
	"gitlab.com/merakilab9/yasuo/pkg/utils"
	"net/http"
)

type CategoryHandlers struct {
	service service.CategoryInterface
}

func NewCategoryHandlers(service service.CategoryInterface) *CategoryHandlers {
	return &CategoryHandlers{service: service}
}

func (h *CategoryHandlers) SaveCate(r *ginext.Request) (*ginext.Response, error) {

	requestModel := model.ShopeeCateData{}
	if err := r.GinCtx.ShouldBind(&requestModel); err != nil {
		return nil, ginext.NewError(http.StatusBadRequest, utils.MessageError()[http.StatusBadRequest])
	}

	if err := h.service.SaveCate(r.GinCtx, requestModel.Data.CategoryList); err != nil {
		return nil, err
	}

	return ginext.NewResponseData(http.StatusOK, requestModel), nil
}

func (h *CategoryHandlers) SaveShopsByCate(r *ginext.Request) (*ginext.Response, error) {

	requestModel := model.DataShopCrawled{}
	if err := r.GinCtx.ShouldBind(&requestModel); err != nil {
		return nil, ginext.NewError(http.StatusBadRequest, utils.MessageError()[http.StatusBadRequest])
	}

	if err := h.service.SaveShopsByCate(r.GinCtx, requestModel); err != nil {
		return nil, err
	}

	return ginext.NewResponseData(http.StatusOK, requestModel.Data), nil
}

func (h *CategoryHandlers) SaveShopsDetail(r *ginext.Request) (*ginext.Response, error) {

	requestModel := model.DataInfoShopCrawled{}
	if err := r.GinCtx.ShouldBind(&requestModel); err != nil {
		return nil, ginext.NewError(http.StatusBadRequest, utils.MessageError()[http.StatusBadRequest])
	}

	if err := h.service.SaveShopsInfo(r.GinCtx, requestModel); err != nil {
		return nil, err
	}

	return ginext.NewResponseData(http.StatusOK, requestModel.DataInfoShop), nil
}
