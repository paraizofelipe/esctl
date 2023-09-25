package cmd

import (
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/paraizofelipe/esctl/internal/client"
	"github.com/urfave/cli/v2"
)

func NewGetIndexAliasCommand() *cli.Command {
	return &cli.Command{
		Name:  "alias",
		Usage: "Get information about a created index alias",
		Action: func(ctx *cli.Context) error {
			indexPatterns := ctx.Args().Slice()
			es := ctx.Context.Value("esClient").(*client.Elastic)
			request := &esapi.IndicesGetAliasRequest{
				Pretty: true,
				Index:  indexPatterns,
			}
			return es.ExecRequest(ctx.Context, request)
		},
	}
}

func NewGetIndexSettingsCommand() *cli.Command {
	return &cli.Command{
		Name:  "settings",
		Usage: "Get settings for a created index",
		Action: func(ctx *cli.Context) error {
			indexPatterns := ctx.Args().Slice()
			es := ctx.Context.Value("esClient").(*client.Elastic)
			request := &esapi.IndicesGetSettingsRequest{
				Pretty: true,
				Index:  indexPatterns,
			}
			return es.ExecRequest(ctx.Context, request)
		},
	}
}

func NewGetIndexMappingCommand() *cli.Command {
	return &cli.Command{
		Name:  "mapping",
		Usage: "Get mapping for a created index",
		Action: func(ctx *cli.Context) error {
			indexPatterns := ctx.Args().Slice()
			es := ctx.Context.Value("esClient").(*client.Elastic)
			request := &esapi.IndicesGetMappingRequest{
				Pretty: true,
				Index:  indexPatterns,
			}
			return es.ExecRequest(ctx.Context, request)
		},
	}
}

func NewGetIndexStatsCommand() *cli.Command {
	return &cli.Command{
		Name:  "stats",
		Usage: "Get stats for a created index",
		Action: func(ctx *cli.Context) error {
			indexPatterns := ctx.Args().Slice()
			es := ctx.Context.Value("esClient").(*client.Elastic)
			request := &esapi.IndicesStatsRequest{
				Pretty: true,
				Index:  indexPatterns,
			}
			return es.ExecRequest(ctx.Context, request)
		},
	}
}

func NewGetIndexCommand() *cli.Command {
	return &cli.Command{
		Name:  "index",
		Usage: "Get information about a created index",
		Action: func(ctx *cli.Context) error {
			es := ctx.Context.Value("esClient").(*client.Elastic)
			indexRequest := &esapi.IndicesGetRequest{
				Pretty: true,
				Index:  []string{ctx.Args().Get(0)},
			}
			return es.ExecRequest(ctx.Context, indexRequest)
		},
		Subcommands: []*cli.Command{
			NewGetIndexAliasCommand(),
			NewGetIndexStatsCommand(),
			NewGetIndexMappingCommand(),
			NewGetIndexSettingsCommand(),
		},
	}
}
