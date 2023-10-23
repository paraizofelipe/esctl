package cmd

import (
	"github.com/urfave/cli/v2"
)

func ChangeCommand() *cli.Command {
	return &cli.Command{
		Name:  "change",
		Usage: "Change resources in Elasticsearch",
		Subcommands: []*cli.Command{
			ChangeIndexAliasCommand(),
			ChangeIndexMappingCommand(),
		},
	}
}
