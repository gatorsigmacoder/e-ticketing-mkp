package main

import "api-e-ticketing/src/database"

func main() {
	database.DatabaseInit()

	database.DropTables()
	database.Migration()

	database.UserSeeder()
}