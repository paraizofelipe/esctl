package cmd

import (
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/paraizofelipe/esctl/internal/client"
	"github.com/urfave/cli/v2"
)

func NewGetTaskCommand() *cli.Command {
	return &cli.Command{
		Name:  "task",
		Usage: "Get task by id",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "id",
				Aliases:  []string{"i"},
				Usage:    "Task id",
				Required: false,
			},
		},
		Action: func(ctx *cli.Context) error {
			es := ctx.Context.Value("esClient").(*client.Elastic)
			request := &esapi.TasksGetRequest{
				Pretty: true,
				TaskID: ctx.String("id"),
			}
			return es.ExecRequest(ctx.Context, request)
		},
	}
}
