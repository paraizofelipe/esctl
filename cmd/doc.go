package cmd

import (
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/paraizofelipe/esctl/internal/client"
	"github.com/urfave/cli/v2"
)

func NewGetDocCommand() *cli.Command {
	return &cli.Command{
		Name:  "doc",
		Usage: "Get document by id",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "id",
				Aliases:  []string{"i"},
				Usage:    "Document id",
				Required: true,
			},
			&cli.StringSliceFlag{
				Name:     "fields",
				Aliases:  []string{"f"},
				Usage:    "Comma-separated list of stored fields to return in the response",
				Required: false,
			},
		},
		Action: func(ctx *cli.Context) error {
			es := ctx.Context.Value("esClient").(*client.Elastic)
			docRequest := &esapi.GetRequest{
				Pretty:         true,
				DocumentID:     ctx.String("id"),
				Index:          ctx.Args().Get(0),
				SourceIncludes: ctx.StringSlice("fields"),
			}
			return es.ExecRequest(ctx.Context, docRequest)
		},
	}
}
