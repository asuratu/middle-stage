package seeders

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"

	"middle/pkg/console"
	"middle/pkg/seed"

	"gorm.io/gorm"
	"middle/app/user/rpc/db/database/factories"
)

func init() {

	seed.Add("SeedLinksTable", func(db *gorm.DB) {

		links := factories.MakeLinks(5)

		result := db.Table("links").Create(&links)

		if err := result.Error; err != nil {
			logx.Error(err)
			return
		}

		console.Success(fmt.Sprintf("Table [%v] %v rows seeded", result.Statement.Table, result.RowsAffected))
	})
}
