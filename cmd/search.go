package cmd

import (
	"strings"

	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/paraizofelipe/esctl/internal/client"
	"github.com/urfave/cli/v2"
)

func NewSearchCommand() *cli.Command {

	appFlags := []cli.Flag{
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
			es := ctx.Context.Value("esClient").(*client.Elastic)
			request := &esapi.SearchRequest{
				Pretty: true,
				Index:  ctx.Args().Slice(),
				Body:   strings.NewReader(ctx.String("query")),
			}

			return es.ExecRequest(ctx.Context, request)
		},
	}
}
