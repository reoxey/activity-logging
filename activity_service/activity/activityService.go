package activity

import (
	"context"
	"time"

	proto "activity_service/proto/activity"
	"activity_service/proto/user"
)

type activityService struct {
	actRepo Repo // Database repository interface
	grpcUser user.UserServiceClient // grpc client connection to User grpc server
}

func (a *activityService) CreateActivity(ctx context.Context, req *proto.CreateActivityReq) (*proto.Activity, error) {
	activity := req.GetActivity()

	// validates activity fields
	if err := validate(activity); err != nil {
		return nil, err
	}

	activity.Timestamp = time.Now().Unix() // Current Timestamp

	activityId, err := a.actRepo.Create(ctx, activity)  // insert command to DB
	if err != nil {
		return nil, err
	}
	activity.Id = int32(activityId)

	activity.User = getUser(ctx, a.grpcUser, activity.User.Id) // gets user data over RPC call

	activity.Label = getLabel(activity) // get activity label based on time duration

	return activity, nil
}

func (a *activityService) OneActivity(ctx context.Context, req *proto.OneActivityReq) (*proto.Activity, error) {
	activityId := req.GetId()

	activity, err := a.actRepo.ById(ctx, int(activityId))  // select 1 query to DB
	if err != nil {
		return nil, err
	}

	activity.User = getUser(ctx, a.grpcUser, activity.User.Id)

	activity.Label = getLabel(activity)

	return activity, nil
}

func (a *activityService) ListUserActivities(req *proto.ListUserActivitiesReq, stream proto.ActivityService_ListUserActivitiesServer) error {

	activities, err := a.actRepo.AllByUser(context.Background(), int(req.GetUserId()))  // select all query to DB
	if err != nil {
		return err
	}

	u := getUser(context.Background(), a.grpcUser, req.GetUserId())

	// grpc supports data streaming
	// chunks of data is streamed and received at client end
	for _, activity := range activities {
		activity.User = u
		activity.Label = getLabel(activity)
		stream.Send(activity)
	}

	return nil
}

// NewActivityService implements ActivityServiceServer Service built by protobuf
// core logic will reside in Service methods and db instructions will be delegated to Repo Interface
// grpc client will call Service methods
func NewActivityService(ar Repo, usc user.UserServiceClient) proto.ActivityServiceServer {
	return &activityService{
		actRepo: ar,
		grpcUser: usc,
	}
}
