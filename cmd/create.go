package cmd

import (
	"github.com/paraizofelipe/esctl/internal/client"
	"github.com/urfave/cli/v2"
)

func CreateCommand(es client.ElasticClient) *cli.Command {
	return &cli.Command{
		Name:  "create",
		Usage: "Initialize various resources in Elasticsearch, such as indexes, with customizable configurations",
		Subcommands: []*cli.Command{
			CreateIndexCommand(es),
		},
	}
}
