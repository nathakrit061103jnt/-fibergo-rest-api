package models

import (
	_ "github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Bookings struct {
	// gorm.Model
	// Bk_id      int    `json:"bk_id"`
	M_id       int    `json:"m_id"`
	Pk_id      int    `json:"pk_id"`
	Bk_start   string `json:"bk_start"`
	Bk_end     string `json:"bk_end"`
	Bk_time    int    `json:"bk_time"`
	Bk_tel     string `json:"bk_tel"`
	Bk_num_car string `json:"bk_num_car"`
	Pk_date    string `json:"pk_date"`
}
