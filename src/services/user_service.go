package services

import (
	"github.com/jgmc3012/bookstore_oauth-api/src/domain/users"
	"github.com/jgmc3012/bookstore_users-api/utils/errors"
)

type UserRepository interface {
	LoginUser(string, string) (*users.User, *errors.RestErr)
}
