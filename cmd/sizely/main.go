package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/gr1m0h/sizely/internal/cli"
)

func main() {
	subcommand := os.Args[1]

	switch subcommand {
	case "points":
		pointsCmdWithArgs(os.Args[2:])
	case "tasks":
		tasksCmd()
	case "help", "-help", "--help":
		cli.ShowHelp()
	default:
		fmt.Printf("Unknown command: %s\n", subcommand)
		cli.ShowHelp()
		os.Exit(1)
	}
}

func pointsCmdWithArgs(args []string) {
	fs := flag.NewFlagSet("points", flag.ExitOnError)
	inputFile := fs.String("file", "", "T-shirt size data from file")
	fs.StringVar(inputFile, "f", "", "T-shirt size data from file")
	inputData := fs.String("data", "", "T-shirt size data from string ")
	fs.StringVar(inputData, "d", "", "T-shirt size data from string")

	if err := fs.Parse(args); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	app := cli.NewApp()

	if *inputFile != "" {
		if err := app.CalculateFromFile(*inputFile); err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
	} else if *inputData != "" {
		if err := app.CalculateFromJSON(*inputData); err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
	} else {
		fmt.Println("Error: points requires either -f/--file or -d/--data")
		fs.Usage()
		os.Exit(1)
	}
}

func tasksCmd() {
	args := os.Args[2:]

	if len(args) < 1 {
		fmt.Println("Error: tasks requires points as first argument")
		fmt.Println("Usage: sizely tasks <points> [-c/--count <tasks>] [-o/--output-json]")
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

	fs := flag.NewFlagSet("tasks", flag.ExitOnError)
	count := fs.Int("count", 15, "Maximum total tasks count")
	fs.IntVar(count, "c", 15, "Maximum total tasks count")
	outputJSON := fs.Bool("output-json", false, "Output results in JSON format")
	fs.BoolVar(outputJSON, "o", false, "Output results in JSON format")

	if err := fs.Parse(args[1:]); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	app := cli.NewApp()

	if err := app.ReverseCalculate(points, *count, *outputJSON); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
