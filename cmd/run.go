package cmd

import (
	"fmt"
	"os"
	"plugin"

	"github.com/urfave/cli/v2"
)

func RunCommand() *cli.Command {
	return &cli.Command{
		Name:  "run",
		Usage: "run a external plugin",
		Action: func(c *cli.Context) error {
			pluginDir := fmt.Sprintf("%s/.config/esctl/plugins", os.Getenv("HOME"))
			pluginName := fmt.Sprintf("esctl-%s", c.Args().First())
			path := fmt.Sprintf("%s/%s/plugin.so", pluginDir, pluginName)

			if _, err := os.Stat(path); os.IsNotExist(err) {
				return fmt.Errorf("plugin %s not found", path)
			}

			p, err := plugin.Open(path)
			if err != nil {
				return err
			}

			symbol, err := p.Lookup("RunPlugin")
			if err != nil {
				return err
			}

			runFunc, ok := symbol.(func())
			if !ok {
				return fmt.Errorf("plugin does not conform to 'RunPlugin' signature")
			}

			runFunc()
			return nil
		},
	}
}
