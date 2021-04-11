package activity

import (
	"context"
	"fmt"
	"time"

	proto "activity_service/proto/activity"
	"activity_service/proto/user"
)

type Sleep struct {
	*proto.Activity
}
type Play struct {
	*proto.Activity
}
type Eat struct {
	*proto.Activity
}
type Read struct {
	*proto.Activity
}

// IActivity interface provides methods for validation and completion of an Activity
type IActivity interface {
	isDone() bool
	isValid() error
}

// validate an Activity's fields
func validate(a *proto.Activity) error {

	if a.User.Id == 0 {
		return fmt.Errorf("invalid user id %d", a.User.Id)
	}

	switch a.ActivityType {
	case proto.Activity_Sleep:
		return Sleep{a}.isValid()
	case proto.Activity_Play:
		return Play{a}.isValid()
	case proto.Activity_Eat:
		return Eat{a}.isValid()
	case proto.Activity_Read:
		return Read{a}.isValid()
	}

	return nil
}

// getLabel returns label based on Activity duration completion
func getLabel(a *proto.Activity) string {

	l := false

	switch a.ActivityType {
	case proto.Activity_Sleep:
		l = Sleep{a}.isDone()
	case proto.Activity_Play:
		l = Play{a}.isDone()
	case proto.Activity_Eat:
		l = Eat{a}.isDone()
	case proto.Activity_Read:
		l = Read{a}.isDone()
	}

	if l {
		return "COMPLETED"
	}
	return "PENDING"
}

func (p Play) isDone() bool {
	return diff(p.Timestamp, p.Duration)
}
func (s Sleep) isDone() bool {
	return diff(s.Timestamp, s.Duration)
}
func (e Eat) isDone() bool {
	return diff(e.Timestamp, e.Duration)
}
func (r Read) isDone() bool {
	return diff(r.Timestamp, r.Duration)
}

// diff checks difference current time and creation time which in turns compared with
// duration of the activity
// returns true if difference is greater
func diff(t, dura int64) bool {
	d := time.Now().Unix() - t

	if d * int64(time.Second) < dura {
		return false
	}

	return true
}

// Sleep duration validation; usually 6h - 8h
func (s Sleep) isValid() error {

	d := time.Duration(s.Duration)

	fmt.Println(d, 6*time.Hour)

	if d < time.Hour * 6 || d > time.Hour * 8 {
		return fmt.Errorf("sleep duration is invalid: %s", d)
	}

	return nil
}

// Play duration validation; 10m - 2h
func (p Play) isValid() error {

	d := time.Duration(p.Duration)

	if d < time.Minute * 10 || d > time.Hour * 2 {
		return fmt.Errorf("play duration is invalid: %s", d)
	}

	return nil
}

// Eat duration validation; 5m - 1h
func (e Eat) isValid() error {

	d := time.Duration(e.Duration)

	if d < time.Minute * 5 || d > time.Hour * 1 {
		return fmt.Errorf("eat duration is invalid: %s", d)
	}

	return nil
}

// Read duration validation; 15m - 2h
func (r Read) isValid() error {

	d := time.Duration(r.Duration)

	if d < time.Minute * 15 || d > time.Hour * 2 {
		return fmt.Errorf("read duration is invalid: %s", d)
	}

	return nil
}

// getUser retrieve User info from user_service over RPC call
func getUser(ctx context.Context, grpcUser user.UserServiceClient, userId int32) *proto.User {

	// can occur only in unit testing
	if grpcUser == nil {
		return &proto.User{Id: userId}
	}

	u, err := grpcUser.OneUser(ctx, &user.OneUserReq{
		Id: userId,
	})
	// if user service is down then send the user id to client
	// avoids cascade failure due dependent service unavailability
	if err != nil {
		return &proto.User{Id: userId}
	}

	return &proto.User{
		Id:    u.GetId(),
		Name:  u.GetName(),
		Email: u.GetEmail(),
		Phone: u.GetPhone(),
	}
}
