package main

import (
	"arosara.com/task-manager/db"
	"arosara.com/task-manager/routes"
)

func main() {
	db.InitDb()
	routes.InitRoutes()

}
