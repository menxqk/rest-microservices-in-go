package items

import (
	"github.com/menxqk/rest-microservices-in-go/common/errors"
	"github.com/menxqk/rest-microservices-in-go/items-microservice/clients/elasticsearch"
)

const (
	ITEMS_INDEX = "items"
)

func (i *Item) Save() errors.RestError {
	result, err := elasticsearch.Client.Index(ITEMS_INDEX, i)
	if err != nil {
		return errors.NewInternalServerError("error when trying to save item", errors.NewError("database error"))
	}

	i.Id = result.Id
	return nil
}
