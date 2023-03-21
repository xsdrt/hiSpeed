package main

import (
	"errors"
	"fmt"
	"time"
)

func doMake(arg2, arg3 string) error {

	switch arg2 {
	case "migration":
		dbType := his.DB.DataType //get the users database type (supporting mysql-postgre-sqlserver)
		if arg3 == "" {           //arg3 will hold the migration name...
			exitGracefully(errors.New("you must give the migration a name"))
		}

		fileName := fmt.Sprintf("%d_%s", time.Now().UnixMicro(), arg3) //create the files for the up migration and down migration...

		upFile := his.RootPath + "/migrations/" + fileName + "." + dbType + ".up.sql"
		downFile := his.RootPath + "/migrations/" + fileName + "." + dbType + ".down.sql"

		//create some templates for migrations so the end users of our framework have something to work/start with...
		err := copyFileFromTemplate("templates/migrations/migration."+dbType+".up.sql", upFile)
		if err != nil {
			exitGracefully(err)
		}

		err = copyFileFromTemplate("templates/migrations/migration."+dbType+".down.sql", downFile)
		if err != nil {
			exitGracefully(err)
		}

	case "auth":
		err := doAuth() //this func will be in its own file due to its lengthy code...
		if err != nil {
			exitGracefully(err)
		}
	}

	return nil
}
