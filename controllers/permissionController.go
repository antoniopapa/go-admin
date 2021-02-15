package controllers

import (
	"../database"
	"../models"
	"github.com/gofiber/fiber"
)

func AllPermissions(c *fiber.Ctx) error {
	var permissions []models.Permission

	database.DB.Find(&permissions)

	return c.JSON(permissions)
}
