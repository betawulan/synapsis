package delivery

import (
	"net/http"

	"github.com/betawulan/synapsis/model"
	"github.com/betawulan/synapsis/service"
	"github.com/labstack/echo"
)

type onlineStoreDelivery struct {
	onlineStoreService service.OnlineStoreService
}

func (o onlineStoreDelivery) fetch(c echo.Context) error {
	filter := model.ProductCategoryFilter{}
	filter.Category = c.QueryParam("category")

	productCategories, err := o.onlineStoreService.Fetch(c.Request().Context(), filter)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, productCategories)
}

func RegisterOnlineStoreRoute(onlineStoreService service.OnlineStoreService, e *echo.Echo) {
	handler := onlineStoreDelivery{
		onlineStoreService: onlineStoreService,
	}

	e.GET("/product", handler.fetch)
}
