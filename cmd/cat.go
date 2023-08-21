package cmd

import (
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/paraizofelipe/elastic_tools/internal/actions"
	"github.com/urfave/cli/v2"
)

func NewCatCommand(esClient *elasticsearch.Client) *cli.Command {

	appFlags := []cli.Flag{
		&cli.BoolFlag{
			Name:    "pretty",
			Aliases: []string{"p"},
			Usage:   "Format response as pretty-printed JSON",
		},
	}

	catManager := actions.NewCatAction(esClient)

	return &cli.Command{
		Name:  "cat",
		Usage: "Manage Elasticsearch aliases",
		Flags: appFlags,
		Subcommands: []*cli.Command{
			{
				Name:   "indices",
				Usage:  "list all incides",
				Action: catManager.Indices,
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name:    "describe",
						Value:   false,
						Aliases: []string{"d"},
						Usage:   "",
					},
				},
			},
			{
				Name:   "aliases",
				Usage:  "list all aliases",
				Action: catManager.Aliases,
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name:    "describe",
						Value:   false,
						Aliases: []string{"d"},
						Usage:   "",
					},
				},
			},
			{
				Name:   "nodes",
				Usage:  "list all nodes",
				Action: catManager.Nodes,
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name:    "describe",
						Value:   false,
						Aliases: []string{"d"},
						Usage:   "",
					},
				},
			},
			{
				Name:  "shards",
				Usage: "list all nodes",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "index",
						Aliases: []string{"i"},
						Usage:   "",
					},
					&cli.StringFlag{
						Name:    "columns",
						Aliases: []string{"c"},
						Usage:   "",
					},
					&cli.BoolFlag{
						Name:    "describe",
						Value:   false,
						Aliases: []string{"d"},
						Usage:   "",
					},
				},
				Action: catManager.Shards,
			},
			{
				Name:  "thread-pool",
				Usage: "list all threads pool",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "thread-pool-pattern",
						Aliases: []string{"p"},
						Usage:   "",
					},
					&cli.StringFlag{
						Name:    "columns",
						Aliases: []string{"c"},
						Usage:   "",
					},
					&cli.BoolFlag{
						Name:    "describe",
						Value:   false,
						Aliases: []string{"d"},
						Usage:   "",
					},
				},
				Action: catManager.ThreadPool,
			},
			{
				Name:   "pending-tasks",
				Usage:  "list pending tasks",
				Action: catManager.PendingTasks,
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name:    "describe",
						Value:   false,
						Aliases: []string{"d"},
						Usage:   "",
					},
				},
			},
		},
	}
}
