package app

import (
	"github.com/gin-gonic/gin"
	"github.com/jgmc3012/bookstore_oauth-api/src/domain/access_token"
	"github.com/jgmc3012/bookstore_oauth-api/src/http"
	"github.com/jgmc3012/bookstore_oauth-api/src/repository/db"
)

var router = gin.Default()

func StartApplication() {
	repository := db.NewRepository()
	atService := access_token.NewService(repository)
	atHandler := http.NewHandler(atService)

	router.GET("/oauth/access-token/:access_token_id", atHandler.GetById)

	router.Run(":8000")
}
