package access_token

import "time"

const expirationTimeInHours = 24

type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserId      int64  `json:"user_id"`
	ClientId    int64  `json:"client_id"`
	Expires     int64  `json:"expires"`
}

// Web frontend - Client-Id: 123
// Android APP - Client-Id: 234

func GetNewAccessToken() AccessToken {
	return AccessToken{
		Expires: time.Now().UTC().Add(expirationTimeInHours * time.Hour).Unix(),
	}
}

func (at AccessToken) IsExpired() bool {
	expirationTime := time.Unix(at.Expires, 0)

	return expirationTime.Before(time.Now().UTC())
}
