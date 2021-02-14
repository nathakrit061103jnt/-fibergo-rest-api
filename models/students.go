package models

import (
	_ "gorm.io/gorm"
)

type Student struct {
	// gorm.Model
	ID         int    `json:"id" gorm:"primaryKey;autoIncrement"`
	First_Name string `json:"first_name"`
	Last_Name  string `json:"last_name"`
	Age        int    `json:"age"`
	Class      string `json:"class"`
}
