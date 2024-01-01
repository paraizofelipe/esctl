package cmd

import (
	"fmt"
	"os"

	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/elastic/go-elasticsearch/v8/esutil"
	"github.com/paraizofelipe/esctl/internal/client"
	"github.com/paraizofelipe/esctl/internal/operation"
	"github.com/paraizofelipe/esctl/internal/out"
	"github.com/urfave/cli/v2"
	"gopkg.in/yaml.v3"
)

func StagesCommand() *cli.Command {
	return &cli.Command{
		Name:  "run-stages",
		Usage: "Execute a series of stages defined in a YAML file",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "file",
				Aliases: []string{"f"},
				Usage:   "Path to the YAML file containing the configurations to be applied",
			},
		},
		Action: func(ctx *cli.Context) error {
			var stageFile operation.StageFile
			es := ctx.Context.Value("esClient").(client.ElasticClient)

			yamlFile, err := os.ReadFile(ctx.String("file"))
			if err != nil {
				return err
			}

			err = yaml.Unmarshal(yamlFile, &stageFile)
			if err != nil {
				return err
			}

			for _, stage := range stageFile.Stages {
				switch stage.Kind {
				case "alias":
					var body operation.BodyAlias
					err = stage.Body.Decode(&body)
					if err != nil {
						return err
					}
					request := &esapi.IndicesUpdateAliasesRequest{
						Body: esutil.NewJSONReader(body),
					}
					jsonBytes, err := es.ExecRequest(ctx.Context, request)
					if err != nil {
						return err
					}
					out.PrintPrettyJSON(jsonBytes)
				case "reindex":
					var body operation.BodyReindex
					err = stage.Body.Decode(&body)
					if err != nil {
						return err
					}
					request := &esapi.ReindexRequest{
						Pretty: true,
						Body:   esutil.NewJSONReader(body),
					}
					jsonBytes, err := es.ExecRequest(ctx.Context, request)
					if err != nil {
						return err
					}
					fmt.Println(string(jsonBytes))
				}
			}

			return nil
		},
	}
}
