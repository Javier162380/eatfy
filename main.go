package main

import (
	"eatfy/app"
)

func main() {
	app := &app.App{}
	app.Initialize("localhost", "5432", "", "", "")
	app.Run(":3000")
}
