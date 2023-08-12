package actions

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
)

func CreateIndex(esClient *elasticsearch.Client, indexName string, pretty bool) {

	var (
		err         error
		res         *esapi.Response
		reqSettings []func(*esapi.IndicesCreateRequest) = []func(*esapi.IndicesCreateRequest){
			esClient.Indices.Create.WithTimeout(5 * time.Second),
			esClient.Indices.Create.WithContext(context.Background()),
		}
	)

	if pretty {
		reqSettings = append(reqSettings, esClient.Indices.Create.WithPretty())
	}

	if res, err = esClient.Indices.Create(
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

func DeleteIndex(esClient *elasticsearch.Client, indexNames []string, pretty bool) {
	var (
		err         error
		res         *esapi.Response
		reqSettings []func(*esapi.IndicesDeleteRequest) = []func(*esapi.IndicesDeleteRequest){
			esClient.Indices.Delete.WithContext(context.Background()),
			esClient.Indices.Delete.WithTimeout(5 * time.Second),
		}
	)

	if pretty {
		reqSettings = append(reqSettings, esClient.Indices.Delete.WithPretty())
	}

	if res, err = esClient.Indices.Delete(
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

func GetIndex(esClient *elasticsearch.Client, indexNames []string, pretty bool) {
	var (
		err         error
		res         *esapi.Response
		reqSettings []func(*esapi.IndicesGetRequest) = []func(*esapi.IndicesGetRequest){
			esClient.Indices.Get.WithContext(context.Background()),
		}
	)

	if pretty {
		reqSettings = append(reqSettings, esClient.Indices.Get.WithPretty())
	}

	if res, err = esClient.Indices.Get(
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
