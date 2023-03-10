package main

import (
	"fmt"
	"desafio/routes"
	
	"desafio/database"
)

func main() {
	//conexao com o banco
	database.ConectaBanco()
	//print para informar que esta executando
	fmt.Println("starting challenge project")
	//rotas handleRequests
	routes.HandleRequests()
}