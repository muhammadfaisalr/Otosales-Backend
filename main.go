package main

import (
	"Otosales/database"
	"Otosales/endpoint/account"
	"Otosales/endpoint/article"
	"github.com/gorilla/mux"
	"net/http"
)

//todo Input Product Barang.

func main() {
	var r = router()

	//Account
	r.HandleFunc("/process-login", account.Login).
		Methods(database.POST)

	r.HandleFunc("/process-register", account.Register).
		Methods(database.POST)

	//Article
	r.HandleFunc("/publish-article", article.Post).
		Methods(database.POST)

	r.HandleFunc("/articles", article.GetAll).
		Methods(database.GET)

	r.HandleFunc("/article/{id}", article.Get).
		Methods(database.GET)

	r.HandleFunc("/article/{id}", article.Delete).
		Methods(database.DELETE)

	http.Handle("/", r)
	err := http.ListenAndServe(":1010", r)

	if err != nil {
		return
	}
}

func router() *mux.Router {
	var router = mux.NewRouter().StrictSlash(true)

	return router
}
