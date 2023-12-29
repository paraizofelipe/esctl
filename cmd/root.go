package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/paraizofelipe/esctl/internal/client"
	"github.com/paraizofelipe/esctl/internal/config"
	"github.com/paraizofelipe/esctl/internal/file"
	"github.com/urfave/cli/v2"
)

const APP_NAME = "esctl"

func NewRootCommand() *cli.App {
	var cluster config.Cluster
	var esConfig elasticsearch.Config

	es, err := client.NewElastic(elasticsearch.Config{})
	if err != nil {
		log.Fatal(err)
	}

	app := cli.NewApp()
	app.Name = APP_NAME
	app.Usage = "A command-line interface for managing and interacting with Elasticsearch clusters"
	app.Version = "1.0.0"

	app.Suggest = true
	app.EnableBashCompletion = true

	app.CommandNotFound = func(ctx *cli.Context, in string) {
		fmt.Printf("Ops, command %s unknown\n", in)
	}

	app.Flags = []cli.Flag{
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
	app.Before = func(ctx *cli.Context) error {
		filePath := ctx.String("config-file")
		setup, err := config.ReadSetup(filePath)
		if err != nil {
			log.Fatalf("Error while loading configuration file: %s", err)
		}

		clusterName := ctx.String("cluster-name")
		if clusterName != "" {
			cluster = setup.ClusterByName(clusterName)
		} else {
			cluster = setup.DefaultCluster()
		}

		esConfig = elasticsearch.Config{
			Addresses: cluster.Address,
			Username:  cluster.Username,
			Password:  cluster.Password,
		}
		es, err = client.NewElastic(esConfig)
		if err != nil {
			log.Fatal(err)
		}
		return nil
	}

	app.Commands = []*cli.Command{
		ApplyCommand(es),
		SearchCommand(es, file.NewTextEditor()),
		GetCommand(es),
		DescribeCommand(es),
		CreateCommand(es),
		ChangeCommand(es),
		DeleteCommand(es),
		TaskCommand(es),
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
