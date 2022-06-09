package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"training/golangrestswag/model"

	"github.com/gorilla/mux"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

// Article godoc
// @Summary Update Article
// @Description Update Article
// @Tags id
// @Accept  json
// @Produce  json
// @Param article body model.Article true "Update Article"
// @Success 200 {object} model.Article
// @Router /update [put]
func UpdateArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	key := vars["Id"]
	var updatearticle model.Article

	reqBody, _ := ioutil.ReadAll(r.Body)
	a := json.Unmarshal(reqBody, &updatearticle)

	if a == nil {

	}

	for index, article := range model.Articles {
		if article.Id == key {
			article.Title = updatearticle.Title
			article.Desc = updatearticle.Desc
			article.Content = updatearticle.Content
			model.Articles[index] = article
			b := json.NewEncoder(w).Encode(article)
			if b != nil {

			}
		}
	}
}

// Delete Article godoc
// @Summary Delete Article
// @Description Delete Article
// @Tags id
// @Accept  json
// @Produce  json
// @Param Id path string true "ID of the article to be deleted"
// @Success 200
// @Router /delete/{Id} [delete]
func DeleteArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	key := params["Id"]

	for index, article := range model.Articles {
		if article.Id == key {
			model.Articles = append(model.Articles[:index], model.Articles[index+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
}

// CreateArticle godoc
// @Summary Create a new Article
// @Description Create a new article with the input paylod
// @Tags Article
// @Accept  json
// @Produce  json
// @Param order body model.Article true "Create article"
// @Success 200 {object} model.Article
// @Router /create [post]
func CreateNewArticle(w http.ResponseWriter, r *http.Request) {
	// get the body of our POST request
	reqBody, _ := ioutil.ReadAll(r.Body)
	// return the string response containing the request body
	fmt.Fprintf(w, "%+v", string(reqBody))

	var article model.Article
	d := json.Unmarshal(reqBody, &article)

	if d == nil {

	}

	// update our global Articles array to include
	// our new Article
	model.Articles = append(model.Articles, article)

	c := json.NewEncoder(w).Encode(article)

	if c == nil {

	}

	fmt.Println("test post")
}

// Return Single Article godoc
// @Summary Return Single Article
// @Description Return Single Article
// @Tags id
// @Accept  json
// @Produce  json
// @Param Id path string true "Return Single Article"
// @Success 200
// @Router /article/{Id} [get]
func ReturnSingleArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	key := params["Id"]

	for _, article := range model.Articles {

		fmt.Println("article id :", article.Id)
		fmt.Println("key : ", key)
		if article.Id == key {
			fmt.Println("Endpoint Hit: ReturnSingleArticle")
			e := json.NewEncoder(w).Encode(article)

			if e == nil {

			}

			return
		}
	}
}

// Return All Article godoc
// @Summary Return All Article
// @Description Return All Article
// @Tags id
// @Accept  json
// @Produce  json
// @Success 200
// @Router /all [get]
func ReturnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllArticles")
	f := json.NewEncoder(w).Encode(model.Articles)

	if f == nil {

	}
}
