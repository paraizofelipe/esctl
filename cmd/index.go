package cmd

import (
	"strings"

	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/elastic/go-elasticsearch/v8/esutil"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/paraizofelipe/esctl/internal/client"
	"github.com/urfave/cli/v2"
)

type AliasBody struct {
	Bodies []types.IndicesAction `json:"bodies"`
}

type Mapping struct {
	M types.Property `json:"mappings"`
}

func DescribeIndexDocCommand() *cli.Command {
	return &cli.Command{
		Name:  "doc",
		Usage: "Retrieve and display detailed information about a document using its unique identifier",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "id",
				Aliases:  []string{"i"},
				Usage:    "Specify the unique identifier of the document to be described",
				Required: true,
			},
			&cli.StringSliceFlag{
				Name:     "fields",
				Aliases:  []string{"f"},
				Usage:    "List of specific stored fields to include in the response, separated by comma",
				Required: false,
			},
		},
		Action: func(ctx *cli.Context) error {
			es := ctx.Context.Value("esClient").(client.ElasticClient)
			docRequest := &esapi.GetRequest{
				Pretty:         true,
				DocumentID:     ctx.String("id"),
				Index:          ctx.Args().Get(0),
				SourceIncludes: ctx.StringSlice("fields"),
			}
			return es.ExecRequest(ctx.Context, docRequest)
		},
	}
}

func DescribeIndexAliasCommand() *cli.Command {
	return &cli.Command{
		Name:  "alias",
		Usage: "Display details about an existing index alias, including its patterns and configurations",
		Action: func(ctx *cli.Context) error {
			es := ctx.Context.Value("esClient").(client.ElasticClient)
			indexPatterns := ctx.Args().Slice()
			request := &esapi.IndicesGetAliasRequest{
				Pretty: true,
				Index:  indexPatterns,
			}
			return es.ExecRequest(ctx.Context, request)
		},
	}
}

func ApplyIndexAlias(ctx *cli.Context, bodies AliasBody) error {
	es := ctx.Context.Value("esClient").(client.ElasticClient)
	body := esutil.NewJSONReader(bodies)
	request := &esapi.IndicesUpdateAliasesRequest{
		Pretty: true,
		Body:   body,
	}
	return es.ExecRequest(ctx.Context, request)
}

func ChangeIndexAliasCommand() *cli.Command {
	return &cli.Command{
		Name:  "alias",
		Usage: "Modify the configuration or pattern of an existing index alias",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "body",
				Usage:    "Provide the new configuration for the index alias in JSON format",
				Aliases:  []string{"b"},
				Required: true,
			},
		},
		Action: func(ctx *cli.Context) error {
			es := ctx.Context.Value("esClient").(client.ElasticClient)
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

func DeleteIndexAliasCommand() *cli.Command {
	return &cli.Command{
		Name:  "alias",
		Usage: "Remove an existing index alias and its associated configurations",
		Flags: []cli.Flag{
			&cli.StringSliceFlag{
				Name:     "name",
				Aliases:  []string{"n"},
				Required: true,
			},
		},
		Action: func(ctx *cli.Context) error {
			es := ctx.Context.Value("esClient").(client.ElasticClient)
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

func DescribeIndexSettingsCommand() *cli.Command {
	return &cli.Command{
		Name:  "settings",
		Usage: "Retrieve and display the current settings of a specified index",
		Action: func(ctx *cli.Context) error {
			es := ctx.Context.Value("esClient").(client.ElasticClient)
			indexPatterns := ctx.Args().Slice()
			request := &esapi.IndicesGetSettingsRequest{
				Pretty: true,
				Index:  indexPatterns,
			}
			return es.ExecRequest(ctx.Context, request)
		},
	}
}

func DescribeIndexMappingCommand() *cli.Command {
	return &cli.Command{
		Name:  "mapping",
		Usage: "Show the mapping details, including field types and index configurations, of a specific index",
		Action: func(ctx *cli.Context) error {
			es := ctx.Context.Value("esClient").(client.ElasticClient)
			indexPatterns := ctx.Args().Slice()
			request := &esapi.IndicesGetMappingRequest{
				Pretty: true,
				Index:  indexPatterns,
			}
			return es.ExecRequest(ctx.Context, request)
		},
	}
}

func ApplyIndexMapping(ctx *cli.Context, indexName []string, bodies types.Property) error {
	es := ctx.Context.Value("esClient").(client.ElasticClient)
	body := esutil.NewJSONReader(bodies)
	request := &esapi.IndicesPutMappingRequest{
		Index:  indexName,
		Pretty: true,
		Body:   body,
	}
	return es.ExecRequest(ctx.Context, request)
}

func ChangeIndexMappingCommand() *cli.Command {
	return &cli.Command{
		Name:  "mapping",
		Usage: "Update or modify the mapping of an existing index with new configurations",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "body",
				Aliases:  []string{"b"},
				Required: true,
			},
		},
		Action: func(ctx *cli.Context) error {
			es := ctx.Context.Value("esClient").(client.ElasticClient)
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

func DescribeIndexStatsCommand() *cli.Command {
	return &cli.Command{
		Name:  "stats",
		Usage: "Access and display statistical data and metrics for a specified index",
		Action: func(ctx *cli.Context) error {
			es := ctx.Context.Value("esClient").(client.ElasticClient)
			indexPatterns := ctx.Args().Slice()
			request := &esapi.IndicesStatsRequest{
				Pretty: true,
				Index:  indexPatterns,
			}
			return es.ExecRequest(ctx.Context, request)
		},
	}
}

func DescribeIndexCommand() *cli.Command {
	return &cli.Command{
		Name:  "index",
		Usage: "Fetch and display comprehensive information about an index, including its documents, aliases, and settings",
		Action: func(ctx *cli.Context) error {
			es := ctx.Context.Value("esClient").(client.ElasticClient)
			indexRequest := &esapi.IndicesGetRequest{
				Pretty: true,
				Index:  []string{ctx.Args().Get(0)},
			}
			return es.ExecRequest(ctx.Context, indexRequest)
		},
		Subcommands: []*cli.Command{
			DescribeIndexDocCommand(),
			DescribeIndexAliasCommand(),
			DescribeIndexStatsCommand(),
			DescribeIndexMappingCommand(),
			DescribeIndexSettingsCommand(),
		},
	}
}

func CreateIndexDocCommand() *cli.Command {
	return &cli.Command{
		Name:  "doc",
		Usage: "Insert a new document into a specified index, with optional custom ID",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "id",
				Aliases: []string{"i"},
			},
			&cli.StringFlag{
				Name:    "body",
				Aliases: []string{"b"},
			},
		},
		Action: func(ctx *cli.Context) error {
			es := ctx.Context.Value("esClient").(client.ElasticClient)
			body := strings.NewReader(ctx.String("body"))
			request := &esapi.IndexRequest{
				Pretty:     true,
				Index:      ctx.Args().Get(0),
				DocumentID: ctx.String("id"),
				Body:       body,
			}
			return es.ExecRequest(ctx.Context, request)
		},
	}
}

func CreateIndexCommand() *cli.Command {
	return &cli.Command{
		Name:  "index",
		Usage: "Initialize a new index in Elasticsearch with customizable settings and mappings",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "body",
				Aliases: []string{"b"},
			},
		},
		Action: func(ctx *cli.Context) error {
			es := ctx.Context.Value("esClient").(client.ElasticClient)
			body := strings.NewReader(ctx.String("body"))
			request := &esapi.IndicesCreateRequest{
				Pretty: true,
				Index:  ctx.Args().Get(0),
				Body:   body,
			}
			return es.ExecRequest(ctx.Context, request)
		},
		Subcommands: []*cli.Command{
			CreateIndexDocCommand(),
		},
	}
}

func DeleteIndexCommand() *cli.Command {
	return &cli.Command{
		Name:  "index",
		Usage: "Permanently remove an existing index and all its associated data from Elasticsearch",
		Action: func(ctx *cli.Context) error {
			es := ctx.Context.Value("esClient").(client.ElasticClient)
			indexPatterns := ctx.Args().Slice()
			request := &esapi.IndicesDeleteRequest{
				Pretty: true,
				Index:  indexPatterns,
			}
			return es.ExecRequest(ctx.Context, request)
		},
		Subcommands: []*cli.Command{
			DeleteIndexAliasCommand(),
		},
	}
}
