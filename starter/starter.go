package starter

import (
	"cliOpn/config"
	"cliOpn/handlers"
	"cliOpn/routes"
	"fmt"
	"log"
	"net/http"
	"os"
	_ "strings"
)

func HandleCLI() {

	if len(os.Args) < 3 || os.Args[1] != "get" || os.Args[2] != "weather" {
		fmt.Println("Uso: go run main.go get weather \n" +
			"<Se o CLI não for configurdao previamente o comando exibira por default dados de Nova York>")
		os.Exit(1)
	}

	weatherInfo, err := handlers.FetchWeatherData(config.AppConfig.Lat, config.AppConfig.Lon)
	if err != nil {
		log.Fatalf("Erro ao buscar dados do tempo: %v", err)
	}

	fmt.Println("Dados do tempo:", weatherInfo)

}

// StartServer inicia o servidor web
func StartServer() {
	log.SetOutput(os.Stdout)
	r := routes.RegisterRoutes()

	log.Println("Servidor iniciado em :9090")
	err := http.ListenAndServe(":9090", r)

	if err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}
