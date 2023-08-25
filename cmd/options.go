package cmd

import (
	"fmt"

	"github.com/c-bata/go-prompt"
	"github.com/urfave/cli/v2"
)

type Option map[string][]prompt.Suggest

func MergeOptions(dst Option, src Option) {
	for key, values := range src {
		if existingValues, ok := dst[key]; ok {
			dst[key] = append(existingValues, values...)
		} else {
			dst[key] = values
		}
	}
}

func LoadOptions(commands []*cli.Command) Option {
	var option Option = Option{}

	for _, cmd := range commands {
		if len(cmd.Flags) > 0 {
			for _, flag := range cmd.Flags {
				option[cmd.Name] = append(
					option[cmd.Name],
					prompt.Suggest{Text: fmt.Sprintf("--%s", flag.Names()[0])},
				)
			}
		}

		if len(cmd.Subcommands) > 0 {
			for _, subCmd := range cmd.Subcommands {
				option[cmd.Name] = append(
					option[cmd.Name],
					prompt.Suggest{Text: subCmd.Name, Description: subCmd.Usage},
				)
			}
			subOptions := LoadOptions(cmd.Subcommands)
			MergeOptions(option, subOptions)
		}

	}

	return option
}
