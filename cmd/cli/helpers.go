package main

import (
	"fmt"
	"os"

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
