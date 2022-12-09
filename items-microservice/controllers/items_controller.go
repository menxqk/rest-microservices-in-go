package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/menxqk/rest-microservices-in-go/common/errors"
	"github.com/menxqk/rest-microservices-in-go/common/oauth"
	"github.com/menxqk/rest-microservices-in-go/items-microservice/domain/items"
	"github.com/menxqk/rest-microservices-in-go/items-microservice/services"
	"github.com/menxqk/rest-microservices-in-go/items-microservice/utils/http_utils"
)

var (
	ItemsController itemsControllerInterface = &itemsController{}
)

type itemsControllerInterface interface {
	Create(http.ResponseWriter, *http.Request)
	Get(http.ResponseWriter, *http.Request)
}

type itemsController struct{}

func (c *itemsController) Create(w http.ResponseWriter, r *http.Request) {
	if err := oauth.AuthenticateRequest(r); err != nil {
		http_utils.RespondError(w, err)
		return
	}

	var itemRequest items.Item
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respErr := errors.NewBadRequestError("invalid request body")
		http_utils.RespondError(w, respErr)
	}
	defer r.Body.Close()

	if err := json.Unmarshal(requestBody, &itemRequest); err != nil {
		respErr := errors.NewBadRequestError("invalid json body")
		http_utils.RespondError(w, respErr)
	}

	itemRequest.Seller = oauth.GetCallerId(r)

	result, restErr := services.ItemsService.Create(itemRequest)
	if restErr != nil {
		http_utils.RespondError(w, restErr)
		return
	}

	http_utils.RespondJson(w, http.StatusCreated, result)
}

func (c *itemsController) Get(w http.ResponseWriter, r *http.Request) {

}
