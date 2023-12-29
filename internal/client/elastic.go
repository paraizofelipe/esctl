package client

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
)

type ClusterElasticClient struct {
	client *elasticsearch.Client
}

type ElasticClient interface {
	ExecRequest(ctx context.Context, request esapi.Request) (err error)
}

func NewElastic(config elasticsearch.Config) (ElasticClient, error) {
	if config.Addresses == nil {
		config = elasticsearch.Config{
			Addresses: []string{"http://127.0.0.1:9200"},
		}
	}

	client, err := elasticsearch.NewClient(config)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error to access elasticsearch: %s", err))
	}

	return &ClusterElasticClient{
		client: client,
	}, nil
}

func (es *ClusterElasticClient) ExecRequest(ctx context.Context, request esapi.Request) (err error) {
	var res *esapi.Response

	if res, err = request.Do(ctx, es.client); err != nil {
		return fmt.Errorf("Error to get document: %s", err)
	}

	defer res.Body.Close()

	if res.IsError() {
		return fmt.Errorf("[REQ-ERROR]: %s", res.String())
	}

	b, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(b))

	return
}
