package cmd

import (
	"github.com/urfave/cli/v2"
)

func NewDescribeCommand() *cli.Command {
	return &cli.Command{
		Name:    "describe",
		Aliases: []string{"desc"},
		Usage:   "Describe information about a created resources",
		Subcommands: []*cli.Command{
			NewGetIndexCommand(),
			NewGetDocCommand(),
			NewGetSourceCommand(),
			NewGetTaskCommand(),
			NewGetCountCommand(),
		},
	}
}
