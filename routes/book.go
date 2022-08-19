package routes

import (
	"waystalk/handlers"
	"waystalk/pkg/mysql"
	"waystalk/repositories"

	"github.com/gofiber/fiber/v2"
)

func UserRoute(r fiber.Router) {
	bookRepository := repositories.RepositoryBook(mysql.DB)
	h := handlers.HandlerBook(bookRepository)

	r.Get("/books", h.FindBooks)
	r.Get("/book/:id", h.GetBook)
	r.Post("/book", h.CreateBook)
	r.Patch("/book/:id", h.UpdateBook)
	r.Delete("/book/:id", h.DeleteBook)
}
