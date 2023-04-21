package main

import (
	"golang-learning-path/go-testify/database"
	"golang-learning-path/go-testify/router"
)

func main() {
	database.StartDB()
	r := router.StartApp()
	r.Run(":8080")
}
