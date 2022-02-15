package main

import (
	"github.com/tiagompalte/api-go-gin/database"
	"github.com/tiagompalte/api-go-gin/routes"
)

func main() {
	database.ConectaDatabase()
	routes.HandleRequests()
}
