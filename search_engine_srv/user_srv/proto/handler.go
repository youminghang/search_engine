package main

import (
	"context"
	user "github/youminghang/search_engine/search_engine_srv/user_srv/proto/kitex_gen/user"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// UserImpl implements the last service interface defined in the IDL.
type UserImpl struct{}

// GetUserList implements the UserImpl interface.
func (s *UserImpl) GetUserList(ctx context.Context, req *user.UserInfoFilter) (resp *user.UserListResponse, err error) {
	// TODO: Your code here...
	return
}

// GetUserById implements the UserImpl interface.
func (s *UserImpl) GetUserById(ctx context.Context, req *user.IdRequest) (resp *user.UserInfoResponse, err error) {
	// TODO: Your code here...
	return
}

// CreateUser implements the UserImpl interface.
func (s *UserImpl) CreateUser(ctx context.Context, req *user.CreateUserInfo) (resp *user.UserInfoResponse, err error) {
	// TODO: Your code here...
	return
}

// UpdateUser implements the UserImpl interface.
func (s *UserImpl) UpdateUser(ctx context.Context, req *user.UpdateUserInfo) (resp *emptypb.Empty, err error) {
	// TODO: Your code here...
	return
}

// CheckPassWord implements the UserImpl interface.
func (s *UserImpl) CheckPassWord(ctx context.Context, req *user.PasswordInfo) (resp *user.CheckResponse, err error) {
	// TODO: Your code here...
	return
}
