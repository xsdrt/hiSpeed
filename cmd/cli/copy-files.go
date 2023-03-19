package main

import (
	"embed"
	"io/ioutil"
)

//go:embed templates
var templateFS embed.FS

func copyFileFromTemplate(templatePath, targetFile string) error {
	// TODO: check to ensure file doesn't already exits...

	data, err := templateFS.ReadFile(templatePath)
	if err != nil {
		exitGracefully(err)
	}

	err = copyDataToFIle(data, targetFile)
	if err != nil {
		exitGracefully(err)
	}

	return nil
}

func copyDataToFIle(data []byte, to string) error {
	err := ioutil.WriteFile(to, data, 0644)
	if err != nil {
		return err
	}
	return nil
}
