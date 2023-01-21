package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"gorm.io/gorm"
)

var _ LinksModel = (*customLinksModel)(nil)

type (
	// LinksModel is an interface to be customized, add more methods here,
	// and implement the added methods in customLinksModel.
	LinksModel interface {
		linksModel
		customLinksLogicModel
	}

	customLinksModel struct {
		*defaultLinksModel
	}

	customLinksLogicModel interface {
	}
)

// NewLinksModel returns a model for the database table.
func NewLinksModel(conn *gorm.DB, c cache.CacheConf) LinksModel {
	return &customLinksModel{
		defaultLinksModel: newLinksModel(conn, c),
	}
}

func (m *defaultLinksModel) customCacheKeys(data *Links) []string {
	if data == nil {
		return []string{}
	}
	return []string{}
}
