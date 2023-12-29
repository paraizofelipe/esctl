package cmd

import (
	"github.com/paraizofelipe/esctl/internal/client"
	"github.com/urfave/cli/v2"
)

func DeleteCommand(es client.ElasticClient) *cli.Command {
	return &cli.Command{
		Name:  "delete",
		Usage: "Remove various resources from Elasticsearch, including indices and security settings",
		Subcommands: []*cli.Command{
			DeleteIndexCommand(es),
			DeleteSecurityCommand(es),
		},
	}
}
