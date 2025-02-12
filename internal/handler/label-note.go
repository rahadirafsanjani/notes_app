package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rahadirafsanjani/notes_app/database"
	"github.com/rahadirafsanjani/notes_app/internal/model"
)

// Get notes by label ID using raw SQL
func GetNotesByLabel(c *fiber.Ctx) error {
	db := database.DB
	var notes []model.Notes

	// Read the labelId param
	labelID := c.Params("labelId")

	// Raw SQL Query
	query := "SELECT n.* FROM notes n INNER JOIN label_notes ln ON n.id = ln.note_id WHERE ln.label_id = ?"
	result := db.Raw(query, labelID).Scan(&notes)

	// Handle errors
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Database error",
			"error":   result.Error.Error(),
		})
	}

	// If no notes found, return 404
	if len(notes) == 0 {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "No notes found for this label",
		})
	}

	// Return response
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Notes found for the label",
		"data":    notes,
	})
}
