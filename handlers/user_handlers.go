package handlers

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
			w.Write([]byte("400 - Invalid Request"))
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
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("400 - Invalid Request"))
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

func CreateUserPreferences(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user_preferences := models.UserPreferences{}
		err := json.NewDecoder(r.Body).Decode(&user_preferences)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("400 - Invalid Request"))
			return
		}
		defer r.Body.Close()

		userCreation := db.Create(&user_preferences)

		if userCreation.Error != nil {
			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte("403 - Problems to store the db objets"))
			return
		}

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("201 - UserPreferences were created succesfully"))
		return

	}
}

func CreateReservation(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		reservation := models.Reservations{}
		err := json.NewDecoder(r.Body).Decode(&reservation)
		
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("400 - Invalid Request"))
			return
		}
		defer r.Body.Close()
		create_reservation := db.Where("user_name = ? AND reservation_date = ? AND meal_type = ?",
										reservation.UserName, reservation.ReservationDate, 
										reservation.MealType).FirstOrCreate(&reservation)

		if create_reservation.Error != nil {
			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte("403 - Problems to store the db objets"))
			return
		}

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("201 - Reservation was created succesfully"))
		return
	}
}


func UpdateReservationStatus(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		reservation_to_update := models.Reservations{}
		err := json.NewDecoder(r.Body).Decode(&reservation_to_update)
		
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("400 - Invalid Request"))
			return
		}

		defer r.Body.Close()
		update_reservation := db.Model(&reservation_to_update).Where(
									 "user_name = ? AND reservation_date = ? AND meal_type = ?",
									 reservation_to_update.UserName, reservation_to_update.ReservationDate, 
									 reservation_to_update.MealType).Update("pick_up_food", true)

		if update_reservation.Error != nil {
			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte("403 - Problems to store the db objets"))
			return
		}

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("201 - Reservation was updated succesfully"))
		return
	}
}
