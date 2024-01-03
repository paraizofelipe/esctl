package cmd

import (
	"os"

	"github.com/paraizofelipe/esctl/internal/client"
	"github.com/paraizofelipe/esctl/internal/operation"
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
				err = stage.Process(ctx, es)
				if err != nil {
					return err
				}
			}

			return nil
		},
	}
}
