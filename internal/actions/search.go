package actions

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/urfave/cli/v2"
)

type Searcher interface {
	SearchDoc(*cli.Context) error
}

type SearchAction struct {
	client *elasticsearch.Client
}

func NewSearchAction(esClient *elasticsearch.Client) Searcher {
	return &SearchAction{
		client: esClient,
	}
}

func (s *SearchAction) SearchDoc(ctx *cli.Context) (err error) {
	var (
		indexName   string = ctx.Args().Get(0)
		query       string = ctx.String("query")
		pretty      bool   = ctx.Bool("pretty")
		res         *esapi.Response
		reqSettings []func(*esapi.SearchRequest) = []func(*esapi.SearchRequest){
			s.client.Search.WithIndex(indexName),
			s.client.Search.WithContext(context.Background()),
			s.client.Search.WithTimeout(5 * time.Second),
		}
	)

	if pretty {
		reqSettings = append(reqSettings, s.client.Search.WithPretty())
	}

	queryBody := strings.NewReader(query)
	reqSettings = append(reqSettings, s.client.Search.WithBody(queryBody))

	if res, err = s.client.Search(reqSettings...); err != nil {
		return fmt.Errorf("Error when querying the index: %s", err)
	}

	defer res.Body.Close()

	if res.IsError() {
		return fmt.Errorf("[REQ-ERROR]: %s", res.String())
	}

	fmt.Println(res.String())

	return
}
