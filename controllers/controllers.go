package controllers

import (
	"eatfy/controllers"
	"eatfy/models"
	"encoding/json"
	"github.com/jinzhu/gorm"
	"net/http"
)

func CreateUser(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := models.User{}
		err := json.NewDecoder(r.Body).Decode(&user)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("500 - Invalid Request"))
			return
		}
		defer r.Body.Close()

		if db.NewRecord(user) == true {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("400 - UserName or UserEmail already created"))
			return
		}

		db.Create(user)

		w.WriteHeader(http.StatusAccepted)
		w.Write([]byte("202 - User created succesfully"))
		return

	}
}
