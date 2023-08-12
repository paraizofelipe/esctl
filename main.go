package main

import (
	"log"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/paraizofelipe/elastic_tools/cmd"
)

func main() {

	var (
		err      error
		esClient *elasticsearch.Client
	)

	cfg := elasticsearch.Config{
		Addresses: []string{"http://192.168.68.119:9200"},
	}

	if esClient, err = elasticsearch.NewClient(cfg); err != nil {
		log.Fatalf("Exec error: %s", err)
	}

	cmd.Execute(esClient)
}
