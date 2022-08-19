package handlers

import (
	"strconv"
	booksdto "waystalk/dto/books"
	"waystalk/models"
	"waystalk/repositories"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type handler struct {
	BookRepository repositories.BookRepository
}

func HandlerBook(BookRepository repositories.BookRepository) *handler {
	return &handler{BookRepository} // to make access at main
}

func (h *handler) FindBooks(c *fiber.Ctx) error {
	books, err := h.BookRepository.FindBooks()
	if err != nil {
		return (c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		}))
	}

	var data []models.Book
	data = append(data, books...)

	return c.JSON(fiber.Map{
		"data": data,
	})
}

func (h *handler) GetBook(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	book, err := h.BookRepository.GetBook(int(id))
	if err != nil {
		return (c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		}))
	}

	return c.JSON(fiber.Map{
		"data": convertResponse(book),
	})
}

func (h *handler) CreateBook(c *fiber.Ctx) error {
	request := new(booksdto.CreateBookRequest) //take pattern data submission
	if err := c.BodyParser(request); err != nil {
		return err
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// data form pattern submit to pattern entity db user
	book := models.Book{
		Title:       request.Title,
		Description: request.Description,
		Price:       request.Price,
	}

	data, err := h.BookRepository.CreateBook(book)
	if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"data": convertResponse(data),
	})
}

func (h *handler) UpdateBook(c *fiber.Ctx) error {
	request := new(booksdto.UpdateBookRequest) //take pattern data submission
	if err := c.BodyParser(request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	id, _ := strconv.Atoi(c.Params("id"))

	book, err := h.BookRepository.GetBook(int(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if request.Title != "" {
		book.Title = request.Title
	}

	if request.Description != "" {
		book.Description = request.Description
	}

	if request.Price > 0 {
		book.Price = request.Price
	}

	data, err := h.BookRepository.UpdateBook(book)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"data": convertResponse(data),
	})
}

func (h *handler) DeleteBook(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	book, err := h.BookRepository.GetBook(int(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	data, err := h.BookRepository.DeleteBook(book)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"data": convertResponse(data),
	})
}

func convertResponse(u models.Book) booksdto.BookResponse {
	return booksdto.BookResponse{
		ID:          u.ID,
		Title:       u.Title,
		Description: u.Description,
		Price:       u.Price,
	}
}
