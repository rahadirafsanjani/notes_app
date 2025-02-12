package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rahadirafsanjani/notes_app/internal/handler"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	// Notes routes
	notes := api.Group("/notes")
	notes.Get("/", handler.GetNotes)       // GET /api/notes (List all notes)
	notes.Get("/:noteId", handler.GetNote) // GET /api/notes/:noteId (Get a single note)

	// Label notes routes
	labels := api.Group("/labels")
	labels.Get("/:labelId/notes", handler.GetNotesByLabel) // GET /api/labels/:labelId/notes (Get notes by label)
}
