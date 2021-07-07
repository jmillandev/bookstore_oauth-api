package db

import (
	"github.com/jgmc3012/bookstore_oauth-api/src/domain/access_token"
	"github.com/jgmc3012/bookstore_oauth-api/src/utils/errors"
)

type dbRepository struct{}

func NewRepository() access_token.Repository {
	return &dbRepository{}
}

func (r dbRepository) GetById(string) (*access_token.AccessToken, *errors.RestErr) {
	return nil, errors.NewInternalServerError("database connection not implemented yet!")
}
