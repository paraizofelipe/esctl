package cmd

import (
	"log"
	"os"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/urfave/cli/v2"
)

func NewRootCommand(esClient *elasticsearch.Client) *cli.App {
	app := &cli.App{}
	app.Name = "elastic_tools"
	app.Usage = "Elasticsearch Tools CLI"
	app.Version = "1.0.0"

	app.Commands = []*cli.Command{
		NewIndexCommand(esClient),
	}

	return app
}

func Execute(esClient *elasticsearch.Client) {
	app := NewRootCommand(esClient)

	err := app.Run(os.Args)
	if err != nil {
		log.Fatalf("Exec error: %s", err)
	}
}
