package actions

import (
	"log"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/paraizofelipe/esctl/internal/config"
)

func CreateClient(setup *config.ConfigFile) (esClient *elasticsearch.Client, err error) {
	cfg := elasticsearch.Config{
		Addresses: setup.Elastic,
		Username:  setup.Username,
		Password:  setup.Password,
	}

	if esClient, err = elasticsearch.NewClient(cfg); err != nil {
		log.Fatalf("Erro to access elasticsearch: %s", err)
	}

	return
}
