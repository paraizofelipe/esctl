package cmd

import (
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/paraizofelipe/elastic_tools/internal/actions"
	"github.com/urfave/cli/v2"
)

func NewAliasCommand(esClient *elasticsearch.Client) *cli.Command {

	appFlags := []cli.Flag{
		&cli.StringFlag{
			Name:    "config",
			Usage:   "JSON string containing the configuration settings",
			Aliases: []string{"c"},
		},
		&cli.StringFlag{
			Name:    "aliases-config-file",
			Usage:   "Path to the file containing the alias settings",
			Aliases: []string{"f"},
		},
		&cli.BoolFlag{
			Name:    "pretty",
			Aliases: []string{"p"},
			Usage:   "Format the response as pretty-printed JSON",
		},
	}

	aliasManager := actions.NewAliasAction(esClient)

	return &cli.Command{
		Name:   "alias",
		Usage:  "Manage Elasticsearch aliases",
		Flags:  appFlags,
		Action: aliasManager.UpdateAlias,
	}
}
