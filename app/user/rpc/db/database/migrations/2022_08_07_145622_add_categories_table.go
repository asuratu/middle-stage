package migrations

import (
	"database/sql"
	"time"

	"middle/pkg/migrate"

	"gorm.io/gorm"
)

func init() {

	type Category struct {
		ID uint64 `gorm:"column:id;primaryKey;autoIncrement;" json:"id,omitempty"`

		Name        string `gorm:"type:varchar(255);not null;index"`
		Description string `gorm:"type:varchar(255);default:null"`

		CreatedAt time.Time `gorm:"column:created_at;index;" json:"created_at,omitempty"`
		UpdatedAt time.Time `gorm:"column:updated_at;index;" json:"updated_at,omitempty"`
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&Category{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropTable(&Category{})
	}

	migrate.Add("2022_08_07_145622_add_categories_table", up, down)
}
