package controllers

import (
	"github.com/jibaru/home-inventory-api/m/internal/app/application/services"
	"github.com/jibaru/home-inventory-api/m/internal/app/infrastructure/responses"
	"github.com/labstack/echo/v4"
	"net/http"
)

type GetItemsController struct {
	assetService *services.AssetService
	itemService  *services.ItemService
}

type GetItemsRequest struct {
	Search  string `query:"search"`
	Page    int    `query:"page"`
	PerPage int    `query:"per_page"`
}

type GetItemsResponse struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Description *string  `json:"description"`
	Unit        string   `json:"unit"`
	Keywords    []string `json:"keywords"`
	Assets      []struct {
		ID  string `json:"id"`
		Url string `json:"url"`
	} `json:"assets"`
}

func NewGetItemsController(
	assetService *services.AssetService,
	itemService *services.ItemService,
) *GetItemsController {
	return &GetItemsController{
		assetService,
		itemService,
	}
}

func (c *GetItemsController) Handle(ctx echo.Context) error {
	userID := ctx.Get("auth_id").(string)
	request := GetItemsRequest{}

	err := (&echo.DefaultBinder{}).BindQueryParams(ctx, &request)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, responses.NewMessageResponse(err.Error()))
	}

	items, err := c.itemService.GetAll(
		request.Search,
		userID,
		services.PageFilter{
			Page: request.Page,
			Size: request.PerPage,
		},
	)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, responses.NewMessageResponse(err.Error()))
	}

	total, err := c.itemService.CountAll(
		request.Search,
		userID,
	)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, responses.NewMessageResponse(err.Error()))
	}

	responseItems := make([]*GetItemsResponse, 0)
	for _, item := range items {
		data := &GetItemsResponse{
			ID:          item.Item.ID,
			Name:        item.Item.Name,
			Description: item.Item.Description,
			Unit:        item.Item.Unit,
			Keywords:    make([]string, 0),
			Assets: make([]struct {
				ID  string `json:"id"`
				Url string `json:"url"`
			}, 0),
		}

		for _, asset := range item.Assets {
			data.Assets = append(data.Assets, struct {
				ID  string `json:"id"`
				Url string `json:"url"`
			}{
				ID:  asset.ID,
				Url: c.assetService.GetUrl(asset),
			})
		}

		for _, keyword := range item.Item.Keywords {
			data.Keywords = append(data.Keywords, keyword.Value)
		}

		responseItems = append(responseItems, data)
	}

	return ctx.JSON(http.StatusOK, responses.NewPaginatedResponse(
		responseItems,
		total,
		request.Page,
		request.PerPage,
		len(items),
		ctx.Request().URL.Path,
	))
}
