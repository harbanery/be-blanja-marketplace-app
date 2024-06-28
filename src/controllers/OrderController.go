package controllers

import (
	"gofiber-marketplace/src/configs"
	"gofiber-marketplace/src/helpers"
	"gofiber-marketplace/src/middlewares"
	"gofiber-marketplace/src/models"
	"gofiber-marketplace/src/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/veritrans/go-midtrans"
)

// func GetAllOrders(c *fiber.Ctx) error {
// 	orders := models.SelectAllOrders()
// 	if len(orders) == 0 {
// 		return c.Status(fiber.StatusOK).JSON(fiber.Map{
// 			"status":     "no content",
// 			"statusCode": 202,
// 			"message":    "Order is empty.",
// 		})
// 	}

// 	resultOrders := make([]*map[string]interface{}, len(orders))
// 	for i, order := range orders {
// 		address := map[string]interface{}{
// 			"id": order.Address.ID,
// 			"name": order.Address.ID,

// 		}
// 		resultOrders[i] = &map[string]interface{}{
// 			"id":           order.ID,
// 			"created_at":   order.CreatedAt,
// 			"updated_at":   order.UpdatedAt,
// 			"user_id":      order.UserID,
// 			"user":         order.User,
// 			"address_id":   order.AddressID,
// 			"address":      order.Address,
// 			"payment":      order.PaymentMethod,
// 			"total_amount": order.TotalPrice,
// 			"status":       order.Status,
// 		}
// 	}

// 	return c.Status(fiber.StatusOK).JSON(fiber.Map{
// 		"status":     "success",
// 		"statusCode": 200,
// 		"data":       resultOrders,
// 	})
// }

func GetAllOrders(c *fiber.Ctx) error {
	orders := models.SelectAllOrders()
	if len(orders) == 0 {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":     "no content",
			"statusCode": 202,
			"message":    "Order is empty.",
		})
	}

	resultOrders := make([]*map[string]interface{}, len(orders))
	for i, order := range orders {
		var checkout models.Checkout
		if err := configs.DB.Unscoped().Where("id = ?", order.CheckoutID).First(&checkout).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"status":     "error",
				"statusCode": 500,
				"message":    "Failed to get checkout.",
				"error":      err.Error(),
			})
		}

		var carts []models.Cart
		if err := configs.DB.Unscoped().Where("checkout_id = ?", checkout.ID).Find(&carts).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"status":     "error",
				"statusCode": 500,
				"message":    "Failed to get carts.",
				"error":      err.Error(),
			})
		}

		cartProducts := []map[string]interface{}{}
		for _, cart := range carts {
			var products []models.Product
			if err := configs.DB.Unscoped().Model(&cart).Association("Products").Find(&products); err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"status":     "error",
					"statusCode": 500,
					"message":    "Failed to get products.",
					"error":      err.Error(),
				})
			}

			for _, product := range products {
				var images []models.Image
				if err := configs.DB.Unscoped().Model(&product).Association("Images").Find(&images); err != nil {
					return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
						"status":     "error",
						"statusCode": 500,
						"message":    "Failed to get images.",
						"error":      err.Error(),
					})
				}

				var imageURL string
				if len(images) > 0 {
					imageURL = images[0].URL
				}

				seller := models.SelectSellerById(int(product.SellerID))

				cartProducts = append(cartProducts, map[string]interface{}{
					"id":          product.ID,
					"name":        product.Name,
					"price":       product.Price,
					"stock":       product.Stock,
					"description": product.Description,
					"seller_id":   product.SellerID,
					"seller_name": seller.Name,
					"image":       imageURL,
				})
			}
		}

		address := map[string]interface{}{
			"id":             order.Address.ID,
			"name":           order.Address.Name,
			"main_address":   order.Address.MainAddress,
			"detail_address": order.Address.DetailAddress,
			"phone":          order.Address.Phone,
			"postal_code":    order.Address.PostalCode,
			"city":           order.Address.City,
		}

		resultOrders[i] = &map[string]interface{}{
			"id":                 order.ID,
			"created_at":         order.CreatedAt,
			"updated_at":         order.UpdatedAt,
			"user_id":            order.UserID,
			"address_id":         order.AddressID,
			"address":            address,
			"transaction_number": order.TransactionNumber,
			"payment_method":     order.PaymentMethod,
			"total_amount":       order.TotalPrice,
			"status":             order.Status,
			"url":                order.TransactionURL,
			"products":           cartProducts,
		}
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":     "success",
		"statusCode": 200,
		"data":       resultOrders,
	})
}

func GetOrdersUser(c *fiber.Ctx) error {
	userId, err := middlewares.JWTAuthorize(c, "")
	if err != nil {
		if fiberErr, ok := err.(*fiber.Error); ok {
			return c.Status(fiberErr.Code).JSON(fiber.Map{
				"status":     fiberErr.Message,
				"statusCode": fiberErr.Code,
				"message":    fiberErr.Message,
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":     "Internal Server Error",
			"statusCode": fiber.StatusInternalServerError,
			"message":    err.Error(),
		})
	}

	orders := models.SelectOrdersbyUserId(int(userId))
	if len(orders) == 0 {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":     "no content",
			"statusCode": 202,
			"message":    "Order is empty.",
		})
	}

	resultOrders := make([]*map[string]interface{}, len(orders))
	for i, order := range orders {
		var checkout models.Checkout
		if err := configs.DB.Unscoped().Where("id = ?", order.CheckoutID).First(&checkout).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"status":     "error",
				"statusCode": 500,
				"message":    "Failed to get checkout.",
				"error":      err.Error(),
			})
		}

		var carts []models.Cart
		if err := configs.DB.Unscoped().Where("checkout_id = ?", checkout.ID).Find(&carts).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"status":     "error",
				"statusCode": 500,
				"message":    "Failed to get carts.",
				"error":      err.Error(),
			})
		}

		cartProducts := []map[string]interface{}{}
		for _, cart := range carts {
			var products []models.Product
			if err := configs.DB.Unscoped().Model(&cart).Association("Products").Find(&products); err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"status":     "error",
					"statusCode": 500,
					"message":    "Failed to get products.",
					"error":      err.Error(),
				})
			}

			for _, product := range products {
				var images []models.Image
				if err := configs.DB.Unscoped().Model(&product).Association("Images").Find(&images); err != nil {
					return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
						"status":     "error",
						"statusCode": 500,
						"message":    "Failed to get images.",
						"error":      err.Error(),
					})
				}

				var imageURL string
				if len(images) > 0 {
					imageURL = images[0].URL
				}

				seller := models.SelectSellerById(int(product.SellerID))

				cartProducts = append(cartProducts, map[string]interface{}{
					"id":          product.ID,
					"name":        product.Name,
					"price":       product.Price,
					"stock":       product.Stock,
					"description": product.Description,
					"seller_id":   product.SellerID,
					"seller_name": seller.Name,
					"image":       imageURL,
				})
			}
		}

		address := map[string]interface{}{
			"id":             order.Address.ID,
			"name":           order.Address.Name,
			"main_address":   order.Address.MainAddress,
			"detail_address": order.Address.DetailAddress,
			"phone":          order.Address.Phone,
			"postal_code":    order.Address.PostalCode,
			"city":           order.Address.City,
		}

		resultOrders[i] = &map[string]interface{}{
			"id":                 order.ID,
			"created_at":         order.CreatedAt,
			"updated_at":         order.UpdatedAt,
			"user_id":            order.UserID,
			"address_id":         order.AddressID,
			"address":            address,
			"transaction_number": order.TransactionNumber,
			"payment_method":     order.PaymentMethod,
			"total_amount":       order.TotalPrice,
			"status":             order.Status,
			"url":                order.TransactionURL,
			"products":           cartProducts,
		}
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":     "success",
		"statusCode": 200,
		"data":       resultOrders,
	})
}

func CreateOrder(c *fiber.Ctx) error {
	userId, err := middlewares.JWTAuthorize(c, "")
	if err != nil {
		if fiberErr, ok := err.(*fiber.Error); ok {
			return c.Status(fiberErr.Code).JSON(fiber.Map{
				"status":     fiberErr.Message,
				"statusCode": fiberErr.Code,
				"message":    fiberErr.Message,
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":     "Internal Server Error",
			"statusCode": fiber.StatusInternalServerError,
			"message":    err.Error(),
		})
	}

	existUser := models.SelectUserById(int(userId))
	if existUser.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":     "not found",
			"statusCode": 404,
			"message":    "User not found",
		})
	}

	var newOrder models.Order

	if err := c.BodyParser(&newOrder); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":     "bad request",
			"statusCode": 400,
			"message":    "Invalid request body",
		})
	}
	transactionNumber := helpers.GenerateTransactionNumber()

	existCheckout := models.SelectCheckoutByIdAndUserId(int(newOrder.CheckoutID), int(existUser.ID))
	if existCheckout.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":     "not found",
			"statusCode": 404,
			"message":    "Checkout not found",
		})
	}

	newOrder.UserID = existUser.ID
	newOrder.Status = "not_yet_paid"
	newOrder.TransactionNumber = transactionNumber

	carts := models.SelectCartbyCheckoutID(int(newOrder.CheckoutID))
	if len(carts) == 0 {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":     "no content",
			"statusCode": 202,
			"message":    "Order is empty.",
		})
	}

	var items []midtrans.ItemDetail
	var totalPrice int64

	for _, cart := range carts {
		product := models.SelectProductById(int(cart.ProductID))
		if product.ID == 0 {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"status":     "not found",
				"statusCode": 404,
				"message":    "Product not found",
			})
		}

		itemTotalPrice := int64(product.Price) * int64(cart.Quantity)
		items = append(items, midtrans.ItemDetail{
			Price: int64(product.Price),
			Qty:   int32(cart.Quantity),
			Name:  product.Name,
		})
		totalPrice += itemTotalPrice
	}

	taxAmount := int64(existCheckout.Delivery)
	items = append(items, midtrans.ItemDetail{
		Price: taxAmount,
		Qty:   1,
		Name:  "Tax",
	})
	totalPrice += taxAmount

	newOrder.TotalPrice = float64(totalPrice)

	order := middlewares.XSSMiddleware(&newOrder).(*models.Order)

	if errors := helpers.StructValidation(order); len(errors) > 0 {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"status":     "unprocessable entity",
			"statusCode": 422,
			"message":    "Validation failed",
			"errors":     errors,
		})
	}

	if err := models.CreateOrder(order); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":     "server error",
			"statusCode": 500,
			"message":    "Failed to create order",
		})
	}

	var userInfo struct {
		Name  string
		Phone string
	}

	if existUser.Role == "seller" {
		seller := models.SelectSellerByUserId(int(existUser.ID))
		userInfo.Name = seller.Name
		userInfo.Phone = seller.Phone
	} else if existUser.Role == "customer" {
		customer := models.SelectCustomerByUserId(int(existUser.ID))
		userInfo.Name = customer.Name
		userInfo.Phone = customer.Phone
	}

	address := models.SelectAddressbyId(int(order.AddressID))

	snapGateway := midtrans.SnapGateway{
		Client: services.Client,
	}

	CustomerAddress := &midtrans.CustAddress{
		FName:    address.Name,
		Phone:    address.Phone,
		Address:  address.MainAddress,
		City:     address.City,
		Postcode: address.PostalCode,
	}

	snapReq := &midtrans.SnapReq{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  transactionNumber,
			GrossAmt: totalPrice,
		},
		CustomerDetail: &midtrans.CustDetail{
			FName:    userInfo.Name,
			Email:    existUser.Email,
			Phone:    userInfo.Phone,
			BillAddr: CustomerAddress,
			ShipAddr: CustomerAddress,
		},
		Items: &items,
	}

	snapResp, err := snapGateway.GetToken(snapReq)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":     "server error",
			"statusCode": 404,
			"message":    "Failed to create transaction with Midtrans",
			"error":      err.Error(),
		})
	}

	if err := models.DeleteCartsByCheckoutID(int(order.CheckoutID)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":     "server error",
			"statusCode": 500,
			"message":    "Failed to delete carts",
		})
	}

	if err := models.DeleteCheckout(int(order.CheckoutID)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":     "server error",
			"statusCode": 500,
			"message":    "Failed to delete checkout",
		})
	}

	if err := models.UpdateURLOrder(int(order.ID), snapResp.RedirectURL); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":     "server error",
			"statusCode": 500,
			"message":    "Failed to input URL to order",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":       "success",
		"statusCode":   201,
		"message":      "Order created successfully. Please paid as soon as possible.",
		"token":        snapResp.Token,
		"redirect_url": snapResp.RedirectURL,
	})
}

func HandlePaymentCallback(c *fiber.Ctx) error {
	var notificationPayload models.Notification
	if err := c.BodyParser(&notificationPayload); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	orderID := notificationPayload.OrderID

	existOrder := models.SelectOrderbyTransactionNumber(orderID)
	if existOrder.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":     "not found",
			"statusCode": 404,
			"message":    "Order not found",
		})
	}

	var paymentMethod string
	if len(notificationPayload.VaNumbers) > 0 {
		paymentMethod = notificationPayload.PaymentType + "-" + notificationPayload.VaNumbers[0].Bank
	} else {
		paymentMethod = notificationPayload.PaymentType
	}

	if notificationPayload.TransactionStatus == "pending" {
		if err := models.UpdateStatusOrder(int(existOrder.ID), "not_yet_paid"); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"status":     "server error",
				"statusCode": 500,
				"message":    "Failed to update status order",
			})
		}
	}

	if notificationPayload.TransactionStatus == "settlement" {
		if err := models.UpdateStatusOrder(int(existOrder.ID), "get_paid"); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"status":     "server error",
				"statusCode": 500,
				"message":    "Failed to update status order",
			})
		}
		if err := models.UpdateURLOrder(int(existOrder.ID), ""); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"status":     "server error",
				"statusCode": 500,
				"message":    "Failed to input URL to order",
			})
		}
	}

	if notificationPayload.TransactionStatus == "cancel" {
		if err := models.UpdateStatusOrder(int(existOrder.ID), "canceled"); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"status":     "server error",
				"statusCode": 500,
				"message":    "Failed to update status order",
			})
		}
		if err := models.UpdateURLOrder(int(existOrder.ID), ""); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"status":     "server error",
				"statusCode": 500,
				"message":    "Failed to input URL to order",
			})
		}
	}

	if notificationPayload.TransactionStatus == "expire" {
		if err := models.UpdateStatusOrder(int(existOrder.ID), "expired"); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"status":     "server error",
				"statusCode": 500,
				"message":    "Failed to update status order",
			})
		}
		if err := models.UpdateURLOrder(int(existOrder.ID), ""); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"status":     "server error",
				"statusCode": 500,
				"message":    "Failed to input URL to order",
			})
		}
	}

	if err := models.UpdatePaymentMethodOrder(int(existOrder.ID), paymentMethod); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":     "server error",
			"statusCode": 500,
			"message":    "Failed to update status order",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":     "ok",
		"statusCode": 200,
		"message":    "Transaction payment and status updated.",
	})
}

func UpdateStatusOrder(c *fiber.Ctx) error {
	userId, err := middlewares.JWTAuthorize(c, "seller")
	if err != nil {
		if fiberErr, ok := err.(*fiber.Error); ok {
			return c.Status(fiberErr.Code).JSON(fiber.Map{
				"status":     fiberErr.Message,
				"statusCode": fiberErr.Code,
				"message":    fiberErr.Message,
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":     "Internal Server Error",
			"statusCode": fiber.StatusInternalServerError,
			"message":    err.Error(),
		})
	}

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":     "bad request",
			"statusCode": 400,
			"message":    "Invalid ID format",
		})
	}

	existorder := models.SelectOrderbyId(id)
	if existorder.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":     "not found",
			"statusCode": 404,
			"message":    "Order not found",
		})
	}

	if existorder.UserID != uint(userId) {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"status":     "forbidden",
			"statusCode": 403,
			"message":    "This order is not same user",
		})
	}

	if existorder.Status == "not_yet_paid" || existorder.Status == "expired" || existorder.Status == "canceled" {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"status":     "forbidden",
			"statusCode": 403,
			"message":    "This order should already paid",
		})
	}

	var updatedOrder models.Order

	if err := c.BodyParser(&updatedOrder); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":     "bad request",
			"statusCode": 400,
			"message":    "Invalid request body",
		})
	}

	order := middlewares.XSSMiddleware(&updatedOrder).(*models.Order)

	if authErrors := helpers.FieldRequiredValidation(order.Status, "required,oneof=not_yet_paid expired get_paid processed sent completed canceled"); authErrors != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"status":     "unprocessable entity",
			"statusCode": 422,
			"message":    "Validation failed",
			"errors":     authErrors,
		})
	}

	if err := models.UpdateStatusOrder(id, order.Status); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":     "server error",
			"statusCode": 500,
			"message":    "Failed to update status order",
		})
	} else {
		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"status":     "success",
			"statusCode": 200,
			"message":    "Status order updated successfully",
		})
	}
}

func DeleteOrder(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":     "bad request",
			"statusCode": 400,
			"message":    "Invalid ID format",
		})
	}

	order := models.SelectOrderbyId(id)
	if order.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":     "not found",
			"statusCode": 404,
			"message":    "order not found",
		})
	}

	if err := models.DeleteOrder(id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":     "server error",
			"statusCode": 500,
			"message":    "Failed to delete order",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":     "success",
		"statusCode": 200,
		"message":    "Order deleted successfully",
	})
}
