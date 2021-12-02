package rest

import (
	"log"
	"os"
	"time"

	"github.com/federicoleon/golang-restclient/rest"

	"github.com/jmillandev/bookstore_oauth-api/src/domain/users"
	"github.com/jmillandev/bookstore_oauth-api/src/services"
	"github.com/jmillandev/bookstore_utils-go/rest_errors"
	"github.com/joho/godotenv"
)

type userRepository struct {
}

func NewUserRepository() services.UserRepository {
	return &userRepository{}
}

var usersRestClient rest.RequestBuilder
var (
	userBaseURL, userLoginEndpoint string
)

func init() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file in user repository")
	}
	userBaseURL = os.Getenv("USERS_SERVICE_BASE_URL")
	userLoginEndpoint = os.Getenv("USERS_SERVICE_LOGIN_ENDPOINT")

	usersRestClient = rest.RequestBuilder{
		BaseURL: userBaseURL,
		Timeout: 100 * time.Millisecond,
	}
}

func (r *userRepository) LoginUser(email string, password string) (*users.User, *rest_errors.RestErr) {
	request := users.UserLoginRequest{
		Email:    email,
		Password: password,
	}
	response := usersRestClient.Post(userLoginEndpoint, request)

	if response == nil || response.Response == nil {
		return nil, rest_errors.NewInternalServerError("invalid restclient response when trying to login user", err)
	}
	if response.StatusCode > 299 {
		var restErr rest_errors.RestErr
		err := response.FillUp(&restErr)
		if err != nil {
			return nil, rest_errors.NewInternalServerError("invalid error interface when trying to login user", err)
		}
		return nil, &restErr
	}
	var user users.User
	if err := response.FillUp(&user); err != nil {
		return nil, rest_errors.NewInternalServerError("error when trying to unmarshal user login response", err)
	}
	return &user, nil
}
