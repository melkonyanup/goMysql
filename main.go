package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	_ "github.com/go-sql-driver/mysql"

	article_api "./apis"
)

// func homePage(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Welcome to the HomePage!")
// 	fmt.Println("Endpoint Hit: homePage")
// }

// func returnAllArticles(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("Endpoint Hit: returnAllArticles")
// 	json.NewEncoder(w).Encode(Articles)
// }

// func returnSingleArticle(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	key := vars["id"]

// 	for _, article := range Articles {
// 		if article.Id == key {
// 			json.NewEncoder(w).Encode(article)
// 		}
// 	}
// }

// func createNewArticle(w http.ResponseWriter, r *http.Request) {

// 	reqBody, _ := ioutil.ReadAll(r.Body)

// 	var article Article
// 	json.Unmarshal(reqBody, &article)

// 	Articles = append(Articles, article)

// 	json.NewEncoder(w).Encode(article)
// 	// fmt.Fprintf(w, "%+v", string(reqBody))
// }

// func deleteArticle(w http.ResponseWriter, r *http.Request) {

// 	vars := mux.Vars(r)
// 	id := vars["id"]

// 	for index, article := range Articles {
// 		if article.Id == id {
// 			Articles = append(Articles[:index], Articles[index+1:]...)
// 		}
// 	}
// }

// func handleRequests() {
// 	// creates a new instance of a mux router
// 	myRouter := mux.NewRouter().StrictSlash(true)

// 	myRouter.HandleFunc("/", homePage)
// 	myRouter.HandleFunc("/articles", returnAllArticles)
// 	myRouter.HandleFunc("/article/{id}", returnSingleArticle)

// 	myRouter.HandleFunc("/article", createNewArticle).Methods("POST")
// 	myRouter.HandleFunc("/article/{id}", deleteArticle).Methods("DELETE")
// 	log.Fatal(http.ListenAndServe(":10000", myRouter))
// }

// let's declare a global Articles array
// that we can then populate in our main function
// to simulate a database
func main() {

	router := mux.NewRouter()

	router.HandleFunc("/api/articles/viewall", article_api.ViewAll).Methods("GET")
	router.HandleFunc("/api/articles/view/{id}", article_api.ViewArticle).Methods("GET")
	router.HandleFunc("/api/articles/create", article_api.CreateArticle).Methods("POST")
	router.HandleFunc("/api/articles/update", article_api.UpdateArticle).Methods("PUT")
	router.HandleFunc("/api/articles/delete/{id}", article_api.DeleteArticle).Methods("DELETE")

	err := http.ListenAndServe(":5000", router)

	if err != nil {
		fmt.Println(err)
	} else {

	}
}
