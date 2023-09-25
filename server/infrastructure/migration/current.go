package migration

import (
	"github.com/digicon-trap1-2023/backend/interfaces/repository/model"
	"github.com/go-gormigrate/gormigrate/v2"
)

func Migrations() []*gormigrate.Migration {
	return []*gormigrate.Migration{
		v1(),
	}
}

func AllTables() []interface{} {
	return []interface{}{
		&model.User{},
		&model.BookMark{},
		&model.Document{},
		&model.TagDocument{},
		&model.Tag{},
		&model.Reference{},
	}
}
