package delivery

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/betawulan/synapsis/model"
	"github.com/betawulan/synapsis/service"
)

type productDelivery struct {
	productService service.ProductService
}

//	@Summary		fetch
//	@Description	view product list by product category
//	@Tags			product
//  @Param category query string false "category"
//	@Success		200 {array} []model.ProductCategory
//	@Failure		500
//	@Router			/product [get]
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
