package operation

import (
	"fmt"

	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/elastic/go-elasticsearch/v8/esutil"
	"github.com/elastic/go-elasticsearch/v8/typedapi/core/reindex"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/paraizofelipe/esctl/internal/client"
	"github.com/paraizofelipe/esctl/internal/output"
	"github.com/urfave/cli/v2"
	"gopkg.in/yaml.v3"
)

type Options struct {
	WaitForCompletion bool
}

type BodyAlias struct {
	Actions []types.IndicesAction `json:"actions"`
}

type BodyReindex reindex.Request

type StageAlias struct {
	Kind    string    `json:"kind"`
	Body    BodyAlias `json:"body"`
	Options Options   `json:"options"`
}

type StageReindex struct {
	Kind    string      `json:"kind"`
	Body    BodyReindex `json:"body"`
	Options Options     `json:"options"`
}

type StageProcessor interface {
	Process(*cli.Context, client.ElasticClient) error
}

type StageFile struct {
	Stages []StageProcessor `json:"stages"`
}

func (s *StageAlias) Process(ctx *cli.Context, es client.ElasticClient) error {
	request := &esapi.IndicesUpdateAliasesRequest{
		Body: esutil.NewJSONReader(s.Body),
	}
	jsonBytes, err := es.ExecRequest(ctx.Context, request)
	if err != nil {
		return err
	}
	output.PrintPrettyJSON(jsonBytes)
	return nil
}

func (s *StageReindex) Process(ctx *cli.Context, es client.ElasticClient) error {
	request := &esapi.ReindexRequest{
		Body: esutil.NewJSONReader(s.Body),
	}
	jsonBytes, err := es.ExecRequest(ctx.Context, request)
	if err != nil {
		return err
	}
	output.PrintPrettyJSON(jsonBytes)
	return nil
}

func (s *StageFile) UnmarshalYAML(node *yaml.Node) error {
	var rawFile struct {
		Stages []yaml.Node `json:"stages"`
	}

	if err := node.Decode(&rawFile); err != nil {
		return err
	}

	for _, rawStage := range rawFile.Stages {
		var kind struct {
			Kind string `yaml:"kind"`
		}
		if err := rawStage.Decode(&kind); err != nil {
			return err
		}

		var stage StageProcessor
		switch kind.Kind {
		case "alias":
			stage = &StageAlias{}
		case "reindex":
			stage = &StageReindex{}
		default:
			return fmt.Errorf("unknown stage kind: %s", kind.Kind)
		}

		if err := rawStage.Decode(stage); err != nil {
			return err
		}

		s.Stages = append(s.Stages, stage)
	}
	return nil
}
