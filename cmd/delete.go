package cmd

import (
	"github.com/urfave/cli/v2"
)

func NewDeleteCommand() *cli.Command {
	return &cli.Command{
		Name:  "delete",
		Usage: "Delete resources in Elasticsearch",
		Subcommands: []*cli.Command{
			NewDeleteIndexCommand(),
			NewDeleteIndexAliasCommand(),
		},
	}
}
