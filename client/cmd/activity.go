package cmd

import (
	"context"
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/spf13/cobra"

	proto "client/proto/activity"
)

// createCmd represents the create command
var createActivityCmd = &cobra.Command{
	Use:   "createactivity",
	Short: "Create a new activity",
	Long: `Create a new activity on the server through gRPC. 
	
	A activity requires a UserId, Type and Duration.`,
	RunE: func(cmd *cobra.Command, args []string) error {

		u, err := cmd.Flags().GetString("user")
		t, err := cmd.Flags().GetString("type")
		d, err := cmd.Flags().GetString("duration")
		if err != nil {
			return err
		}

		user, err := strconv.Atoi(u)
		if err != nil {
			return err
		}

		duration, err := time.ParseDuration(d)
		if err != nil {
			return err
		}

		typ, ok := proto.Activity_ActivityType_value[t] // get Activity int32 value
		if !ok {
			return fmt.Errorf("%s type is invalid, must be in [Play, Sleep, Eat, Read]", t)
		}

		a := &proto.Activity{
			User:  &proto.User{Id: int32(user)},
			ActivityType: 	proto.Activity_ActivityType(typ),
			Duration:  int64(duration),
		}
		activity, err := activityClient.CreateActivity(
			context.Background(),
			&proto.CreateActivityReq{
				Activity: a,
			},
		)
		if err != nil {
			return err
		}
		fmt.Printf("Activity created: %d\n", activity.Id)
		return nil
	},
}

// strActivity embeds proto.Activity struct to extend its functionality
type strActivity struct {
	*proto.Activity
}

// String formats the proto.Activity data embedded in strActivity
func (x strActivity) String() string {
	typ := proto.Activity_ActivityType_name[int32(x.ActivityType)] // get Activity string type

	return fmt.Sprintf("#%d => User: %d %s, Activity: %s, CreateAt: %s, Duration: %s, Status: %s",
		x.Id, x.User.Id, x.User.Name, typ, time.Unix(x.Timestamp, 0).Format("02/01/2006, 15:04:05"), time.Duration(x.Duration), x.Label)
}

// oneActivityCmd command will read one Activity by id
var oneActivityCmd = &cobra.Command{
	Use:   "oneactivity",
	Short: "Find a activity by its ID",
	Long: `Find a activity by it's unique identifier.
	
	If no activity is found for the ID it will return a 'Not Found' error`,
	RunE: func(cmd *cobra.Command, args []string) error {

		id, err := cmd.Flags().GetString("id")
		if err != nil {
			return err
		}

		actId, err := strconv.Atoi(id)
		if err != nil {
			return err
		}

		req := &proto.OneActivityReq{
			Id: int32(actId),
		}
		activity, err := activityClient.OneActivity(context.Background(), req)
		if err != nil {
			return err
		}

		fmt.Println(strActivity{activity})
		return nil
	},
}

// listActivityCmd command will stream all activities by user id
var listActivityCmd = &cobra.Command{
	Use:   "listactivities",
	Short: "List all activities by this userid",
	RunE: func(cmd *cobra.Command, args []string) error {

		u, err := cmd.Flags().GetString("userid")
		if err != nil {
			return err
		}

		user, err := strconv.Atoi(u)
		if err != nil {
			return err
		}

		req := &proto.ListUserActivitiesReq{
			UserId: int32(user),
		}

		stream, err := activityClient.ListUserActivities(context.Background(), req)

		if err != nil {
			return err
		}

		for {
			// stream.Recv returns a pointer to an Activity at the current iteration
			activity, err := stream.Recv()
			// If end of stream, break the loop
			if err == io.EOF {
				break
			}
			// if err, return an error
			if err != nil {
				return err
			}
			// If everything went well, print the activity object
			fmt.Println(strActivity{activity})
		}
		return nil
	},
}


func init() {
	createActivityCmd.Flags().StringP("user", "u", "", "A user of the activity")
	createActivityCmd.Flags().StringP("type", "t", "", "Type of activity. Must be in [`Play`,`Sleep`,`Eat`,`Read`]")
	createActivityCmd.Flags().StringP("duration", "d", "", "Add duration of the activity [`1m` minute, `1h` hour]")
	createActivityCmd.MarkFlagRequired("user")
	createActivityCmd.MarkFlagRequired("type")
	createActivityCmd.MarkFlagRequired("duration")
	rootCmd.AddCommand(createActivityCmd)

	oneActivityCmd.Flags().StringP("id", "i", "", "The id of the activity")
	oneActivityCmd.MarkFlagRequired("id")
	rootCmd.AddCommand(oneActivityCmd)

	listActivityCmd.Flags().StringP("userid", "u", "", "List of all activities by this user")
	listActivityCmd.MarkFlagRequired("userid")
	rootCmd.AddCommand(listActivityCmd)
}
