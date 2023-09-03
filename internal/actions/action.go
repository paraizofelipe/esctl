package actions

import (
	"log"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/paraizofelipe/esctl/internal/config"
)

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
