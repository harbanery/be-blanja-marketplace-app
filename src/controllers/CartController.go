package controllers

import (
	"gofiber-marketplace/src/configs"
	"gofiber-marketplace/src/helpers"
	"gofiber-marketplace/src/middlewares"
	"gofiber-marketplace/src/models"
	// "strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

func GetCart(c *fiber.Ctx) error {
	carts := models.SelectAllCart()
	if len(carts) == 0 {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":     "no content",
			"statusCode": 200,
			"message":    "Cart is empty. You should create cart",
		})
	}
	resultCarts := make([]map[string]interface{}, len(carts))
	for i, cart := range carts {
		products := make([]map[string]interface{}, len(cart.Products))
		for j, product := range cart.Products {
			var cartProduct models.CartProduct
			if err := configs.DB.Where("cart_id = ? AND product_id = ?", cart.ID, product.ID).First(&cartProduct).Error; err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve quantity"})
			}
			var image *models.Image
			configs.DB.First(&image, "product_id = ?", product.ID)
			products[j] = map[string]interface{}{
				"id":         product.ID,
				"created_at": product.CreatedAt,
				"updated_at": product.UpdatedAt,
				"name":       product.Name,
				"price":      product.Price,
				"photo":      image.URL,
				"size":       product.Sizes,
				"color":      product.Colors,
				"quantity":   cartProduct.Quantity,
				"rating":     product.Rating,
			}
		}

		resultCarts[i] = map[string]interface{}{
			"id":          cart.ID,
			"user_id":     cart.UserID,
			"seller_id":   cart.SellerID,
			"seller_name": cart.Seller.Name,
			"created_at":  cart.CreatedAt,
			"updated_at":  cart.UpdatedAt,
			"products":    products,
		}
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":     "success",
		"statusCode": 200,
		"data":       resultCarts,
	})
}

func CreateCart(c *fiber.Ctx) error {
	var newCart models.Cart
	var user = c.Locals("user").(jwt.MapClaims)
	if err := c.BodyParser(&newCart); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":     "bad request",
			"statusCode": 400,
			"message":    "Invalid request body",
		})
	}
	userID := user["id"].(float64)
	newCart.UserID = uint(userID)

	cart := middlewares.XSSMiddleware(&newCart).(*models.Cart)

	if errors := helpers.StructValidation(cart); len(errors) > 0 {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"status":     "unprocessable entity",
			"statusCode": 422,
			"message":    "There are bad request or validation",
			"errors":     errors,
		})
	}

	// Use a transaction to ensure atomicity
	err := configs.DB.Transaction(func(tx *gorm.DB) error {
		var existingCart models.Cart
		if err := tx.Where("user_id = ? AND product_id = ?", newCart.UserID, newCart.ProductID).First(&existingCart).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				// Cart with the same user and product not found, create a new cart
				if err := tx.Create(cart).Error; err != nil {
					return err
				}
				// Assign new cart ID for the next step
				newCart.ID = cart.ID
			} else {
				return err
			}
		} else {
			// Cart with the same user and product found, update the quantity
			existingCart.Quantity += newCart.Quantity
			if err := tx.Save(&existingCart).Error; err != nil {
				return err
			}
			// Update request struct for existing cart
			newCart.ID = existingCart.ID
			newCart.Quantity = existingCart.Quantity
		}

		var request struct {
			CartID    uint `json:"cart_id" binding:"required"`
			ProductID uint `json:"product_id" binding:"required"`
			Quantity  uint `json:"quantity" binding:"required,min=1"`
		}

		// Use the existing cart if it exists, otherwise use the new cart
		request.CartID = newCart.ID
		request.Quantity = newCart.Quantity
		request.ProductID = newCart.ProductID

		var newCartProduct models.CartProduct
		cartProduct := middlewares.XSSMiddleware(&newCartProduct).(*models.CartProduct)

		errCP := tx.Where("cart_id = ? AND product_id = ?", request.CartID, request.ProductID).First(&cartProduct).Error

		if errCP != nil {
			if errCP == gorm.ErrRecordNotFound {
				// Create new entry
				cartProduct = &models.CartProduct{
					CartID:    request.CartID,
					ProductID: request.ProductID,
					Quantity:  request.Quantity,
				}
				if err := tx.Create(&cartProduct).Error; err != nil {
					return err
				}
			} else {
				return errCP
			}
		} else {
			// Update existing entry
			cartProduct.Quantity = newCart.Quantity
			if err := tx.Save(&cartProduct).Error; err != nil {
				return err
			}
		}
		return nil
	})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":     "server error",
			"statusCode": 500,
			"message":    "Failed to create cart",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":     "success",
		"statusCode": 200,
		"message":    "Cart created successfully",
		"cart_id": newCart.ID,
	})
}




func GetCartByUserID(c *fiber.Ctx) error {
	var user = c.Locals("user").(jwt.MapClaims)
	userID := user["id"].(float64)
	cartUser := models.SelectCartById(int(userID))
	resultCarts := make([]map[string]interface{}, len(cartUser))
	for i, cart := range cartUser {
		products := make([]map[string]interface{}, len(cart.Products))
		for j, product := range cart.Products {
			var cartProduct models.CartProduct
			if err := configs.DB.Where("cart_id = ? AND product_id = ?", cart.ID, product.ID).First(&cartProduct).Error; err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve quantity"})
			}
			var image *models.Image
			configs.DB.First(&image, "product_id = ?", product.ID)
			products[j] = map[string]interface{}{
				"id":         product.ID,
				"created_at": product.CreatedAt,
				"updated_at": product.UpdatedAt,
				"name":       product.Name,
				"price":      product.Price,
				"photo":      image.URL,
				"size":       product.Sizes,
				"color":      product.Colors,
				"quantity":   cartProduct.Quantity,
				"rating":     product.Rating,
			}
		}

		resultCarts[i] = map[string]interface{}{
			"id":          cart.ID,
			"user_id":     cart.UserID,
			"seller_id":   cart.SellerID,
			"seller_name": cart.Seller.Name,
			"created_at":  cart.CreatedAt,
			"updated_at":  cart.UpdatedAt,
			"products":    products,
		}
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":     "success",
		"statusCode": 200,
		"message":    "Show Cart successfully",
		"data": resultCarts,
	})
}
// func DeleteProductFromCart (c *fiber.Ctx) error {
// 	id, err := strconv.Atoi(c.Params("id"))
// 	if err != nil {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"status":     "bad request",
// 			"statusCode": 400,
// 			"message":    "Invalid ID format",
// 		})
// 	}
// }