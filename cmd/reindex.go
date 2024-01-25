package cmd

import (
	"fmt"
	"strings"

	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/paraizofelipe/esctl/internal/client"
	"github.com/paraizofelipe/esctl/internal/output"
	"github.com/urfave/cli/v2"
)

func ReindexCommand() *cli.Command {
	return &cli.Command{
		Name:  "reindex",
		Usage: "Reindex data from one index to another",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "source",
				Aliases:  []string{"s"},
				Usage:    "Specify the source index to reindex from",
				Required: false,
			},
			&cli.StringFlag{
				Name:     "dest",
				Aliases:  []string{"d"},
				Usage:    "Specify the destination index to reindex to",
				Required: false,
			},
			&cli.StringFlag{
				Name:     "body",
				Aliases:  []string{"b"},
				Usage:    "Specify the body of the reindex request",
				Required: false,
			},
			&cli.BoolFlag{
				Name:     "wait-for-completion",
				Aliases:  []string{"w"},
				Usage:    "Specify whether to wait for the reindex operation to complete",
				Required: false,
			},
		},
		Action: func(ctx *cli.Context) error {
			es := ctx.Context.Value("esClient").(client.ElasticClient)
			body := ctx.String("body")
			wait := ctx.Bool("wait-for-completion")
			if body == "" {
				body = fmt.Sprintf(`{"source": {"index": "%s"}, "dest": {"index": "%s"}}`, ctx.String("source"), ctx.String("dest"))
			}

			request := &esapi.ReindexRequest{
				Pretty:            true,
				Body:              strings.NewReader(body),
				WaitForCompletion: &wait,
			}
			jsonBytes, err := es.ExecRequest(ctx.Context, request)
			if err != nil {
				return err
			}
			output.PrintPrettyJSON(jsonBytes)
			return nil
		},
	}
}
