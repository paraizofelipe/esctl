package cmd

import (
	"github.com/urfave/cli/v2"
)

func DescribeCommand() *cli.Command {
	return &cli.Command{
		Name:    "describe",
		Aliases: []string{"desc"},
		Usage:   "Provide detailed descriptions and status of various Elasticsearch resource",
		Subcommands: []*cli.Command{
			DescribeIndexCommand(),
			DescribeTaskCommand(),
			DescribeCountCommand(),
			DescribeSecurityCommand(),
		},
	}
}
