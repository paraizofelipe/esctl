package cmd

import (
	"fmt"

	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/paraizofelipe/esctl/internal/client"
	"github.com/urfave/cli/v2"
)

func GetCommand() *cli.Command {

	return &cli.Command{
		Name:  "get",
		Usage: "Retrieve detailed information about various Elasticsearch cluster resources",
		Subcommands: []*cli.Command{
			{
				Name:  "indices",
				Usage: "List and provide detailed information about all indices in Elasticsearc",
				Flags: []cli.Flag{
					&cli.StringSliceFlag{
						Name:    "columns",
						Aliases: []string{"c"},
						Usage:   "Specify the columns to include in the output, separated by commas",
					},
					&cli.BoolFlag{
						Name:    "describe",
						Value:   false,
						Aliases: []string{"d"},
						Usage:   "Include detailed descriptions for each column in the output",
					},
				},
				Action: func(ctx *cli.Context) error {
					es := ctx.Context.Value("esClient").(client.ElasticClient)
					columns := ctx.StringSlice("columns")
					indices := ctx.Args().Slice()
					if len(indices) == 0 {
						indices = []string{"*"}
					}
					request := &esapi.CatIndicesRequest{
						Index:  indices,
						V:      esapi.BoolPtr(true),
						Pretty: true,
						H:      columns,
						Help:   esapi.BoolPtr(ctx.Bool("describe")),
					}
					textBytes, err := es.ExecRequest(ctx.Context, request)
					fmt.Println(string(textBytes))
					return err
				},
			},
			{
				Name:  "aliases",
				Usage: "Display a list of all aliases configured in Elasticsearch",
				Flags: []cli.Flag{
					&cli.StringSliceFlag{
						Name:    "columns",
						Aliases: []string{"c"},
						Usage:   "Specify the columns to include in the output, separated by commas",
					},
					&cli.BoolFlag{
						Name:    "describe",
						Value:   false,
						Aliases: []string{"d"},
						Usage:   "Include detailed descriptions for each column in the output",
					},
				},
				Action: func(ctx *cli.Context) error {
					es := ctx.Context.Value("esClient").(client.ElasticClient)
					columns := ctx.StringSlice("columns")
					request := &esapi.CatAliasesRequest{
						V:      esapi.BoolPtr(true),
						Pretty: true,
						H:      columns,
						Help:   esapi.BoolPtr(ctx.Bool("describe")),
					}
					textBytes, err := es.ExecRequest(ctx.Context, request)
					fmt.Println(string(textBytes))
					return err
				},
			},
			{
				Name:  "nodes",
				Usage: "List all nodes in Elasticsearch cluster",
				Flags: []cli.Flag{
					&cli.StringSliceFlag{
						Name:    "columns",
						Aliases: []string{"c"},
						Usage:   "Specify the columns to include in the output, separated by commas",
					},
					&cli.BoolFlag{
						Name:    "describe",
						Value:   false,
						Aliases: []string{"d"},
						Usage:   "Include detailed descriptions for each column in the output",
					},
				},
				Action: func(ctx *cli.Context) error {
					es := ctx.Context.Value("esClient").(client.ElasticClient)
					columns := ctx.StringSlice("columns")
					request := &esapi.CatNodesRequest{
						V:      esapi.BoolPtr(true),
						Pretty: true,
						H:      columns,
						Help:   esapi.BoolPtr(ctx.Bool("describe")),
					}
					textBytes, err := es.ExecRequest(ctx.Context, request)
					fmt.Println(string(textBytes))
					return err
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
						Usage:   "Specify the columns to include in the output, separated by commas",
					},
					&cli.BoolFlag{
						Name:    "describe",
						Value:   false,
						Aliases: []string{"d"},
						Usage:   "Include detailed descriptions for each column in the output",
					},
				},
				Action: func(ctx *cli.Context) error {
					es := ctx.Context.Value("esClient").(client.ElasticClient)
					columns := ctx.StringSlice("columns")

					request := &esapi.CatShardsRequest{
						V:      esapi.BoolPtr(true),
						Pretty: true,
						Index:  ctx.StringSlice("index"),
						H:      columns,
						Help:   esapi.BoolPtr(ctx.Bool("describe")),
					}
					textBytes, err := es.ExecRequest(ctx.Context, request)
					fmt.Println(string(textBytes))
					return err
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
						Usage:   "Specify the columns to include in the output, separated by commas",
					},
					&cli.BoolFlag{
						Name:    "describe",
						Value:   false,
						Aliases: []string{"d"},
						Usage:   "Include detailed descriptions for each column in the output",
					},
				},
				Action: func(ctx *cli.Context) error {
					es := ctx.Context.Value("esClient").(client.ElasticClient)
					threadPoolPatterns := ctx.StringSlice("thread-pool-pattern")
					columns := ctx.StringSlice("columns")
					request := &esapi.CatThreadPoolRequest{
						Pretty:             true,
						ThreadPoolPatterns: threadPoolPatterns,
						H:                  columns,
						Help:               esapi.BoolPtr(ctx.Bool("describe")),
					}
					textBytes, err := es.ExecRequest(ctx.Context, request)
					fmt.Println(string(textBytes))
					return err
				},
			},
			{
				Name:  "pending-tasks",
				Usage: "List pending tasks in Elasticsearch",
				Flags: []cli.Flag{
					&cli.StringSliceFlag{
						Name:    "columns",
						Aliases: []string{"c"},
						Usage:   "Specify the columns to include in the output, separated by commas",
					},
					&cli.BoolFlag{
						Name:    "describe",
						Value:   false,
						Aliases: []string{"d"},
						Usage:   "Include detailed descriptions for each column in the output",
					},
				},
				Action: func(ctx *cli.Context) error {
					es := ctx.Context.Value("esClient").(client.ElasticClient)
					columns := ctx.StringSlice("columns")
					request := &esapi.CatPendingTasksRequest{
						V:      esapi.BoolPtr(true),
						Pretty: true,
						H:      columns,
						Help:   esapi.BoolPtr(ctx.Bool("describe")),
					}
					textBytes, err := es.ExecRequest(ctx.Context, request)
					fmt.Println(string(textBytes))
					return err
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
						Usage:   "Include detailed descriptions for each column in the output",
					},
					&cli.StringSliceFlag{
						Name:    "columns",
						Aliases: []string{"c"},
						Usage:   "Specify the columns to include in the output, separated by commas",
					},
				},
				Action: func(ctx *cli.Context) error {
					es := ctx.Context.Value("esClient").(client.ElasticClient)
					columns := ctx.StringSlice("columns")
					request := &esapi.CatTasksRequest{
						V:      esapi.BoolPtr(true),
						Pretty: true,
						H:      columns,
						Help:   esapi.BoolPtr(ctx.Bool("describe")),
					}
					textBytes, err := es.ExecRequest(ctx.Context, request)
					fmt.Println(string(textBytes))
					return err
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
						Usage:   "Include detailed descriptions for each column in the output",
					},
					&cli.StringSliceFlag{
						Name:    "columns",
						Aliases: []string{"c"},
						Usage:   "Specify the columns to include in the output, separated by commas",
					},
				},
				Action: func(ctx *cli.Context) error {
					es := ctx.Context.Value("esClient").(client.ElasticClient)
					columns := ctx.StringSlice("columns")
					request := &esapi.CatHealthRequest{
						V:      esapi.BoolPtr(true),
						Pretty: true,
						H:      columns,
						Help:   esapi.BoolPtr(ctx.Bool("describe")),
					}
					textBytes, err := es.ExecRequest(ctx.Context, request)
					fmt.Println(string(textBytes))
					return err
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
						Usage:   "Include detailed descriptions for each column in the output",
					},
					&cli.StringSliceFlag{
						Name:    "columns",
						Aliases: []string{"c"},
						Usage:   "Specify the columns to include in the output, separated by commas",
					},
				},
				Action: func(ctx *cli.Context) error {
					es := ctx.Context.Value("esClient").(client.ElasticClient)
					columns := ctx.StringSlice("columns")
					request := &esapi.CatRepositoriesRequest{
						V:      esapi.BoolPtr(true),
						Pretty: true,
						H:      columns,
						Help:   esapi.BoolPtr(ctx.Bool("describe")),
					}
					textBytes, err := es.ExecRequest(ctx.Context, request)
					fmt.Println(string(textBytes))
					return err
				},
			},
			{
				Name:  "allocation",
				Usage: "Returns information about disk usage and shard allocation",
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name:    "describe",
						Value:   false,
						Aliases: []string{"d"},
						Usage:   "Include detailed descriptions for each column in the output",
					},
					&cli.StringSliceFlag{
						Name:    "columns",
						Aliases: []string{"c"},
						Usage:   "Specify the columns to include in the output, separated by commas",
					},
					&cli.StringSliceFlag{
						Name:    "node",
						Aliases: []string{"n"},
						Usage:   "Comma-separated list of node IDs or names to limit the returned information",
					},
				},
				Action: func(ctx *cli.Context) error {
					es := ctx.Context.Value("esClient").(client.ElasticClient)
					columns := ctx.StringSlice("columns")
					request := &esapi.CatAllocationRequest{
						V:      esapi.BoolPtr(true),
						H:      columns,
						Help:   esapi.BoolPtr(ctx.Bool("describe")),
						Pretty: true,
						NodeID: ctx.StringSlice("node"),
					}
					textBytes, err := es.ExecRequest(ctx.Context, request)
					fmt.Println(string(textBytes))
					return err
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
						Usage:   "Include detailed descriptions for each column in the output",
					},
					&cli.StringSliceFlag{
						Name:    "columns",
						Aliases: []string{"c"},
						Usage:   "Specify the columns to include in the output, separated by commas",
					},
					&cli.StringSliceFlag{
						Name:    "repository",
						Aliases: []string{"r"},
						Usage:   "Comma-separated list of snapshot repositories used",
					},
				},
				Action: func(ctx *cli.Context) error {
					es := ctx.Context.Value("esClient").(client.ElasticClient)
					columns := ctx.StringSlice("columns")
					request := &esapi.CatSnapshotsRequest{
						V:          esapi.BoolPtr(true),
						H:          columns,
						Help:       esapi.BoolPtr(ctx.Bool("describe")),
						Pretty:     true,
						Repository: ctx.StringSlice("repository"),
					}
					textBytes, err := es.ExecRequest(ctx.Context, request)
					fmt.Println(string(textBytes))
					return err
				},
			},
			GetConfigCommand(),
		},
	}
}
