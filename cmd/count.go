package cmd

import (
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/paraizofelipe/esctl/internal/client"
	"github.com/urfave/cli/v2"
)

func DescribeCountCommand() *cli.Command {
	return &cli.Command{
		Name:  "count",
		Usage: "Count documents matching a query",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "query",
				Aliases:  []string{"q"},
				Usage:    "Query to be executed",
				Required: false,
			},
		},
		Action: func(ctx *cli.Context) error {
			indexPatterns := ctx.Args().Slice()
			es := ctx.Context.Value("esClient").(*client.Elastic)
			request := &esapi.CountRequest{
				Pretty: true,
				Index:  indexPatterns,
				Query:  ctx.String("query"),
			}
			return es.ExecRequest(ctx.Context, request)
		},
	}
}
