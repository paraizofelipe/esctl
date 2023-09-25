package cmd

import (
	"github.com/urfave/cli/v2"
)

func NewGetCommand() *cli.Command {
	return &cli.Command{
		Name:  "get",
		Usage: "Get information about a created resources",
		Subcommands: []*cli.Command{
			NewGetIndexCommand(),
			NewGetDocCommand(),
			NewGetSourceCommand(),
			NewGetTaskCommand(),
			NewGetCountCommand(),
		},
	}
}
