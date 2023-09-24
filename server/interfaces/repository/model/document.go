package model

type Document struct {
	Id          string `gorm:"type:char(36);not null;primaryKey"`
	Title       string `gorm:"type:varchar(40)"`
	Description string `gorm:"type:varchar(200)"`
	FileId      string `gorm:"type:char(36)"`
}