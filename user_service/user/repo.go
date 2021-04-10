package user

import (
	"context"

	proto "user_service/proto/user"
)

type Users []*proto.User

// Repo Database Repository Interface provides methods to implemented by Database client
type Repo interface {
	Create(ctx context.Context, user *proto.User) (userId int, err error)
	ById(ctx context.Context, id int) (user *proto.User, err error)
	All(ctx context.Context) (users Users, err error)
}
