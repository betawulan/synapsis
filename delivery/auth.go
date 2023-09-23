package delivery

import (
	"net/http"

	"github.com/betawulan/synapsis/model"
	"github.com/betawulan/synapsis/service"
	"github.com/labstack/echo"
)

type authDelivery struct {
	authService service.AuthService
}

type credential struct {
	Role     string `json:"role"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type successLogin struct {
	Token string `json:"token"`
}

func (a authDelivery) register(c echo.Context) error {
	var user model.User

	err := c.Bind(&user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	err = a.authService.Register(c.Request().Context(), user)
	if err != nil {

		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, "success")
}

func (a authDelivery) login(c echo.Context) error {
	cred := credential{}

	err := c.Bind(&cred)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	token, err := a.authService.Login(c.Request().Context(), cred.Role, cred.Email, cred.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, successLogin{Token: token})
}

func AddAuthRoute(authService service.AuthService, e *echo.Echo) {
	handler := authDelivery{
		authService: authService,
	}

	e.POST("/auth/register", handler.register)
	e.POST("/auth/login", handler.login)
}
