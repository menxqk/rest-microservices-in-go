package myoauth

import "net/http"

const (
	headerXPublic   = "X-Public"
	headerXClientId = "X-Client-Id"
	headerXUserId   = "X-User-Id"
)

type oauthInterface interface {
}

type oatuhClient struct {
}

func IsPublic(request *http.Request) bool {
	if request == nil {
		return true
	}
	return request.Header.Get(headerXPublic) == "true"
}

func AuthenticateRequest(request *http.Request) {
	if request == nil {
		return
	}
}
