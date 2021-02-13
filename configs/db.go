package configs

import (
	"fmt"

	"github.com/fiber_go_api/models"
	"github.com/jinzhu/gorm"
)

func InitDB() *gorm.DB {
	db, err := gorm.Open("mysql", "root:@tcp(127.0.0.1:3307)/smart-parking?charset=utf8&parseTime=True")
	if err != nil {
		panic("Error")
	}
	fmt.Println("Database Connected")
	db.AutoMigrate(&models.Student{})
	return db
}
