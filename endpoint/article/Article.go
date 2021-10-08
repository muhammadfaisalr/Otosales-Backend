package article

import (
	"Otosales/database"
	"Otosales/endpoint/account"
	"Otosales/exception"
	"Otosales/helper"
	"Otosales/models/article"
	"Otosales/rest"
	"Otosales/table"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

func Post(w http.ResponseWriter, r *http.Request) {
	db, err := database.Connect()

	if err != nil {
		exception.Throw(w, http.StatusInternalServerError, err.Error())
		return
	}

	var articleRequest article.Request

	err = helper.ParseFromRaw(r, &articleRequest)
	if err != nil {
		exception.Throw(w, http.StatusInternalServerError, err.Error())
	}

	var articleCommon = article.Response{
		Title:      articleRequest.Title,
		Content:    articleRequest.Content,
		AuthorId:   articleRequest.AuthorId,
		PostDate:   time.Now(),
		UpdateDate: nil,
	}

	authorName, err := account.AuthorName(db, articleRequest.AuthorId)
	if err != nil {
		exception.Throw(w, http.StatusUnauthorized, err.Error())
		return
	}

	articleCommon.AuthorName = authorName

	execute := db.Table(table.Article).Create(&articleCommon)

	err = execute.Error
	if err != nil {
		exception.Throw(w, http.StatusInternalServerError, err.Error())
		return
	}

	err = rest.WriteJson(w, http.StatusOK, http.StatusText(http.StatusOK), http.StatusText(http.StatusOK))
	if err != nil {
		exception.Throw(w, http.StatusInternalServerError, err.Error())
		return
	}
}

func GetAll(w http.ResponseWriter, r *http.Request) {
	db, err := database.Connect()

	if err != nil {
		exception.Throw(w, http.StatusInternalServerError, err.Error())
		return
	}

	var articleResponse []article.Response

	result, err := db.Table(table.Article).Rows()

	if err != nil {
		exception.Throw(w, http.StatusInternalServerError, err.Error())
		return
	}

	for result.Next() {
		var each = article.Response{}
		var err = result.Scan(&each.Id, &each.Title, &each.Content, &each.AuthorId, &each.AuthorName, &each.PostDate, &each.UpdateDate)

		if err != nil {
			exception.Throw(w, http.StatusInternalServerError, err.Error())
			return
		}

		articleResponse = append(articleResponse, each)
	}

	if err != nil {
		exception.Throw(w, http.StatusInternalServerError, err.Error())
		return
	}

	err = rest.WriteJson(w, http.StatusOK, http.StatusText(http.StatusOK), articleResponse)

	if err != nil {
		exception.Throw(w, http.StatusInternalServerError, err.Error())
		return
	}
}

func Get(w http.ResponseWriter, r *http.Request) {
	db, err := database.Connect()

	if err != nil {
		exception.InternalServerError(w, err.Error())
		return
	}

	id := mux.Vars(r)["id"]

	var response *article.Response = nil

	db.Raw("SELECT * FROM s_article WHERE id = ?", id).Scan(&response)

	if response == nil {
		rest.WriteJson(w, http.StatusNotFound, http.StatusText(http.StatusNotFound), nil)
		return
	}

	err = rest.WriteJson(w, http.StatusOK, http.StatusText(http.StatusOK), response)
	if err != nil {
		exception.InternalServerError(w, err.Error())
		return
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {
	db, err := database.Connect()

	if err != nil {
		exception.InternalServerError(w, err.Error())
		return
	}

	id := mux.Vars(r)["id"]

	var response *article.Response = nil

	db.Raw("SELECT * FROM s_article WHERE id = ?", id).Scan(&response)

	if response == nil {
		rest.WriteJson(w, http.StatusNotFound, http.StatusText(http.StatusNotFound), nil)
		return
	}

	db.Table(table.Article).Delete(response, id)

	err = rest.WriteJson(w, http.StatusOK, http.StatusText(http.StatusOK), nil)

	if err != nil {
		exception.InternalServerError(w, err.Error())
		return
	}
}
