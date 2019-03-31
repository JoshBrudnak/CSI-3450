# Budget Planner

Budget Planner is a simple budgeting application that allows users to create, configure, and monitor their budgets over time. The users will also be able to update their progress in the various budget categories and be provided with a simple A database is essential for this application in storing the application data. This will include the user’s metadata such as their name, password, and possibly settings. The application will also need to store the budget data for each user.

# Setup

### Server

1. Install [go](https://golang.org/doc/install).
2. Add database connection information to the [database.json](./Server/database.json).
3. Install necessary go packages using the [setupServer.sh](./Server/setupServer.sh).
4. Setup the test data `go run ./testdata/setupTestData.go`.
5. Startup the server `go run server.go endpoints.go`.
