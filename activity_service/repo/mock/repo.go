package mock

import (
	"context"
	"time"

	"activity_service/activity"
	proto "activity_service/proto/activity"
)

type repo struct {

}

func (r repo) Create(ctx context.Context, activity *proto.Activity) (activityId int, err error) {

	return 1, nil
}

func (r repo) ById(ctx context.Context, id int) (activity *proto.Activity, err error) {

	return &proto.Activity{
		Id:           1,
		User:         &proto.User{Id: 1},
		ActivityType: proto.Activity_Sleep,
		Timestamp:    time.Now().Unix(),
		Duration:     int64(time.Hour),
	}, nil
}

func (r repo) AllByUser(ctx context.Context, userId int) (activities activity.Activities, err error) {
	return activity.Activities{
		{
			Id:           1,
			User:         &proto.User{Id: 1},
			ActivityType: proto.Activity_Sleep,
			Timestamp:    1617014265,
			Duration:     int64(time.Hour),
		},
		{
			Id:           2,
			User:         &proto.User{Id: 1},
			ActivityType: proto.Activity_Sleep,
			Timestamp:    1617016030,
			Duration:     int64(time.Hour),
		},
		{
			Id:           3,
			User:         &proto.User{Id: 1},
			ActivityType: proto.Activity_Sleep,
			Timestamp:    1617014204,
			Duration:     int64(time.Hour),
		},
	}, nil
}

func NewMock() activity.Repo {
	return &repo{}
}
