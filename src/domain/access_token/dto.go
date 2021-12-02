package access_token

import (
	"fmt"
	"strings"
	"time"

	"github.com/jmillandev/bookstore_utils-go/cripto_utils"
	"github.com/jmillandev/bookstore_utils-go/rest_errors"
)

const (
	expirationTimeInHours = 24
	grantTypePassword     = "password"
	grantTypeClientCreds  = "client_credentials"
)

type AccessTokenRequest struct {
	GrantType string `json:"grant_type"`
	Scope     string `json:"scope"`

	// Used for password grant type
	Email    string `json:"email"`
	Password string `json:"password"`

	// Used for client_credentials grant type
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

func (atr *AccessTokenRequest) Validate() *rest_errors.RestErr {
	switch atr.GrantType {
	case grantTypePassword:
		break
	case grantTypeClientCreds:
		break
	default:
		return rest_errors.NewBadRequestError("invalid grant_type parameter")
	}
	return nil
}

type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserId      int64  `json:"user_id"`
	ClientId    int64  `json:"client_id"`
	Expires     int64  `json:"expires"`
}

func (at *AccessToken) Validate() *rest_errors.RestErr {
	at.AccessToken = strings.TrimSpace(at.AccessToken)
	if len(at.AccessToken) == 0 {
		return rest_errors.NewBadRequestError("invalid access token id")
	}
	if at.UserId <= 0 {
		return rest_errors.NewBadRequestError("invalid user id")
	}
	if at.ClientId <= 0 {
		return rest_errors.NewBadRequestError("invalid client id")
	}
	if at.Expires <= 0 {
		return rest_errors.NewBadRequestError("invalid expiration time")
	}
	return nil
}

func GetNewAccessToken(userId int64) AccessToken {
	return AccessToken{
		UserId:  userId,
		Expires: time.Now().UTC().Add(expirationTimeInHours * time.Hour).Unix(),
	}
}

func (at AccessToken) IsExpired() bool {
	expirationTime := time.Unix(at.Expires, 0)

	return expirationTime.Before(time.Now().UTC())
}

func (at *AccessToken) Generate() {
	at.AccessToken = cripto_utils.GetMd5(fmt.Sprintf("at-%d-%d-ran", at.UserId, at.Expires))
}
