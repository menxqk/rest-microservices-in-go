package access_token

import (
	"testing"
	"time"
)

func TestAccessTokenConstants(t *testing.T) {
	if EXPIRATION_TIME != 24*time.Hour {
		t.Error("expiration time should be 24 hours")
	}
}

func TestGetNewAccessToken(t *testing.T) {
	at := GetNewAccessToken()

	if at == nil {
		t.Fatal("brand new access token should not be nil")
	}
	if at.IsExpired() {
		t.Error("brand new access token should not be expired")
	}
	if at.AccessToken != "" {
		t.Error("new access token should not have defined access token id")
	}
	if at.UserId != 0 {
		t.Error("new access token should not have an associated user id")
	}
}

func TestAccessTokenIsExpired(t *testing.T) {
	at := AccessToken{}

	if !at.IsExpired() {
		t.Error("empty access token should be expired by default")
	}

	at.Expires = time.Now().UTC().Add(3 * time.Hour).Unix()

	if at.IsExpired() {
		t.Error("access token should not be expired")
	}
}
