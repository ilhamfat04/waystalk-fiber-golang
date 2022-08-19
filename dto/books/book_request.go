package booksdto

type CreateBookRequest struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"desc" validate:"required"`
	Price       int    `json:"price" validate:"required"`
}

type UpdateBookRequest struct {
	Title       string `json:"title"`
	Description string `json:"desc"`
	Price       int    `json:"price"`
}
