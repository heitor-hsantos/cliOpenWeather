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
		fmt.Println(config.AppConfig.Lon)
		fmt.Println(config.AppConfig.Lat)
		fmt.Println(config.AppConfig.ExcludedFields)
		if err := config.ReadConfig(); err != nil {
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
	if len(args) < 1 || args[0] != "weather" {
		printHelp()
		os.Exit(1)
	}

	fmt.Println("Fetching weather data...")

	data, err := handlers.FetchWeatherData(config.AppConfig.Lat, config.AppConfig.Lon)
	if err != nil {
		fmt.Printf("Error fetching weather data: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Weather data:", data)

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
			os.Exit(1)
		}

		lon, err := strconv.ParseFloat(args[2], 64)
		if err != nil {
			log.Fatal("Invalid longitude value: %s\n", args[2])
			os.Exit(1)
		}

		if err := config.UpdateCoordinates(lat, lon); err != nil {
			log.Fatalf("Error updating coordinates: %v\n", err)
		}
		fmt.Println("Updated default coordinates:", lat, lon)

		fmt.Printf("Latitude updated to: %f\n", config.AppConfig.Lat)
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
