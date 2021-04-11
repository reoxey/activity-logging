package main

import (
	"log"
	"net"
	"os"
	"os/signal"

	"google.golang.org/grpc"

	"activity_service/activity"
	proto "activity_service/proto/activity"
	"activity_service/proto/user"
	"activity_service/repo/mysql"
)

func main() {

	// environment variables
	grpcPort := os.Getenv("ACTIVITY_PORT")
	userGrpcPort := os.Getenv("USER_PORT")
	dsn := os.Getenv("DB_DSN")
	dbTable := os.Getenv("DB_TABLE")

	// network listener for grpc on a specified port
	l, err := net.Listen("tcp", grpcPort)
	if err != nil {
		log.Println(err)
	}
	defer l.Close()

	grpcOb := grpc.NewServer()

	// grpc connection for UserClient, connect to User Server grpc
	conn, err := grpc.Dial(userGrpcPort, grpc.WithInsecure())
	if err != nil {
		log.Println(err)
	}
	defer conn.Close()

	// Mysql repository which implements Repo interface. This can be replaced by any Database
	// without affecting the business logic implemented in Services.
	dbRepo, err := mysql.NewRepo(dsn, dbTable, 10)
	if err != nil {
		log.Fatal(err)
	}

	go func() {

		// Activity Service initialized with Repo, grpc calls will be intercepted by this service.
		// UserClient is passed in the service which is used to find user by id
		s := activity.NewActivityService(dbRepo, user.NewUserServiceClient(conn))

		proto.RegisterActivityServiceServer(grpcOb, s)

		log.Println("gRPC started "+grpcPort)
		if err = grpcOb.Serve(l); err != nil {
			log.Fatal(err)
		}
	}()

	// SIGTERM is received by the channel before terminated the application
	sigC := make(chan os.Signal)
	signal.Notify(sigC, os.Interrupt)
	signal.Notify(sigC, os.Kill)

	log.Println("Terminated", <-sigC)

	// ensures Graceful Stop before terminating
	grpcOb.GracefulStop()
}
