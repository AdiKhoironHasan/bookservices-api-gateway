package client

import (
	"context"

	protoBook "github.com/AdiKhoironHasan/bookservice-protobank/proto/book"
	"github.com/AdiKhoironHasan/bookservices-api-gateway/domain/dto"
)

func (r GRPCClient) BookList(ctx context.Context) (*protoBook.BookListRes, error) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	books, err := r.book.List(ctx, &protoBook.BookListReq{})
	if err != nil {
		return nil, err
	}

	return books, nil
}

func (r GRPCClient) BookStore(ctx context.Context, bookReq *dto.BookReqDTO) (*protoBook.BookStoreRes, error) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	protoBookReq := &protoBook.BookStoreReq{
		Title:       bookReq.Title,
		Description: bookReq.Description,
	}

	_, err := r.book.Store(ctx, protoBookReq)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (r GRPCClient) BookDetail(ctx context.Context, id int) (*protoBook.BookDetailRes, error) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	bookReq := &protoBook.BookDetailReq{
		Id: int64(id),
	}

	book, err := r.book.Detail(ctx, bookReq)
	if err != nil {
		return nil, err
	}

	return book, nil
}
