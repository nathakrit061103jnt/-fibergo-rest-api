package main

import (
	"github.com/fiber_go_api/controllers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger" // new
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	// db = configs.InitDB()
	app := fiber.New()

	app.Use(logger.New()) // new

	// Route oluşturuldu.
	app.Get("/api/students", controllers.GetStudents)
	app.Get("/api/student", controllers.NewStudent)
	// app.Get("/studentt", controllers.DeleteStudent)
	app.Get("/api/studentUp", controllers.UpdateStudent)

	app.Get("/api/parkings", controllers.GetParkings)

	app.Listen(":9001") // 3000. porttan serv edildi

}
