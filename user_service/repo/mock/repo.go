package mock

import (
	"context"

	proto "user_service/proto/user"
	"user_service/user"
)

type repo struct {

}

func (r *repo) Create(ctx context.Context, user *proto.User) (userId int, err error) {
	return 1, nil
}

func (r *repo) ById(ctx context.Context, id int) (user *proto.User, err error) {
	return &proto.User{
		Id:    1,
		Name:  "Test",
		Email: "test@test.com",
		Phone: "+91 (123) 456-7799",
	}, nil
}

func (r *repo) All(ctx context.Context) (users user.Users, err error) {
	return user.Users{
		{
			Id:    1,
			Name:  "Test",
			Email: "test@test.com",
			Phone: "+91 (123) 456-7799",
		},
		{
			Id:    2,
			Name:  "Test",
			Email: "test@test.com",
			Phone: "+91 (123) 456-7799",
		},
		{
			Id:    3,
			Name:  "Test",
			Email: "test@test.com",
			Phone: "+91 (123) 456-7799",
		},
	}, nil
}

func NewMock() user.Repo {
	return &repo{}
}
