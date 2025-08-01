package main

import (
	"cliOpn/handlers"
	"cliOpn/routes"
	"fmt"
	"log"
	"os"
	"strings"
)

func handleCLI() {
	if len(os.Args) < 3 || os.Args[1] != "get" || os.Args[2] != "weather" {
		fmt.Println("Uso: go run main.go get weather <cidade>")
		os.Exit(1)
	}

	// Junta todos os argumentos após "get weather" para formar o nome da cidade
	city := strings.Join(os.Args[3:], " ")
	if city == "" {
		fmt.Println("Nome da cidade é obrigatório.")
		fmt.Println("Uso: go run main.go get weather <cidade>")
		os.Exit(1)
	}

	weatherInfo, err := handlers.FetchWeather(city)
	if err != nil {
		log.Fatalf("Erro ao buscar dados do tempo: %v", err)
	}

	fmt.Println(weatherInfo)
}

// startServer inicia o servidor web
func startServer() {
	log.SetOutput(os.Stdout)
	r := routes.RegisterRoutes()
	log.Println("Servidor iniciado em :9090")
	err := http.ListenAndServe(":9090", r)
	if err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}
