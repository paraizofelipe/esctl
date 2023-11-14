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
		Usage: "Reroute shards",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "body",
				Aliases:  []string{"b"},
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
		Usage: "Cluster commands",
		Subcommands: []*cli.Command{
			ClusterRerouteCommand(),
		},
	}
}