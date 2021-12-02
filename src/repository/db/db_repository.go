package db

import (
	"log"

	"github.com/gocql/gocql"
	"github.com/jmillandev/bookstore_oauth-api/src/clients/cassandra"
	"github.com/jmillandev/bookstore_oauth-api/src/domain/access_token"
	"github.com/jmillandev/bookstore_oauth-api/src/services"
	"github.com/jmillandev/bookstore_utils-go/rest_errors"
)

const (
	queryGetAccessToken    = "SELECT access_token, user_id, client_id, expires FROM access_tokens WHERE access_token = ?;"
	queryCreateAccessToken = "INSERT INTO access_tokens(access_token, user_id, client_id, expires) VALUES (?, ?, ?, ?);"
	queryUpdateExpiration  = "UPDATE access_tokens SET expires = ? WHERE access_token = ?;"
)

func NewAccessTokenRepository() services.AccessTokenRepository {
	return &accessTokenRepository{}
}

type accessTokenRepository struct{}

func (r accessTokenRepository) GetById(id string) (*access_token.AccessToken, *rest_errors.RestErr) {
	session := cassandra.GetSession()
	var result access_token.AccessToken
	if err := session.Query(queryGetAccessToken, id).Scan(
		&result.AccessToken,
		&result.UserId,
		&result.ClientId,
		&result.Expires,
	); err != nil {

		if err == gocql.ErrNotFound {
			return nil, rest_errors.NewNotFoundError("no access token found with given id")
		}
		log.Printf("Error getting access token: %s\n", err.Error())
		return nil, rest_errors.NewInternalServerError("error when trying to get access token", err)
	}

	return &result, nil
}

func (r accessTokenRepository) Create(at access_token.AccessToken) *rest_errors.RestErr {
	session := cassandra.GetSession()

	if err := session.Query(
		queryCreateAccessToken,
		at.AccessToken,
		at.UserId,
		at.ClientId,
		at.Expires,
	).Exec(); err != nil {
		log.Printf("Error creating access token: %s\n", err.Error())
		return rest_errors.NewInternalServerError("error when trying to create access token", err)
	}
	return nil
}

func (r accessTokenRepository) UpdateExpirationTime(at access_token.AccessToken) *rest_errors.RestErr {
	session := cassandra.GetSession()

	if err := session.Query(
		queryUpdateExpiration,
		at.Expires,
		at.AccessToken,
	).Exec(); err != nil {
		log.Printf("Error updating expires access token: %s\n", err.Error())
		return rest_errors.NewInternalServerError("error when trying to update expires access token", err)
	}
	return nil
}
