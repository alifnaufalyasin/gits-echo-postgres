package main

import (
	"gits-echo-boilerplate/database"
	"gits-echo-boilerplate/routes"
)

func main() {

	database.Init()
	e := routes.Init()

	e.Logger.Fatal(e.Start(":4132"))
}
