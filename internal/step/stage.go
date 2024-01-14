package step

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

type StepAlias struct {
	Kind    string    `json:"kind"`
	Body    BodyAlias `json:"body"`
	Options Options   `json:"options"`
}

type StepReindex struct {
	Kind    string      `json:"kind"`
	Body    BodyReindex `json:"body"`
	Options Options     `json:"options"`
}

type StepProcessor interface {
	Process(*cli.Context, client.ElasticClient) error
}

type StepFile struct {
	Steps []StepProcessor `json:"steps"`
}

func (s *StepAlias) Process(ctx *cli.Context, es client.ElasticClient) error {
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

func (s *StepReindex) Process(ctx *cli.Context, es client.ElasticClient) error {
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

func (s *StepFile) UnmarshalYAML(node *yaml.Node) error {
	var rawFile struct {
		Steps []yaml.Node `json:"steps"`
	}

	if err := node.Decode(&rawFile); err != nil {
		return err
	}

	for _, rawStep := range rawFile.Steps {
		var kind struct {
			Kind string `yaml:"kind"`
		}
		if err := rawStep.Decode(&kind); err != nil {
			return err
		}

		var step StepProcessor
		switch kind.Kind {
		case "alias":
			step = &StepAlias{}
		case "reindex":
			step = &StepReindex{}
		default:
			return fmt.Errorf("unknown step kind: %s", kind.Kind)
		}

		if err := rawStep.Decode(step); err != nil {
			return err
		}

		s.Steps = append(s.Steps, step)
	}
	return nil
}
