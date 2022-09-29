package main

import (
	"src/database"
)

func main() {
	database.Init()
	for {
		database.Menu()
	}
}
