package es

import (
	"log"
	"middle/app/user/rpc/internal/config"
	"strings"

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
	exists, err := es.Client.Indices.Exists([]string{indexName})
	if err != nil {
		log.Fatalf("es index exists error: %s", err)
	}
	if exists.StatusCode != 200 {
		return false
	}
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

	if response.StatusCode != 200 {
		log.Fatalf("es create index %s error: %v", indexName, response)
	}
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
	if response.StatusCode != 200 {
		log.Fatalf("es delete index %s error: %v", indexName, response)
	}
}
