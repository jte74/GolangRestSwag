package main

import (
	"fmt"
	"log"
	"net/http"

	"training/golangrestswag/controller"
	_ "training/golangrestswag/docs"
	"training/golangrestswag/model"

	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/gorilla/mux"
)

var Articles []model.Article

// @title Article API
// @version 1.0
// @description This is a sample service
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email soberkoder@gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:1323
// @BasePath /
func main() {

	fmt.Println("Rest API v2.0 - Mux Routers")

	// On vient remplir le tableau d'articles avec 2 articles
	Articles = []model.Article{
		{Id: "1", Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		{Id: "2", Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
	}

	myRouter := mux.NewRouter()

	myRouter.HandleFunc("/", controller.HomePage)

	myRouter.HandleFunc("/all", controller.ReturnAllArticles)
	myRouter.HandleFunc("/article/{Id}", controller.ReturnSingleArticle)
	// Pour toutes méthodes POST, spécifier le comme ci dessous
	myRouter.HandleFunc("/create", controller.CreateNewArticle).Methods("POST")
	// Pour toutes méthodes DELETE, spécifier le comme ci dessous
	myRouter.HandleFunc("/delete/{Id}", controller.DeleteArticle).Methods("DELETE")
	// Pour toutes méthodes PUT, spécifier le comme ci dessous
	myRouter.HandleFunc("/update", controller.UpdateArticle).Methods("PUT")

	myRouter.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)

	log.Fatal(http.ListenAndServe(":1323", myRouter))
}


