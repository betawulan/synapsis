package delivery

import (
	"net/http"
	"strings"

	"github.com/labstack/echo"

	"github.com/betawulan/synapsis/error_message"
	"github.com/betawulan/synapsis/service"
)

type transactionDelivery struct {
	transactionService service.TransactionService
}

type inputProductCategoryIDs struct {
	ProductCategoryIDs []int `json:"product_category_ids"`
}

//	@Summary		checkout
//	@Description	checkout product
//	@Tags			transaction
//	@Param			Authorization header string true "Bearer token"
//  @Param          productCategoryIDs body inputProductCategoryIDs true "request"
//	@Success		201 {object} model.TransactionResponse
//	@Failure		500
//	@Router			/checkout [post]
func (t transactionDelivery) checkout(c echo.Context) error {
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

	var input inputProductCategoryIDs

	err := c.Bind(&input)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	sumPrice, err := t.transactionService.Checkout(c.Request().Context(), tokens[1], input.ProductCategoryIDs)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, sumPrice)
}

func AddTransactionRoute(transactionService service.TransactionService, e *echo.Echo) {
	handler := transactionDelivery{
		transactionService: transactionService,
	}

	e.POST("/checkout", handler.checkout)

}
