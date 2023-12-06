package cmd

import (
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/paraizofelipe/esctl/internal/client"
	"github.com/urfave/cli/v2"
)

func DescribeCountCommand() *cli.Command {
	return &cli.Command{
		Name:  "count",
		Usage: "Determine the number of documents in Elasticsearch matching a specified query",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "query",
				Aliases:  []string{"q"},
				Usage:    "Provide the Elasticsearch query in JSON format to filter the documents for counting",
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
