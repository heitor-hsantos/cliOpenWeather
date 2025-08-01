package main

import (
	"cliOpn/starter"
	"os"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	if len(os.Args) > 1 {
		// Modo CLI
		starter.HandleCLI()
	} else {
		// Modo Servidor
		starter.StartServer()
	}

}
