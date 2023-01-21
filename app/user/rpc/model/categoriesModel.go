package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"gorm.io/gorm"
)

var _ CategoriesModel = (*customCategoriesModel)(nil)

type (
	// CategoriesModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCategoriesModel.
	CategoriesModel interface {
		categoriesModel
		customCategoriesLogicModel
	}

	customCategoriesModel struct {
		*defaultCategoriesModel
	}

	customCategoriesLogicModel interface {
	}
)

// NewCategoriesModel returns a model for the database table.
func NewCategoriesModel(conn *gorm.DB, c cache.CacheConf) CategoriesModel {
	return &customCategoriesModel{
		defaultCategoriesModel: newCategoriesModel(conn, c),
	}
}

func (m *defaultCategoriesModel) customCacheKeys(data *Categories) []string {
	if data == nil {
		return []string{}
	}
	return []string{}
}
