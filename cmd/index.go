package cmd

import (
	"strings"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/paraizofelipe/elastic_tools/internal/actions"
	"github.com/urfave/cli/v2"
)

func NewIndexCommand(esClient *elasticsearch.Client) *cli.Command {

	appFlags := []cli.Flag{
		&cli.BoolFlag{
			Name:  "p",
			Usage: "Format response as pretty-printed JSON",
		},
	}

	return &cli.Command{
		Name:  "index",
		Usage: "Manage Elasticsearch indices",
		Flags: appFlags,
		Subcommands: []*cli.Command{
			{
				Name:  "create",
				Usage: "Create a new index",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "n",
						Usage: "Name of the index to be created",
					},
				},
				Action: func(c *cli.Context) error {
					indexName := c.String("n")
					pretty := c.Bool("p")
					actions.CreateIndex(esClient, indexName, pretty)
					return nil
				},
			},
			{
				Name:  "delete",
				Usage: "Delete a new index",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "n",
						Usage: "List of names for creating indices",
					},
				},
				Action: func(c *cli.Context) error {
					names := c.String("n")
					pretty := c.Bool("p")
					indexNames := strings.Split(names, ",")
					actions.DeleteIndex(esClient, indexNames, pretty)
					return nil
				},
			},
			{
				Name:  "get",
				Usage: "get a created indice",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "n",
						Usage: "List of names for information",
					},
				},
				Action: func(c *cli.Context) error {
					names := c.String("n")
					pretty := c.Bool("p")
					indexNames := strings.Split(names, ",")
					actions.GetIndex(esClient, indexNames, pretty)
					return nil
				},
			},
		},
	}
}
