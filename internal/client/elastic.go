package client

import (
	"context"
	"fmt"
	"log"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/paraizofelipe/esctl/internal/config"
)

type Elastic struct {
	client *elasticsearch.Client
}

func CreateClient(host config.Host) (esClient *elasticsearch.Client, err error) {
	cfg := elasticsearch.Config{
		Addresses: host.Address,
		Username:  host.Username,
		Password:  host.Password,
	}

	if esClient, err = elasticsearch.NewClient(cfg); err != nil {
		log.Fatalf("Erro to access elasticsearch: %s", err)
	}

	return
}

func NewElastic(config elasticsearch.Config) *Elastic {
	client, err := elasticsearch.NewClient(config)
	if err != nil {
		log.Fatalf("Erro to access elasticsearch: %s", err)
	}

	return &Elastic{
		client: client,
	}
}

func (es *Elastic) ExecRequest(ctx context.Context, request esapi.Request) (err error) {
	var res *esapi.Response

	if res, err = request.Do(ctx, es.client); err != nil {
		return fmt.Errorf("Error to get document: %s", err)
	}

	defer res.Body.Close()

	if res.IsError() {
		return fmt.Errorf("[REQ-ERROR]: %s", res.String())
	}

	fmt.Println(res.String())

	return
}
