package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/gr1m0h/sizely/internal/cli"
)

func main() {
	var (
		calculateCmd = flag.Bool("calc", false, "Calculate total points from T-shirt sizes")
		reverseCmd   = flag.Bool("reverse", false, "Find all combinations for given points")
		inputFile    = flag.String("input", "", "JSON file containing T-shirt size counts")
		inputJSON    = flag.String("json", "", "JSON string containing T-shirt size counts")
		points       = flag.Int("points", 0, "Target points for reverse calculation")
		maxTasks     = flag.Int("max", 15, "Maximum total tasks for reverse calculation")
		help         = flag.Bool("help", false, "Show help")
	)

	flag.Parse()

	if *help {
		cli.ShowHelp()
		return
	}

	app := cli.NewApp()

	if *calculateCmd {
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
			fmt.Println("Error: -calc requires either -input of -json")
			os.Exit(1)
		}
	} else if *reverseCmd {
		if *points <= 0 {
			fmt.Println("Error: -reverse requires -points with a positive value")
			os.Exit(1)
		}
		if err := app.ReverseCalculate(*points, *maxTasks); err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
	} else {
		cli.ShowHelp()
	}
}
