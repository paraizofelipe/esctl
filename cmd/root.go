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
	app.Usage = "A command-line interface for managing and interacting with Elasticsearch clusters"
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
			Usage:      "Specify the path to the configuration file for esctl",
		},
		&cli.StringFlag{
			Name:    "cluster-name",
			Aliases: []string{"n"},
			Usage:   "Select a specific Elasticsearch cluster by name for executing commands",
		},
		&cli.StringFlag{
			Name:    "address",
			Aliases: []string{"a"},
			Usage:   "Set the address of the Elasticsearch cluster to connect to",
		},
		&cli.StringFlag{
			Name:    "username",
			Aliases: []string{"u"},
			Usage:   "Username for authentication with the Elasticsearch cluster",
		},
		&cli.StringFlag{
			Name:    "password",
			Aliases: []string{"p"},
			Usage:   "Password for authentication with the Elasticsearch cluster",
		},
	}
	app.Flags = flags
	app.Before = func(ctx *cli.Context) error {
		filePath := ctx.String("config-file")
		setup, err := config.ReadSetup(filePath)
		if err != nil {
			log.Fatalf("Error while loading configuration file: %s", err)
		}

		var cluster config.Cluster
		var config elasticsearch.Config

		clusterName := ctx.String("cluster-name")
		if clusterName != "" {
			cluster = setup.ClusterByName(clusterName)
		} else {
			cluster = setup.DefaultCluster()
		}

		config = elasticsearch.Config{
			Addresses: cluster.Address,
			Username:  cluster.Username,
			Password:  cluster.Password,
		}
		es := client.NewElastic(config)
		ctx.Context = context.WithValue(ctx.Context, "esClient", es)

		return nil
	}

	app.Commands = []*cli.Command{
		ApplyCommand(),
		SearchCommand(),
		GetCommand(),
		DescribeCommand(),
		CreateCommand(),
		ChangeCommand(),
		DeleteCommand(),
		TaskCommand(),
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
