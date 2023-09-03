package cmd

import (
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/paraizofelipe/esctl/internal/actions"
	"github.com/urfave/cli/v2"
)

func NewCatCommand() *cli.Command {

	appFlags := []cli.Flag{
		&cli.BoolFlag{
			Name:    "pretty",
			Aliases: []string{"p"},
			Usage:   "Format response as pretty-printed JSON",
		},
	}

	return &cli.Command{
		Name:  "cat",
		Usage: "Manage Elasticsearch aliases and diagnostic information",
		Flags: appFlags,
		Subcommands: []*cli.Command{
			{
				Name:  "indices",
				Usage: "List all indices in Elasticsearch",
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
						Usage:   "Show detailed information about indices",
					},
				},
				Action: func(ctx *cli.Context) error {
					esClient := ctx.Context.Value("esClient").(*elasticsearch.Client)
					catManager := actions.NewCatAction(esClient)
					return catManager.Indices(ctx)
				},
			},
			{
				Name:  "aliases",
				Usage: "List all aliases in Elasticsearch",
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
						Usage:   "Show detailed information about aliases",
					},
				},
				Action: func(ctx *cli.Context) error {
					esClient := ctx.Context.Value("esClient").(*elasticsearch.Client)
					catManager := actions.NewCatAction(esClient)
					return catManager.Aliases(ctx)
				},
			},
			{
				Name:  "nodes",
				Usage: "List all nodes in Elasticsearch cluster",
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
				Action: func(ctx *cli.Context) error {
					esClient := ctx.Context.Value("esClient").(*elasticsearch.Client)
					catManager := actions.NewCatAction(esClient)
					return catManager.Nodes(ctx)
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
				Action: func(ctx *cli.Context) error {
					esClient := ctx.Context.Value("esClient").(*elasticsearch.Client)
					catManager := actions.NewCatAction(esClient)
					return catManager.Shards(ctx)
				},
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
				Action: func(ctx *cli.Context) error {
					esClient := ctx.Context.Value("esClient").(*elasticsearch.Client)
					catManager := actions.NewCatAction(esClient)
					return catManager.ThreadPool(ctx)
				},
			},
			{
				Name:  "pending-tasks",
				Usage: "List pending tasks in Elasticsearch",
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name:    "describe",
						Value:   false,
						Aliases: []string{"d"},
						Usage:   "Show detailed information about pending tasks",
					},
				},
				Action: func(ctx *cli.Context) error {
					esClient := ctx.Context.Value("esClient").(*elasticsearch.Client)
					catManager := actions.NewCatAction(esClient)
					return catManager.PendingTasks(ctx)
				},
			},
			{
				Name:  "tasks",
				Usage: "Returns information about tasks currently executing in the cluster, similar to the task management API",
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name:    "describe",
						Value:   false,
						Aliases: []string{"d"},
						Usage:   "Show detailed information about pending tasks",
					},
					&cli.StringFlag{
						Name:    "columns",
						Aliases: []string{"c"},
						Usage:   "Comma-separated list of columns to display",
					},
				},
				Action: func(ctx *cli.Context) error {
					esClient := ctx.Context.Value("esClient").(*elasticsearch.Client)
					catManager := actions.NewCatAction(esClient)
					return catManager.Tasks(ctx)
				},
			},
			{
				Name:  "health",
				Usage: "Returns the health status of a cluster, similar to the cluster health API",
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name:    "describe",
						Value:   false,
						Aliases: []string{"d"},
						Usage:   "Show detailed information about pending tasks",
					},
					&cli.StringFlag{
						Name:    "columns",
						Aliases: []string{"c"},
						Usage:   "Comma-separated list of columns to display",
					},
				},
				Action: func(ctx *cli.Context) error {
					esClient := ctx.Context.Value("esClient").(*elasticsearch.Client)
					catManager := actions.NewCatAction(esClient)
					return catManager.Health(ctx)
				},
			},
			{
				Name:  "repositories",
				Usage: "Returns the snapshot repositories for a cluster",
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name:    "describe",
						Value:   false,
						Aliases: []string{"d"},
						Usage:   "Show detailed information about pending tasks",
					},
					&cli.StringFlag{
						Name:    "columns",
						Aliases: []string{"c"},
						Usage:   "Comma-separated list of columns to display",
					},
				},
				Action: func(ctx *cli.Context) error {
					esClient := ctx.Context.Value("esClient").(*elasticsearch.Client)
					catManager := actions.NewCatAction(esClient)
					return catManager.Repositories(ctx)
				},
			},
			{
				Name:  "snapshots",
				Usage: "Returns information about the snapshots stored in one or more repositories",
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name:    "describe",
						Value:   false,
						Aliases: []string{"d"},
						Usage:   "Show detailed information about pending tasks",
					},
					&cli.StringFlag{
						Name:    "columns",
						Aliases: []string{"c"},
						Usage:   "Comma-separated list of columns to display",
					},
					&cli.StringFlag{
						Name:    "repository",
						Aliases: []string{"r"},
						Usage:   "Comma-separated list of snapshot repositories used",
					},
				},
				Action: func(ctx *cli.Context) error {
					esClient := ctx.Context.Value("esClient").(*elasticsearch.Client)
					catManager := actions.NewCatAction(esClient)
					return catManager.Snapshots(ctx)
				},
			},
		},
	}
}
