package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strings"
	"time"

	"github.com/gertd/go-pluralize"
	"github.com/iancoleman/strcase"
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

	case "handler": //have to make sure that the 3rd arg is not an empty string...
		if arg3 == "" {
			exitGracefully(errors.New("you must give the handler a name"))
		}
		//Build a fileName for the new handler, copy some template file into that file and then write the file...
		fileName := his.RootPath + "/handlers/" + strings.ToLower(arg3) + ".go"
		if fileExists(fileName) { //check if exists and then stop if does so we do not overwrite users data...
			exitGracefully(errors.New(fileName + " already exists!"))
		}

		data, err := templateFS.ReadFile("templates/handlers/handler.go.txt")
		if err != nil {
			exitGracefully(err)
		}

		handler := string(data)
		handler = strings.ReplaceAll(handler, "$HANDLERNAME$", strcase.ToCamel(arg3))

		err = ioutil.WriteFile(fileName, []byte(handler), 0644)
		if err != nil {
			exitGracefully(err)
		}

	case "model": // if some one were to type in the cli ./hiSpeed.exe make model it would make a model for them...
		if arg3 == "" {
			exitGracefully(errors.New("you must give the model a name"))
		}

		data, err := templateFS.ReadFile("templates/data/model.go.txt")
		if err != nil {
			exitGracefully(err)
		}

		model := string(data)

		plur := pluralize.NewClient()

		var modelName = arg3
		var tableName = arg3

		if plur.IsPlural(arg3) {
			modelName = plur.Singular(arg3)
			tableName = strings.ToLower(tableName)
		} else {
			tableName = strings.ToLower(plur.Plural(arg3))
		}

		fileName := his.RootPath + "/data/" + strings.ToLower(modelName) + ".go"
		if fileExists(fileName) { //check if exists and then stop if does so we do not overwrite users data...
			exitGracefully(errors.New(fileName + " already exists!"))
		}

		model = strings.ReplaceAll(model, "$MODELNAME$", strcase.ToCamel(modelName))
		model = strings.ReplaceAll(model, "$TABLENAME$", tableName)

		err = copyDataToFIle([]byte(model), fileName)
		if err != nil {
			exitGracefully(err)
		}

	}

	return nil
}
