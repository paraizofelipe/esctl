package cmd

import (
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/paraizofelipe/esctl/internal/actions"
	"github.com/urfave/cli/v2"
)

func NewSearchCommand(esClient *elasticsearch.Client) *cli.Command {

	appFlags := []cli.Flag{
		&cli.BoolFlag{
			Name:    "pretty",
			Aliases: []string{"p"},
			Usage:   "Format the response as pretty-printed JSON",
		},
		&cli.StringFlag{
			Name:     "query",
			Aliases:  []string{"q"},
			Usage:    "The search query string",
			Required: true,
		},
	}

	searcher := actions.NewSearchAction(esClient)

	return &cli.Command{
		Name:   "search",
		Usage:  "Search documents in an index",
		Flags:  appFlags,
		Action: searcher.SearchDoc,
	}
}
