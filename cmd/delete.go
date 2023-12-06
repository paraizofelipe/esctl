package cmd

import (
	"github.com/urfave/cli/v2"
)

func DeleteCommand() *cli.Command {
	return &cli.Command{
		Name:  "delete",
		Usage: "Remove various resources from Elasticsearch, including indices and security settings",
		Subcommands: []*cli.Command{
			DeleteIndexCommand(),
			DeleteSecurityCommand(),
		},
	}
}
