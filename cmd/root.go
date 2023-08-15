package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/pelletier/go-toml"
	"github.com/urfave/cli/v2"
)

type Config struct {
	Elastic []string `toml:"elastic"`
}

func loadConfig() (*Config, error) {
	configFile, err := toml.LoadFile("~/.config/elastic_tools/config.toml")
	if err != nil {
		return nil, err
	}

	var config Config
	if err := configFile.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}

func NewRootCommand(esClient *elasticsearch.Client) *cli.App {
	app := cli.NewApp()
	app.Name = "elastic_tools"
	app.Usage = "Elasticsearch Tools CLI"
	app.Version = "1.0.0"

	flags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "elastic",
			Aliases: []string{"e"},
			Value:   cli.NewStringSlice("http://localhost:9200"),
		},
		&cli.StringFlag{
			Name:       "config-file",
			Aliases:    []string{"f"},
			Value:      fmt.Sprintf("%s/.config/elastic_tools/config.toml", os.Getenv("HOME")),
			HasBeenSet: true,
		},
	}

	app.Action = func(c *cli.Context) error {
		fmt.Println(c.StringSlice("e"))
		return nil
	}

	app.Flags = flags
	app.Commands = []*cli.Command{
		NewIndexCommand(esClient),
		NewDocCommand(esClient),
	}

	return app
}

func Execute() {

	var (
		err      error
		esClient *elasticsearch.Client
	)

	cfg := elasticsearch.Config{
		Addresses: []string{"http://192.168.68.119:9200"},
	}

	if esClient, err = elasticsearch.NewClient(cfg); err != nil {
		log.Fatalf("Exec error: %s", err)
	}

	app := NewRootCommand(esClient)
	err = app.Run(os.Args)
	if err != nil {
		log.Fatalf("Exec error: %s", err)
	}
}
