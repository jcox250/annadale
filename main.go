package main

import (
	"log"
	"net/http"
	"os"

	"github.com/jcox250/annadale/adapters/repository"
	"github.com/jcox250/annadale/adapters/service"
	"github.com/jcox250/annadale/usecases"

	"github.com/jcox250/annadale/infrastructure"
)

var DB_TYPE = os.Getenv("DB_TYPE")
var DB_URL = os.Getenv("DATABASE_URL")
var PORT = os.Getenv("PORT")

func main() {
	// Enable line numbers for logging
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	sqlHandler := infrastructure.NewSQLHandler(DB_TYPE, DB_URL)
	postRepo := repository.NewPostRepo(sqlHandler)
	postInteractor := usecases.NewPostInteractor(postRepo)
	postService := service.NewPostService(postInteractor)

	homeInteractor := usecases.NewHomeInteractor(nil)
	homeService := service.NewHomeService(homeInteractor)

	adapters := map[string]http.Handler{
		infrastructure.PostService: postService,
		infrastructure.HomeService: homeService,
	}

	server := infrastructure.NewHTTPServer(PORT, adapters)
	if err := server.Serve(); err != nil {
		log.Fatal(err)
	}
}
