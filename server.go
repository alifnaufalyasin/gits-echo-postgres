package main

import "gits-echo-boilerplate/routes"

func main() {

	e := routes.Init()

	e.Logger.Fatal(e.Start(":4132"))
}
