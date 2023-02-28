package es

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"middle/app/user/rpc/internal/config"
	"strings"

	"github.com/elastic/go-elasticsearch/v7/esapi"

	"github.com/zeromicro/go-zero/core/logx"

	elastic "github.com/elastic/go-elasticsearch/v7"
)

type Es struct {
	Client *elastic.Client
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
	}
}

// IndexExists 检查索引是否存在
func (es *Es) IndexExists(indexName string) bool {
	response, err := es.Client.Indices.Exists([]string{indexName})

	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(response.Body)

	if err != nil {
		log.Fatalf("es index exists error: %s", err)
	}

	isError(response)

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

	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(response.Body)

	if err != nil {
		log.Fatalf("es create index error: %s", err)
	}

	isError(response)
}

// DeleteIndex 删除索引
func (es *Es) DeleteIndex(indexName string) {
	exists := es.IndexExists(indexName)
	if !exists {
		return
	}
	response, err := es.Client.Indices.Delete([]string{indexName})

	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(response.Body)

	if err != nil {
		log.Fatalf("es delete index %s error: %v", indexName, err)
	}

	isError(response)
}

// Insert 索引一条数据
func (es *Es) Insert(indexName, id string, body *bytes.Buffer) error {
	response, err := es.Client.Create(indexName, id, body)

	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(response.Body)

	if err != nil {
		logx.Errorf("es index: %s id: %s body %v insert error: %v", indexName, id, body, err)
		return err
	}

	isError(response)

	return nil
}

func isError(response *esapi.Response) {
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
