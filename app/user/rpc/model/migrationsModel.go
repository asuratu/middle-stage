package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"gorm.io/gorm"
)

var _ MigrationsModel = (*customMigrationsModel)(nil)

type (
	// MigrationsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customMigrationsModel.
	MigrationsModel interface {
		migrationsModel
		customMigrationsLogicModel
	}

	customMigrationsModel struct {
		*defaultMigrationsModel
	}

	customMigrationsLogicModel interface {
	}
)

// NewMigrationsModel returns a model for the database table.
func NewMigrationsModel(conn *gorm.DB, c cache.CacheConf) MigrationsModel {
	return &customMigrationsModel{
		defaultMigrationsModel: newMigrationsModel(conn, c),
	}
}

func (m *defaultMigrationsModel) customCacheKeys(data *Migrations) []string {
	if data == nil {
		return []string{}
	}
	return []string{}
}
