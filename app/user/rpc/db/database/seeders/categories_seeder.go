package seeders

import (
    "fmt"

    "middle/app/user/rpc/db/database/factories"
    "middle/pkg/console"
    "middle/pkg/seed"

    "github.com/zeromicro/go-zero/core/logx"
    "gorm.io/gorm"
)

func init() {

    seed.Add("SeedCategoriesTable", func(db *gorm.DB) {

        categories := factories.MakeCategories(10)

        result := db.Table("categories").Create(&categories)

        if err := result.Error; err != nil {
            logx.Error(err)
            return
        }

        console.Success(fmt.Sprintf("Table [%v] %v rows seeded", result.Statement.Table, result.RowsAffected))
    })
}
