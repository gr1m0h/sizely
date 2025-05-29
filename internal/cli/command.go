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
	fmt.Println(`sizely

USAGE:
  capacity-calc [OPTIONS]

OPTIONS:
  -calc          Calculate total points from T-shirt sizes
  -reverse       Find all combinations for given points
  -input FILE    JSON file containing task counts
  -json STRING   JSON string containing task counts
  -points INT    Target points for reverse calculation
  -max INT       Maximum total tasks for reverse calculation (default: 15)
  -version       Show version information
  -help          Show this help

T-SHIRT SIZE POINTS:
  XS: 1 point  (30min - 4hrs)
  S:  3 points (4hrs - 1 day)
  M:  5 points (2-3 days)
  L:  10 points (1 week)

EXAMPLES:
  # Calculate points from JSON file
  capacity-calc -calc -input tasks.json

  # Calculate points from JSON string
  capacity-calc -calc -json '{"xs":3,"s":2,"m":1,"l":1}'

  # Find all combinations for 33 points
  capacity-calc -reverse -points 33

  # Find combinations with max 10 tasks
  capacity-calc -reverse -points 33 -max 10

JSON FORMAT:
  {
    "xs": 2,  // XS tasks count
    "s": 3,   // S tasks count
    "m": 1,   // M tasks count
    "l": 2    // L tasks count
  }

SRE BEST PRACTICES:
  - Reserve 20-30% capacity for incident response
  - Balance large strategic work (L) with quick wins (XS/S)
  - Consider using ScrumBan for operational flexibility
  - Track actual vs estimated time for continuous improvement

For more information, visit: https://github.com/gr1m0h/sizely`)
}
