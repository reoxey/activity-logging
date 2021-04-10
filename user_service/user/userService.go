package user

import (
	"context"

	proto "user_service/proto/user"
)

type userService struct {
	userRepo 	Repo // Database repository interface
}

func (u *userService) CreateUser(ctx context.Context, req *proto.CreateUserReq) (*proto.User, error) {
	user := req.GetUser()

	v := validateUser{user}

	// validates user fields; name, email & phone
	if err := v.Validate(); err != nil {
		return nil, err
	}

	userId, err := u.userRepo.Create(ctx, user) // insert command to DB
	if err != nil {
		return nil, err
	}
	user.Id = int32(userId)

	return user, nil
}

func (u *userService) OneUser(ctx context.Context, req *proto.OneUserReq) (*proto.User, error) {
	userId := req.GetId()

	user, err := u.userRepo.ById(ctx, int(userId))  // select 1 query to DB
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userService) ListUsers(req *proto.ListUsersReq, stream proto.UserService_ListUsersServer) error {

	users, err := u.userRepo.All(context.Background()) // select all query by user to DB
	if err != nil {
		return err
	}

	// grpc supports data streaming
	// chunks of data is streamed and received at client end
	for _, user := range users {
		stream.Send(user)
	}

	return nil
}

// NewUserService implements UserServiceServer Service built by protobuf
// core logic will reside in Service methods and db instructions will be delegated to Repo Interface
// grpc client will call Service methods
func NewUserService(ur Repo) proto.UserServiceServer {

	return &userService{
		userRepo: ur,
	}
}
