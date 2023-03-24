package handler

// import (
// 	"context"

// 	"github.com/AdiKhoironHasan/bookservices/domain/entity"
// 	"github.com/AdiKhoironHasan/bookservices/proto/book"
// )

// // Ping is a function
// func (c *Handler) List(_ context.Context, _ *book.BookRequest) (*book.BookResponse, error) {
// 	books := []entity.Book{}

// 	rows, err := c.repo.DB.Table("public.books").Select("id, title, description, created_at, updated_at").Rows()
// 	if err != nil {
// 		return nil, err
// 	}

// 	defer rows.Close()
// 	for rows.Next() {
// 		book := entity.Book{}
// 		rows.Scan(&book.ID, &book.Title, &book.Description, &book.CreatedAt, &book.UpdatedAt)
// 		books = append(books, book)
// 	}

// 	ch := make(chan []*book.Book)
// 	defer close(ch)

// 	go func(books []entity.Book, ch chan<- []*book.Book) {
// 		value := []*book.Book{}
// 		for _, val := range books {
// 			value = append(value, &book.Book{
// 				Id:          val.ID,
// 				Title:       val.Title,
// 				Description: val.Description,
// 				CreatedAt:   val.CreatedAt.String(),
// 				UpdatedAt:   val.UpdatedAt.String(),
// 			})
// 		}

// 		ch <- value
// 	}(books, ch)

// 	bookResp := <-ch

// 	res := &book.BookResponse{
// 		Books: bookResp,
// 	}

// 	return res, nil
// }
