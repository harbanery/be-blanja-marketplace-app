package helpers

import (
	"gofiber-marketplace/src/configs"
	"gofiber-marketplace/src/models"
	"log"
)

func Migration() {
	err := configs.DB.AutoMigrate(
		&models.User{},
		&models.UserVerification{},
		&models.Seller{},
		&models.Customer{},
		&models.Product{},
		&models.Image{},
		&models.Size{},
		&models.Color{},
		&models.Category{},
		&models.Address{},
		&models.Cart{},
		&models.CartProduct{},
		&models.Checkout{},
		&models.Order{},
	)

	if err != nil {
		log.Fatalf("Failed to auto migrate: %v", err)
	}
}
