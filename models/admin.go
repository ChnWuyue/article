package models

// Admin Model
type Admin struct {
	ID       string `json:"id" gorm:"uniqueIndex;Type:varchar(20)"`
	Password string `json:"password" gorm:"NOT NULL;Type:varchar(20)"`
}
