package cmd

import (
	"github.com/urfave/cli/v2"
)

func CreateCommand() *cli.Command {
	return &cli.Command{
		Name:  "create",
		Usage: "Create resources in Elasticsearch",
		Subcommands: []*cli.Command{
			CreateIndexCommand(),
		},
	}
}
