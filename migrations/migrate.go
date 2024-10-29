package migrations

import (
	"resto_app_api/entity"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(
		&entity.User{},
	)
	if err != nil {
		return err
	}

	return nil
}
