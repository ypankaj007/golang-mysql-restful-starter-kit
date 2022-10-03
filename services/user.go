package services

import (
	"context"
	"errors"
	"golang-mysql-restful-starter-kit/dao"
	"golang-mysql-restful-starter-kit/helpers"
	"golang-mysql-restful-starter-kit/models"
	"log"
)

type IUserService interface {
	CreateUser(context.Context, *models.User) error
	FindUserByID(context.Context, int32) (*helpers.ResponseUser, error)
	Login(context.Context, helpers.LoginPayload) (*helpers.LoginResponse, error)
	FindAll(context.Context) ([]*helpers.ResponseUser, error)
}

type UserService struct {
	userDao dao.IUserDao
}

func NewUserService(userDao dao.IUserDao) IUserService {
	return &UserService{userDao: userDao}
}

func (userSrv *UserService) CreateUser(ctx context.Context, user *models.User) error {

	tempUser, err := userSrv.userDao.FindByEmail(ctx, user.Email)
	if err != nil && err.Error() != helpers.ErrUserNotFound {
		log.Printf("Error %s when preparing SQL statement", err)
		return err
	}
	if tempUser != nil {
		log.Printf("Duplicate entry %s, %v", user.Email, tempUser)
		return errors.New(helpers.ErrDuplicateEntry)
	}
	err = userSrv.userDao.Insert(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

func (userSrv *UserService) FindUserByID(ctx context.Context, ID int32) (*helpers.ResponseUser, error) {
	return userSrv.userDao.FindByID(ctx, ID)
}

func (userSrv *UserService) Login(ctx context.Context, payload helpers.LoginPayload) (*helpers.LoginResponse, error) {
	user, err := userSrv.userDao.FindByEmail(ctx, payload.Email)

	if err != nil || user == nil {
		return nil, errors.New(helpers.ErrUserNotFound)
	}
	err = user.ComparePassword(payload.Password)
	if err != nil {
		return nil, errors.New(helpers.ErrUserNotFound)
	}
	token, err := helpers.GenerateJWT(user.ID)
	if err != nil {
		return nil, errors.New(helpers.ErrBadRequest)
	}

	return &helpers.LoginResponse{
		Token: token,
		User: &helpers.ResponseUser{
			ID:        user.ID,
			Name:      user.Email,
			Email:     user.Email,
			CreatedAT: user.CreatedAT,
		},
	}, nil

}

func (userSrv *UserService) FindAll(ctx context.Context) ([]*helpers.ResponseUser, error) {
	return userSrv.userDao.FindAll(ctx)
}
