# activity-logging

simple user activity logging app using cobra, grpc and docker-compose

This project stores user activities like Play, Sleep, Eat & Read with their durations. The project consist of 3 services.
I.e 2 hexagonal services and 1 Cobra cli client.

- User service
- Activity service
- Cobra client

User and activity service runs on port `:9001` & `:9002` ports respectively on gRPC server. Whereas, Cobra client interacts with these services as gRPC client.

### Run the project

`docker-compose build` will build the Dockerfiles of all 3 services

`docker-compose run client sh` will start the `Cobra cli` in interactive terminal.

### Cli commands

```text
+++++++++++++++++++++++++++++++++++++
a gRPC client to communicate with the UserActivity server.
        You can use this client to create and read user activities.

Usage:
  cli [command]

Available Commands:
  createactivity Create a new activity
  createuser     Create a new user
  help           Help about any command
  listactivities List all activities by this userid
  listusers      List all users
  oneactivity    Find a activity by its ID
  oneuser        Find a user by its ID

Flags:
  -h, --help   help for cli

Use "cli [command] --help" for more information about a command.
```

#### Create a new user

`./cli createuser -n "Test" -e "test@test.com" -p "+91 999 888 3333"`

```text
+++++++++++++++++++++++++++++++++++++
User created: 17
```

#### Show one user

`./cli go run main.go oneuser -i 17`

```text
+++++++++++++++++++++++++++++++++++++
#17 => Name: Test, Email: test@test.com, Phone: +91 999 888 3333
```

#### Show all users

`./cli listusers`

```text
+++++++++++++++++++++++++++++++++++++
#15 => Name: Testt, Email: test@test.com, Phone: 1234567890
#16 => Name: Testt, Email: test@test.com, Phone: 123 456 7890
#17 => Name: Test, Email: test@test.com, Phone: +91 999 888 3333
#18 => Name: Test, Email: test@test.com, Phone: +91 (999) 888 3333
#19 => Name: Test, Email: test@test.com, Phone: +91 999-888-3333
#20 => Name: Test, Email: test@test.com, Phone: 999-888-3333
```

#### Create a new activity

`./cli createactivity -u 18 -t Sleep -d 6h`

```text
+++++++++++++++++++++++++++++++++++++
Activity created: 6
```

#### Show one activity

`./cli oneactivity -i 6`

```text
+++++++++++++++++++++++++++++++++++++
#6 => User: 18 Test, Activity: Sleep, CreateAt: 30/03/2021, 08:40:43, Duration: 6h0m0s, Status: PENDING
```

#### Show all activities by a user

`./cli listactivities -u 4`

```text
+++++++++++++++++++++++++++++++++++++
#1 => User: 4 Test26, Activity: Play, CreateAt: 29/03/2021, 16:06:44, Duration: 1h0m0s, Status: COMPLETED
#3 => User: 4 Test26, Activity: Sleep, CreateAt: 29/03/2021, 16:37:10, Duration: 6h0m0s, Status: COMPLETED
#7 => User: 4 Test26, Activity: Eat, CreateAt: 30/03/2021, 08:43:49, Duration: 30m0s, Status: PENDING
```
