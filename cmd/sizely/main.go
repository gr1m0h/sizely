package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/gr1m0h/sizely/internal/cli"
)

func main() {
	// If no args or first arg starts with -, default to estimate
	if len(os.Args) < 2 || (len(os.Args) >= 2 && os.Args[1][0] == '-') {
		estimateCmdWithArgs(os.Args[1:])
		return
	}

	subcommand := os.Args[1]

	switch subcommand {
	case "estimate":
		estimateCmdWithArgs(os.Args[2:])
	case "breakdown":
		breakdownCmd()
	case "help", "-help", "--help":
		cli.ShowHelp()
	default:
		fmt.Printf("Unknown command: %s\n", subcommand)
		cli.ShowHelp()
		os.Exit(1)
	}
}

func estimateCmdWithArgs(args []string) {
	fs := flag.NewFlagSet("estimate", flag.ExitOnError)
	inputFile := fs.String("input", "", "JSON file containing T-shirt size counts")
	inputJSON := fs.String("json", "", "JSON string containing T-shirt size counts")

	fs.Parse(args)

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
		fmt.Println("Error: estimate requires either -input or -json")
		fs.Usage()
		os.Exit(1)
	}
}

func breakdownCmd() {
	args := os.Args[2:]
	
	if len(args) < 1 {
		fmt.Println("Error: breakdown requires points as first argument")
		fmt.Println("Usage: sizely breakdown <points> [-max <max_tasks>]")
		os.Exit(1)
	}

	var points int
	if _, err := fmt.Sscanf(args[0], "%d", &points); err != nil {
		fmt.Printf("Error: invalid points value '%s', must be a positive integer\n", args[0])
		os.Exit(1)
	}

	if points <= 0 {
		fmt.Println("Error: points must be positive")
		os.Exit(1)
	}

	fs := flag.NewFlagSet("breakdown", flag.ExitOnError)
	maxTasks := fs.Int("max", 15, "Maximum total tasks for reverse calculation")

	fs.Parse(args[1:])

	app := cli.NewApp()

	if err := app.ReverseCalculate(points, *maxTasks); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
