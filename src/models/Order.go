package models

import (
	"gofiber-marketplace/src/configs"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	UserID            uint     `json:"user_id" validate:"required"`
	User              User     `gorm:"foreignKey:UserID" validate:"-"`
	AddressID         uint     `json:"address_id" validate:"required"`
	Address           Address  `gorm:"foreignKey:AddressID" validate:"-"`
	CheckoutID        uint     `json:"checkout_id" validate:"required"`
	Checkout          Checkout `gorm:"foreignKey:CheckoutID" validate:"-"`
	TransactionNumber string   `json:"transaction_number"`
	PaymentMethod     string   `json:"payment_method"`
	TotalPrice        float64  `json:"total_price"`
	Status            string   `json:"status" validate:"required,oneof=not_yet_paid expired get_paid processed sent completed canceled"`
	TransactionURL    string   `json:"transaction_url"`
}

type VaNumber struct {
	Bank     string `json:"bank"`
	VaNumber string `json:"va_number"`
}

type Notification struct {
	TransactionStatus string     `json:"transaction_status"`
	OrderID           string     `json:"order_id"`
	GrossAmount       string     `json:"gross_amount"`
	PaymentType       string     `json:"payment_type"`
	TransactionTime   string     `json:"transaction_time"`
	TransactionID     string     `json:"transaction_id"`
	FraudStatus       string     `json:"fraud_status"`
	VaNumbers         []VaNumber `json:"va_numbers"`
}

func SelectAllOrders() []*Order {
	var orders []*Order
	configs.DB.Preload("User").Preload("Address").Preload("Checkout").Find(&orders)
	return orders
}

func SelectOrderbyId(id int) *Order {
	var order Order
	configs.DB.Preload("User").Preload("Address").Preload("Checkout").First(&order, "id = ?", id)
	return &order
}

func SelectOrdersbyUserId(userID int) []*Order {
	var orders []*Order
	configs.DB.Preload("User").Preload("Address").Preload("Checkout").Find(&orders, "user_id = ?", userID)
	return orders
}

func SelectOrderbyTransactionNumber(trans_number string) *Order {
	var order Order
	configs.DB.Preload("User").Preload("Address").Preload("Checkout").First(&order, "transaction_number = ?", trans_number)
	return &order
}

// func SelectOrderbyStatus(status string) *Order {
// 	var order Order
// 	configs.DB.Preload("User").First(&order, "status = ?", status)
// 	return &order
// }

// func SelectOrdersbyStatus(status string) []*Order {
// 	var orders []*Order
// 	configs.DB.Preload("User").Find(&orders, "status = ?", status)
// 	return orders
// }

func CreateOrder(order *Order) error {
	result := configs.DB.Create(&order)
	return result.Error
}

func UpdateOrder(id int, updatedOrder *Order) error {
	result := configs.DB.Model(&Order{}).Where("id = ?", id).Updates(updatedOrder)
	return result.Error
}

func UpdateStatusOrder(id int, status string) error {
	result := configs.DB.Model(&Order{}).Where("id = ?", id).Update("status", status)
	return result.Error
}

func UpdatePaymentMethodOrder(id int, paymentMethod string) error {
	result := configs.DB.Model(&Order{}).Where("id = ?", id).Update("payment_method", paymentMethod)
	return result.Error
}

func UpdateURLOrder(id int, url string) error {
	result := configs.DB.Model(&Order{}).Where("id = ?", id).Update("transaction_url", url)
	return result.Error
}

func DeleteOrder(id int) error {
	result := configs.DB.Delete(&Order{}, "id = ?", id)
	return result.Error
}
