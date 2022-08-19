package booksdto

type BookResponse struct {
	ID          int    `json:"id"`
	Title       string `json:"title" form:"name"`
	Description string `json:"desc" form:"email"`
	Price       int    `json:"price" form:"password"`
}
