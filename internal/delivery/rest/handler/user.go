package handler

import (
	"errors"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/yudgxe/sima-rest-api/internal/model"
	"github.com/yudgxe/sima-rest-api/internal/service"
)

type User struct {
	service service.UserService
}

func NewUserHanlder(us service.UserService) *User {
	return &User{
		service: us,
	}
}

func (us *User) Init(e *echo.Group) {
	e.GET("/users/:login", us.handleGet())
	e.POST("/users", us.handleCreateWithPermission())
	e.DELETE("/users/:login", us.handleDelete())
	e.PUT("/users/:login", us.handleUpdate())
}

func (us *User) handleGet() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		if ctx.Get("permission") == "banned" {
			return ctx.JSON(http.StatusForbidden, newErrorResponce("forbidden", "you got banned"))
		}

		login := ctx.Param("login")
		user, err := us.service.GetUser(login)
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, newErrorResponce("error", err.Error()))
		}

		return ctx.JSON(http.StatusOK, newSuccessResponce("success", user))
	}
}

func (us *User) handleDelete() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		if ctx.Get("permission") != "admin" {
			return ctx.JSON(http.StatusForbidden, newErrorResponce("forbidden", "you don't have enough rights"))
		}
		login := ctx.Param("login")
		if err := us.service.DeleteUser(login); err != nil {
			return ctx.JSON(http.StatusInternalServerError, newErrorResponce("error", err.Error()))
		}
		return ctx.JSON(http.StatusOK, newSuccessResponce("success", nil))
	}
}

func (us *User) handleUpdate() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		if ctx.Get("permission") != "admin" {
			return ctx.JSON(http.StatusForbidden, newErrorResponce("forbidden", "you don't have enough rights"))
		}
		user, err := bindAndParseUser(ctx)
		if err != nil {
			return ctx.JSON(http.StatusBadRequest, newErrorResponce("error", err.Error()))
		}

		login := ctx.Param("login")
		if err := us.service.UpdateUser(user, login); err != nil {
			return ctx.JSON(http.StatusInternalServerError, newErrorResponce("error", err.Error()))
		}

		return ctx.JSON(http.StatusOK, newSuccessResponce("success", nil))
	}
}

func (us *User) handleCreateWithPermission() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		if ctx.Get("permission") != "admin" {
			return ctx.JSON(http.StatusForbidden, newErrorResponce("forbidden", "you don't have enough rights"))
		}

		req := &UserWithPermissionInput{}

		if err := ctx.Bind(req); err != nil {
			if he, ok := err.(*echo.HTTPError); ok {
				return ctx.JSON(http.StatusBadRequest, newErrorResponce("error", he.Message.(string)))
			}

			return ctx.JSON(http.StatusBadRequest, newErrorResponce("error", err.Error()))
		}

		date, err := time.Parse("01/02/2006", req.Birthdate)
		if err != nil {
			return ctx.JSON(http.StatusBadRequest, newErrorResponce("error", err.Error()))
		}

		user := &model.User{
			Name:      req.Name,
			Surname:   req.Surname,
			Login:     req.Login,
			Password:  req.Password,
			Birthdate: date,
		}

		if err := us.service.CreateWithPermission(user, req.Permission); err != nil {
			return ctx.JSON(http.StatusBadRequest, newErrorResponce("error", err.Error()))
		}

		return ctx.JSON(http.StatusCreated, newSuccessResponce("success", user))
	}
}

func bindAndParseUser(ctx echo.Context) (*model.User, error) {
	req := &UserInput{}

	if err := ctx.Bind(req); err != nil {
		if he, ok := err.(*echo.HTTPError); ok {
			return nil, errors.New(he.Message.(string))
		}

		return nil, err
	}

	date, err := time.Parse("01/02/2006", req.Birthdate)
	if err != nil {
		return nil, err
	}

	u := &model.User{
		Name:      req.Name,
		Surname:   req.Surname,
		Login:     req.Login,
		Password:  req.Password,
		Birthdate: date,
	}

	return u, nil
}
