# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

Sizely is a Go CLI tool for calculating sprint capacity using T-shirt size estimation (XS, S, M, L) and finding optimal task combinations for target sprint points. It's designed for SRE teams practicing ScrumBan methodology.

## Common Development Commands

### Build and Test
```bash
# Build the binary
make build

# Run all tests with race detection
make test

# Run all checks (format, vet, test)
make check

# Format code
make fmt

# Run linter (requires golangci-lint)
make lint

# Run go vet
make vet

# Generate test coverage report
make coverage
```

### Development Dependencies
```bash
# Download and tidy dependencies
make deps

# Install golangci-lint for linting
curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.54.2
```

### Running Examples
```bash
# Run example calculations
make run-example

# Test calculation mode
./bin/sizely -calc -input examples/basic/tasks.json
./bin/sizely -calc -json '{"xs":3,"s":2,"m":1,"l":1}'

# Test reverse calculation mode  
./bin/sizely -reverse -points 33
./bin/sizely -reverse -points 33 -max 10
```

## Architecture

### Package Structure
- `cmd/sizely/` - Main CLI entry point with flag parsing
- `internal/models/` - Core data types and T-shirt size point mappings
- `internal/calculator/` - Sprint capacity calculation logic and combination algorithms
- `internal/cli/` - CLI application logic and output formatting

### Core Types
- `TaskCount` - Input structure for T-shirt size counts (xs, s, m, l)
- `SprintCapacity` - Complete sprint calculation with breakdown
- `Combination` - Task combination result for reverse calculations
- `CombinationResult` - Full reverse calculation result with recommendations

### T-shirt Size Points Mapping
- XS: 1 point (30min - 4hrs)
- S: 3 points (4hrs - 1 day) 
- M: 5 points (2-3 days)
- L: 10 points (1 week)

### Key Algorithms
- Forward calculation: Sum of (count × points) for each size
- Reverse calculation: Brute-force enumeration with constraints, sorted by task count
- Recommendations: Balanced workflow analysis and SRE-specific capacity advice

## Testing Strategy

Tests are located adjacent to source files (`*_test.go`). The project uses:
- `github.com/stretchr/testify` for assertions
- Race detection enabled in test runs
- Coverage reporting with HTML output

## Build Configuration

- Go 1.21+ required
- Uses ldflags for version injection
- Cross-platform builds supported via `make build-all`
- Docker support available