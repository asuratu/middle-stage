package migrations

import (
	"database/sql"
	"time"

	"middle/pkg/migrate"

	"gorm.io/gorm"
)

func init() {

	type Link struct {
		ID uint64 `gorm:"column:id;primaryKey;autoIncrement;" json:"id,omitempty"`

		Name string `gorm:"type:varchar(255);not null"`
		URL  string `gorm:"type:varchar(255);default:null"`

		CreatedAt time.Time `gorm:"column:created_at;index;" json:"created_at,omitempty"`
		UpdatedAt time.Time `gorm:"column:updated_at;index;" json:"updated_at,omitempty"`
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&Link{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropTable(&Link{})
	}

	migrate.Add("2022_08_07_164130_add_links_table", up, down)
}