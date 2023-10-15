package cmd

import (
	"strings"

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

func NewChangeIndexAliasCommand() *cli.Command {
	return &cli.Command{
		Name:  "alias",
		Usage: "Change index alias",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "body",
				Aliases:  []string{"b"},
				Required: true,
			},
		},
		Action: func(ctx *cli.Context) error {
			es := ctx.Context.Value("esClient").(*client.Elastic)
			indexPatterns := ctx.Args().Slice()
			body := strings.NewReader(ctx.String("body"))
			request := &esapi.IndicesPutAliasRequest{
				Pretty: true,
				Index:  indexPatterns,
				Body:   body,
			}
			return es.ExecRequest(ctx.Context, request)
		},
	}
}

func NewDeleteIndexAliasCommand() *cli.Command {
	return &cli.Command{
		Name:  "alias",
		Usage: "Delete index alias",
		Action: func(ctx *cli.Context) error {
			es := ctx.Context.Value("esClient").(*client.Elastic)
			indexPatterns := ctx.Args().Slice()
			names := ctx.StringSlice("name")
			request := &esapi.IndicesDeleteAliasRequest{
				Pretty: true,
				Index:  indexPatterns,
				Name:   names,
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

func NewChangeIndexMappingCommand() *cli.Command {
	return &cli.Command{
		Name:  "mapping",
		Usage: "Change index mapping",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "body",
				Aliases:  []string{"b"},
				Required: true,
			},
		},
		Action: func(ctx *cli.Context) error {
			es := ctx.Context.Value("esClient").(*client.Elastic)
			indexPatterns := ctx.Args().Slice()
			body := strings.NewReader(ctx.String("body"))
			request := &esapi.IndicesPutMappingRequest{
				Pretty: true,
				Index:  indexPatterns,
				Body:   body,
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

func NewCreateIndexCommand() *cli.Command {
	return &cli.Command{
		Name:  "index",
		Usage: "Create index",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "body",
				Aliases: []string{"b"},
			},
		},
		Action: func(ctx *cli.Context) error {
			es := ctx.Context.Value("esClient").(*client.Elastic)
			body := strings.NewReader(ctx.String("body"))
			request := &esapi.IndicesCreateRequest{
				Pretty: true,
				Index:  ctx.Args().Get(0),
				Body:   body,
			}
			return es.ExecRequest(ctx.Context, request)
		},
	}
}

func NewDeleteIndexCommand() *cli.Command {
	return &cli.Command{
		Name:  "index",
		Usage: "Delete index",
		Action: func(ctx *cli.Context) error {
			es := ctx.Context.Value("esClient").(*client.Elastic)
			indexPatterns := ctx.Args().Slice()
			request := &esapi.IndicesDeleteRequest{
				Pretty: true,
				Index:  indexPatterns,
			}
			return es.ExecRequest(ctx.Context, request)
		},
	}
}
