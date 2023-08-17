package cmd

import (
	"fmt"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/paraizofelipe/elastic_tools/internal/actions"
	"github.com/paraizofelipe/elastic_tools/internal/file"
	"github.com/urfave/cli/v2"
)

func NewIndexCommand(esClient *elasticsearch.Client) *cli.Command {

	appFlags := []cli.Flag{
		&cli.BoolFlag{
			Name:    "pretty",
			Aliases: []string{"p"},
			Usage:   "Format response as pretty-printed JSON",
		},
	}

	indexer := actions.NewIndexAction(esClient)

	return &cli.Command{
		Name:  "index",
		Usage: "Manage Elasticsearch indices",
		Flags: appFlags,
		Subcommands: []*cli.Command{
			{
				Name:   "create",
				Usage:  "Create a new index",
				Action: indexer.CreateIndex,
			},
			{
				Name:   "delete",
				Usage:  "Delete a new index",
				Action: indexer.DeleteIndex,
			},
			{
				Name:   "get",
				Usage:  "get a created indice",
				Action: indexer.GetIndex,
			},
			{
				Name:  "add",
				Usage: "add document inside index",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "document",
						Usage:   "JSON string with the contents of the document",
						Aliases: []string{"d"},
						Action: func(ctx *cli.Context, doc string) (err error) {
							if !file.IsContentValid(doc) {
								return fmt.Errorf("JSON string invalid!")
							}
							return
						},
					},
					&cli.StringFlag{
						Name:    "document-file",
						Usage:   "path of the JSON file with the document",
						Aliases: []string{"f"},
						Action: func(ctx *cli.Context, docFile string) (err error) {
							if !file.Exists(docFile) {
								return fmt.Errorf("file %s not found!", docFile)
							}
							return
						},
					},
				},
				Action: indexer.AddDoc,
			},
			{
				Name:  "list",
				Usage: "List documents in index",
				Flags: []cli.Flag{
					&cli.IntFlag{
						Name:    "size",
						Usage:   "number of hits to return",
						Aliases: []string{"s"},
					},
				},
				Action: indexer.ListDoc,
			},
		},
	}
}
