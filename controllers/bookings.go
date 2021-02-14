package controllers

import (
	"fmt"

	"github.com/fiber_go_api/models"
	"github.com/gofiber/fiber/v2"
)

func NewBooking(ctx *fiber.Ctx) error {

	var student models.Student
	student.Age = 14
	student.Class = "12"
	student.First_Name = "Ali"
	student.Last_Name = "Kalim"
	db.Create(&student)
	fmt.Println("Yazıldı")
	return ctx.JSON(&student)
	// defer db.Close()
}
