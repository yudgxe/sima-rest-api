package handler

import (
	"errors"
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/yudgxe/sima-rest-api/internal/service"
)

type Deps struct {
	UserService service.UserService
	AuthService service.AuthService
}

type Handler struct {
	User *User
	Auth *Auth
}

func NewHandler(d *Deps) *Handler {
	return &Handler{
		User: NewUserHanlder(d.UserService),
		Auth: NewAuthHandler(d.AuthService),
	}
}

func (h *Handler) Init(e *echo.Echo, key string) {
	auth := e.Group("")
	h.Auth.Init(auth)

	config := middleware.JWTConfig{
		ParseTokenFunc: func(auth string, c echo.Context) (interface{}, error) {
			token, err := jwt.Parse(auth, func(t *jwt.Token) (interface{}, error) {
				if t.Method.Alg() != "HS256" {
					return nil, fmt.Errorf("unexpected jwt signing method=%v", t.Header["alg"])
				}
				return []byte(key), nil
			})
			if err != nil {
				return nil, err
			}

			if !token.Valid {
				return nil, errors.New("invalid token")
			}

			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok {
				return nil, errors.New("token claims are not of type jwt.MapClaims")
			}
			c.Set("permission", claims["permission"])
			return token, nil
		},
	}

	api := auth.Group("/api", middleware.JWTWithConfig(config))
	h.User.Init(api)
}
