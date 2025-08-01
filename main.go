package main

import (
	"cliOpn/routes"
	"log"
	"net/http"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	r := routes.RegisterRoutes()
	log.Println("Servidor iniciado em :8080")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		return
	}
}
