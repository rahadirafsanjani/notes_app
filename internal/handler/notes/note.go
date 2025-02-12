package notehandler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rahadirafsanjani/notes_app/database"
	"github.com/rahadirafsanjani/notes_app/internal/model"
)

// Get all notes
func GetNotes(c *fiber.Ctx) error {
	db := database.DB
	var notes []model.Notes

	// Fetch all notes
	result := db.Find(&notes)
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Database error", "error": result.Error.Error()})
	}

	// If no notes exist, return 404
	if len(notes) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No notes found"})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "List of notes", "data": notes})
}

// Get a single note by ID
func GetNote(c *fiber.Ctx) error {
	db := database.DB
	var note model.Notes

	// Read the noteId param
	id := c.Params("noteId")

	// Find the note by ID
	result := db.First(&note, "id = ?", id)

	// If note not found
	if result.RowsAffected == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No note found"})
	}

	// If database error occurs
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Database error", "error": result.Error.Error()})
	}

	// Return the note
	return c.JSON(fiber.Map{"status": "success", "message": "Note found", "data": note})
}
