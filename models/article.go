package models

// Article Model
type Article struct {
	ID         int    `json:"id" gorm:"primaryKey;uniqueIndex;AUTO_INCREMENT"`
	Title      string `json:"title" gorm:"index;not null;Type:VARCHAR(255)"`
	CreateTime string `json:"create_time" gorm:"Type:varchar(20)"`
	UpdateTime string `json:"update_time" gorm:"Type:varchar(20)"`
	Html       string `json:"html" gorm:"Type:text"`
	Author     string `json:"author" gorm:"not null;Type:varchar(20)"`
	Views      int    `json:"views" gorm:"Type:BIGINT"`
	Link       string `json:"link" gorm:"Type:varchar(255)"`
	File       string `json:"file" gorm:"Type:varchar(100)"`
}
