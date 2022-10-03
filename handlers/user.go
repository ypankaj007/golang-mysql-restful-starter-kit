package handlers

import (
	"context"
	"golang-mysql-restful-starter-kit/helpers"
	"golang-mysql-restful-starter-kit/services"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type UserHandler struct {
	userSrv services.IUserService
}

func NewUserHandler(userSrv services.IUserService) *UserHandler {
	return &UserHandler{userSrv: userSrv}
}

func (userHandler *UserHandler) UserDetails(w http.ResponseWriter, r *http.Request) {
	strID := mux.Vars(r)["id"]
	ID, err := strconv.ParseInt(strID, 10, 32)
	if err != nil || ID < 1 {
		helpers.SendRespond(w, http.StatusBadRequest, nil, helpers.ErrBadRequest)
		return
	}
	user, err := userHandler.userSrv.FindUserByID(context.TODO(), int32(ID))
	if err != nil {
		helpers.SendRespond(w, http.StatusBadRequest, nil, err.Error())
		return
	}
	helpers.SendRespond(w, http.StatusOK, user, "")
}

func (userHandler *UserHandler) AllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := userHandler.userSrv.FindAll(context.TODO())
	if err != nil {
		helpers.SendRespond(w, http.StatusBadRequest, nil, err.Error())
		return
	}
	helpers.SendRespond(w, http.StatusOK, users, "")
}
