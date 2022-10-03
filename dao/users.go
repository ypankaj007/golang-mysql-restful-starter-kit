package dao

import (
	"context"
	"database/sql"
	"errors"
	"log"

	"golang-mysql-restful-starter-kit/helpers"
	"golang-mysql-restful-starter-kit/models"
)

type IUserDao interface {
	Insert(context.Context, *models.User) error
	FindByID(context.Context, int32) (*helpers.ResponseUser, error)
	FindByEmail(context.Context, string) (*models.User, error)
	FindAll(context.Context) ([]*helpers.ResponseUser, error)
}

type userDao struct {
	db *sql.DB
}

func NewUserDao(db *sql.DB) IUserDao {
	return &userDao{db: db}
}

func (ud *userDao) Insert(ctx context.Context, user *models.User) error {
	query := "INSERT INTO users(name, email, password, createdAt) VALUES (?, ?, ?, ?)"
	stmt, err := ud.db.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when preparing SQL statement", err)
		return err
	}
	defer stmt.Close()
	res, err := stmt.ExecContext(ctx, user.Name, user.Email, user.Password, user.CreatedAT.Format(helpers.Layout))
	if err != nil {
		log.Printf("Error %s when inserting row into products table", err)
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error %s when finding rows affected", err)
		return err
	}
	log.Printf("%d User created ", rows)
	return nil
}

func (ud *userDao) FindByID(ctx context.Context, ID int32) (*helpers.ResponseUser, error) {
	var user helpers.ResponseUser
	var date string
	result, err := ud.db.Query("SELECT id, name, email, createdAt FROM users WHERE id = ?", ID)
	if err != nil {
		log.Printf("Error %s when finding by ID", err)
		return nil, err
	}
	defer result.Close()
	if result.Next() {

		err = result.Scan(&user.ID, &user.Name, &user.Email, &date)

		if err != nil {
			log.Printf("Error %s when finding by ID", err)
			return nil, err
		}
		if date != "" && user.ID > 0 {
			user.CreatedAT = helpers.StrToTime(date)
		}

	}
	if user.ID < 1 {
		return nil, errors.New(helpers.ErrUserNotFound)
	}
	return &user, nil
}

func (ud *userDao) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	var date string
	result, err := ud.db.Query("SELECT id, name, email, password, createdAt FROM users WHERE email = ?", email)
	if err != nil {
		log.Printf("Error %s when finding by email", err)
		return nil, err
	}
	defer result.Close()
	if result.Next() {

		err = result.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &date)

		if err != nil {
			log.Printf("Error %s when finding by email", err)
			return nil, err
		}
		if date != "" && user.ID > 0 {
			user.CreatedAT = helpers.StrToTime(date)
		}

	}
	if user.ID < 1 {
		return nil, errors.New(helpers.ErrUserNotFound)
	}
	return &user, nil
}

func (ud *userDao) FindAll(ctx context.Context) ([]*helpers.ResponseUser, error) {
	var users []*helpers.ResponseUser
	var date string
	result, err := ud.db.Query("SELECT id, name, email, createdAt FROM users")
	if err != nil {
		log.Printf("Error %s when finding by ID", err)
		return nil, err
	}
	defer result.Close()
	if result.Next() {
		var user helpers.ResponseUser
		err = result.Scan(&user.ID, &user.Name, &user.Email, &date)

		if err != nil {
			log.Printf("Error %s when finding by ID", err)
			return nil, err
		}
		if date != "" && user.ID > 0 {
			user.CreatedAT = helpers.StrToTime(date)
		}
		users = append(users, &user)

	}
	return users, nil
}
