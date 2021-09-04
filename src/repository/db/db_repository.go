package db

import (
	"log"

	"github.com/jgmc3012/bookstore_oauth-api/src/clients/cassandra"
	"github.com/jgmc3012/bookstore_oauth-api/src/domain/access_token"
	"github.com/jgmc3012/bookstore_oauth-api/src/utils/errors"
)

type dbRepository struct{}

func NewRepository() access_token.Repository {
	return &dbRepository{}
}

func (r dbRepository) GetById(string) (*access_token.AccessToken, *errors.RestErr) {
	session, err := cassandra.GetSession()
	if err != nil {
		log.Printf("Error getting session: %s\n", err.Error())
		return nil, errors.NewInternalServerError("error when trying to connect db")
	}

	defer session.Close()

	// TODO: implement get access token by id from db
	return nil, errors.NewInternalServerError("database connection not implemented yet!")
}
