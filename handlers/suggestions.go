package handlers

import (
	"shopping-list/db"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// GetSuggestions returns item name suggestions for auto-completion
func GetSuggestions(c *fiber.Ctx) error {
	query := c.Query("q")
	limitStr := c.Query("limit", "10")

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		limit = 10
	}

	// If no query, return all suggestions (for offline cache)
	if query == "" {
		suggestions, err := db.GetAllItemSuggestions(limit)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch suggestions"})
		}
		if suggestions == nil {
			suggestions = []db.ItemSuggestion{}
		}
		return c.JSON(suggestions)
	}

	suggestions, err := db.GetItemSuggestions(query, limit)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch suggestions"})
	}

	if suggestions == nil {
		suggestions = []db.ItemSuggestion{}
	}

	return c.JSON(suggestions)
}
