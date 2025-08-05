package cmd

import (
	"cliOpn/config"
	"cliOpn/handlers"
	"fmt"
	"log"
	"os"
	"strconv"
)

func Execute() {
	if len(os.Args) < 2 {
		printHelp()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "get":
		handleGetCommand(os.Args[2:])
	case "set":
		handleConfigCommand(os.Args[2:])
	case "show":
		fmt.Println("Current configuration:")
		if err, _ := config.GetConfig(); err != nil {
			log.Fatalf("Error reading config: %v", err)
		}
	default:
		printHelp()
		os.Exit(1)
	}

}

func printHelp() {
	fmt.Println("Usage: cliOpn <get>[weather] or to config default location cliOpn <set> [coordinates|excluded] <value> \n" +
		"Example: cliOpn get weather \n" +
		"Example: cliOpn set coordinates 40.7128 -74.0060 \n" +
		"Example: cliOpn set excluded minutely hourly daily alerts \n" +
		"Example: cliOpn show {display current json config}")
}
func handleGetCommand(args []string) {
	if len(args) < 1 {
		printHelp()
		os.Exit(1)
	}

	switch args[0] {
	case "weather":
		fmt.Println("Fetching weather data... ")
		data, err := handlers.FetchWeatherDataWithJson()
		if err != nil {
			fmt.Printf("Error fetching weather data: %v\n", err)
			return
		}
		if data == nil {
			fmt.Println("No weather data found.")
			return
		}

	case "coordinate":
		if len(args) < 3 {
			fmt.Println("Usage: cliOpn get coordinate <lat> <lon>")
			os.Exit(1)
		}
		lat, err1 := strconv.ParseFloat(args[1], 64)
		lon, err2 := strconv.ParseFloat(args[2], 64)
		if err1 != nil || err2 != nil {
			fmt.Println("Invalid coordinates. Use: cliOpn get coordinate <lat> <lon>")
			os.Exit(1)
		}
		fmt.Printf("Fetching weather data for coordinates: %f, %f...\n", lat, lon)
		cfg, err := config.GetConfig()

		data, err := handlers.FetchWeatherDataWithCoordinates(lat, lon, cfg.ExcludedFields)
		if err != nil {
			fmt.Printf("Error fetching weather data: %v\n", err)
			return
		}
		if data == nil {
			fmt.Println("No weather data found for the specified coordinates.")
			return
		}
		fmt.Println("Weather data for coordinates:", lat, lon)
	default:
		printHelp()
		os.Exit(1)
	}
}
func handleConfigCommand(args []string) {
	if len(args) < 2 {
		printHelp()
		os.Exit(1)
	}

	switch args[0] {
	case "coordinates":
		if len(args) < 2 {
			printHelp()
			os.Exit(1)
		}
		lat, err := strconv.ParseFloat(args[1], 64)

		if err != nil {
			log.Fatal("Invalid latitude value: %s\n", args[1])
		}

		lon, err := strconv.ParseFloat(args[2], 64)
		if err != nil {
			log.Fatal("Invalid longitude value: %s\n", args[2])
		}

		if err := config.UpdateCoordinates(lat, lon); err != nil {
			log.Fatalf("Error updating coordinates: %v\n", err)
		}
		fmt.Println("Updated default coordinates:", lat, lon)

	case "excluded":
		if len(args) < 2 {
			err := config.UpdatexcludedFields(args[1:])
			if err != nil {
				return
			}
			os.Exit(1)
		}

	default:
		printHelp()
		os.Exit(1)
	}
}
