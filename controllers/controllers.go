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

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.UserPassword), 8)
		user.UserPassword = string(hashedPassword)
		userCreation := db.Create(&user)

		if userCreation.Error != nil {
			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte("403 - Username or Useremail already used"))
			return
		}

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("201 - User was created succesfully"))
		return
	}
}

func Login(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user_login := models.UserLogin{}
		err := json.NewDecoder(r.Body).Decode(&user_login)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("500 - Invalid Request"))
			return
		}
		defer r.Body.Close()

		user := models.User{}
		login_query := db.Where("(user_name = ? or user_email = ?)",
			user_login.UserIdentifier, user_login.UserIdentifier).First(&user)

		if login_query.Error != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("401 - Invalid UserName or UserEmail"))
			return
		}

		if bcrypt.CompareHashAndPassword([]byte(user.UserPassword),
			[]byte(user_login.UserPassword)) != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("401 - Invalid Password"))
			return
		}

		w.WriteHeader(http.StatusAccepted)
		w.Write([]byte("202 - User was login succesfully"))
		return

	}
}
