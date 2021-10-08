package account

import (
	"Otosales/database"
	"Otosales/exception"
	"Otosales/helper"
	"Otosales/message"
	"Otosales/models/profile"
	"Otosales/models/signin"
	"Otosales/models/signup"
	"Otosales/rest"
	"Otosales/table"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	db, err := database.Connect()
	if err != nil {
		exception.Throw(
			w,
			http.StatusInternalServerError,
			err.Error(),
		)
	}

	var signInRequest signin.Request
	err = helper.ParseFromRaw(r, &signInRequest)

	if err != nil {
		exception.Throw(
			w,
			http.StatusInternalServerError,
			err.Error())
	}

	var response *profile.Response = nil

	db.Raw("SELECT * FROM s_user WHERE phone_number = ?", &signInRequest.PhoneNumber).
		Scan(&response)

	if response == nil {
		err = rest.WriteJson(w, http.StatusUnauthorized, message.PhoneNumberNotFound, nil)
		return
	}

	err = rest.WriteJson(w, http.StatusOK, http.StatusText(http.StatusOK), response)

	if err != nil {
		exception.Throw(
			w,
			http.StatusInternalServerError,
			err.Error())
	}
}

func Register(w http.ResponseWriter, r *http.Request) {
	db, err := database.Connect()

	if err != nil {
		exception.Throw(
			w,
			http.StatusInternalServerError,
			err.Error(),
		)
	}

	var signupRequest signup.Request
	err = helper.ParseFromRaw(r, &signupRequest)

	if err != nil {
		exception.Throw(
			w,
			http.StatusInternalServerError,
			err.Error())
	}

	create := db.Table(table.User).Create(&signupRequest)

	err = create.Error
	if err != nil {
		rest.WriteJson(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	err = rest.WriteJson(w, http.StatusOK, http.StatusText(http.StatusOK), nil)

	if err != nil {
		rest.WriteJson(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}
}
