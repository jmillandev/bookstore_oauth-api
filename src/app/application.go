package app

import (
	"github.com/gin-gonic/gin"
	"github.com/jmillandev/bookstore_oauth-api/src/clients/cassandra"
	"github.com/jmillandev/bookstore_oauth-api/src/http"
	"github.com/jmillandev/bookstore_oauth-api/src/repository/db"
	"github.com/jmillandev/bookstore_oauth-api/src/repository/rest"
	"github.com/jmillandev/bookstore_oauth-api/src/services"
)

var router = gin.Default()

func StartApplication() {
	cassandra.GetSession()

	atService := services.NewAccessTokenService(db.NewAccessTokenRepository(), rest.NewUserRepository())
	atHandler := http.NewHandler(atService)

	router.GET("/oauth/access-token/:access_token_id", atHandler.GetById)
	router.POST("/oauth/access-token/", atHandler.Create)

	router.Run(":8000")
}
