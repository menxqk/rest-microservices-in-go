package db

import (
	"github.com/menxqk/rest-microservices-in-go/bookstore_oauth-api/src/domain/access_token"
	"github.com/menxqk/rest-microservices-in-go/bookstore_oauth-api/src/utils/errors"
)

type DbRepository interface {
	GetById(string) (*access_token.AccessToken, *errors.RestError)
}

func NewRepository() DbRepository {
	return &dbRepository{}
}

type dbRepository struct {
}

func (dr *dbRepository) GetById(accessTokenId string) (*access_token.AccessToken, *errors.RestError) {
	return nil, errors.NewInternalServerError("database connection not implemented yet")
}
