package dto

type BookReqDTO struct {
	AuthorID    int64  `json:"author_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type BookResDTO struct {
	Id          int64  `json:"id"`
	AuthorID    int64  `json:"author_id"`
	Author      string `json:"author"`
	Title       string `json:"title"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}
