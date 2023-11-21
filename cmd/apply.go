package cmd

import (
	"encoding/json"
	"errors"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/paraizofelipe/esctl/internal/file"
	"github.com/urfave/cli/v2"
)

type ApplyFile struct {
	Kind  string          `json:"kind"`
	Index json.RawMessage `json:"index,omitempty"`
	Body  json.RawMessage `json:"body"`
}

func ApplyCommand() *cli.Command {
	return &cli.Command{
		Name:  "apply",
		Usage: "",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "file",
				Aliases: []string{"f"},
			},
		},
		Action: func(ctx *cli.Context) error {
			var applyFile ApplyFile

			filePath := ctx.String("file")
			contentFile, err := file.ReadJSONFile(filePath)
			if err != nil {
				return err
			}

			if err = json.Unmarshal([]byte(contentFile), &applyFile); err != nil {
				return err
			}

			switch applyFile.Kind {
			case "SecurityUser":
				var securityUser []SecurityUser
				if err := json.Unmarshal(applyFile.Body, &securityUser); err != nil {
					return err
				}
				err = ApplySecurityUsers(ctx, securityUser)
				if err != nil {
					return err
				}
			case "ClusterReroute":
				var rerouteCommand types.Command
				if err := json.Unmarshal(applyFile.Body, &rerouteCommand); err != nil {
					return err
				}
				err = ApplyClusterReroute(ctx, rerouteCommand)
				if err != nil {
					return err
				}
			case "IndexAlias":
				var bodyActions AliasAction
				if err := json.Unmarshal(applyFile.Body, &bodyActions); err != nil {
					return err
				}
				err = ApplyIndexAlias(ctx, bodyActions)
				if err != nil {
					return err
				}
			case "IndexMapping":
				var (
					index          []string
					bodyProperties types.Property
				)
				if err := json.Unmarshal(applyFile.Index, &index); err != nil {
					return err
				}
				if err := json.Unmarshal(applyFile.Body, &bodyProperties); err != nil {
					return err
				}
				err = ApplyIndexMapping(ctx, index, bodyProperties)
				if err != nil {
					return err
				}
			default:
				return errors.New("Unknown kind")
			}

			return nil
		},
	}
}
