package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/c-bata/go-prompt"
	"github.com/paraizofelipe/elastic_tools/internal/actions"
	"github.com/paraizofelipe/elastic_tools/internal/config"
	"github.com/urfave/cli/v2"
)

const APP_NAME = "elastic_tools"

type Config struct {
	Elastic []string `toml:"elastic"`
}

func NewRootCommand(setup *config.ConfigFile) *cli.App {
	app := cli.NewApp()
	app.Name = APP_NAME
	app.Usage = "Elasticsearch Tools CLI"
	app.Version = "1.0.0"

	app.EnableBashCompletion = true

	esNodes := setup.Elastic
	esClient, err := actions.CreateClient(esNodes)
	if err != nil {
		log.Fatalf("Error to create client: %s", err)
	}

	app.Commands = []*cli.Command{
		NewIndexCommand(esClient),
		NewSearchCommand(esClient),
		NewAliasCommand(esClient),
		NewCatCommand(esClient),
	}

	app.CommandNotFound = func(ctx *cli.Context, in string) {
		fmt.Printf("Ops, command %s unknown\n", in)
	}

	flags := []cli.Flag{
		&cli.StringFlag{
			Name:       "config-file",
			Aliases:    []string{"f"},
			Value:      fmt.Sprintf("%s/.config/elastic_tools/config.toml", os.Getenv("HOME")),
			HasBeenSet: true,
		},
		&cli.BoolFlag{
			Name:    "prompt-mode",
			Aliases: []string{"m"},
			Value:   false,
			Action: func(ctx *cli.Context, b bool) error {
				pt := NewPrompt(app)
				pt.Run()
				return nil
			},
		},
	}
	app.Flags = flags

	return app
}

func parseCLIArguments(input string) []string {
	parts := strings.Fields(input)
	var args []string
	var quotedArg string

	for _, part := range parts {
		if strings.HasPrefix(part, "'") {
			part = strings.TrimPrefix(part, "'")
			quotedArg = part
		} else if quotedArg != "" {
			quotedArg += " " + part
			if strings.HasSuffix(part, "'") {
				quotedArg = strings.TrimSuffix(quotedArg, "'")
				args = append(args, quotedArg)
				quotedArg = ""
			}
		} else {
			args = append(args, part)
		}
	}

	return append([]string{APP_NAME}, args...)
}

func NewPrompt(app *cli.App) (pt *prompt.Prompt) {
	options := LoadOptions(app.Commands)

	pt = prompt.New(
		func(in string) {
			if in == "exit" {
				fmt.Println("Exiting REPL...")
				os.Exit(0)
			}

			cliArguments := parseCLIArguments(in)
			if len(cliArguments) > 0 {
				err := app.Run(cliArguments)
				if err != nil {
					fmt.Println(err)
				}
			}
		},
		func(d prompt.Document) (suggest []prompt.Suggest) {
			if d.TextBeforeCursor() == "" {
				return []prompt.Suggest{}
			}
			args := strings.Split(d.TextBeforeCursor(), " ")
			w := d.GetWordBeforeCursor()

			if len(args) < 2 {
				for _, cmd := range app.Commands {
					suggest = append(suggest, prompt.Suggest{Text: cmd.Name, Description: cmd.Usage})
				}
				return prompt.FilterHasPrefix(suggest, w, true)
			}

			for _, arg := range args {
				if strings.HasPrefix(arg, "-") || strings.HasPrefix(arg, "'") {
					continue
				}
				if aux, ok := options[arg]; ok {
					suggest = aux
				}
			}

			return prompt.FilterHasPrefix(suggest, w, true)
		},
		prompt.OptionPrefix(">>> "),
		prompt.OptionTitle("Interactive CLI"),
	)
	return
}

func Execute() {

	config, err := config.Load()
	if err != nil {
		log.Fatalf("Error while loading configuration file: %s", err)
	}

	app := NewRootCommand(config)
	err = app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}
