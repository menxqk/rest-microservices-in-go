package access_token

import (
	"strings"
	"time"

	"github.com/menxqk/rest-microservices-in-go/common/errors"
)

const (
	EXPIRATION_TIME = 24 * time.Hour
)

type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserId      int64  `json:"user_id"`
	ClientId    int64  `json:"client_id,omitempty"`
	Expires     int64  `json:"expires"`
}

func GetNewAccessToken() *AccessToken {
	return &AccessToken{
		Expires: time.Now().UTC().Add(EXPIRATION_TIME).Unix(),
	}
}

func (at *AccessToken) Validate() errors.RestError {
	at.AccessToken = strings.TrimSpace(at.AccessToken)
	if at.AccessToken == "" {
		return errors.NewBadRequestError("invalid access token id")
	}
	if at.UserId <= 0 {
		return errors.NewBadRequestError("invalid user id")
	}
	if at.ClientId <= 0 {
		return errors.NewBadRequestError("invalid client id")
	}
	if at.Expires <= 0 {
		return errors.NewBadRequestError("invalid expiration time")
	}

	return nil
}

func (at *AccessToken) IsExpired() bool {
	now := time.Now().UTC()
	expirationTime := time.Unix(at.Expires, 0)
	return now.After(expirationTime)
}
