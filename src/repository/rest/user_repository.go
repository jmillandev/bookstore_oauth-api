package rest

import (
	"log"
	"os"
	"time"

	"github.com/federicoleon/golang-restclient/rest"

	"github.com/jgmc3012/bookstore_oauth-api/src/domain/users"
	"github.com/jgmc3012/bookstore_oauth-api/src/utils/errors"
	"github.com/joho/godotenv"
)

type RestUsersRepository interface {
	LoginUser(string, string) (*users.User, *errors.RestErr)
}

type userRepository struct {
}

func NewRepository() RestUsersRepository {
	return &userRepository{}
}

var usersRestClient rest.RequestBuilder
var (
	userBaseURL, userLoginEndpoint string
)

func init() {
	err := godotenv.Load("../../.env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	userBaseURL = os.Getenv("USERS_SERVICE_BASE_URL")
	userLoginEndpoint = os.Getenv("USERS_SERVICE_LOGIN_ENDPOINT")

	usersRestClient = rest.RequestBuilder{
		BaseURL: userBaseURL,
		Timeout: 100 * time.Millisecond,
	}
}

func (r *userRepository) LoginUser(email string, password string) (*users.User, *errors.RestErr) {
	request := users.UserLoginRequest{
		Email:    email,
		Password: password,
	}
	response := usersRestClient.Post(userLoginEndpoint, request)

	if response == nil || response.Response == nil {
		return nil, errors.NewInternalServerError("invalid restclient response when trying to login user")
	}
	if response.StatusCode > 299 {
		var restErr errors.RestErr
		err := response.FillUp(&restErr)
		if err != nil {
			return nil, errors.NewInternalServerError("invalid error interface when trying to login user")
		}
		return nil, &restErr
	}
	var user users.User
	if err := response.FillUp(&user); err != nil {
		return nil, errors.NewInternalServerError("error when trying to unmarshal user login response")
	}
	return &user, nil
}
