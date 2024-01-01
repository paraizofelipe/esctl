package cmd

import (
	"fmt"

	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/paraizofelipe/esctl/internal/client"
	"github.com/urfave/cli/v2"
)

func NewDescribeSourceCommand() *cli.Command {
	return &cli.Command{
		Name:  "source",
		Usage: "Retrieve the source of a document in Elasticsearch using its ID",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "id",
				Aliases:  []string{"i"},
				Usage:    "Specify the unique identifier of the document whose source is to be retrieved",
				Required: true,
			},
			&cli.StringSliceFlag{
				Name:     "fields",
				Aliases:  []string{"f"},
				Usage:    "List specific stored fields to include in the response, separated by commas",
				Required: false,
			},
		},
		Action: func(ctx *cli.Context) error {
			es := ctx.Context.Value("esClient").(*client.ClusterElasticClient)
			docRequest := &esapi.GetSourceRequest{
				Index:          ctx.Args().Get(0),
				Pretty:         true,
				DocumentID:     ctx.String("id"),
				SourceIncludes: ctx.StringSlice("fields"),
			}
			jsonBytes, err := es.ExecRequest(ctx.Context, docRequest)
			fmt.Println(string(jsonBytes))
			return err
		},
	}
}
