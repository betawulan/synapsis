package delivery

import (
	"net/http"

	"github.com/labstack/echo"

	"github.com/betawulan/synapsis/model"
	"github.com/betawulan/synapsis/service"
)

type productDelivery struct {
	productService service.ProductService
}

func (p productDelivery) fetch(c echo.Context) error {
	filter := model.ProductCategoryFilter{}
	filter.Category = c.QueryParam("category")

	productCategories, err := p.productService.Fetch(c.Request().Context(), filter)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, productCategories)
}

func AddProductRoute(productService service.ProductService, e *echo.Echo) {
	handler := productDelivery{
		productService: productService,
	}

	e.GET("/product", handler.fetch)
}
