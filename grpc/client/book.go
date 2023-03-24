package client

import (
	"context"

	protoBook "github.com/AdiKhoironHasan/bookservice-protobank/proto/book"
	"github.com/AdiKhoironHasan/bookservices-api-gateway/domain/dto"
)

func (r GRPCClient) BooksPing(ctx context.Context) (*protoBook.PingRes, error) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	data, err := r.book.Ping(ctx, &protoBook.PingReq{})
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (r GRPCClient) BookList(ctx context.Context, dataReq *dto.BookReqDTO) (*protoBook.BookListRes, error) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	protoBookReq := &protoBook.BookListReq{
		Title: dataReq.Title,
	}

	books, err := r.book.List(ctx, protoBookReq)
	if err != nil {
		return nil, err
	}

	return books, nil
}

func (r GRPCClient) BookStore(ctx context.Context, bookReq *dto.BookReqDTO) (*protoBook.BookStoreRes, error) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	protoBookReq := &protoBook.BookStoreReq{
		AuthorId:    bookReq.AuthorID,
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

func (r GRPCClient) BookUpdate(ctx context.Context, bookReq *dto.BookReqDTO, id int) (*protoBook.BookDetailRes, error) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	protoBookReq := &protoBook.BookUpdateReq{
		Id:          int64(id),
		AuthorId:    bookReq.AuthorID,
		Title:       bookReq.Title,
		Description: bookReq.Description,
	}

	_, err := r.book.Update(ctx, protoBookReq)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (r GRPCClient) BookDelete(ctx context.Context, id int) (*protoBook.BookDeleteRes, error) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	bookReq := &protoBook.BookDeleteReq{
		Id: int64(id),
	}

	_, err := r.book.Delete(ctx, bookReq)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
