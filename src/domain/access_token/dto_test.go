package access_token

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAccessTokenConstants(t *testing.T) {
	assert.EqualValues(t, 24, expirationTimeInHours, "expiration time should be 24 hours")
}

func TestGetNewAccessToken(t *testing.T) {
	at := GetNewAccessToken(1)

	assert.False(t, at.IsExpired(), "brand new access token should not be expired")
	assert.Greater(t, len(at.AccessToken), 0, "brand new access token should have defines access token id")
	assert.EqualValues(t, 1, at.UserId, "brand new access token should have associated an user id")
}

func TestAccessTokenIsExpired(t *testing.T) {
	at := AccessToken{}

	assert.True(t, at.IsExpired(), "empty access token should be expired by default")

	at.Expires = time.Now().UTC().Add(time.Hour * 3).Unix()

	assert.False(t, at.IsExpired(), "empty access token expiring three hours from now should NOT be expired")

}
