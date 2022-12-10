package oauth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/menxqk/rest-microservices-in-go/common/errors"
	"github.com/mercadolibre/golang-restclient/rest"
)

const (
	headerXPublic   = "X-Public"
	headerXClientId = "X-Client-Id"
	headerXCallerId = "X-Caller-Id"

	paramAccessToken = "access_token"
)

var (
	oauthRestClient = rest.RequestBuilder{
		BaseURL: "localhost:8080",
		Timeout: 200 * time.Millisecond,
	}
)

type accessToken struct {
	Id       string `json:"id"`
	UserId   int64  `json:"user_id"`
	ClientId int64  `json:"client_id"`
}

func IsPublic(request *http.Request) bool {
	if request == nil {
		return true
	}
	return request.Header.Get(headerXPublic) == "true"
}

func GetCallerId(request *http.Request) int64 {
	if request == nil {
		return 0
	}

	headerCallerId := request.Header.Get(headerXCallerId)
	callerId, err := strconv.ParseInt(headerCallerId, 10, 64)
	if err != nil {
		return 0
	}

	return callerId
}

func GetClientId(request *http.Request) int64 {
	if request == nil {
		return 0
	}

	headerClientId := request.Header.Get(headerXClientId)
	clientId, err := strconv.ParseInt(headerClientId, 10, 64)
	if err != nil {
		return 0
	}

	return clientId
}

func AuthenticateRequest(request *http.Request) errors.RestError {
	if request == nil {
		return nil
	}

	cleanRequest(request)

	accessTokenId := strings.TrimSpace(request.URL.Query().Get(paramAccessToken))
	if accessTokenId == "" {
		return nil
	}

	at, err := getAccessToken(accessTokenId)
	if err != nil {
		if err.Status() == http.StatusNotFound {
			return nil
		}
		return err
	}

	request.Header.Add(headerXClientId, fmt.Sprintf("%d", at.ClientId))
	request.Header.Add(headerXCallerId, fmt.Sprintf("%d", at.UserId))

	return nil
}

func cleanRequest(request *http.Request) {
	if request == nil {
		return
	}

	request.Header.Del(headerXClientId)
	request.Header.Del(headerXCallerId)
}

func getAccessToken(accessTokenId string) (*accessToken, errors.RestError) {
	response := oauthRestClient.Get(fmt.Sprintf("/oauth/access_token/%s", accessTokenId))
	if response == nil || response.Response == nil {
		return nil, errors.NewInternalServerError("error when trying to getaccess token from request", errors.NewError("invalid restclient response when trying to get access token"))
	}

	if response.StatusCode > 299 {
		var restErr errors.RestError
		err := json.Unmarshal(response.Bytes(), &restErr)
		if err != nil {
			return nil, errors.NewInternalServerError("error when trying to getaccess token from request", errors.NewError("invalid error interface whentrying yo get access token"))
		}

		return nil, restErr
	}

	var at accessToken
	if err := json.Unmarshal(response.Bytes(), &at); err != nil {
		return nil, errors.NewInternalServerError("error when trying to getaccess token from request", errors.NewError("error when trying to unmarshal access token response"))
	}

	return &at, nil
}
