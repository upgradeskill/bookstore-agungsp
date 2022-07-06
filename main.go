package main

import "bookstore/system"

type App struct {
	server system.Server
}

func main() {

	var app App
	
	app.server.Init()
	app.server.Start()

}
