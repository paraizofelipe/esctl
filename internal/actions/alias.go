package actions

import (
	"context"
	"fmt"
	"strings"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/paraizofelipe/elastic_tools/internal/file"
	"github.com/urfave/cli/v2"
)

type AliasesManager interface {
	UpdateAlias(ctx *cli.Context) error
}

type AliasesAction struct {
	client *elasticsearch.Client
}

func NewAliasAction(esClient *elasticsearch.Client) AliasesManager {
	return &AliasesAction{
		client: esClient,
	}
}

func (a *AliasesAction) UpdateAlias(ctx *cli.Context) (err error) {
	var (
		pretty      bool   = ctx.Bool("pretty")
		actions     string = ctx.String("config")
		actionsFile string = ctx.String("aliases-config-file")
		reqSettings        = []func(*esapi.IndicesUpdateAliasesRequest){
			a.client.Indices.UpdateAliases.WithContext(context.Background()),
		}
		res *esapi.Response
	)

	if actionsFile != "" {
		if actions, err = file.ReadJSONFile(actionsFile); err != nil {
			return fmt.Errorf("Error to load document path: %s", actionsFile)
		}
	}

	if pretty {
		reqSettings = append(reqSettings, a.client.Indices.UpdateAliases.WithPretty())
	}

	updateRequest := strings.NewReader(actions)

	if res, err = a.client.Indices.UpdateAliases(updateRequest, reqSettings...); err != nil {
		return fmt.Errorf("Erro ao executar actions de aliases: %s", err)
	}

	defer res.Body.Close()

	if res.IsError() {
		return fmt.Errorf("[REQ-ERROR]: %s", res.String())
	}

	fmt.Println(res.String())

	return
}
