package cmd

import (
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/paraizofelipe/esctl/internal/client"
	"github.com/urfave/cli/v2"
)

func DescribeTaskCommand() *cli.Command {
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

func CancelTaskCommand() *cli.Command {
	return &cli.Command{
		Name:  "cancel",
		Usage: "Cancel task by id",
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
		Usage: "List tasks",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "nodes",
				Aliases:  []string{"n"},
				Usage:    "Node id",
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
		Usage: "Manage tasks",
		Subcommands: []*cli.Command{
			DescribeTaskCommand(),
			CancelTaskCommand(),
			GetTaskCommand(),
		},
	}
}
