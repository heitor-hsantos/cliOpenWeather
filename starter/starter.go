package starter

import (
	"cliOpn/cmd"
	"cliOpn/config"
	"fmt"
	_ "strings"
)

func HandleCLI() {
	cmd.Execute()
}

// StartServer inicia o servidor web
func StartServer() {
	fmt.Println("Debug mode")
	fmt.Println(config.GetConfig())

	//fmt.Println(config.AppConfig.Lat, config.AppConfig.Lon, config.AppConfig.ExcludedFields)
	//fmt.Print(os.Getenv("API_KEY"))
	//fmt.Print(os.Getenv("API_URL"))

}
