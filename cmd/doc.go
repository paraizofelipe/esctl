package cmd

import (
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/paraizofelipe/elastic_tools/internal/actions"
	"github.com/urfave/cli/v2"
)

func NewDocCommand(esClient *elasticsearch.Client) *cli.Command {

	appFlags := []cli.Flag{
		&cli.BoolFlag{
			Name:  "p",
			Usage: "Format response as pretty-printed JSON",
		},
	}

	actions := actions.NewDocAction(esClient)

	return &cli.Command{
		Name:  "doc",
		Usage: "Manage Elasticsearch document",
		Flags: appFlags,
		Subcommands: []*cli.Command{
			{
				Name:  "list",
				Usage: "List documents in index",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "i",
						Usage: "Name of Indice that will be consulted",
					},
				},
				Action: func(c *cli.Context) error {
					indexName := c.String("i")
					pretty := c.Bool("p")
					actions.ListDoc(indexName, pretty)
					return nil
				},
			},
			{
				Name:  "add",
				Usage: "Add documents in index",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "i",
						Usage: "Name of Indice that will be consulted",
					},
					&cli.StringFlag{
						Name:  "d",
						Usage: "Document",
					},
				},
				Action: func(c *cli.Context) error {
					indexName := c.String("i")
					document := c.String("d")
					actions.AddDoc(indexName, document)
					return nil
				},
			},
		},
	}
}
