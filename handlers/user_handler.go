package handlers

import (
	"fiber-api/config"
	"fiber-api/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func GetUsers(c *fiber.Ctx) error {
	//_, err := config.DB.Exec(c.Context(), "SET search_path TO account")

	rows, err := config.DB.Query(c.Context(), "SELECT id, user_name, user_email FROM account.users")
	if err != nil {
		log.Error(err)
		return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch users"})
	}
	defer rows.Close()

	users := []models.User{}

	for rows.Next() {
		var user models.User
		rows.Scan(&user.ID, &user.Name, &user.Email)
		users = append(users, user)
	}

	return c.JSON(users)
}

// GetUser mengambil user berdasarkan ID
func GetUser(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		log.Error("Invalid ID format")
		return c.Status(400).JSON(fiber.Map{"error": "Invalid user ID"})
	}

	var user models.User
	err = config.DB.QueryRow(c.Context(), "SELECT id, user_name, user_email FROM account.users WHERE id = $1", id).
		Scan(&user.ID, &user.Name, &user.Email)

	if err != nil {
		log.Error(err)
		return c.Status(404).JSON(fiber.Map{"error": "User not found"})
	}

	return c.JSON(user)
}

func CreateUser(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	_, err := config.DB.Exec(c.Context(), "INSERT INTO account.users (user_name, user_email) VALUES ($1, $2)", user.Name, user.Email)
	if err != nil {
		log.Error(err)
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create user"})
	}

	return c.Status(201).JSON(user)
}

func UpdateUser(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	_, err := config.DB.Exec(c.Context(), "UPDATE account.users SET user_name = $1, user_email = $2 WHERE id = $3", user.Name, user.Email, user.ID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to update users"})
	}

	return c.JSON(user)
}

func DeleteUser(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	_, err := config.DB.Exec(c.Context(), "DELETE FROM account.users WHERE id = $1", user.ID)
	if err != nil {
		log.Error(err)
		return c.Status(500).JSON(fiber.Map{"error": "Failed to delete users"})
	}

	return c.JSON(user)
}
