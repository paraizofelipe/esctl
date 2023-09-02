package cmd

import (
	"fmt"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/paraizofelipe/esctl/internal/actions"
	"github.com/paraizofelipe/esctl/internal/file"
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
				Name:  "create",
				Usage: "Create a new index in Elasticsearch",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "settings-file",
						Usage:   "Path to the configuration file for the index",
						Aliases: []string{"f"},
					},
				},
				Action: indexer.CreateIndex,
			},
			{
				Name:   "delete",
				Usage:  "Delete an existing index from Elasticsearch",
				Action: indexer.DeleteIndex,
			},
			{
				Name:   "foce-merge",
				Usage:  "Forces a merge on the shards of one or more indices",
				Action: indexer.ForceMerge,
			},
			{
				Name:   "get",
				Usage:  "Get information about a created index",
				Action: indexer.GetIndex,
			},
			{
				Name:  "add",
				Usage: "Add a document to an index",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "document",
						Usage:   "JSON string with the contents of the document",
						Aliases: []string{"d"},
						Action: func(ctx *cli.Context, doc string) (err error) {
							if !file.IsContentValid(doc) {
								return fmt.Errorf("Invalid JSON string!")
							}
							return
						},
					},
					&cli.StringFlag{
						Name:    "document-file",
						Usage:   "Path to the JSON file with the document",
						Aliases: []string{"f"},
						Action: func(ctx *cli.Context, docFile string) (err error) {
							if !file.Exists(docFile) {
								return fmt.Errorf("File %s not found!", docFile)
							}
							return
						},
					},
				},
				Action: indexer.AddDoc,
			},
			{
				Name:  "bulk",
				Usage: "Perform bulk operations on an index",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "bulk-file",
						Usage:    "Path to the JSON file with bulk operations",
						Aliases:  []string{"f"},
						Required: true,
						Action: func(ctx *cli.Context, bulkFile string) (err error) {
							if !file.Exists(bulkFile) {
								return fmt.Errorf("File %s not found!", bulkFile)
							}
							return
						},
					},
				},
				Action: indexer.ExecBulkOperation,
			},
			{
				Name:  "list",
				Usage: "List documents in an index",
				Flags: []cli.Flag{
					&cli.IntFlag{
						Name:    "size",
						Usage:   "Number of hits to return",
						Aliases: []string{"s"},
					},
				},
				Action: indexer.ListDoc,
			},
		},
	}
}
