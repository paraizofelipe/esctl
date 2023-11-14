package cmd

import (
	"fmt"
	"log"

	"github.com/paraizofelipe/esctl/internal/config"
	"github.com/paraizofelipe/esctl/internal/table"
	"github.com/urfave/cli/v2"
)

func GetConfigCommand() *cli.Command {
	return &cli.Command{
		Name:  "config",
		Usage: "Show CLI configurations",
		Action: func(ctx *cli.Context) error {
			filePath := ctx.String("config-file")
			setup, err := config.ReadSetup(filePath)
			if err != nil {
				log.Fatalf("Error while loading configuration file: %s", err)
			}
			outTable := table.NewTable("CURRENT", "CLUSTER", "USERNAME", "ADDRESS")
			for _, cluster := range setup.Cluster {
				if cluster.Default {
					outTable.AddRow("->", cluster.Name, cluster.Username, cluster.Address)
					continue
				}
				outTable.AddRow("", cluster.Name, cluster.Username, cluster.Address)
			}
			fmt.Println(outTable)

			return nil
		},
	}
}

func ChangeConfigCommand() *cli.Command {
	return &cli.Command{
		Name:  "config",
		Usage: "Change the default cluster",
		Action: func(ctx *cli.Context) error {
			filePath := ctx.String("config-file")
			setup, err := config.ReadSetup(filePath)
			if err != nil {
				log.Fatalf("Error while loading configuration file: %s", err)
			}

			clusterName := ctx.Args().First()
			cluster := setup.ClusterByName(clusterName)
			if cluster.Name == "" {
				return fmt.Errorf("Cluster %s not found", clusterName)
			}

			for i, h := range setup.Cluster {
				if h.Default {
					setup.Cluster[i].Default = false
				}
				if h.Name == clusterName {
					setup.Cluster[i].Default = true
				}
			}

			if err := config.WriteSetup(setup, filePath); err != nil {
				return err
			}

			return nil
		},
	}
}
