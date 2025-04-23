package models

// ErrorCode mendefinisikan struktur error yang akan disimpan di database
type ErrorCode struct {
	ID       uint   `gorm:"primary_key;auto_increment" json:"id"`
	ErrorKey string `gorm:"type:varchar(255);unique_index;not null" json:"error_key"`
	Code     string `gorm:"type:varchar(20);not null" json:"code"`
	Message  string `gorm:"type:text;not null" json:"message"`
}
