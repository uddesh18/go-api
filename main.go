package main

import (
	"go-api/config"
	"go-api/helper"
	"go-api/modules/tags"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Info().Msg("Started Server!")
	db := config.DatabaseConnection()
	validate := validator.New()

	repository := tags.NewTagsREpositoryImpl(db)
	// Database
	tagsservice := tags.NewTagsServiceImpl(repository, validate)
	tagsController := tags.NewTagsController(tagsservice)
	router := gin.Default()

	routes := tags.NewRouter(tagsController, router)

	router.Run(":8080")

	server := &http.Server{
		Addr:    ":8888",
		Handler: routes,
	}

	err := server.ListenAndServe()
	helper.ErrorPanic(err)
}
