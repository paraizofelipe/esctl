package cmd

import (
	"github.com/urfave/cli/v2"
)

func DescribeCommand() *cli.Command {
	return &cli.Command{
		Name:    "describe",
		Aliases: []string{"desc"},
		Usage:   "Describe information about a created resources",
		Subcommands: []*cli.Command{
			DescribeIndexCommand(),
			DescribeTaskCommand(),
			DescribeCountCommand(),
			DescribeSecurityCommand(),
		},
	}
}
