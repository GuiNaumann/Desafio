package main

import (
	"fmt"
	"desafio/routes"
	
	"desafio/database"
)

func main() {

	database.ConectaBanco()
	fmt.Println("starting challenge project")
	routes.HandleRequests()
}