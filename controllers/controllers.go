package controllers

import (
	"eatfy/models"
	"encoding/json"
	"net/http"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
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

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.User_password), 8)
		user.User_password = string(hashedPassword)
		userCreation, err := db.Debug().Create(&user)

		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte("403 - Username or Useremail already used"))
		}
		userjson, err := json.Marshal(user)
		if err != nil {
		}
		w.WriteHeader(http.StatusCreated)
		w.Write(userjson)
	}
}
