package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"gorm.io/gorm"
)

var _ LikesModel = (*customLikesModel)(nil)

type (
	// LikesModel is an interface to be customized, add more methods here,
	// and implement the added methods in customLikesModel.
	LikesModel interface {
		likesModel
		customLikesLogicModel
	}

	customLikesModel struct {
		*defaultLikesModel
	}

	customLikesLogicModel interface {
	}
)

// NewLikesModel returns a model for the database table.
func NewLikesModel(conn *gorm.DB, c cache.CacheConf) LikesModel {
	return &customLikesModel{
		defaultLikesModel: newLikesModel(conn, c),
	}
}

func (m *defaultLikesModel) customCacheKeys(data *Likes) []string {
	if data == nil {
		return []string{}
	}
	return []string{}
}
