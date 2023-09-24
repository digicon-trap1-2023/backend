package migration

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB, tables []interface{}) (init bool, err error) {
	m := gormigrate.New(db, gormigrate.DefaultOptions, Migrations())

	m.InitSchema(func(db *gorm.DB) error {
		init = true

		return db.AutoMigrate(tables...)
	})
	err = m.Migrate()
	return
}
