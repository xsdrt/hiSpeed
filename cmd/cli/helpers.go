package main

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/joho/godotenv"
)

func setup() {
	err := godotenv.Load()
	if err != nil {
		exitGracefully(err)
	}

	path, err := os.Getwd() //need to get the root path of the web app thats using HiSpeed
	if err != nil {
		exitGracefully(err)
	}

	his.RootPath = path
	his.DB.DataType = os.Getenv("DATABASE_TYPE")
}

func getDSN() string {
	dbType := his.DB.DataType

	if dbType == "pgx" {
		dbType = "postgres"
	}

	if dbType == "postgres" {
		var dsn string
		if os.Getenv("DATABASE_PASS") != "" {
			dsn = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
				os.Getenv("DATABASE_USER"),
				os.Getenv("DATABASE_PASS"),
				os.Getenv("DATABASE_HOST"),
				os.Getenv("DATABASE_PORT"),
				os.Getenv("DATABASE_NAME"),
				os.Getenv("DATABASE_SSL_MODE"))
		} else {
			dsn = fmt.Sprintf("postgres://%s@%s:%s/%s?sslmode=%s",
				os.Getenv("DATABASE_USER"),
				os.Getenv("DATABASE_HOST"),
				os.Getenv("DATABASE_PORT"),
				os.Getenv("DATABASE_NAME"),
				os.Getenv("DATABASE_SSL_MODE"))
		}
		return dsn //If the db is postgres then return the dsn (postgres is the default db)...
	} else if dbType == "mysql" {
		return "mysql://" + his.BuildDSN()
	} else {
		return "mssql://" + his.BuildDSN()
	}
}

// Moved showHelp from main.go
func showHelp() {
	color.Yellow(`Available commands:

	help			      - show the help commands
	version 		      - print application version
	migrate               - runs all up migrations that have not been run previously
	migrate down          - reverses the most recent migration
	migrate reset         - runs all down migrations in reverse order, and then all up migrations
	make migration <name> - creates (2) two new up and down migrations in the migrations folder
	make auth             - create and run migrations for authentication tables, also creates models and middleware
	make handler <name>   - create a stub handler in the handlers directory
	make model <name>     - create a new model in the models directory
	
	`)
}
