package app

import (
	"github.com/gin-gonic/gin"
	"github.com/jgmc3012/bookstore_oauth-api/src/clients/cassandra"
	"github.com/jgmc3012/bookstore_oauth-api/src/domain/access_token"
	"github.com/jgmc3012/bookstore_oauth-api/src/http"
	"github.com/jgmc3012/bookstore_oauth-api/src/repository/db"
)

var router = gin.Default()

func StartApplication() {
	session, dbErr := cassandra.GetSession()
	if dbErr != nil {
		panic(dbErr)
	}
	session.Close()

	repository := db.NewRepository()
	atService := access_token.NewService(repository)
	atHandler := http.NewHandler(atService)

	router.GET("/oauth/access-token/:access_token_id", atHandler.GetById)
	router.POST("/oauth/access-token/", atHandler.Create)

	router.Run(":8000")
}
