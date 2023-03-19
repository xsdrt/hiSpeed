package main

import (
	"errors"
	"os"

	"github.com/fatih/color"
	"github.com/xsdrt/hiSpeed"
)

const version = "1.0.0"

var his hiSpeed.HiSpeed

func main() {
	var message string
	arg1, arg2, arg3, err := validateInput()
	if err != nil {
		exitGracefully(err)
	}

	setup()

	switch arg1 {
	case "help":
		showHelp()

	case "version":
		color.Yellow("Application version: " + version)

	case "migrate":
		if arg2 == "" { //if arg2 is equal to nothing , assume they are running an up migration...
			arg2 = "up"
		}
		err = doMigrate(arg2, arg3)
		if err != nil {
			exitGracefully(err)
		}
		message = "Migrations completed!"

	case "make":
		if arg2 == "" { //if blank or empty string exit with error
			exitGracefully(errors.New("make requires a subcommand: (migration|model|handler)"))
		}
		err = doMake(arg2, arg3) //refer to the make.go file in the cli folder...
		if err != nil {
			exitGracefully(err)
		}

	default:
		showHelp()
	}
	exitGracefully(nil, message)
}

func validateInput() (string, string, string, error) {
	var arg1, arg2, arg3 string

	if len(os.Args) > 1 {
		arg1 = os.Args[1]

		if len(os.Args) >= 3 {
			arg2 = os.Args[2]
		}

		if len(os.Args) >= 4 {
			arg3 = os.Args[3]
		}
	} else {
		color.Red("Error: command required")
		showHelp()
		return "", "", "", errors.New("command required")
	}

	return arg1, arg2, arg3, nil
}

func showHelp() {
	color.Yellow(`Available commands:
	help			- show the help commands
	version 		- print application version
	
	`)
}

func exitGracefully(err error, msg ...string) {
	message := ""
	if len(msg) > 0 {
		message = msg[0]
	}

	if err != nil {
		color.Red("Error: %v\n", err)
	}

	if len(message) > 0 {
		color.Yellow(message)
	} else {
		color.Green("Finished!")
	}

	os.Exit(0)
}
