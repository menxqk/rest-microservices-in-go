package db

import (
	"github.com/gocql/gocql"
	"github.com/menxqk/rest-microservices-in-go/common/errors"
	"github.com/menxqk/rest-microservices-in-go/oauth-microservice/src/clients/cassandra"
	"github.com/menxqk/rest-microservices-in-go/oauth-microservice/src/domain/access_token"
)

const (
	queryGetAccessToken    = "SELECT access_token, user_id, client_id, expires FROM access_tokens WHERE access_token=?;"
	queryCreateAccessToken = "INSERT INTO access_tokens (access_token, user_id, client_id, expires) VALUES(?, ?, ?, ?);"
	queryUpdateExpires     = "UPDATE access_tokens SET expires=? WHERE access_token=?;"
)

type DbRepository interface {
	GetById(string) (*access_token.AccessToken, *errors.RestError)
	Create(access_token.AccessToken) *errors.RestError
	UpdateExpirationTime(access_token.AccessToken) *errors.RestError
}

func NewRepository() DbRepository {
	return &dbRepository{}
}

type dbRepository struct {
}

func (dr *dbRepository) GetById(id string) (*access_token.AccessToken, *errors.RestError) {
	session := cassandra.GetSession()

	var result access_token.AccessToken
	err := session.Query(queryGetAccessToken, id).Scan(&result.AccessToken, &result.UserId, &result.ClientId, &result.Expires)
	if err != nil {
		if err == gocql.ErrNotFound {
			return nil, errors.NewNotFoundError("no access token found with given id")
		}
		return nil, errors.NewInternalServerError(err.Error())
	}

	return &result, nil
}

func (dr *dbRepository) Create(at access_token.AccessToken) *errors.RestError {
	session := cassandra.GetSession()

	err := session.Query(queryCreateAccessToken, at.AccessToken, at.UserId, at.ClientId, at.Expires).Exec()
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	return nil
}

func (dr *dbRepository) UpdateExpirationTime(at access_token.AccessToken) *errors.RestError {
	session := cassandra.GetSession()

	err := session.Query(queryUpdateExpires, at.Expires, at.AccessToken).Exec()
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	return nil
}
