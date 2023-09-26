package model

type Request struct {
	Id          string `gorm:"type:char(36);not null;primaryKey"`
	Title       string `gorm:"type:varchar(36)"`
	Description string `gorm:"type:varchar(200)"`
	UserId      string `gorm:"type:char(36)"`
}

type TagRequest struct {
	RequestId   string `gorm:"type:char(36);not null;primaryKey"`
	TagId       string `gorm:"type:char(36);not null;primaryKey"`
}
