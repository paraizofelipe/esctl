package cmd

import (
	"github.com/urfave/cli/v2"
)

func ChangeCommand() *cli.Command {
	return &cli.Command{
		Name:  "change",
		Usage: "Modify various Elasticsearch resources, including index aliases, mappings, security settings, and CLI configurations",
		Subcommands: []*cli.Command{
			ChangeIndexAliasCommand(),
			ChangeIndexMappingCommand(),
			ChangeSecurityCommand(),
			ChangeConfigCommand(),
		},
	}
}
