package model

type Tag struct {
	Id   string `gorm:"type:char(36);not null;primaryKey"`
	Name string `gorm:"type:varchar(40)"`
}

type TagDocument struct {
	Id         string `gorm:"type:char(36);not null;primaryKey"`
	TagId      string `gorm:"type:char(36)"`
	DocumentId string `gorm:"type:char(36)"`
}