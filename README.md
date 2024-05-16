# Football Training Tracker API

This API allows users to create new training sessions for football teams, save player information, and retrieve team information. It provides endpoints for managing trainings, teams, and players.

## Installation

To run this API, you need to have Go installed on your machine. You can download and install it from the [official Go website](https://golang.org/dl/).

1. Clone the repository:

   `git clone git@github.com:jrealpe_meli/trainin-tracker.git`

2. Generate Swagger documentation:
   `swag init -g ./cmd/api/main.go -o cmd/api/docs`
3. Run `go mod tidy` to ensure all dependencies are up to date.
      
4. Run the application:
   `go run cmd/api/main.go`
5. Access the API documentation:
      You can access the Swagger documentation by navigating to [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html).

## Endpoints

### Training
- `POST /training`: Create a new training session.

### Team
- `GET /team`: Get information about a specific team.

### Players
- `POST /players`: Save player information.
- `GET /players`: Get a list of all players.

## Dependencies

The application uses a Mysql database to store player and training information. Make sure you have a PostgreSQL database running before starting the application.
### Validate MySQL database
`mysql --version`
#### If you do not have MySQL installed, please follow the installation guide below.
> [Installation guide]([https://github.com/melisource/fury_go-mini#supported-tags](https://www.devart.com/dbforge/mysql/how-to-install-mysql-on-macos/)).
#### The next step is to add a .env file to include the configuration variables for the DB.
### Example
>DB_NAME=name
DB_USER=trainingtracker
DB_PASSWORD=123456
DB_HOST=localhost
DB_PORT=3306
APP=trainingtracker
### The next step is to execute the script to create the database.
>trainin-tracker/database/script_db.sql
## 
## Database Configuration

To connect to the MySQL database in the cloud, you can use the following configuration variables in your .env file:

>DB_NAME=training_tracker
DB_USER=admin
DB_PASSWORD=12345678
DB_HOST=test-jr.c7y2ge44ua6a.us-east-2.rds.amazonaws.com
DB_PORT=3306
APP=trainingtracker
> 
Make sure to replace the values with your actual database credentials.

## Contributing

If you would like to contribute to this project, feel free to fork the repository and submit a pull request. Your contributions are highly appreciated.

Thank you for using Football Training Tracker API! 🚀🏈