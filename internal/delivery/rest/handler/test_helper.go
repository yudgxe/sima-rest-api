package handler

import (
	"testing"

	"github.com/labstack/echo/v4"
)

func GetServerAndGroups(t *testing.T) (*echo.Echo, *echo.Group, *echo.Group) {
	t.Helper()

	e := echo.New()
	admin := e.Group("/api", func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("permission", "admin")
			return next(c)
		}
	})
	read := e.Group("/api", func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("permission", "read")
			return next(c)
		}
	})

	return e, admin, read
}
