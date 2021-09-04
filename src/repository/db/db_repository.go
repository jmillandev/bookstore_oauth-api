package db

import (
	"log"

	"github.com/gocql/gocql"
	"github.com/jgmc3012/bookstore_oauth-api/src/clients/cassandra"
	"github.com/jgmc3012/bookstore_oauth-api/src/domain/access_token"
	"github.com/jgmc3012/bookstore_oauth-api/src/utils/errors"
)

const (
	queryGetAccessToken    = "SELECT access_token, user_id, client_id, expires FROM access_tokens WHERE access_token = ?;"
	queryCreateAccessToken = "INSERT INTO access_tokens(access_token, user_id, client_id, expires) VALUES (?, ?, ?, ?);"
	queryUpdateExpiration  = "UPDATE access_tokens SET expires = ? WHERE access_token = ?;"
)

func NewRepository() access_token.Repository {
	return &dbRepository{}
}

type dbRepository struct{}

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

func (r dbRepository) Create(at access_token.AccessToken) *errors.RestErr {
	session, err := cassandra.GetSession()
	if err != nil {
		log.Printf("Error getting session: %s\n", err.Error())
		return errors.NewInternalServerError("error when trying to connect db")
	}
	defer session.Close()

	if err := session.Query(
		queryCreateAccessToken,
		at.AccessToken,
		at.UserId,
		at.ClientId,
		at.Expires,
	).Exec(); err != nil {
		log.Printf("Error creating access token: %s\n", err.Error())
		return errors.NewInternalServerError("error when trying to create access token")
	}
	return nil
}

func (r dbRepository) UpdateExpirationTime(at access_token.AccessToken) *errors.RestErr {
	session, err := cassandra.GetSession()
	if err != nil {
		log.Printf("Error getting session: %s\n", err.Error())
		return errors.NewInternalServerError("error when trying to connect db")
	}
	defer session.Close()

	if err := session.Query(
		queryUpdateExpiration,
		at.Expires,
		at.AccessToken,
	).Exec(); err != nil {
		log.Printf("Error updating expires access token: %s\n", err.Error())
		return errors.NewInternalServerError("error when trying to update expires access token")
	}
	return nil
}
