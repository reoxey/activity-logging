package main_test

import (
	"context"
	"io"
	"log"
	"net"
	"testing"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"

	"activity_service/activity"
	proto "activity_service/proto/activity"
	"activity_service/repo/mock"
)

const buffer = 1024 * 1024 // buffer size

var l *bufconn.Listener

func bufferDialer(context.Context, string) (net.Conn, error) {
	return l.Dial()
}

func init() {
	// creates a dummy grpc server using buffer which emulates real connections
	l = bufconn.Listen(buffer)
	grpcOb := grpc.NewServer()

	dbRepo := mock.NewMock() // Mock DB Repo for CRUD operations

	s := activity.NewActivityService(dbRepo, nil) // user grpc client is set nil for this unit test

	proto.RegisterActivityServiceServer(grpcOb, s)

	go func() {
		if err := grpcOb.Serve(l); err != nil {
			log.Fatal(err)
		}
	}()
}

func TestUser(t *testing.T) {

	// connects to buffer grpc server
	conn, err := grpc.DialContext(context.Background(), "buffernet",
		grpc.WithContextDialer(bufferDialer),
		grpc.WithInsecure())
	if err != nil {
		t.Fatalf("failed to dial buffer net: %v", err)
	}
	defer conn.Close()

	client := proto.NewActivityServiceClient(conn)

	// Create a new Activity
	t.Run("Should create a new activity for user #1", func(t *testing.T) {

		a := &proto.Activity{
			User:         &proto.User{Id: 1},
			ActivityType: proto.Activity_Sleep,
			Duration:     int64(6 * time.Hour),
		}
		act, err := client.CreateActivity(
			context.Background(),
			&proto.CreateActivityReq{
				Activity: a,
			},
		)
		if err != nil {
			t.Errorf("failed to create a new activity for user #1: %v", err)
		}
		t.Logf("Activity created: %d\n", act.Id)
	})

	// Query an Activity by id
	t.Run("Should query an activity with id 1", func(t *testing.T) {

		req := &proto.OneActivityReq{
			Id: 1,
		}
		u, err := client.OneActivity(context.Background(), req)
		if err != nil {
			t.Errorf("failed to query an activity: %v", err)
		}

		t.Logf("Activity: %v\n", u)
	})

	// Stream all the Activities by a User
	t.Run("Should query all activities for a user 1", func(t *testing.T) {

		req := &proto.ListUserActivitiesReq{
			UserId: 1,
		}
		stream, err := client.ListUserActivities(context.Background(), req)
		if err != nil {
			t.Errorf("failed to query all activities: %v", err)
		}

		for {
			u, err := stream.Recv() // Stream data till EOF
			if err == io.EOF {
				break
			}
			if err != nil {
				t.Errorf("failed to stream: %v", err)
			}
			t.Logf("Activity: %v\n", u)
		}
	})
}
