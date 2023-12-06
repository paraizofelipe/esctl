package cmd

import (
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/paraizofelipe/esctl/internal/client"
	"github.com/urfave/cli/v2"
)

func DescribeTaskCommand() *cli.Command {
	return &cli.Command{
		Name:  "task",
		Usage: "Retrieve detailed information about a specific task using its ID",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "id",
				Aliases:  []string{"i"},
				Usage:    "Specify the unique identifier of the task to describe",
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

func CancelTaskCommand() *cli.Command {
	return &cli.Command{
		Name:  "cancel",
		Usage: "Cancel a specified task using its ID",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "id",
				Aliases:  []string{"i"},
				Usage:    "Provide the unique identifier of the task to be cancelled",
				Required: false,
			},
		},
		Action: func(ctx *cli.Context) error {
			es := ctx.Context.Value("esClient").(*client.Elastic)
			request := &esapi.TasksCancelRequest{
				Pretty: true,
				TaskID: ctx.String("id"),
			}
			return es.ExecRequest(ctx.Context, request)
		},
	}
}

func GetTaskCommand() *cli.Command {
	return &cli.Command{
		Name:  "list",
		Usage: "List all tasks or filter tasks by node IDs",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "nodes",
				Aliases:  []string{"n"},
				Usage:    "Filter tasks by specific node IDs, separated by commas",
				Required: false,
			},
		},
		Action: func(ctx *cli.Context) error {
			es := ctx.Context.Value("esClient").(*client.Elastic)
			request := &esapi.TasksListRequest{
				Pretty: true,
				Nodes:  ctx.StringSlice("nodes"),
			}
			return es.ExecRequest(ctx.Context, request)
		},
	}
}

func TaskCommand() *cli.Command {
	return &cli.Command{
		Name:  "task",
		Usage: "Commands to manage and interact with Elasticsearch tasks",
		Subcommands: []*cli.Command{
			DescribeTaskCommand(),
			CancelTaskCommand(),
			GetTaskCommand(),
		},
	}
}
