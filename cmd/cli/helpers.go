package main

import (
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
