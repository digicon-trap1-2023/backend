package model

type BookMark struct {
	Id         string `gorm:"type:char(36);not null;primaryKey"`
	UserId     string `gorm:"type:char(36)"`
	DocumentId string `gorm:"type:char(36)"`
}