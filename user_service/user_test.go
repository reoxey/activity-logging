package main_test

import (
	"context"
	"io"
	"log"
	"net"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"

	proto "user_service/proto/user"
	"user_service/repo/mock"
	"user_service/user"
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

	s := user.NewUserService(dbRepo)

	proto.RegisterUserServiceServer(grpcOb, s)

	go func() {
		if err := grpcOb.Serve(l); err != nil {
			log.Fatal(err)
		}
	}()
}

func TestUser(t *testing.T) {

	// connects to buffer grpc server
	conn, err := grpc.DialContext( context.Background(), "buffernet", grpc.WithContextDialer(bufferDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("failed to dial buffer net: %v", err)
	}
	defer conn.Close()

	client := proto.NewUserServiceClient(conn)

	// Create a new User
	t.Run("Should create a new user", func(t *testing.T) {

		u := &proto.User{
			Name: 	"Test",
			Email:  "test@test.com",
			Phone:  "+91 (123) 456-7799",
		}
		us, err := client.CreateUser(
			context.Background(),
			&proto.CreateUserReq{
				User: u,
			},
		)
		if err != nil {
			t.Errorf("failed to create a new user: %v", err)
		}
		t.Logf("User created: %d\n", us.Id)
	})

	// Query a User by id
	t.Run("Should query a user with id 1", func(t *testing.T) {

		req := &proto.OneUserReq{
			Id: 1,
		}
		u, err := client.OneUser(context.Background(), req)
		if err != nil {
			t.Errorf("failed to query a user: %v", err)
		}

		t.Logf("User: %v\n", u)
	})

	// Stream all the Users
	t.Run("Should query all user", func(t *testing.T) {

		req := &proto.ListUsersReq{}
		stream, err := client.ListUsers(context.Background(), req)
		if err != nil {
			t.Errorf("failed to query all users: %v", err)
		}

		for {
			u, err := stream.Recv() // Stream data till EOF
			if err == io.EOF {
				break
			}
			if err != nil {
				t.Errorf("failed to stream: %v", err)
			}
			t.Logf("User: %v\n", u)
		}
	})
}
