package app

import (
	"eatfy/handlers"
	"eatfy/middlewares"
	"eatfy/models"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

func (a *App) Initialize() {

	err := godotenv.Load()

	DB, err := gorm.Open("postgres", os.Getenv("POSTGRES_URI"))
	if err != nil {
		log.Fatal("Could not connect database %s", err)
	}
	DB.LogMode(true)

	a.DB = models.DBMigration(DB)
	a.Router = mux.NewRouter()
	a.SetMiddlewares()
	a.SetRoutes()
}

func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}

func (a *App) SetMiddlewares() {
	a.Router.Use(middlewares.AuthMiddleware, middlewares.LoggingMiddleware)

}

func (a *App) SetRoutes() {

	a.Router.HandleFunc("/users", handlers.CreateUser(a.DB)).Methods("POST")
	a.Router.HandleFunc("/login", handlers.Login(a.DB)).Methods("GET")
	a.Router.HandleFunc("/userpreferences", handlers.CreateUserPreferences(a.DB)).Methods("POST")
	a.Router.HandleFunc("/reservation", handlers.CreateReservation(a.DB)).Methods("POST")
	a.Router.HandleFunc("/reservation", handlers.UpdateReservationStatus(a.DB)).Methods("UPDATE")
}
