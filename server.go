package main

import "golang-mysql-restful-starter-kit/api"

func main() {
	app := api.NewApp()
	app.Run()
}
