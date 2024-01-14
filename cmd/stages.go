package cmd

import (
	"os"

	"github.com/paraizofelipe/esctl/internal/client"
	"github.com/paraizofelipe/esctl/internal/step"
	"github.com/urfave/cli/v2"
	"gopkg.in/yaml.v3"
)

func StepsCommand() *cli.Command {
	return &cli.Command{
		Name:  "run-steps",
		Usage: "Execute a series of steps defined in a YAML file",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "file",
				Aliases: []string{"f"},
				Usage:   "Path to the YAML file containing the configurations to be applied",
			},
		},
		Action: func(ctx *cli.Context) error {
			var stepFile step.StepFile
			es := ctx.Context.Value("esClient").(client.ElasticClient)

			yamlFile, err := os.ReadFile(ctx.String("file"))
			if err != nil {
				return err
			}

			err = yaml.Unmarshal(yamlFile, &stepFile)
			if err != nil {
				return err
			}

			for _, step := range stepFile.Steps {
				err = step.Process(ctx, es)
				if err != nil {
					return err
				}
			}

			return nil
		},
	}
}
