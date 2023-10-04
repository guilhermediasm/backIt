package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"back-It/routes"
)

func main() {
	// Crie um roteador
	router := mux.NewRouter()

	// Configure as rotas
	routes.SetupRoutes(router)

	// Inicie o servidor na porta 8080
	log.Fatal(http.ListenAndServe(":8080", router))
}
