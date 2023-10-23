package cmd

import (
	"strings"

	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/paraizofelipe/esctl/internal/client"
	"github.com/paraizofelipe/esctl/internal/file"
	"github.com/urfave/cli/v2"
)

func SearchCommand() *cli.Command {

	appFlags := []cli.Flag{
		&cli.StringFlag{
			Name:     "query",
			Aliases:  []string{"q"},
			Usage:    "The search query string",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "file",
			Aliases:  []string{"f"},
			Usage:    "The file with the search query",
			Required: false,
		},
	}

	return &cli.Command{
		Name:  "search",
		Usage: "Search documents in an index",
		Flags: appFlags,
		Action: func(ctx *cli.Context) error {
			es := ctx.Context.Value("esClient").(*client.Elastic)
			filePath := ctx.String("file")
			if filePath != "" {
				jsonQuery, err := file.ReadJSONFile(filePath)
				if err != nil {
					return err
				}
				if jsonQuery != "" {
					ctx.Set("query", jsonQuery)
				}
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
