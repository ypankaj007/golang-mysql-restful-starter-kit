package api

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"

	"golang-mysql-restful-starter-kit/config"
	"golang-mysql-restful-starter-kit/dao"
	"golang-mysql-restful-starter-kit/handlers"
	"golang-mysql-restful-starter-kit/middleware"
	"golang-mysql-restful-starter-kit/services"
)

type IApp interface {
	Run()
}

type app struct {
	router *mux.Router
	db     *sql.DB
}

func NewApp() IApp {
	a := &app{}
	connStr := fmt.Sprintf("%s:%s@tcp(%s)/%s", os.Getenv(config.DB_USERNAME), os.Getenv(config.DB_PASSWORD), os.Getenv(config.DB_HOST), os.Getenv(config.DB_NAME))
	log.Println("Connection String ", connStr)
	db, err := sql.Open("mysql", connStr)
	if err != nil {
		log.Printf("Error %s when opening DB\n", err)
	}
	log.Println("DB connected")
	a.db = db
	a.router = mux.NewRouter()
	a.initRoutes()
	return a
}

func (a *app) Run() {
	log.Println(":" + os.Getenv(config.APP_PORT))
	log.Fatal(http.ListenAndServe(":"+os.Getenv(config.APP_PORT), a.router))
}

func (a *app) initRoutes() {

	userDao := dao.NewUserDao(a.db)
	userSrv := services.NewUserService(userDao)

	userHandlsers := handlers.NewUserHandler(userSrv)
	authHandlsers := handlers.NewAuthHandler(userSrv)

	// *********************************** Auth ******************************************
	a.router.HandleFunc("/auth/register", authHandlsers.Register).Methods(http.MethodPost)
	a.router.HandleFunc("/auth/login", authHandlsers.Login).Methods(http.MethodPost)

	// ********************************* User ******************************************
	a.router.HandleFunc("/users/{id:[0-9]+}", middleware.IsAuthorized(userHandlsers.UserDetails)).Methods(http.MethodGet)
	a.router.HandleFunc("/users", middleware.IsAuthorized(userHandlsers.AllUsers)).Methods(http.MethodGet)
}
