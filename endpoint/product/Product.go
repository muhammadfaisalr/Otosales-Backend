package product

import (
	"Otosales/database"
	"Otosales/exception"
	"Otosales/message"
	"Otosales/models/product"
	"Otosales/rest"
	"net/http"
)

func Get(w http.ResponseWriter, r *http.Request) {
	db, err := database.Connect()

	if err != nil {
		exception.Throw(w, http.StatusInternalServerError, err.Error())
		return
	}

	var response *product.Response = nil
	var id = "1"

	db.Raw("SELECT * FROM m_product WHERE id = ?", id).Scan(&response)

	if response == nil {
		rest.WriteJson(w, http.StatusNotFound, message.ProductNotFound, nil)
		return
	}

	err = rest.WriteJson(w, http.StatusOK, http.StatusText(http.StatusOK), response)

	if err != nil {
		exception.Throw(w, http.StatusInternalServerError, err.Error())
		return
	}
}
