package cmd

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/paraizofelipe/esctl/internal/client"
	"github.com/paraizofelipe/esctl/internal/config"

	"github.com/urfave/cli/v2"
)

const APP_NAME = "esctl"

func NewRootCommand() *cli.App {
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
		&cli.StringFlag{
			Name:    "host-name",
			Aliases: []string{"n"},
		},
		&cli.StringFlag{
			Name:    "address",
			Aliases: []string{"a"},
		},
		&cli.StringFlag{
			Name:    "username",
			Aliases: []string{"u"},
		},
		&cli.StringFlag{
			Name:    "password",
			Aliases: []string{"p"},
		},
	}
	app.Flags = flags
	app.Before = func(ctx *cli.Context) error {
		filePath := ctx.String("config-file")
		setup, err := config.Load(filePath)
		if err != nil {
			log.Fatalf("Error while loading configuration file: %s", err)
		}

		var host config.Host
		var config elasticsearch.Config

		hostName := ctx.String("host-name")
		if hostName != "" {
			host = setup.HostByName(hostName)
		} else {
			host = setup.DefaultHost()
		}

		config = elasticsearch.Config{
			Addresses: host.Address,
			Username:  host.Username,
			Password:  host.Password,
		}
		es := client.NewElastic(config)
		ctx.Context = context.WithValue(ctx.Context, "esClient", es)

		return nil
	}

	app.Commands = []*cli.Command{
		NewSearchCommand(),
		NewGetCommand(),
		NewDescribeCommand(),
		NewCreateCommand(),
		NewChangeCommand(),
		NewDeleteCommand(),
		NewTaskCommand(),
	}

	return app
}

func Execute() {
	app := NewRootCommand()
	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
	}
}
