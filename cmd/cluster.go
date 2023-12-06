package cmd

import (
	"strings"

	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/elastic/go-elasticsearch/v8/esutil"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/paraizofelipe/esctl/internal/client"
	"github.com/urfave/cli/v2"
)

func ClusterRerouteCommand() *cli.Command {
	return &cli.Command{
		Name:  "reroute",
		Usage: "Manually reroute shards in the cluster to optimize distribution or repair issues",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "body",
				Aliases:  []string{"b"},
				Usage:    "JSON body specifying the reroute actions to be taken",
				Required: true,
			},
		},
		Action: func(ctx *cli.Context) error {
			es := ctx.Context.Value("esClient").(*client.Elastic)
			body := strings.NewReader(ctx.String("body"))
			request := &esapi.ClusterRerouteRequest{
				Pretty: true,
				Metric: nil,
				Body:   body,
			}

			return es.ExecRequest(ctx.Context, request)
		},
	}
}

func ApplyClusterReroute(ctx *cli.Context, commands types.Command) error {
	es := ctx.Context.Value("esClient").(*client.Elastic)
	body := esutil.NewJSONReader(commands)
	request := &esapi.ClusterRerouteRequest{
		Pretty: true,
		Metric: nil,
		Body:   body,
	}
	return es.ExecRequest(ctx.Context, request)
}

func ClusterCommand() *cli.Command {
	return &cli.Command{
		Name:  "cluster",
		Usage: "Commands for managing and interacting with the Elasticsearch cluster",
		Subcommands: []*cli.Command{
			ClusterRerouteCommand(),
		},
	}
}
