package activity

import (
	"context"

	proto "activity_service/proto/activity"
)

type Activities []*proto.Activity

// Repo Database Repository Interface provides methods to implemented by Database client
type Repo interface {
	Create(ctx context.Context, activity *proto.Activity) (activityId int, err error)
	ById(ctx context.Context, id int) (activity *proto.Activity, err error)
	AllByUser(ctx context.Context, userId int) (activities Activities, err error)
}
