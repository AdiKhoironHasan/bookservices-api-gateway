package client

import (
	"context"

	protoUser "github.com/AdiKhoironHasan/bookservice-protobank/proto/user"
	"github.com/AdiKhoironHasan/bookservices-api-gateway/domain/dto"
)

func (r GRPCClient) UserPing(ctx context.Context) (*protoUser.PingRes, error) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	data, err := r.user.Ping(ctx, &protoUser.PingReq{})
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (r GRPCClient) UserList(ctx context.Context, dataReq *dto.UserReqDTO) (*protoUser.UserListRes, error) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	protoUserReq := &protoUser.UserListReq{
		Name: dataReq.Name,
		Role: dataReq.Role,
	}

	User, err := r.user.List(ctx, protoUserReq)
	if err != nil {
		return nil, err
	}

	return User, nil
}

func (r GRPCClient) Usertore(ctx context.Context, UserReq *dto.UserReqDTO) (*protoUser.UserStoreRes, error) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	protoUserReq := &protoUser.UserStoreReq{
		Name: UserReq.Name,
		Role: UserReq.Role,
	}

	_, err := r.user.Store(ctx, protoUserReq)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (r GRPCClient) UserDetail(ctx context.Context, id int) (*protoUser.UserDetailRes, error) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	UserReq := &protoUser.UserDetailReq{
		Id: int64(id),
	}

	User, err := r.user.Detail(ctx, UserReq)
	if err != nil {
		return nil, err
	}

	return User, nil
}

func (r GRPCClient) UserUpdate(ctx context.Context, UserReq *dto.UserReqDTO, id int) (*protoUser.UserDetailRes, error) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	protoUserReq := &protoUser.UserUpdateReq{
		Id:   int64(id),
		Name: UserReq.Name,
		Role: UserReq.Role,
	}

	_, err := r.user.Update(ctx, protoUserReq)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (r GRPCClient) UserDelete(ctx context.Context, id int) (*protoUser.UserDeleteRes, error) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	UserReq := &protoUser.UserDeleteReq{
		Id: int64(id),
	}

	_, err := r.user.Delete(ctx, UserReq)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
