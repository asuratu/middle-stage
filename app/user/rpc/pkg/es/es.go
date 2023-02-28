package es

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"log"
	"middle/app/user/rpc/internal/config"
	"strings"

	"github.com/zeromicro/go-zero/core/service"

	"github.com/elastic/go-elasticsearch/v7/esapi"

	"github.com/zeromicro/go-zero/core/logx"

	elastic "github.com/elastic/go-elasticsearch/v7"
)

type Es struct {
	Client *elastic.Client
	Config config.Config
}

func Boot(c config.Config) *Es {
	esClient := NewEsClient(c)
	// 检查索引是否存在
	userIndexExists := esClient.IndexExists(UserIndexName)
	if !userIndexExists {
		// 创建索引
		esClient.CreateIndex(UserIndexName, UserMappings)
	}
	return esClient
}

// NewEsClient create es client.
func NewEsClient(c config.Config) (es *Es) {
	client, err := elastic.NewClient(elastic.Config{
		Addresses: []string{
			c.Es.Address,
		},
		Username: c.Es.Username,
		Password: c.Es.Password,
	})
	if err != nil {
		log.Fatal("es client error: ", err)
	}
	return &Es{
		Client: client,
		Config: c,
	}
}

// IndexExists 检查索引是否存在
func (es *Es) IndexExists(indexName string) bool {
	response, err := es.Client.Indices.Exists([]string{indexName})

	if err != nil {
		log.Fatalf("es index exists error: %s", err)
	}

	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(response.Body)

	IsError(response)

	return true
}

// CreateIndex 创建索引
func (es *Es) CreateIndex(indexName, indexBody string) {
	exists := es.IndexExists(indexName)
	if exists {
		return
	}

	logx.Infof("------------- es create index: %s", indexName)

	// 创建索引
	indexBody = UserMappings
	response, err := es.Client.Indices.Create(indexName, es.Client.Indices.Create.WithBody(strings.NewReader(UserMappings)))

	if err != nil {
		log.Fatalf("es create index error: %s", err)
	}

	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(response.Body)

	IsError(response)
}

// DeleteIndex 删除索引
func (es *Es) DeleteIndex(indexName string) {
	exists := es.IndexExists(indexName)
	if !exists {
		return
	}
	response, err := es.Client.Indices.Delete([]string{indexName})

	if err != nil {
		log.Fatalf("es delete index %s error: %v", indexName, err)
	}

	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(response.Body)

	IsError(response)
}

// Insert 索引一条数据
func (es *Es) Insert(indexName, id string, body *bytes.Buffer) error {
	response, err := es.Client.Create(indexName, id, body)

	if err != nil {
		logx.Errorf("es index: %s id: %s body %v insert error: %v", indexName, id, body, err)
		return err
	}

	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(response.Body)

	IsError(response)

	return nil
}

// 全量更新数据

// 增量更新数据

func (es *Es) Search(ctx context.Context, indexName string, body *bytes.Buffer) (rsp []map[string]interface{}, err error) {
	var response *esapi.Response
	response, err = es.Client.Search(
		es.Client.Search.WithContext(ctx),
		es.Client.Search.WithIndex(indexName),
		es.Client.Search.WithBody(body),
		es.Client.Search.WithTrackTotalHits(true),
		es.Client.Search.WithPretty(),
	)

	if err != nil {
		logx.WithContext(ctx).Errorf("es index: %s body %v search error: %v", indexName, body, err)
		return nil, err
	}

	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(response.Body)

	IsError(response)

	esRsp := make(map[string]interface{})
	if err := json.NewDecoder(response.Body).Decode(&esRsp); err != nil {
		logx.WithContext(ctx).Errorf("Error parsing the response body, index: %s body %v search error: %v", indexName, body, err)
		return nil, err
	}

	// 如果不是生产环境，打印日志
	if es.Config.Mode != service.ProMode {
		logx.Infof("[%s] %d hits; took: %dms",
			response.Status(),
			int(esRsp["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64)),
			int(esRsp["took"].(float64)),
		)
		logx.Info(strings.Repeat("=", 37))
	}

	rsp = make([]map[string]interface{}, 0)
	for _, hit := range esRsp["hits"].(map[string]interface{})["hits"].([]interface{}) {
		item := make(map[string]interface{})
		item["id"] = hit.(map[string]interface{})["_id"]
		item["source"] = hit.(map[string]interface{})["_source"]
		logx.Infof(" * ID=%s, %s", hit.(map[string]interface{})["_id"], hit.(map[string]interface{})["_source"])
		rsp = append(rsp, item)
	}

	return rsp, nil
}

func IsError(response *esapi.Response) {
	if response.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(response.Body).Decode(&e); err != nil {
			log.Fatalf("Error parsing the response body: %s", err)
		} else {
			log.Fatalf("[%s] %s: %s",
				response.Status(),
				e["error"].(map[string]interface{})["type"],
				e["error"].(map[string]interface{})["reason"],
			)
		}
	}
}
