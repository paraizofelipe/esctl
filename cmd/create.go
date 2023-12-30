package cmd

import (
	"github.com/urfave/cli/v2"
)

func CreateCommand() *cli.Command {
	return &cli.Command{
		Name:  "create",
		Usage: "Initialize various resources in Elasticsearch, such as indexes, with customizable configurations",
		Subcommands: []*cli.Command{
			CreateIndexCommand(),
		},
	}
}
