package cmd

import (
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/paraizofelipe/esctl/internal/actions"
	"github.com/urfave/cli/v2"
)

func NewSearchCommand() *cli.Command {

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

	return &cli.Command{
		Name:  "search",
		Usage: "Search documents in an index",
		Flags: appFlags,
		Action: func(ctx *cli.Context) error {
			esClient := ctx.Context.Value("esClient").(*elasticsearch.Client)
			searcher := actions.NewSearchAction(esClient)
			return searcher.SearchDoc(ctx)
		},
	}
}
