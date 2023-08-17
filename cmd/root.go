package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/paraizofelipe/elastic_tools/internal/actions"
	"github.com/paraizofelipe/elastic_tools/internal/config"
	"github.com/urfave/cli/v2"
	"github.com/urfave/cli/v2/altsrc"
)

type Config struct {
	Elastic []string `toml:"elastic"`
}

func NewRootCommand(setup *config.ConfigFile) *cli.App {
	app := cli.NewApp()
	app.Name = "elastic_tools"
	app.Usage = "Elasticsearch Tools CLI"
	app.Version = "1.0.0"

	flags := []cli.Flag{
		altsrc.NewStringSliceFlag(&cli.StringSliceFlag{
			Name:    "elastic",
			Aliases: []string{"e"},
			Value:   cli.NewStringSlice("http://localhost:9200"),
		}),
		&cli.StringFlag{
			Name:       "config-file",
			Aliases:    []string{"f"},
			Value:      fmt.Sprintf("%s/.config/elastic_tools/config.toml", os.Getenv("HOME")),
			HasBeenSet: true,
		},
	}

	app.Before = altsrc.InitInputSourceWithContext(flags, altsrc.NewTomlSourceFromFlagFunc("config-file"))

	esNodes := setup.Elastic
	esClient, err := actions.CreateClient(esNodes)
	if err != nil {
		log.Fatalf("Error to create client: %s", err)
	}

	app.Flags = flags
	app.Commands = []*cli.Command{
		NewIndexCommand(esClient),
	}

	return app
}

func Execute() {

	config, err := config.Load()
	if err != nil {
		log.Fatalf("Error while loading configuration file: %s", err)
	}

	app := NewRootCommand(config)
	err = app.Run(os.Args)
	if err != nil {
		log.Fatalf("Exec error: %s", err)
	}
}
