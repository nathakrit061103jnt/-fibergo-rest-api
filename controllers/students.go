package controllers

import (
	"fmt"

	"github.com/fiber_go_api/configs"
	"github.com/fiber_go_api/models"
	"github.com/gofiber/fiber/v2" // new
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func GetStudents(c *fiber.Ctx) error {
	db = configs.InitDB()
	var students []models.Student
	db.Preload("Items").Find(&students)
	c.JSON(&students)

	// defer db.Close()
	return c.JSON(&students)
}

func UpdateStudent(c *fiber.Ctx) error {
	db = configs.InitDB()
	var student models.Student
	student.Age = 16
	student.Class = "111"
	student.First_Name = "Ali"
	student.Last_Name = "Kal"

	db.Update(&student)
	c.JSON(&student)

	// defer db.Close()
	return c.JSON(&student)
}

// parametre olarak id yi alıp işlem yapmadım.
// func DeleteStudent(c *fiber.Ctx) error {
// 	db = initDB()
// 	db.Where("ID = ?", 1).Delete(&Student{})
// 	defer db.Close()
// 	// defer db.Close()
// 	return c.JSON(&student)
// }

func NewStudent(c *fiber.Ctx) error {

	var student models.Student
	student.Age = 14
	student.Class = "12"
	student.First_Name = "Ali"
	student.Last_Name = "Kalim"
	db.Create(&student)
	fmt.Println("Yazıldı")
	c.JSON(&student)
	// defer db.Close()
	return c.JSON(&student)

}
