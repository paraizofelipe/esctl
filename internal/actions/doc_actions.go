package actions

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"strings"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
)

type Documenter interface {
	ListDoc(string, bool)
	AddDoc(string, string)
}

type DocAction struct {
	client *elasticsearch.Client
}

func NewDocAction(esClient *elasticsearch.Client) Documenter {
	return &DocAction{
		client: esClient,
	}
}

func MapToJSONReader(data map[string]interface{}) io.Reader {
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("Error converting map to JSON: %s", err)
	}
	return strings.NewReader(string(jsonData))
}

func (d *DocAction) ListDoc(indexName string, pretty bool) {

	var (
		err         error
		res         *esapi.Response
		reqSettings []func(*esapi.SearchRequest) = []func(*esapi.SearchRequest){
			d.client.Search.WithContext(context.Background()),
			d.client.Search.WithIndex(indexName),
			d.client.Search.WithSize(10),
			d.client.Search.WithTimeout(5 * time.Second),
		}
	)

	if pretty {
		reqSettings = append(reqSettings, d.client.Search.WithPretty())
	}

	if res, err = d.client.Search(reqSettings...); err != nil {
		log.Fatalf("Error when querying the index: %s", err)
	}

	defer res.Body.Close()

	if res.IsError() {
		log.Fatalf("[REQ-ERROR]: %s", res.String())
	}

	fmt.Println(res.String())

}

func (d *DocAction) AddDoc(indexName string, doc string) {
	var (
		err         error
		res         *esapi.Response
		reqSettings []func(*esapi.IndexRequest) = []func(*esapi.IndexRequest){}
	)

	// if pretty {
	// 	reqSettings = append(reqSettings, d.client.Search.WithPretty())
	// }

	if res, err = d.client.Index(indexName, strings.NewReader(doc), reqSettings...); err != nil {
		log.Fatalf("Error when querying the index: %s", err)
	}

	defer res.Body.Close()

	if res.IsError() {
		log.Fatalf("[REQ-ERROR]: %s", res.String())
	}

	fmt.Println(res.String())
}
