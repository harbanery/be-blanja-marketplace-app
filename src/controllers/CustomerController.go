package controllers

import (
	"gofiber-marketplace/src/helpers"
	"gofiber-marketplace/src/middlewares"
	"gofiber-marketplace/src/models"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

func GetCustomers(c *fiber.Ctx) error {
	customers := models.SelectAllCustomers()
	if len(customers) == 0 {
		return c.Status(fiber.StatusNoContent).JSON(fiber.Map{
			"status":     "no content",
			"statusCode": 202,
			"message":    "Customer is empty.",
		})
	}

	resultCustomers := make([]map[string]interface{}, len(customers))
	for i, customer := range customers {

		resultCustomers[i] = map[string]interface{}{
			"id":            customer.ID,
			"created_at":    customer.CreatedAt,
			"updated_at":    customer.UpdatedAt,
			"name":          customer.Name,
			"user_id":       customer.User.ID,
			"email":         customer.User.Email,
			"image":         customer.Image,
			"phone":         customer.Phone,
			"gender":        customer.Gender,
			"date_of_birth": customer.DateOfBirth,
			"role":          customer.User.Role,
		}
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":     "success",
		"statusCode": 200,
		"data":       resultCustomers,
	})
}

func GetDetailCustomer(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":     "bad request",
			"statusCode": 400,
			"message":    "Invalid ID format",
		})
	}

	customer := models.SelectCustomerById(id)
	if customer.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":     "not found",
			"statusCode": 404,
			"message":    "Customer not found",
		})
	}

	if customer.User.Role != "customer" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":     "bad request",
			"statusCode": 400,
			"message":    "Role of this user is not customer",
		})
	}

	resultCustomer := map[string]interface{}{
		"id":            customer.ID,
		"created_at":    customer.CreatedAt,
		"updated_at":    customer.UpdatedAt,
		"name":          customer.Name,
		"user_id":       customer.User.ID,
		"email":         customer.User.Email,
		"image":         customer.Image,
		"phone":         customer.Phone,
		"gender":        customer.Gender,
		"date_of_birth": customer.DateOfBirth,
		"role":          customer.User.Role,
	}

	// return c.JSON(product)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":     "success",
		"statusCode": 200,
		"data":       resultCustomer,
	})
}

func GetCustomerProfile(c *fiber.Ctx) error {
	id, err := middlewares.JWTAuthorize(c, "customer")
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

	customer := models.SelectCustomerByUserId(int(id))
	if customer.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":     "not found",
			"statusCode": 404,
			"message":    "Customer not found",
		})
	}

	resultAddresses := models.SelectAddressesbyUserID(int(id))
	addresses := make([]map[string]interface{}, len(resultAddresses))
	for i, address := range resultAddresses {
		addresses[i] = map[string]interface{}{
			"id":             address.ID,
			"created_at":     address.CreatedAt,
			"updated_at":     address.UpdatedAt,
			"main_address":   address.MainAddress,
			"detail_address": address.DetailAddress,
			"postal_code":    address.PostalCode,
			"name":           address.Name,
			"phone":          address.Phone,
			"primary":        address.Primary,
			"city":           address.City,
		}
	}

	resultCustomer := map[string]interface{}{
		"id":            customer.ID,
		"created_at":    customer.CreatedAt,
		"updated_at":    customer.UpdatedAt,
		"name":          customer.Name,
		"user_id":       customer.User.ID,
		"email":         customer.User.Email,
		"image":         customer.Image,
		"phone":         customer.Phone,
		"gender":        customer.Gender,
		"date_of_birth": customer.DateOfBirth.Format("2006-01-02"),
		"role":          customer.User.Role,
		"addresses":     addresses,
	}

	// return c.JSON(product)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":     "success",
		"statusCode": 200,
		"data":       resultCustomer,
	})
}

func UpdateCustomerProfile(c *fiber.Ctx) error {
	var profileData models.CustomerProfile

	id, err := middlewares.JWTAuthorize(c, "customer")
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

	customer := models.SelectCustomerByUserId(int(id))
	if customer.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":     "not found",
			"statusCode": 404,
			"message":    "Customer not found",
		})
	}

	if err := c.BodyParser(&profileData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":     "bad request",
			"statusCode": 400,
			"message":    "Invalid request body",
		})
	}

	user := middlewares.XSSMiddleware(&profileData).(*models.CustomerProfile)
	if errors := helpers.StructValidation(user); len(errors) > 0 {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"status":     "unprocessable entity",
			"statusCode": 422,
			"message":    "Validation failed",
			"errors":     errors,
		})
	}

	if customer.User.Email != user.Email {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":     "bad request",
			"statusCode": 400,
			"message":    "Email already exists",
		})
	}

	parsedDate, err := time.Parse("2006-01-02", user.DateOfBirth)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":     "bad request",
			"statusCode": 400,
			"message":    "Date format invalid",
		})
	}

	updatedUser := models.User{
		Email: user.Email,
	}

	updatedCustomer := models.Customer{
		Name:        user.Name,
		Image:       user.Image,
		Phone:       user.Phone,
		Gender:      user.Gender,
		DateOfBirth: parsedDate,
	}

	if err := models.UpdateUser(int(id), &updatedUser); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":     "server error",
			"statusCode": 500,
			"message":    "Failed to update user",
		})
	}

	if err := models.UpdateCustomer(int(customer.ID), &updatedCustomer); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":     "server error",
			"statusCode": 500,
			"message":    "Failed to update customer",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":     "success",
		"statusCode": 200,
		"message":    "Profile updated successfully",
	})
}

func DeleteCustomer(c *fiber.Ctx) error {
	id, err := middlewares.JWTAuthorize(c, "customer")
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

	customer := models.SelectCustomerByUserId(int(id))
	if customer.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":     "not found",
			"statusCode": 404,
			"message":    "Customer not found",
		})
	}

	if err := models.DeleteCustomer(int(customer.ID)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":     "server error",
			"statusCode": 500,
			"message":    "Failed to delete customer",
		})
	}

	if err := models.DeleteUser(int(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":     "server error",
			"statusCode": 500,
			"message":    "Failed to delete user",
		})
	} else {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":     "success",
			"statusCode": 200,
			"message":    "User deleted successfully",
		})
	}
}
