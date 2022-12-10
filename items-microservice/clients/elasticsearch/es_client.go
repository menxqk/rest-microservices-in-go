package elasticsearch

import (
	"context"
	"fmt"
	"time"

	"github.com/menxqk/rest-microservices-in-go/common/logger"
	"github.com/olivere/elastic"
)

var (
	Client esClientInterface = &esClient{}
)

type esClientInterface interface {
	setClient(c *elastic.Client)
	Index(string, interface{}) (*elastic.IndexResponse, error)
}

type esClient struct {
	client *elastic.Client
}

func (ec *esClient) Index(index string, doc interface{}) (*elastic.IndexResponse, error) {
	ctx := context.Background()
	result, err := ec.client.Index().Index(index).BodyJson(doc).Do(ctx)
	if err != nil {
		logger.Error(fmt.Sprintf("error when trying to index document in %s", index), err)
		return nil, err
	}
	return result, err

}

func (ec *esClient) setClient(c *elastic.Client) {
	ec.client = c
}

func Init() {
	log := logger.GetLogger()
	client, err := elastic.NewClient(
		elastic.SetURL("http://127.0.0.1:9200"),
		elastic.SetHealthcheckInterval(10*time.Second),
		elastic.SetErrorLog(log),
		elastic.SetInfoLog(log),
	)
	if err != nil {
		panic(err)
	}

	Client.setClient(client)
}
