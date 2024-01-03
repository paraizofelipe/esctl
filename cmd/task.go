package cmd

import (
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/paraizofelipe/esctl/internal/client"
	"github.com/paraizofelipe/esctl/internal/output"
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
			es := ctx.Context.Value("esClient").(client.ElasticClient)
			request := &esapi.TasksGetRequest{
				TaskID: ctx.String("id"),
			}
			jsonBytes, err := es.ExecRequest(ctx.Context, request)
			output.PrintPrettyJSON(jsonBytes)
			return err
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
			es := ctx.Context.Value("esClient").(client.ElasticClient)
			request := &esapi.TasksCancelRequest{
				TaskID: ctx.String("id"),
			}
			jsonBytes, err := es.ExecRequest(ctx.Context, request)
			output.PrintPrettyJSON(jsonBytes)
			return err
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
			es := ctx.Context.Value("esClient").(client.ElasticClient)
			request := &esapi.TasksListRequest{
				Nodes: ctx.StringSlice("nodes"),
			}
			jsonBytes, err := es.ExecRequest(ctx.Context, request)
			output.PrintPrettyJSON(jsonBytes)
			return err
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
