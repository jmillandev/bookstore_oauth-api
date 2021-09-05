package rest

import (
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/federicoleon/golang-restclient/rest"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	rest.StartMockupServer()
	os.Exit(m.Run())
}

func TestLoginUserTimeoutFromAPI(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		HTTPMethod:   http.MethodPost,
		URL:          userBaseURL + userLoginEndpoint,
		ReqBody:      `{"email":"email@getnada.com","password":"TheBigPassword"}`,
		RespBody:     `{}`,
		RespHTTPCode: -1,
		RespHeaders: http.Header{
			"Content-Type": []string{"application/json"},
		},
	})
	respository := userRepository{}
	user, err := respository.LoginUser("email@getnada.com", "TheBigPassword")

	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "invalid restclient response when trying to login user", err.Message)
}

func TestLoginUserInvalidErrorInterface(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		HTTPMethod:   http.MethodPost,
		URL:          fmt.Sprintf("%s%s", userBaseURL, userLoginEndpoint),
		ReqBody:      `{"email":"email@getnada.com","password":"TheBigPassword"}`,
		RespBody:     `{"message": "Invalid credentials", "status": "400", "error": "bad_request"}`,
		RespHTTPCode: http.StatusBadRequest,
		RespHeaders: http.Header{
			"Content-Type": []string{"application/json"},
		},
	})
	respository := userRepository{}
	user, err := respository.LoginUser("email@getnada.com", "TheBigPassword")

	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "invalid error interface when trying to login user", err.Message)
}

func TestLoginUserInvalidCredentials(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		HTTPMethod:   http.MethodPost,
		URL:          fmt.Sprintf("%s%s", userBaseURL, userLoginEndpoint),
		ReqBody:      `{"email":"email@getnada.com","password":"TheBigPassword"}`,
		RespBody:     `{"message": "Invalid credentials", "status": 400, "error": "bad_request"}`,
		RespHTTPCode: http.StatusBadRequest,
		RespHeaders: http.Header{
			"Content-Type": []string{"application/json"},
		},
	})
	respository := userRepository{}
	user, err := respository.LoginUser("email@getnada.com", "TheBigPassword")

	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusBadRequest, err.Status)
	assert.EqualValues(t, "Invalid credentials", err.Message)
}

func TestLoginUserInvalidJsonResponse(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		HTTPMethod:   http.MethodPost,
		URL:          fmt.Sprintf("%s%s", userBaseURL, userLoginEndpoint),
		ReqBody:      `{"email":"email@getnada.com","password":"TheBigPassword"}`,
		RespBody:     `{"id": "1","first_name": "Peter","last_name": "Castillo","email": "email@getnada.com","date_created": "2021-09-05T17:05:46Z","status": "active"}`,
		RespHTTPCode: http.StatusOK,
		RespHeaders: http.Header{
			"Content-Type": []string{"application/json"},
		},
	})
	respository := userRepository{}
	user, err := respository.LoginUser("email@getnada.com", "TheBigPassword")

	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "invalid error interface when trying to login user", err.Message)
}

func TestLoginUserSuccess(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		HTTPMethod:   http.MethodPost,
		URL:          fmt.Sprintf("%s%s", userBaseURL, userLoginEndpoint),
		ReqBody:      `{"email":"email@getnada.com","password":"TheBigPassword"}`,
		RespBody:     `{"id": 1,"first_name": "Peter","last_name": "Castillo","email": "email@getnada.com","date_created": "2021-09-05T17:05:46Z","status": "active"}`,
		RespHTTPCode: 200,
		RespHeaders: http.Header{
			"Content-Type": []string{"application/json"},
		},
	})
	respository := userRepository{}
	user, err := respository.LoginUser("email@getnada.com", "TheBigPassword")

	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.EqualValues(t, user.Id, 1)
	assert.EqualValues(t, user.Email, "email@getnada.com")
	assert.EqualValues(t, user.FirstName, "Peter")
	assert.EqualValues(t, user.LastName, "Castillo")

}
