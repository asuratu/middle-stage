package migrations

import (
	"database/sql"
	"time"

	"middle/pkg/migrate"

	"gorm.io/gorm"
)

func init() {

	type User struct {
		ID uint64 `gorm:"column:id;primaryKey;autoIncrement;" json:"id,omitempty"`
	}
	type Category struct {
		ID uint64 `gorm:"column:id;primaryKey;autoIncrement;" json:"id,omitempty"`
	}

	type Topic struct {
		ID uint64 `gorm:"column:id;primaryKey;autoIncrement;" json:"id,omitempty"`

		Title      string `gorm:"type:varchar(255);not null;index"`
		Body       string `gorm:"type:longtext;not null"`
		UserID     string `gorm:"type:bigint;not null;index"`
		CategoryID string `gorm:"type:bigint;not null;index"`

		// 会创建 user_id 和 category_id 外键的约束
		User     User
		Category Category

		CreatedAt time.Time `gorm:"column:created_at;index;" json:"created_at,omitempty"`
		UpdatedAt time.Time `gorm:"column:updated_at;index;" json:"updated_at,omitempty"`
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&Topic{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropTable(&Topic{})
	}

	migrate.Add("2022_08_07_153652_add_topics_table", up, down)
}
