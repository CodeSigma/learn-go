package models

import (
	"github.com/jinzhu/gorm"
)

var db *gorm.DB
var err *error

type Model struct {
	ID        int `gorm:"primary_key" json:"id"`
	CreatedAt int `json:"created_at"`
	UpdatedAt int `json:"updated_at"`
}

func initModel() {
	db, err : = gorm.Open(
		config.DatabaseSetting.type,
		

	)
}
