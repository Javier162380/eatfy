package app

import (
	"fmt"
	"log"
	"net/http"
	"eatfy/models"
	"github.com/jinzhu/gorm"

	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

func (a *App) Initialize(host, port, user, password, dbname string) {
	connectionString :=
		fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s",
			host, port, user, password, dbname)

	DB, err := gorm.Open("postgres", connectionString)
	if err != nil {
		log.Fatal("Could not connect database")
	}

	a.DB = DB
	a.Router = mux.NewRouter()
}

func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}
