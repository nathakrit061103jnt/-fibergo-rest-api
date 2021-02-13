package models

import (
	"github.com/jinzhu/gorm"
)

type Student struct {
	gorm.Model
	First_Name string `json:"first_name"`
	Last_Name  string `json:"last_name"`
	Age        int    `json:"age"`
	Class      string `json:"class"`
}
