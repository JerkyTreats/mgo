package main

import (
	"fmt"

	"gopkg.in/night-codes/summer.v2"
)

type (
	obj map[string]interface{}
	arr []interface{}
)

var (
	panel = summer.Create(summer.Settings{
		Title:       "Admin",
		Port:        9000,
		DefaultPage: "hello",
		Path:        "admin", // application path
		DBName:      "admin",
		Views:       "templates/main",
		ViewsDoT:    "templates/dot", // doT.js templates
		FirstStart: func() { // some DB migrations etc.
			fmt.Println("Application is running for the first time!")
		},
		Debug: summer.Env("production", "") == "", // set env. var "production" for Debug:false
		JS:    []string{},                         // add custom JS files to template
		CSS:   []string{},                         // add custom CSS files to template
	})
)

func main() {
	fmt.Println("Application started at http://localhost:9000/admin")
	summer.Wait()
}
