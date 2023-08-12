package actions

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
)

type Indexer interface {
	CreateIndex(string, bool)
	DeleteIndex([]string, bool)
	GetIndex([]string, bool)
}

type IndexAction struct {
	client *elasticsearch.Client
}

func NewIndexAction(esClient *elasticsearch.Client) Indexer {
	return &IndexAction{
		client: esClient,
	}
}

func (i *IndexAction) CreateIndex(indexName string, pretty bool) {

	var (
		err         error
		res         *esapi.Response
		reqSettings []func(*esapi.IndicesCreateRequest) = []func(*esapi.IndicesCreateRequest){
			i.client.Indices.Create.WithTimeout(5 * time.Second),
			i.client.Indices.Create.WithContext(context.Background()),
		}
	)

	if pretty {
		reqSettings = append(reqSettings, i.client.Indices.Create.WithPretty())
	}

	if res, err = i.client.Indices.Create(
		indexName,
		reqSettings...,
	); err != nil {
		log.Fatalf("Error to create index: %s", err)
	}

	defer res.Body.Close()

	if res.IsError() {
		log.Fatalf("[REQ-ERROR]: %s", res.String())
	}

	fmt.Println(res.String())
}

func (i *IndexAction) DeleteIndex(indexNames []string, pretty bool) {
	var (
		err         error
		res         *esapi.Response
		reqSettings []func(*esapi.IndicesDeleteRequest) = []func(*esapi.IndicesDeleteRequest){
			i.client.Indices.Delete.WithContext(context.Background()),
			i.client.Indices.Delete.WithTimeout(5 * time.Second),
		}
	)

	if pretty {
		reqSettings = append(reqSettings, i.client.Indices.Delete.WithPretty())
	}

	if res, err = i.client.Indices.Delete(
		indexNames,
		reqSettings...,
	); err != nil {
		log.Fatalf("Error to delete index: %s", err)
	}

	defer res.Body.Close()

	if res.IsError() {
		log.Fatalf("[REQ-ERROR]: %s", res.String())
	}

	fmt.Println(res.String())
}

func (i *IndexAction) GetIndex(indexNames []string, pretty bool) {
	var (
		err         error
		res         *esapi.Response
		reqSettings []func(*esapi.IndicesGetRequest) = []func(*esapi.IndicesGetRequest){
			i.client.Indices.Get.WithContext(context.Background()),
		}
	)

	if pretty {
		reqSettings = append(reqSettings, i.client.Indices.Get.WithPretty())
	}

	if res, err = i.client.Indices.Get(
		indexNames,
		reqSettings...,
	); err != nil {
		log.Fatalf("Error to read index: %s", err)
	}

	defer res.Body.Close()

	if res.IsError() {
		log.Fatalf("[REQ-ERROR]: %s", res.String())
	}

	fmt.Println(res.String())
}
