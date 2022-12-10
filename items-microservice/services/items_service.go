package services

import (
	"net/http"

	"github.com/menxqk/rest-microservices-in-go/common/errors"
	"github.com/menxqk/rest-microservices-in-go/items-microservice/domain/items"
)

var (
	ItemsService itemServiceInterface = &itemsService{}
)

type itemServiceInterface interface {
	Create(items.Item) (*items.Item, *errors.RestError)
	Get(string) (*items.Item, *errors.RestError)
}

type itemsService struct{}

func (s *itemsService) Create(item items.Item) (*items.Item, *errors.RestError) {
	if err := item.Save(); err != nil {
		return nil, err
	}

	return &item, nil
}

func (s *itemsService) Get(id string) (*items.Item, *errors.RestError) {

	return nil, errors.NewRestError("not implemented", http.StatusNotImplemented, "not_implemented", nil)
}
