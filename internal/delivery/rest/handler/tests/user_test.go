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
	"github.com/yudgxe/sima-rest-api/internal/model"
	"github.com/yudgxe/sima-rest-api/internal/service/mock"
)

func TestHandleGet(t *testing.T) {
	e := echo.New()
	api := e.Group("/api")

	us := new(mock.UserService)

	handler.NewUserHanlder(us).Init(api)

	user := model.GetTestUser(t)
	us.On("GetUser", user.Login).Return(user, nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/users/login", nil)

	e.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}

func TestHandleDeleteUser(t *testing.T) {
	e, admin, read := handler.GetServerAndGroups(t)

	us := new(mock.UserService)
	user := model.GetTestUser(t)

	us.On("DeleteUser", user.Login).Return(nil)

	handler.NewUserHanlder(us).Init(admin)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/api/users/login", nil)

	e.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

	handler.NewUserHanlder(us).Init(read)
	w = httptest.NewRecorder()

	e.ServeHTTP(w, req)
	assert.Equal(t, 403, w.Code)
}

func TestHandleUpdateUser(t *testing.T) {
	user := model.GetTestUser(t)
	in := &handler.UserInput{
		Name:      user.Name,
		Surname:   user.Surname,
		Login:     user.Login,
		Password:  user.Password,
		Birthdate: user.Birthdate.Format("01/02/2006"),
	}

	body, err := json.Marshal(in)
	assert.NoError(t, err)

	e, admin, read := handler.GetServerAndGroups(t)

	us := new(mock.UserService)
	us.On("UpdateUser", user, user.Login).Return(nil)

	handler.NewUserHanlder(us).Init(admin)
	w := httptest.NewRecorder()

	req, _ := http.NewRequest("PUT", "/api/users/login", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	e.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

	handler.NewUserHanlder(us).Init(read)
	w = httptest.NewRecorder()

	e.ServeHTTP(w, req)
	assert.Equal(t, 403, w.Code)

}
func TestHandleCreateWithPermission(t *testing.T) {
	user := model.GetTestUser(t)
	in := &handler.UserWithPermissionInput{
		Name:       user.Name,
		Surname:    user.Surname,
		Login:      user.Login,
		Password:   user.Password,
		Birthdate:  user.Birthdate.Format("01/02/2006"),
		Permission: "admin",
	}

	body, err := json.Marshal(in)
	assert.NoError(t, err)

	e, admin, read := handler.GetServerAndGroups(t)

	us := new(mock.UserService)
	us.On("CreateWithPermission", user, in.Permission).Return(nil)

	handler.NewUserHanlder(us).Init(admin)
	w := httptest.NewRecorder()

	req, _ := http.NewRequest("POST", "/api/users", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	e.ServeHTTP(w, req)
	assert.Equal(t, 201, w.Code)

	handler.NewUserHanlder(us).Init(read)
	w = httptest.NewRecorder()

	e.ServeHTTP(w, req)
	assert.Equal(t, 403, w.Code)
}
