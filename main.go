package main

import (
	"log"
	"net/http"
	"os"

	"github.com/jcox250/annadale/adapters/repository"
	"github.com/jcox250/annadale/adapters/service"
	"github.com/jcox250/annadale/infrastructure/middleware"
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

	adminService := service.NewAdminService(postInteractor)

	loginInteractor := usecases.NewLoginInteractor(nil)
	loginService := service.NewLoginService(loginInteractor)

	adapters := map[int]http.Handler{
		infrastructure.PostService:  postService,
		infrastructure.HomeService:  homeService,
		infrastructure.AdminService: adminService,
		infrastructure.LoginService: loginService,
	}

	logger := middleware.Logger{}

	server := infrastructure.NewHTTPServer(PORT, adapters, logger)
	if err := server.Serve(); err != nil {
		log.Fatal(err)
	}
}
