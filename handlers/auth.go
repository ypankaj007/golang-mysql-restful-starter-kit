package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"golang-mysql-restful-starter-kit/helpers"
	"golang-mysql-restful-starter-kit/models"
	"golang-mysql-restful-starter-kit/services"
)

type AuthHandler struct {
	userSrv services.IUserService
}

func NewAuthHandler(userSrv services.IUserService) *AuthHandler {
	return &AuthHandler{userSrv: userSrv}
}

func (userHandler *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var payload helpers.CreateUser
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&payload); err != nil {
		helpers.SendRespond(w, http.StatusBadRequest, nil, helpers.ErrBadRequest)
		return
	}
	user, err := models.NewUser(payload.Name, payload.Email, payload.Password)
	if err != nil {
		helpers.SendRespond(w, http.StatusBadRequest, nil, err.Error())
		return
	}
	defer r.Body.Close()
	err = userHandler.userSrv.CreateUser(context.TODO(), user)
	if err != nil {
		helpers.SendRespond(w, http.StatusBadRequest, nil, err.Error())
		return
	}
	helpers.SendRespond(w, http.StatusCreated, nil, helpers.UserCreated)

}

func (userHandler *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var payload helpers.LoginPayload
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&payload); err != nil {
		helpers.SendRespond(w, http.StatusBadRequest, nil, helpers.ErrBadRequest)
		return
	}
	defer r.Body.Close()
	resp, err := userHandler.userSrv.Login(context.TODO(), payload)
	if err != nil {
		helpers.SendRespond(w, http.StatusBadRequest, nil, err.Error())
		return
	}
	helpers.SendRespond(w, http.StatusOK, resp, helpers.SuccessLogin)

}
