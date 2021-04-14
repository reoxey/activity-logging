package cmd

import (
	"context"
	"fmt"
	"io"
	"strconv"

	"github.com/spf13/cobra"

	proto "client/proto/user"
)

// createCmd represents the create command
var createUserCmd = &cobra.Command{
	Use:   "createuser",
	Short: "Create a new user",
	Long: `Create a new user on the server through gRPC. 
	
	A user requires a Name, Email and Phone.`,
	RunE: func(cmd *cobra.Command, args []string) error {

		name, err := cmd.Flags().GetString("name")
		email, err := cmd.Flags().GetString("email")
		phone, err := cmd.Flags().GetString("phone")
		if err != nil {
			return err
		}

		u := &proto.User{
			Name: 	name,
			Email:  email,
			Phone:  phone,
		}
		user, err := userClient.CreateUser(
			context.Background(),
			&proto.CreateUserReq{
				User: u,
			},
		)
		if err != nil {
			return err
		}
		fmt.Printf("User created: %d\n", user.Id)
		return nil
	},
}

// strUser embeds proto.User struct to extend its functionality
type strUser struct {
	*proto.User
}

// String formats the proto.User data embedded in strUser
func (x strUser) String() string {
	return fmt.Sprintf("#%d => Name: %s, Email: %s, Phone: %s",
		x.Id, x.Name, x.Email, x.Phone)
}

// oneUserCmd command will read one User by id
var oneUserCmd = &cobra.Command{
	Use:   "oneuser",
	Short: "Find a user by its ID",
	Long: `Find a user by it's unique identifier.
	
	If no user is found for the ID it will return a 'Not Found' error`,
	RunE: func(cmd *cobra.Command, args []string) error {

		id, err := cmd.Flags().GetString("id")
		if err != nil {
			return err
		}

		userId, err := strconv.Atoi(id)
		if err != nil {
			return err
		}

		req := &proto.OneUserReq{
			Id: int32(userId),
		}

		user, err := userClient.OneUser(context.Background(), req)
		if err != nil {
			return err
		}

		fmt.Println(strUser{user})
		return nil
	},
}

// listUserCmd command will stream all Users
var listUserCmd = &cobra.Command{
	Use:   "listusers",
	Short: "List all users",
	RunE: func(cmd *cobra.Command, args []string) error {

		req := &proto.ListUsersReq{}

		stream, err := userClient.ListUsers(context.Background(), req)

		if err != nil {
			return err
		}

		for {
			// stream.Recv returns a pointer to a User at the current iteration
			user, err := stream.Recv()
			// If end of stream, break the loop
			if err == io.EOF {
				break
			}
			// if err, return an error
			if err != nil {
				return err
			}
			// If everything went well, print the user object
			fmt.Println(strUser{user})
		}
		return nil
	},
}


func init() {
	createUserCmd.Flags().StringP("name", "n", "", "Add a name")
	createUserCmd.Flags().StringP("email", "e", "", "An email of the user")
	createUserCmd.Flags().StringP("phone", "p", "", "Add a phone")
	createUserCmd.MarkFlagRequired("name")
	createUserCmd.MarkFlagRequired("email")
	createUserCmd.MarkFlagRequired("phone")
	rootCmd.AddCommand(createUserCmd)

	oneUserCmd.Flags().StringP("id", "i", "", "The id of the user")
	oneUserCmd.MarkFlagRequired("id")
	rootCmd.AddCommand(oneUserCmd)

	rootCmd.AddCommand(listUserCmd)
}
