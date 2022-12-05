package access_token

import "github.com/menxqk/rest-microservices-in-go/bookstore_oauth-api/src/utils/errors"

type Repository interface {
	GetById(string) (*AccessToken, *errors.RestError)
}

type Service interface {
	GetById(string) (*AccessToken, *errors.RestError)
}

func NewService(repo Repository) Service {
	return &service{repository: repo}
}

type service struct {
	repository Repository
}

func (s *service) GetById(accessTokenId string) (*AccessToken, *errors.RestError) {
	return s.repository.GetById(accessTokenId)
}
