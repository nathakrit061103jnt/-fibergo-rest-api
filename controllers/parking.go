package controllers

import (
	"github.com/fiber_go_api/configs"
	"github.com/fiber_go_api/models"
	"github.com/gofiber/fiber/v2"
)

var err error

func GetParkings(c *fiber.Ctx) error {
	db := configs.InitDB()
	var parking []models.Tbl_parkings
	db.Find(&parking)
	// db.Exec("SELECT * FROM tbl_parkings").Joins("JOIN tbl_status USING(st_id)").Scan(&parking)
	db.Table("tbl_parkings").Select("*").Joins("JOIN tbl_status USING(st_id)").Scan(&parking)
	// defer db.Close()
	return c.JSON(&parking)
}

func GetParkingsId(c *fiber.Ctx) error {
	db := configs.InitDB()
	pk_id := c.Params("id")
	var parking []models.Tbl_parkings
	db.Find(&parking)
	// db.Exec("SELECT * FROM tbl_parkings").Joins("JOIN tbl_status USING(st_id)").Scan(&parking)
	db.Table("tbl_parkings").Select("*").Joins("JOIN tbl_status USING(st_id)").Where("pk_id = ?", pk_id).Scan(&parking)
	// defer db.Close()
	return c.JSON(&parking)
}
