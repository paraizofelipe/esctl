package operation

import (
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/elastic/go-elasticsearch/v8/esutil"
	"github.com/elastic/go-elasticsearch/v8/typedapi/core/reindex"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/paraizofelipe/esctl/internal/client"
	"github.com/paraizofelipe/esctl/internal/out"
	"github.com/urfave/cli/v2"
	"gopkg.in/yaml.v3"
)

type Options struct {
	WaitForCompletion bool
}

type BodyAlias struct {
	Actions []types.IndicesAction `json:"actions"`
}

type Body interface {
	BodyAlias | BodyReindex
}

type BodyReindex reindex.Request

type Stage struct {
	Kind    string    `json:"kind"`
	Body    yaml.Node `json:"body"`
	Options Options   `json:"options"`
}

type StageFile struct {
	Stages []Stage `json:"stages"`
}

func EexecStages[T Body](ctx *cli.Context, es client.ElasticClient, body T) error {
	request := &esapi.IndicesUpdateAliasesRequest{
		Body: esutil.NewJSONReader(body),
	}
	jsonBytes, err := es.ExecRequest(ctx.Context, request)
	if err != nil {
		return err
	}
	out.PrintPrettyJSON(jsonBytes)
	return nil
}
