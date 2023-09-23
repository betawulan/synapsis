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

//	@Summary		create
//	@Description	add product to shopping cart
//	@Tags			shopping cart
//	@Param			Authorization header string true "Bearer token"
//  @Param          shoppingCart body model.ShoppingCart true "request"
//	@Success		201				
//	@Failure		500			
//	@Router			/shopping-cart [post]
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

//	@Summary		delete
//	@Description	delete product list in shopping cart
//	@Tags			shopping cart
//	@Param			Authorization header string true "Bearer token"
//  @Param          id path integer true "id of shopping cart"
//	@Success		204				
//	@Failure		500			
//	@Router			/shopping-cart/{id} [delete]
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

// @Summary 	read 
// @Description see a list of products that have been added to the shopping cart
// @Tags 		shopping cart
//	@Param			Authorization header string true "Bearer token"
// @Success 	200 {array} []model.ShoppingCart
// @Failure 	500 
// @Router 		/shopping-cart [get]
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
