package cmd

import (
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/paraizofelipe/esctl/internal/client"
	"github.com/urfave/cli/v2"
)

func NewCatCommand() *cli.Command {

	return &cli.Command{
		Name:  "cat",
		Usage: "Show information about Elasticsearch cluster in text mode",
		Subcommands: []*cli.Command{
			{
				Name:  "indices",
				Usage: "List all indices in Elasticsearch",
				Flags: []cli.Flag{
					&cli.StringSliceFlag{
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
					es := ctx.Context.Value("esClient").(*client.Elastic)
					columns := ctx.StringSlice("columns")
					request := &esapi.CatIndicesRequest{
						V:      esapi.BoolPtr(true),
						Pretty: true,
						H:      columns,
						Help:   esapi.BoolPtr(ctx.Bool("describe")),
					}
					return es.ExecRequest(ctx.Context, request)
				},
			},
			{
				Name:  "aliases",
				Usage: "List all aliases in Elasticsearch",
				Flags: []cli.Flag{
					&cli.StringSliceFlag{
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
					es := ctx.Context.Value("esClient").(*client.Elastic)
					columns := ctx.StringSlice("columns")
					request := &esapi.CatAliasesRequest{
						V:      esapi.BoolPtr(true),
						Pretty: true,
						H:      columns,
						Help:   esapi.BoolPtr(ctx.Bool("describe")),
					}
					return es.ExecRequest(ctx.Context, request)
				},
			},
			{
				Name:  "nodes",
				Usage: "List all nodes in Elasticsearch cluster",
				Flags: []cli.Flag{
					&cli.StringSliceFlag{
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
					es := ctx.Context.Value("esClient").(*client.Elastic)
					columns := ctx.StringSlice("columns")
					request := &esapi.CatNodesRequest{
						V:      esapi.BoolPtr(true),
						Pretty: true,
						H:      columns,
						Help:   esapi.BoolPtr(ctx.Bool("describe")),
					}
					return es.ExecRequest(ctx.Context, request)
				},
			},
			{
				Name:  "shards",
				Usage: "List shard information for indices",
				Flags: []cli.Flag{
					&cli.StringSliceFlag{
						Name:    "index",
						Aliases: []string{"i"},
						Usage:   "Filter by index name",
					},
					&cli.StringSliceFlag{
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
					es := ctx.Context.Value("esClient").(*client.Elastic)
					columns := ctx.StringSlice("columns")

					request := &esapi.CatShardsRequest{
						V:      esapi.BoolPtr(true),
						Pretty: true,
						Index:  ctx.StringSlice("index"),
						H:      columns,
						Help:   esapi.BoolPtr(ctx.Bool("describe")),
					}
					return es.ExecRequest(ctx.Context, request)
				},
			},
			{
				Name:  "thread-pool",
				Usage: "List thread pool statistics",
				Flags: []cli.Flag{
					&cli.StringSliceFlag{
						Name:    "thread-pool-pattern",
						Aliases: []string{"p"},
						Usage:   "Filter by thread pool pattern",
					},
					&cli.StringSliceFlag{
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
					threadPoolPatterns := ctx.StringSlice("thread-pool-pattern")
					es := ctx.Context.Value("esClient").(*client.Elastic)
					columns := ctx.StringSlice("columns")
					request := &esapi.CatThreadPoolRequest{
						Pretty:             true,
						ThreadPoolPatterns: threadPoolPatterns,
						H:                  columns,
						Help:               esapi.BoolPtr(ctx.Bool("describe")),
					}
					return es.ExecRequest(ctx.Context, request)
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
					&cli.StringSliceFlag{
						Name:    "columns",
						Aliases: []string{"c"},
						Usage:   "Comma-separated list of columns to display",
					},
				},
				Action: func(ctx *cli.Context) error {
					es := ctx.Context.Value("esClient").(*client.Elastic)
					columns := ctx.StringSlice("columns")
					request := &esapi.CatPendingTasksRequest{
						V:      esapi.BoolPtr(true),
						Pretty: true,
						H:      columns,
						Help:   esapi.BoolPtr(ctx.Bool("describe")),
					}
					return es.ExecRequest(ctx.Context, request)
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
					&cli.StringSliceFlag{
						Name:    "columns",
						Aliases: []string{"c"},
						Usage:   "Comma-separated list of columns to display",
					},
				},
				Action: func(ctx *cli.Context) error {
					es := ctx.Context.Value("esClient").(*client.Elastic)
					columns := ctx.StringSlice("columns")
					request := &esapi.CatTasksRequest{
						V:      esapi.BoolPtr(true),
						Pretty: true,
						H:      columns,
						Help:   esapi.BoolPtr(ctx.Bool("describe")),
					}
					return es.ExecRequest(ctx.Context, request)
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
					&cli.StringSliceFlag{
						Name:    "columns",
						Aliases: []string{"c"},
						Usage:   "Comma-separated list of columns to display",
					},
				},
				Action: func(ctx *cli.Context) error {
					es := ctx.Context.Value("esClient").(*client.Elastic)
					columns := ctx.StringSlice("columns")
					request := &esapi.CatHealthRequest{
						V:      esapi.BoolPtr(true),
						Pretty: true,
						H:      columns,
						Help:   esapi.BoolPtr(ctx.Bool("describe")),
					}
					return es.ExecRequest(ctx.Context, request)
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
					&cli.StringSliceFlag{
						Name:    "columns",
						Aliases: []string{"c"},
						Usage:   "Comma-separated list of columns to display",
					},
				},
				Action: func(ctx *cli.Context) error {
					es := ctx.Context.Value("esClient").(*client.Elastic)
					columns := ctx.StringSlice("columns")
					request := &esapi.CatRepositoriesRequest{
						V:      esapi.BoolPtr(true),
						Pretty: true,
						H:      columns,
						Help:   esapi.BoolPtr(ctx.Bool("describe")),
					}
					return es.ExecRequest(ctx.Context, request)
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
					&cli.StringSliceFlag{
						Name:    "columns",
						Aliases: []string{"c"},
						Usage:   "Comma-separated list of columns to display",
					},
					&cli.StringSliceFlag{
						Name:    "repository",
						Aliases: []string{"r"},
						Usage:   "Comma-separated list of snapshot repositories used",
					},
				},
				Action: func(ctx *cli.Context) error {
					es := ctx.Context.Value("esClient").(*client.Elastic)
					columns := ctx.StringSlice("columns")
					request := &esapi.CatSnapshotsRequest{
						V:          esapi.BoolPtr(true),
						H:          columns,
						Help:       esapi.BoolPtr(ctx.Bool("describe")),
						Pretty:     true,
						Repository: ctx.StringSlice("repository"),
					}
					return es.ExecRequest(ctx.Context, request)
				},
			},
		},
	}
}
