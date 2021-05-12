package article_api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"../entities"
	"../models"
	"github.com/gorilla/mux"

	"../config"
)

func ViewAll(response http.ResponseWriter, request *http.Request) {
	db, err := config.GetDB()

	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
	} else {
		articleModel := models.ArticleModel{

			Db: db,
		}
		articles, err2 := articleModel.ViewAll()
		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, err2.Error())
		} else {
			respondWithJson(response, http.StatusOK, articles)
		}
	}
}
func ViewArticle(response http.ResponseWriter, request *http.Request) {

	vars := mux.Vars(request)
	sid := vars["id"]
	id, _ := strconv.ParseInt(sid, 10, 64)

	db, err := config.GetDB()

	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
	} else {
		articleModel := models.ArticleModel{

			Db: db,
		}
		articles, err2 := articleModel.ViewArticle(id)
		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, err2.Error())
		} else {
			respondWithJson(response, http.StatusOK, articles)
		}
	}
}

func CreateArticle(response http.ResponseWriter, request *http.Request) {
	var article entities.Article
	err := json.NewDecoder(request.Body).Decode(&article)
	db, err := config.GetDB()

	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
	} else {
		articleModel := models.ArticleModel{

			Db: db,
		}
		err2 := articleModel.CreateArticle(&article)
		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, err2.Error())
		} else {
			respondWithJson(response, http.StatusOK, article)
		}
	}
}

func UpdateArticle(response http.ResponseWriter, request *http.Request) {
	var article entities.Article
	err := json.NewDecoder(request.Body).Decode(&article)
	db, err := config.GetDB()

	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
	} else {
		articleModel := models.ArticleModel{

			Db: db,
		}
		_, err2 := articleModel.UpdateArticle(&article)
		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, err2.Error())
		} else {
			respondWithJson(response, http.StatusOK, article)
		}
	}
}
func DeleteArticle(response http.ResponseWriter, request *http.Request) {

	vars := mux.Vars(request)
	sid := vars["id"]
	id, _ := strconv.ParseInt(sid, 10, 64)
	db, err := config.GetDB()

	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
	} else {
		articleModel := models.ArticleModel{

			Db: db,
		}
		_, err2 := articleModel.DeleteArticle(id)
		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, err2.Error())
		} else {
			respondWithJson(response, http.StatusOK, nil)
		}
	}
}
func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
