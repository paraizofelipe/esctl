package cmd

import (
	"github.com/paraizofelipe/esctl/internal/client"
	"github.com/urfave/cli/v2"
)

func DescribeCommand(es client.ElasticClient) *cli.Command {
	return &cli.Command{
		Name:    "describe",
		Aliases: []string{"desc"},
		Usage:   "Provide detailed descriptions and status of various Elasticsearch resource",
		Subcommands: []*cli.Command{
			DescribeIndexCommand(es),
			DescribeTaskCommand(es),
			DescribeCountCommand(),
			DescribeSecurityCommand(es),
		},
	}
}
