package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/gr1m0h/sizely/internal/cli"
)

func main() {
	if len(os.Args) < 2 {
		cli.ShowHelp()
		return
	}

	subcommand := os.Args[1]
	
	switch subcommand {
	case "calc":
		calcCmd()
	case "reverse":
		reverseCmd()
	case "help", "-help", "--help":
		cli.ShowHelp()
	default:
		fmt.Printf("Unknown command: %s\n", subcommand)
		cli.ShowHelp()
		os.Exit(1)
	}
}

func calcCmd() {
	fs := flag.NewFlagSet("calc", flag.ExitOnError)
	inputFile := fs.String("input", "", "JSON file containing T-shirt size counts")
	inputJSON := fs.String("json", "", "JSON string containing T-shirt size counts")
	
	fs.Parse(os.Args[2:])
	
	app := cli.NewApp()
	
	if *inputFile != "" {
		if err := app.CalculateFromFile(*inputFile); err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
	} else if *inputJSON != "" {
		if err := app.CalculateFromJSON(*inputJSON); err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
	} else {
		fmt.Println("Error: calc requires either -input or -json")
		fs.Usage()
		os.Exit(1)
	}
}

func reverseCmd() {
	fs := flag.NewFlagSet("reverse", flag.ExitOnError)
	points := fs.Int("points", 0, "Target points for reverse calculation")
	maxTasks := fs.Int("max", 15, "Maximum total tasks for reverse calculation")
	
	fs.Parse(os.Args[2:])
	
	if *points <= 0 {
		fmt.Println("Error: reverse requires -points with a positive value")
		fs.Usage()
		os.Exit(1)
	}
	
	app := cli.NewApp()
	
	if err := app.ReverseCalculate(*points, *maxTasks); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
