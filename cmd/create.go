package cmd

import (
	"github.com/urfave/cli/v2"
)

func NewCreateCommand() *cli.Command {
	return &cli.Command{
		Name:  "create",
		Usage: "Create resources in Elasticsearch",
		Subcommands: []*cli.Command{
			NewCreateIndexCommand(),
		},
	}
}
