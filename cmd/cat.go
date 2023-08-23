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
		Usage: "Manage Elasticsearch aliases and diagnostic information",
		Flags: appFlags,
		Subcommands: []*cli.Command{
			{
				Name:   "indices",
				Usage:  "List all indices in Elasticsearch",
				Action: catManager.Indices,
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name:    "describe",
						Value:   false,
						Aliases: []string{"d"},
						Usage:   "Show detailed information about indices",
					},
				},
			},
			{
				Name:   "aliases",
				Usage:  "List all aliases in Elasticsearch",
				Action: catManager.Aliases,
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name:    "describe",
						Value:   false,
						Aliases: []string{"d"},
						Usage:   "Show detailed information about aliases",
					},
				},
			},
			{
				Name:   "nodes",
				Usage:  "List all nodes in Elasticsearch cluster",
				Action: catManager.Nodes,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "columns",
						Aliases: []string{"c"},
						Usage:   "Comma-separated list of columns to display",
					},
					&cli.BoolFlag{
						Name:    "describe",
						Value:   false,
						Aliases: []string{"d"},
						Usage:   "Show detailed information about nodes",
					},
				},
			},
			{
				Name:  "shards",
				Usage: "List shard information for indices",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "index",
						Aliases: []string{"i"},
						Usage:   "Filter by index name",
					},
					&cli.StringFlag{
						Name:    "columns",
						Aliases: []string{"c"},
						Usage:   "Comma-separated list of columns to display",
					},
					&cli.BoolFlag{
						Name:    "describe",
						Value:   false,
						Aliases: []string{"d"},
						Usage:   "Show detailed information about shards",
					},
				},
				Action: catManager.Shards,
			},
			{
				Name:  "thread-pool",
				Usage: "List thread pool statistics",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "thread-pool-pattern",
						Aliases: []string{"p"},
						Usage:   "Filter by thread pool pattern",
					},
					&cli.StringFlag{
						Name:    "columns",
						Aliases: []string{"c"},
						Usage:   "Comma-separated list of columns to display",
					},
					&cli.BoolFlag{
						Name:    "describe",
						Value:   false,
						Aliases: []string{"d"},
						Usage:   "Show detailed information about thread pools",
					},
				},
				Action: catManager.ThreadPool,
			},
			{
				Name:   "pending-tasks",
				Usage:  "List pending tasks in Elasticsearch",
				Action: catManager.PendingTasks,
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name:    "describe",
						Value:   false,
						Aliases: []string{"d"},
						Usage:   "Show detailed information about pending tasks",
					},
				},
			},
		},
	}
}
