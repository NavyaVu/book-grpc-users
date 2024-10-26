package users

import (
	"context"
)

type UserService struct {
	UnimplementedUsersServer
}

func (u *UserService) GetUser(ctx context.Context, m *UserGetRequest) (*UserGetReply, error) {

	x := User{
		Id:        m.Id,
		FirstName: "Vinod",
		LastName:  "p",
		Age:       32,
	}
	return &UserGetReply{User: &x}, nil
}
