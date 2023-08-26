package actions

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/paraizofelipe/elastic_tools/internal/file"
	"github.com/urfave/cli/v2"
)

type Indexer interface {
	CreateIndex(*cli.Context) error
	DeleteIndex(*cli.Context) error
	GetIndex(*cli.Context) error
	AddDoc(*cli.Context) error
	ListDoc(*cli.Context) error
	ForceMerge(*cli.Context) error
	ExecBulkOperation(*cli.Context) error
}

type IndexAction struct {
	client *elasticsearch.Client
}

func NewIndexAction(esClient *elasticsearch.Client) Indexer {
	return &IndexAction{
		client: esClient,
	}
}

func (i *IndexAction) CreateIndex(ctx *cli.Context) (err error) {

	var (
		indexName    string = ctx.Args().Get(0)
		pretty       bool   = ctx.Bool("pretty")
		settingsFile string = ctx.String("settings-file")
		res          *esapi.Response
		reqSettings  []func(*esapi.IndicesCreateRequest) = []func(*esapi.IndicesCreateRequest){
			i.client.Indices.Create.WithTimeout(5 * time.Second),
			i.client.Indices.Create.WithContext(context.Background()),
		}
		settingsIndex string
	)

	if pretty {
		reqSettings = append(reqSettings, i.client.Indices.Create.WithPretty())
	}

	if settingsFile != "" {
		settingsIndex, err = file.ReadJSONFile(settingsFile)
		if err != nil {
			return fmt.Errorf("Erro to read settings file: %s", err)
		}
		settingsBody := strings.NewReader(settingsIndex)
		reqSettings = append(reqSettings, i.client.Indices.Create.WithBody(settingsBody))
	}

	if res, err = i.client.Indices.Create(
		indexName,
		reqSettings...,
	); err != nil {
		return fmt.Errorf("Error to create index: %s", err)
	}

	defer res.Body.Close()

	if res.IsError() {
		return fmt.Errorf("[REQ-ERROR]: %s", res.String())
	}

	fmt.Println(res.String())

	return
}

func (i *IndexAction) DeleteIndex(ctx *cli.Context) (err error) {
	var (
		args        string = ctx.Args().Get(0)
		pretty      bool   = ctx.Bool("pretty")
		res         *esapi.Response
		reqSettings []func(*esapi.IndicesDeleteRequest) = []func(*esapi.IndicesDeleteRequest){
			i.client.Indices.Delete.WithContext(context.Background()),
			i.client.Indices.Delete.WithTimeout(5 * time.Second),
		}
	)

	indexNames := strings.Split(args, ",")

	if pretty {
		reqSettings = append(reqSettings, i.client.Indices.Delete.WithPretty())
	}

	if res, err = i.client.Indices.Delete(
		indexNames,
		reqSettings...,
	); err != nil {
		return fmt.Errorf("Error to delete index: %s", err)
	}

	defer res.Body.Close()

	if res.IsError() {
		return fmt.Errorf("[REQ-ERROR]: %s", res.String())
	}

	fmt.Println(res.String())

	return
}

func (i *IndexAction) ForceMerge(ctx *cli.Context) (err error) {
	var (
		args        string = ctx.Args().Get(0)
		indexNames         = strings.Split(args, ",")
		pretty      bool   = ctx.Bool("pretty")
		res         *esapi.Response
		reqSettings = []func(*esapi.IndicesForcemergeRequest){
			i.client.Indices.Forcemerge.WithIndex(indexNames...),
			i.client.Indices.Forcemerge.WithContext(ctx.Context),
		}
	)

	if pretty {
		reqSettings = append(reqSettings, i.client.Indices.Forcemerge.WithPretty())
	}

	if res, err = i.client.Indices.Forcemerge(
		reqSettings...,
	); err != nil {
		return fmt.Errorf("Error to force merge in index: %s", err)
	}

	defer res.Body.Close()

	if res.IsError() {
		return fmt.Errorf("[REQ-ERROR]: %s", res.String())
	}

	fmt.Println(res.String())

	return
}

func (i *IndexAction) GetIndex(ctx *cli.Context) (err error) {
	var (
		arg         string = ctx.Args().Get(0)
		pretty      bool   = ctx.Bool("pretty")
		res         *esapi.Response
		reqSettings []func(*esapi.IndicesGetRequest) = []func(*esapi.IndicesGetRequest){
			i.client.Indices.Get.WithContext(context.Background()),
		}
	)

	indexNames := strings.Split(arg, ",")

	if pretty {
		reqSettings = append(reqSettings, i.client.Indices.Get.WithPretty())
	}

	if res, err = i.client.Indices.Get(
		indexNames,
		reqSettings...,
	); err != nil {
		return fmt.Errorf("Error to read index: %s", err)
	}

	defer res.Body.Close()

	if res.IsError() {
		return fmt.Errorf("[REQ-ERROR]: %s", res.String())
	}

	fmt.Println(res.String())

	return
}

func (i *IndexAction) AddDoc(ctx *cli.Context) (err error) {
	var (
		indexName   string = ctx.Args().Get(0)
		pretty      bool   = ctx.Bool("pretty")
		doc         string = ctx.String("document")
		docFile     string = ctx.String("document-file")
		res         *esapi.Response
		reqSettings []func(*esapi.IndexRequest) = []func(*esapi.IndexRequest){}
	)

	if docFile != "" {
		if doc, err = file.ReadJSONFile(docFile); err != nil {
			return fmt.Errorf("Error to load document path: %s", docFile)
		}
	}

	if pretty {
		reqSettings = append(reqSettings, i.client.Index.WithPretty())
	}

	if res, err = i.client.Index(indexName, strings.NewReader(doc), reqSettings...); err != nil {
		return fmt.Errorf("Error when querying the index: %s", err)
	}

	defer res.Body.Close()

	if res.IsError() {
		return fmt.Errorf("[REQ-ERROR]: %s", res.String())
	}

	fmt.Println(res.String())

	return
}

func (i *IndexAction) ExecBulkOperation(ctx *cli.Context) (err error) {

	var (
		pretty         bool   = ctx.Bool("pretty")
		bulkFile       string = ctx.String("bulk-file")
		res            *esapi.Response
		reqSettings    []func(*esapi.BulkRequest) = []func(*esapi.BulkRequest){}
		bulkOperations string
	)

	if pretty {
		reqSettings = append(reqSettings, i.client.Bulk.WithPretty())
	}

	bulkOperations, err = file.ReadJSONFile(bulkFile)
	if err != nil {
		return fmt.Errorf("Error to laod bulk file: %s", err)
	}

	if res, err = i.client.Bulk(
		strings.NewReader(bulkOperations),
		reqSettings...,
	); err != nil {
		return fmt.Errorf("")
	}

	defer res.Body.Close()

	if res.IsError() {
		return fmt.Errorf("[REQ-ERROR]: %s", res.String())
	}

	fmt.Println(res.String())

	return
}

func (i *IndexAction) ListDoc(ctx *cli.Context) (err error) {

	var (
		indexName   string = ctx.Args().Get(0)
		size        int    = ctx.Int("size")
		pretty      bool   = ctx.Bool("pretty")
		res         *esapi.Response
		reqSettings []func(*esapi.SearchRequest) = []func(*esapi.SearchRequest){
			i.client.Search.WithContext(context.Background()),
			i.client.Search.WithIndex(indexName),
			i.client.Search.WithTimeout(5 * time.Second),
		}
	)

	pageSize := i.client.Search.WithSize(10)
	if size > 0 {
		pageSize = i.client.Search.WithSize(size)
	}
	reqSettings = append(reqSettings, pageSize)

	if pretty {
		reqSettings = append(reqSettings, i.client.Search.WithPretty())
	}

	if res, err = i.client.Search(reqSettings...); err != nil {
		return fmt.Errorf("Error when querying the index: %s", err)
	}

	defer res.Body.Close()

	if res.IsError() {
		return fmt.Errorf("[REQ-ERROR]: %s", res.String())
	}

	fmt.Println(res.String())

	return
}
