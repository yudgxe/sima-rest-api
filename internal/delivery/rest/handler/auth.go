package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/yudgxe/sima-rest-api/internal/service"
)

type Auth struct {
	service service.AuthService
}

func NewAuthHandler(s service.AuthService) *Auth {
	return &Auth{
		service: s,
	}
}

func (a *Auth) Init(e *echo.Group) {
	e.POST("/sing-in", a.handleSingIn())
}

func (a *Auth) handleSingIn() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		req := &AuthSingInInput{}

		if err := ctx.Bind(req); err != nil {
			if he, ok := err.(*echo.HTTPError); ok {
				return ctx.JSON(http.StatusBadRequest, newErrorResponce("error", he.Message.(string)))
			}

			return ctx.JSON(http.StatusBadRequest, newErrorResponce("error", err.Error()))
		}

		token, err := a.service.GenerateToken(req.Login, req.Password)
		if err != nil {
			return ctx.JSON(http.StatusUnauthorized, newErrorResponce("error", err.Error()))
		}

		return ctx.JSON(http.StatusOK, newSuccessResponce("Success", map[string]string{
			"token": token,
		}))
	}
}
