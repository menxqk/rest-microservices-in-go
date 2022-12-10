package access_token

import (
	"strings"

	"github.com/menxqk/rest-microservices-in-go/common/errors"
)

type Repository interface {
	GetById(string) (*AccessToken, errors.RestError)
	Create(AccessToken) errors.RestError
	UpdateExpirationTime(AccessToken) errors.RestError
}

type Service interface {
	GetById(string) (*AccessToken, errors.RestError)
	Create(AccessToken) errors.RestError
	UpdateExpirationTime(AccessToken) errors.RestError
}

func NewService(repo Repository) Service {
	return &service{repository: repo}
}

type service struct {
	repository Repository
}

func (s *service) GetById(accessTokenId string) (*AccessToken, errors.RestError) {
	accessTokenId = strings.TrimSpace(accessTokenId)
	if accessTokenId == "" {
		return nil, errors.NewBadRequestError("invalid access token id")
	}

	accessToken, err := s.repository.GetById(accessTokenId)
	if err != nil {
		return nil, err
	}

	return accessToken, nil
}

func (s *service) Create(at AccessToken) errors.RestError {
	if err := at.Validate(); err != nil {
		return err
	}

	return s.repository.Create(at)
}

func (s *service) UpdateExpirationTime(at AccessToken) errors.RestError {
	if err := at.Validate(); err != nil {
		return nil
	}

	return s.repository.UpdateExpirationTime(at)
}
