package delivery

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo"

	"github.com/betawulan/synapsis/error_message"
	"github.com/betawulan/synapsis/model"
	"github.com/betawulan/synapsis/service"
)

type shoppingCartDelivery struct {
	shoppingCartService service.ShoppingCartService
}

func (s shoppingCartDelivery) create(c echo.Context) error {
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

	err = s.shoppingCartService.Create(c.Request().Context(), tokens[1], shoppingCart)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, "success")
}

func (s shoppingCartDelivery) delete(c echo.Context) error {
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

	ID := c.Param("id")
	IDint, err := strconv.Atoi(ID)
	if err != nil {
		return err
	}

	err = s.shoppingCartService.Delete(c.Request().Context(), tokens[1], int64(IDint))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusNoContent, nil)
}

func (s shoppingCartDelivery) read(c echo.Context) error {
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

	shoppingCart, err := s.shoppingCartService.Read(c.Request().Context(), tokens[1])
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, shoppingCart)
}

func AddShoppingCartRoute(shoppingCartService service.ShoppingCartService, e *echo.Echo) {
	handler := shoppingCartDelivery{
		shoppingCartService: shoppingCartService,
	}

	e.GET("/shopping-cart", handler.read)
	e.POST("/shopping-cart", handler.create)
	e.DELETE("/shopping-cart/:id", handler.delete)
}
