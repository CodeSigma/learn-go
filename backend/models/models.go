package models

import (
	"fmt"

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
	var err error

	db, err = gorm.Open(
		"mysql",
		fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
			"root",
			"root",
			"localhost:3306",
			"go_tweet"))

}
