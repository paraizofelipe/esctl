package actions

import (
	"bytes"
	"fmt"
	"io"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/urfave/cli/v2"
)

type CatManager interface {
	Aliases(ctx *cli.Context) error
	Indices(ctx *cli.Context) error
	Nodes(ctx *cli.Context) error
	PendingTasks(ctx *cli.Context) error
	Shards(ctx *cli.Context) error
	ThreadPool(ctx *cli.Context) error
}

type CatAction struct {
	client *elasticsearch.Client
}

func NewCatAction(esClient *elasticsearch.Client) CatManager {
	return &CatAction{
		client: esClient,
	}
}

func readerCloserToString(rc io.ReadCloser) (string, error) {
	defer rc.Close()

	var buf bytes.Buffer
	_, err := io.Copy(&buf, rc)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

func (c *CatAction) Aliases(ctx *cli.Context) (err error) {
	var (
		res         *esapi.Response
		pretty      bool = ctx.Bool("pretty")
		reqSettings      = []func(*esapi.CatAliasesRequest){
			c.client.Cat.Aliases.WithV(true),
		}
		resBody string
	)

	reqSettings = append(reqSettings, c.client.Cat.Aliases.WithHelp(ctx.Bool("describe")))

	if pretty {
		reqSettings = append(reqSettings, c.client.Cat.Aliases.WithPretty())
	}

	if res, err = c.client.Cat.Aliases(reqSettings...); err != nil {
		return fmt.Errorf("%s", err)
	}

	defer res.Body.Close()

	if res.IsError() {
		return fmt.Errorf("[REQ-ERROR]: %s", res.String())
	}

	if resBody, err = readerCloserToString(res.Body); err != nil {
		return fmt.Errorf("%s", err)
	}

	fmt.Println(resBody)

	return
}

func (c *CatAction) Indices(ctx *cli.Context) (err error) {
	var (
		res         *esapi.Response
		reqSettings = []func(*esapi.CatIndicesRequest){
			c.client.Cat.Indices.WithV(true),
		}
		pretty  bool = ctx.Bool("pretty")
		resBody string
	)

	reqSettings = append(reqSettings, c.client.Cat.Indices.WithHelp(ctx.Bool("describe")))

	if pretty {
		reqSettings = append(reqSettings, c.client.Cat.Indices.WithPretty())
	}

	if res, err = c.client.Cat.Indices(reqSettings...); err != nil {
		return fmt.Errorf("%s", err)
	}

	defer res.Body.Close()

	if res.IsError() {
		return fmt.Errorf("[REQ-ERROR]: %s", res.String())
	}

	if resBody, err = readerCloserToString(res.Body); err != nil {
		return fmt.Errorf("%s", err)
	}

	fmt.Println(resBody)

	return
}

func (c *CatAction) Nodes(ctx *cli.Context) (err error) {
	var (
		res         *esapi.Response
		reqSettings = []func(*esapi.CatNodesRequest){
			c.client.Cat.Nodes.WithV(true),
		}
		columns      = ctx.String("columns")
		pretty  bool = ctx.Bool("pretty")
		resBody string
	)

	reqSettings = append(reqSettings, c.client.Cat.Nodes.WithHelp(ctx.Bool("describe")))

	if pretty {
		reqSettings = append(reqSettings, c.client.Cat.Nodes.WithPretty())
	}

	if columns != "" {
		reqSettings = append(reqSettings, c.client.Cat.Nodes.WithH(columns))
	}

	if res, err = c.client.Cat.Nodes(reqSettings...); err != nil {
		return fmt.Errorf("%s", err)
	}

	defer res.Body.Close()

	if res.IsError() {
		return fmt.Errorf("[REQ-ERROR]: %s", res.String())
	}

	if resBody, err = readerCloserToString(res.Body); err != nil {
		return fmt.Errorf("%s", err)
	}

	fmt.Println(resBody)

	return
}

func (c *CatAction) Shards(ctx *cli.Context) (err error) {
	var (
		columns     = ctx.String("columns")
		indexName   = ctx.String("index")
		res         *esapi.Response
		reqSettings = []func(*esapi.CatShardsRequest){
			c.client.Cat.Shards.WithV(true),
			c.client.Cat.Shards.WithIndex(indexName),
		}
		pretty  = ctx.Bool("pretty")
		resBody string
	)

	reqSettings = append(reqSettings, c.client.Cat.Shards.WithHelp(ctx.Bool("describe")))

	if columns != "" {
		reqSettings = append(reqSettings, c.client.Cat.Shards.WithH(columns))
	}

	if pretty {
		reqSettings = append(reqSettings, c.client.Cat.Shards.WithPretty())
	}

	if res, err = c.client.Cat.Shards(reqSettings...); err != nil {
		return fmt.Errorf("%s", err)
	}

	defer res.Body.Close()

	if res.IsError() {
		return fmt.Errorf("[REQ-ERROR]: %s", res.String())
	}

	if resBody, err = readerCloserToString(res.Body); err != nil {
		return fmt.Errorf("%s", err)
	}

	fmt.Println(resBody)

	return
}

func (c *CatAction) PendingTasks(ctx *cli.Context) (err error) {
	var (
		columns     = ctx.String("columns")
		res         *esapi.Response
		reqSettings = []func(*esapi.CatPendingTasksRequest){
			c.client.Cat.PendingTasks.WithV(true),
		}
		pretty  = ctx.Bool("pretty")
		resBody string
	)

	reqSettings = append(reqSettings, c.client.Cat.PendingTasks.WithHelp(ctx.Bool("describe")))

	if columns != "" {
		reqSettings = append(reqSettings, c.client.Cat.PendingTasks.WithH(columns))
	}

	if pretty {
		reqSettings = append(reqSettings, c.client.Cat.PendingTasks.WithPretty())
	}

	if res, err = c.client.Cat.PendingTasks(reqSettings...); err != nil {
		return fmt.Errorf("%s", err)
	}

	defer res.Body.Close()

	if res.IsError() {
		return fmt.Errorf("[REQ-ERROR]: %s", res.String())
	}

	if resBody, err = readerCloserToString(res.Body); err != nil {
		return fmt.Errorf("%s", err)
	}

	fmt.Println(resBody)

	return
}

func (c *CatAction) ThreadPool(ctx *cli.Context) (err error) {
	var (
		columns        = ctx.String("columns")
		nameThreadPool = ctx.String("thread-pool-pattern")
		res            *esapi.Response
		reqSettings    = []func(*esapi.CatThreadPoolRequest){
			c.client.Cat.ThreadPool.WithThreadPoolPatterns(nameThreadPool),
			c.client.Cat.ThreadPool.WithV(true),
		}
		resBody string
	)

	if columns != "" {
		reqSettings = append(reqSettings, c.client.Cat.ThreadPool.WithH(columns))
	}

	if ctx.Bool("pretty") {
		reqSettings = append(reqSettings, c.client.Cat.ThreadPool.WithPretty())
	}

	reqSettings = append(reqSettings, c.client.Cat.ThreadPool.WithHelp(ctx.Bool("describe")))

	if res, err = c.client.Cat.ThreadPool(reqSettings...); err != nil {
		return fmt.Errorf("%s", err)
	}

	defer res.Body.Close()

	if res.IsError() {
		return fmt.Errorf("[REQ-ERROR]: %s", res.String())
	}

	if resBody, err = readerCloserToString(res.Body); err != nil {
		return fmt.Errorf("%s", err)
	}

	fmt.Println(resBody)

	return
}
