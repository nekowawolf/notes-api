package url

import (
	"github.com/nekowawolf/notes-api/controller"

	"github.com/gofiber/fiber/v2"
)

func Web(page *fiber.App) {

	page.Get("/", controller.Sink)
	page.Post("/", controller.Sink)
	page.Put("/", controller.Sink)
	page.Patch("/", controller.Sink)
	page.Delete("/", controller.Sink)
	page.Options("/", controller.Sink)
	
	page.Get("/notes", controller.GetAllNotes)
	page.Get("/notes/:id", controller.GetNotesByID)
	page.Post("/insert", controller.InsertNotes)
	page.Put("/update/:id", controller.UpdateNotesByID)
	page.Delete("/delete/:id", controller.DeleteNotesByID)
}
