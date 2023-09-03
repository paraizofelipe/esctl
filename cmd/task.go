package cmd

import (
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/urfave/cli/v2"
)

func NewTaskCommand(esClient *elasticsearch.Client) *cli.Command {

	appFlags := []cli.Flag{
		&cli.BoolFlag{
			Name:    "pretty",
			Aliases: []string{"p"},
			Usage:   "Format response as pretty-printed JSON",
		},
	}

	return &cli.Command{
		Name:  "task",
		Usage: "Manage Elasticsearch background tasks",
		Flags: appFlags,
		Subcommands: []*cli.Command{
			{
				Name:  "list",
				Usage: "List all indices in Elasticsearch",
				// Action: catManager.Indices,
			},
			{
				Name:  "get",
				Usage: "List all indices in Elasticsearch",
				// Action: catManager.Indices,
			},
			{
				Name:  "cancel",
				Usage: "List all indices in Elasticsearch",
				// Action: catManager.Indices,
			},
		},
	}
}
