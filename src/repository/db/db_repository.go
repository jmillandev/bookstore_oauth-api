package db

import (
	"log"

	"github.com/gocql/gocql"
	"github.com/jgmc3012/bookstore_oauth-api/src/clients/cassandra"
	"github.com/jgmc3012/bookstore_oauth-api/src/domain/access_token"
	"github.com/jgmc3012/bookstore_oauth-api/src/utils/errors"
)

type dbRepository struct{}

const (
	queryGetAccessToken = "SELECT access_token, user_id, client_id, expires FROM access_tokens WHERE access_token = ?"
)

func NewRepository() access_token.Repository {
	return &dbRepository{}
}

func (r dbRepository) GetById(id string) (*access_token.AccessToken, *errors.RestErr) {
	session, err := cassandra.GetSession()
	if err != nil {
		log.Printf("Error getting session: %s\n", err.Error())
		return nil, errors.NewInternalServerError("error when trying to connect db")
	}
	defer session.Close()

	var result access_token.AccessToken
	if err := session.Query(queryGetAccessToken, id).Scan(
		&result.AccessToken,
		&result.UserId,
		&result.ClientId,
		&result.Expires,
	); err != nil {

		if err == gocql.ErrNotFound {
			return nil, errors.NewNotFoundError("no access token found with given id")
		}
		log.Printf("Error getting access token: %s\n", err.Error())
		return nil, errors.NewInternalServerError("error when trying to get access token")
	}

	return &result, nil
}
