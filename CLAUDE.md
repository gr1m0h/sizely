# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

sizely is a Go command-line tool for T-shirt size estimation and sprint capacity planning. It calculates sprint points from T-shirt sizes (XS=1, S=3, M=5, L=10) and performs reverse calculations to find optimal task combinations for target points.

## Development Commands

### Build and Test
- `make build` - Build the sizely binary to ./bin/sizely
- `make test` - Run all tests with race detection
- `make check` - Run format, vet, and test checks
- `make coverage` - Generate test coverage report as coverage.html
- `make lint` - Run golangci-lint (requires golangci-lint installed)

### Development
- `make fmt` - Format Go code using gofmt
- `make vet` - Run go vet analysis
- `make deps` - Download and tidy dependencies
- `make clean` - Remove build artifacts and coverage files

### Running Examples
- `make run-example` - Build and run examples with sample data
- `./bin/sizely -input examples/basic/tasks.json` - Calculate from JSON file (estimate is default)
- `./bin/sizely estimate -input examples/basic/tasks.json` - Explicit estimate command
- `./bin/sizely -json '{"xs":3,"s":2,"m":1,"l":1}'` - Calculate from JSON string
- `./bin/sizely breakdown 33` - Find combinations for 33 points
- `./bin/sizely breakdown 33 -max 10` - Find combinations with max 10 tasks

## CLI Commands

### Commands
- `estimate` (default) - Calculate total points from T-shirt sizes
- `breakdown` - Find all combinations for given points
- `help` - Show help information

### Estimate Options
- `-i, -input FILE` - JSON file containing task counts
- `-j, -json STRING` - JSON string containing task counts

### Breakdown Options
- `<points>` - Target points for reverse calculation (required positional argument)
- `-m, -max INT` - Maximum total tasks for reverse calculation (default: 15)

## Architecture

### Core Components
- **cmd/sizely/main.go** - CLI entry point with flag parsing
- **internal/calculator/** - Core calculation logic for points and combinations
- **internal/cli/** - CLI application logic and output formatting
- **internal/models/** - Data structures (TaskCount, Combination, SprintCapacity)

### Key Types
- `TaskCount` - Represents counts of XS/S/M/L tasks from JSON input
- `SprintCapacity` - Complete capacity calculation with breakdown
- `Combination` - Task combination result for reverse calculations
- `CombinationResult` - Complete reverse calculation with recommendations

### T-shirt Size Point System
- XS: 1 point (30min - 4hrs)
- S: 3 points (4hrs - 1 day)
- M: 5 points (2-3 days)
- L: 10 points (1 week)

## Testing

The project uses testify for assertions. Tests are located alongside source files with `_test.go` suffix. Run `make test` to execute all tests with race detection.

## JSON Input Format

```json
{
  "xs": 2,
  "s": 3,
  "m": 1,
  "l": 2
}
```

## Key Dependencies

- Go 1.21+
- github.com/stretchr/testify v1.8.4 (testing only)
- No external runtime dependencies - builds static binary