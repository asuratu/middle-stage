package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"gorm.io/gorm"
)

var _ TopicsModel = (*customTopicsModel)(nil)

type (
	// TopicsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTopicsModel.
	TopicsModel interface {
		topicsModel
		customTopicsLogicModel
	}

	customTopicsModel struct {
		*defaultTopicsModel
	}

	customTopicsLogicModel interface {
	}
)

// NewTopicsModel returns a model for the database table.
func NewTopicsModel(conn *gorm.DB, c cache.CacheConf) TopicsModel {
	return &customTopicsModel{
		defaultTopicsModel: newTopicsModel(conn, c),
	}
}

func (m *defaultTopicsModel) customCacheKeys(data *Topics) []string {
	if data == nil {
		return []string{}
	}
	return []string{}
}
