package main

import (
	"src/database"
)

func main() {
	database.Initialisation()
	for {
		database.Menu()
	}
}
