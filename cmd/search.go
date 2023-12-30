package cmd

import (
	"errors"
	"strings"

	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/paraizofelipe/esctl/internal/client"
	"github.com/paraizofelipe/esctl/internal/file"
	"github.com/urfave/cli/v2"
)

func SearchCommand(textEditor file.Editor) *cli.Command {
	appFlags := []cli.Flag{
		&cli.StringFlag{
			Name:     "query",
			Aliases:  []string{"q"},
			Usage:    "Enter the search query string to be executed against the index",
			Required: false,
		},
		&cli.BoolFlag{
			Name:     "editor",
			Aliases:  []string{"e"},
			Usage:    "",
			Required: false,
		},
		&cli.StringFlag{
			Name:     "file",
			Aliases:  []string{"f"},
			Usage:    "Specify a file path containing the search query in JSON format",
			Required: false,
		},
	}

	return &cli.Command{
		Name:  "search",
		Usage: "Execute a search query against specified Elasticsearch indices",
		Flags: appFlags,
		Action: func(ctx *cli.Context) error {
			es := ctx.Context.Value("esClient").(client.ElasticClient)
			if ctx.Bool("editor") {
				content, err := textEditor.ExecEditor(ctx.String("file"))
				if err != nil {
					return err
				}
				ctx.Set("query", content)
			}

			if filePath := ctx.String("file"); filePath != "" && !ctx.Bool("editor") {
				jsonQuery, err := file.ReadJSONFile(filePath)
				if err != nil {
					return err
				}
				if jsonQuery != "" {
					ctx.Set("query", jsonQuery)
				}
			}

			if ctx.String("query") == "" {
				return errors.New("No query specified")
			}

			request := &esapi.SearchRequest{
				Pretty: true,
				Index:  ctx.Args().Slice(),
				Body:   strings.NewReader(ctx.String("query")),
			}

			return es.ExecRequest(ctx.Context, request)
		},
	}
}
