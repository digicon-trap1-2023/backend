package model

type User struct {
	Id   string `gorm:"type:char(36);not null;primaryKey"`
	Name string `gorm:"type:varchar(32)"`
}