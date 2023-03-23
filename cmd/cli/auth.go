package main

import (
	"fmt"
	"log"
	"time"
)

func doAuth() error { // This func called from the make.go file case:auth statement...
	// going to need migrations
	dbType := his.DB.DataType

	fileName := fmt.Sprintf("%d_create_auth_tables", time.Now().UnixMicro()) // base file name
	upFile := his.RootPath + "/migrations/" + fileName + ".up.sql"
	downFile := his.RootPath + "/migrations/" + fileName + ".down.sql"

	log.Println(dbType, upFile, downFile)

	err := copyFileFromTemplate("templates/migrations/auth_tables."+dbType+".sql", upFile)
	if err != nil {
		exitGracefully(err)
	}

	err = copyDataToFIle([]byte("drop table if exists users cascade; drop table if exists tokens cascade; drop table if exists remember_tokens;"), downFile)
	if err != nil {
		exitGracefully(err)
	}

	// need to run migrations
	err = doMigrate("up", "")
	if err != nil {
		exitGracefully(err)
	}

	err = copyFileFromTemplate("templates/data/user.go.txt", his.RootPath+"/data/user.go")
	if err != nil {
		exitGracefully(err)
	}

	err = copyFileFromTemplate("templates/data/token.go.txt", his.RootPath+"/data/token.go")
	if err != nil {
		exitGracefully(err)
	}

	// need to copy files over prob need more but this is a start...

	return nil
}
