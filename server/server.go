package main

import (
	"log"
	"net/http"

	"github.com/Ryan-18-system/desafio-client-server-goexpert/server/database"
	"github.com/Ryan-18-system/desafio-client-server-goexpert/server/service"
)

func main() {
	database.InitDB()

	defer func() {
		if database.DB != nil {
			database.DB.Close()
		}
	}()
	server := http.NewServeMux()
	server.HandleFunc("/cotacao", service.ProcessCurrentQuote)
	log.Println("Servidor rodando na porta 8080...")
	err := http.ListenAndServe(":8080", server)
	if err != nil {
		log.Fatal("Erro ao iniciar servidor:", err)
	}
}
