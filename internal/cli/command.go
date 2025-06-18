package cli

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/gr1m0h/sizely/internal/calculator"
	"github.com/gr1m0h/sizely/internal/models"
)

// App represents the CLI application
type App struct {
	calculator *calculator.Calculator
	output     *OutputFormatter
}

// NewApp creates a new CLI application instance
func NewApp() *App {
	return &App{
		calculator: calculator.NewCalculator(),
		output:     NewOutputFormatter(),
	}
}

// CalculateFromFile calculates capacity from a JSON file
func (a *App) CalculateFromFile(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("reading file: %w", err)
	}

	return a.CalculateFromJSON(string(data))
}

// CalculateFromJSON calculates capacity from a JSON string
func (a *App) CalculateFromJSON(jsonStr string) error {
	var tasks models.TaskCount
	if err := json.Unmarshal([]byte(jsonStr), &tasks); err != nil {
		return fmt.Errorf("parsing JSON: %w", err)
	}

	capacity := a.calculator.CalculateSprintCapacity(tasks)
	a.output.PrintCapacity(capacity)

	return nil
}

// ReverseCalculate finds all combinations for given points
func (a *App) ReverseCalculate(points, maxTasks int) error {
	if points <= 0 {
		return fmt.Errorf("points must be positive")
	}

	if maxTasks <= 0 {
		return fmt.Errorf("max tasks must be positive")
	}

	result := a.calculator.FindCombinations(points, maxTasks)
	a.output.PrintCombinations(result)

	return nil
}

// ShowHelp displays help information
func ShowHelp() {
	fmt.Println(`sizely - T-shirt size estimation and sprint capacity planning tool

DESCRIPTION:
  sizely calculates sprint points from T-shirt sizes and performs reverse calculations
  to find optimal task combinations for target points.

USAGE:
  sizely <command> [options]

COMMANDS:
  estimate           Calculate total sprint points from T-shirt size counts (default)
  breakdown          Find all possible task combinations for a target point value
  help               Show this help information

ESTIMATE OPTIONS:
  -i, -input FILE    Path to JSON file containing T-shirt size task counts
  -j, -json STRING   JSON string containing T-shirt size task counts

BREAKDOWN OPTIONS:
  <points>           Target points for reverse calculation (required positional argument)
  -m, -max INT       Maximum number of total tasks allowed in combinations (default: 15)

T-SHIRT SIZE POINT SYSTEM:
  XS: 1 point   (30 minutes - 4 hours)
  S:  3 points  (4 hours - 1 day)
  M:  5 points  (2-3 days)
  L:  10 points (1 week)

EXAMPLES:
  # Calculate total points from JSON file (estimate is the default command)
  sizely -input examples/basic/tasks.json
  sizely estimate -input examples/basic/tasks.json

  # Calculate points from inline JSON string
  sizely -json '{"xs":3,"s":2,"m":1,"l":1}'

  # Find all task combinations that sum to 33 points
  sizely breakdown 33

  # Find combinations with maximum 10 total tasks
  sizely breakdown 33 -max 10

JSON INPUT FORMAT:
  {
    "xs": 2,  // Number of XS tasks
    "s": 3,   // Number of S tasks  
    "m": 1,   // Number of M tasks
    "l": 2    // Number of L tasks
  }

For more information, visit: https://github.com/gr1m0h/sizely`)
}
