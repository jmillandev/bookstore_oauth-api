package services

import (
	"github.com/jmillandev/bookstore_oauth-api/src/domain/users"
	"github.com/jmillandev/bookstore_utils-go/rest_errors"
)

type UserRepository interface {
	LoginUser(string, string) (*users.User, *rest_errors.RestErr)
}
