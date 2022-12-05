package access_token

import (
	"time"
)

const (
	EXPIRATION_TIME = 24 * time.Hour
)

type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserId      int64  `json:"user_id"`
	ClientId    int64  `json:"client_id"`
	Expires     int64  `json:"expires"`
}

func GetNewAccessToken() *AccessToken {
	return &AccessToken{
		Expires: time.Now().UTC().Add(EXPIRATION_TIME).Unix(),
	}
}

func (at *AccessToken) IsExpired() bool {
	now := time.Now().UTC()
	expirationTime := time.Unix(at.Expires, 0)
	return now.After(expirationTime)
}
