package main

import (
	"finalproject/database"
	"finalproject/router"
)

// @title Tag Service API
// @version 1.0
// @description A Tag Service

// @host localhost:8080
// @BasePath /api

func main() {
	database.StartDB()
	r := router.StartApp()
	r.Run(":8080")
}
