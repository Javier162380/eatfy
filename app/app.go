package app

import (
	"eatfy/controllers"
	"eatfy/models"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

func (a *App) Initialize() {

	DB, err := gorm.Open("postgres", "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		log.Fatal("Could not connect database %s", err)
	}
	DB.LogMode(true)

	a.DB = models.DBMigration(DB)
	a.Router = mux.NewRouter()
	a.SetRoutes()
}

func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}

func (a *App) SetRoutes() {

	a.Router.HandleFunc("/users", controllers.CreateUser(a.DB)).Methods("POST")
	a.Router.HandleFunc("/login", controllers.Login(a.DB)).Methods("GET")
	a.Router.HandleFunc("/userpreferences"), controllers.CreateUserPreferences(a.DB)).Methods("POST")
}
