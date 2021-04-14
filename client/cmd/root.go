package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"

	"client/proto/activity"
	"client/proto/user"
)

// grpc clients set as global
var activityClient activity.ActivityServiceClient
var userClient user.UserServiceClient

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cli",
	Short: "a gRPC client to communicate with the UserActivity server",
	Long: `a gRPC client to communicate with the UserActivity server.
	You can use this client to create and read user activities.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		fmt.Println("+++++++++++++++++++++++++++++++++++++")
		os.Exit(1)
	}
}

func init()  {
	cobra.OnInitialize(initConfig) // initialize cobra with config

	conn, err := grpc.Dial(os.Getenv("USER_PORT"), grpc.WithInsecure())
	if err != nil {
		log.Println(err)
	}

	// Instantiate the NewUserServiceClient with our client connection to the server
	userClient = user.NewUserServiceClient(conn)

	conn, err = grpc.Dial(os.Getenv("ACTIVITY_PORT"), grpc.WithInsecure())
	if err != nil {
		log.Println(err)
	}

	// defer postpones the execution of a function until the surrounding function returns
	// conn.Close() will not be called until the end of main()
	// The arguments are evaluated immediately but not executed
	// defer conn.Close()

	// Instantiate the NewActivityServiceClient with our client connection to the server
	activityClient = activity.NewActivityServiceClient(conn)

	fmt.Println("+++++++++++++++++++++++++++++++++++++")
}

// initConfig reads in config file and ENV variables if set.
// no required in our case
func initConfig() {
}
