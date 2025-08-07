package main

import (
	"cliOpn/starter"
	"github.com/joho/godotenv"
	"log"
	"os"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao carregar o arquivo .env")
	}

	if len(os.Args) > 1 {
		// Modo CLI
		starter.HandleCLI()
	}
}
