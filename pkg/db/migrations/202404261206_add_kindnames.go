package migrations

import (
	"gorm.io/gorm"

	"github.com/go-gormigrate/gormigrate/v2"
)

func addKindNames() *gormigrate.Migration {
	type KindName struct {
		Model
	}

	return &gormigrate.Migration{
		ID: "202404261206",
		Migrate: func(tx *gorm.DB) error {
			return tx.AutoMigrate(&KindName{})
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable(&KindName{})
		},
	}
}
