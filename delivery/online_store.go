package delivery

import (
	"net/http"
	"strings"

	"github.com/betawulan/synapsis/error_message"
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

func (o onlineStoreDelivery) create(c echo.Context) error {
	token := c.Request().Header.Get("Authorization")
	if token == "" {
		return c.JSON(http.StatusUnauthorized, error_message.Unauthorized{Message: "format token invalid"})
	}

	tokens := strings.Split(token, " ")
	if len(tokens) < 2 {
		return c.JSON(http.StatusUnauthorized, error_message.Unauthorized{Message: "format token invalid"})
	}

	if tokens[0] != "Bearer" {
		return c.JSON(http.StatusUnauthorized, error_message.Unauthorized{Message: "no Bearer"})
	}

	var shoppingCart model.ShoppingCart

	err := c.Bind(&shoppingCart)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	shoppingCart, err = o.onlineStoreService.Create(c.Request().Context(), tokens[1], shoppingCart)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, shoppingCart)
}

func RegisterOnlineStoreRoute(onlineStoreService service.OnlineStoreService, e *echo.Echo) {
	handler := onlineStoreDelivery{
		onlineStoreService: onlineStoreService,
	}

	e.GET("/product", handler.fetch)
	e.POST("/product", handler.create)
}
