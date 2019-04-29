package main

import (
	"eatfy/app"
)

func main() {
	app := &app.App{}
	app.Initialize()
	app.Run(":3000")
}
