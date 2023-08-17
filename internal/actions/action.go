package actions

import (
	"log"

	"github.com/elastic/go-elasticsearch/v8"
)

func CreateClient(esNodes []string) (esClient *elasticsearch.Client, err error) {
	cfg := elasticsearch.Config{
		Addresses: esNodes,
	}

	if esClient, err = elasticsearch.NewClient(cfg); err != nil {
		log.Fatalf("Erro to access elasticsearch: %s", err)
	}

	return
}
