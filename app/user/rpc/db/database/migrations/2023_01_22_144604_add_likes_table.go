package migrations

import (
	"database/sql"
	"time"

	"middle/pkg/migrate"

	"gorm.io/gorm"
)

func init() {

	type Likes struct {
		ID uint64 `gorm:"column:id;primaryKey;autoIncrement;" json:"id,omitempty"`

		Name     string `gorm:"type:varchar(255);not null;index"`
		Email    string `gorm:"type:varchar(255);index;default:null"`
		Phone    string `gorm:"type:varchar(20);index;default:null"`
		Password string `gorm:"type:varchar(255)"`

		CreatedAt time.Time `gorm:"column:created_at;index;" json:"created_at,omitempty"`
		UpdatedAt time.Time `gorm:"column:updated_at;index;" json:"updated_at,omitempty"`
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&Likes{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropTable(&Likes{})
	}

	migrate.Add("2023_01_22_144604_add_likes_table", up, down)
}
