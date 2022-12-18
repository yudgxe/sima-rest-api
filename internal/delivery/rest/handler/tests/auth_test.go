package handler_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/yudgxe/sima-rest-api/internal/delivery/rest/handler"
	"github.com/yudgxe/sima-rest-api/internal/service/mock"
)

func TestHandleSingIn(t *testing.T) {
	e := echo.New()
	auth := e.Group("")
	as := new(mock.AuthService)

	handler.NewAuthHandler(as).Init(auth)

	in := handler.AuthSingInInput{
		Login:    "login",
		Password: "password",
	}

	body, err := json.Marshal(in)
	assert.NoError(t, err)

	as.On("GenerateToken", in.Login, in.Password).Return("token", nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/sing-in", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	e.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}
