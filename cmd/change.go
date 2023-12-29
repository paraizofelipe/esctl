package cmd

import (
	"github.com/paraizofelipe/esctl/internal/client"
	"github.com/urfave/cli/v2"
)

func ChangeCommand(es client.ElasticClient) *cli.Command {
	return &cli.Command{
		Name:  "change",
		Usage: "Modify various Elasticsearch resources, including index aliases, mappings, security settings, and CLI configurations",
		Subcommands: []*cli.Command{
			ChangeIndexAliasCommand(es),
			ChangeIndexMappingCommand(es),
			ChangeSecurityCommand(es),
			ChangeConfigCommand(),
		},
	}
}
