package router

import (
	"github.com/gofiber/fiber/v2"
	notehandler "github.com/rahadirafsanjani/notes_app/internal/handler/notes"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	notes := api.Group("/notes")
	notes.Get("/", notehandler.GetNotes)       // GET /api/notes (List all notes)
	notes.Get("/:noteId", notehandler.GetNote) // GET /api/notes/:noteId (Get a single note)
}
