package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/paraizofelipe/esctl/internal/actions"
	"github.com/paraizofelipe/esctl/internal/config"
	"github.com/urfave/cli/v2"
)

const APP_NAME = "esctl"

func NewRootCommand(setup *config.ConfigFile) *cli.App {
	app := cli.NewApp()
	app.Name = APP_NAME
	app.Usage = "Elasticsearch Tools CLI"
	app.Version = "1.0.0"

	app.Suggest = true
	app.EnableBashCompletion = true

	app.CommandNotFound = func(ctx *cli.Context, in string) {
		fmt.Printf("Ops, command %s unknown\n", in)
	}

	flags := []cli.Flag{
		&cli.StringFlag{
			Name:       "config-file",
			Aliases:    []string{"f"},
			Value:      fmt.Sprintf("%s/.config/esctl/config.toml", os.Getenv("HOME")),
			HasBeenSet: true,
		},
	}
	app.Flags = flags

	esClient, err := actions.CreateClient(setup)
	if err != nil {
		log.Fatalf("Error to create client: %s", err)
	}

	app.Commands = []*cli.Command{
		NewIndexCommand(esClient),
		NewSearchCommand(esClient),
		NewAliasCommand(esClient),
		NewCatCommand(esClient),
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
		fmt.Println(err)
	}
}
