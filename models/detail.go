package models

// Detail Model
type Detail struct {
	ID         int    `json:"id" gorm:"primaryKey;uniqueIndex;AUTO_INCREMENT"`
	Title      string `json:"title" gorm:"index;not null;Type:VARCHAR(255)"`
	CreateTime string `json:"create_time" gorm:"Type:varchar(20)"`
	UpdateTime string `json:"update_time" gorm:"Type:varchar(20)"`
	Html       string `json:"html" gorm:"Type:text;not null"`
	Abstract   string `json:"abstract" gorm:"Type:text"`
	Display    string `json:"display" gorm:"default:0"`
	Cover      string `json:"cover" gorm:"Type:varchar(100)"`
}
