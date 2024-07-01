package main

import (
	"gofiber-marketplace/src/configs"
	"gofiber-marketplace/src/helpers"
	"gofiber-marketplace/src/routes"
	"gofiber-marketplace/src/services"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/joho/godotenv"
)

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	return "0.0.0.0:" + port

}

func main() {
	if _, err := os.Stat(".env"); err == nil {
		if err := godotenv.Load(); err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	app := fiber.New()

	app.Use(helmet.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins:  "http://localhost:3000, http://localhost:5173, https://blanja-website-project.netlify.app, https://be-blanja-marketplace-app-develop.up.railway.app",
		AllowMethods:  "GET,POST,PUT,DELETE",
		AllowHeaders:  "*",
		ExposeHeaders: "Content-Length",
	}))

	configs.InitDB()
	services.InitMidtrans()
	helpers.Migration()
	routes.Router(app)

	if err := app.Listen(getPort()); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
