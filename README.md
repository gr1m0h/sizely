# sizely

[![CI](https://github.com/gr1m0h/sizely/workflows/CI/badge.svg)](https://github.com/gr1m0h/sizely/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/gr1m0h/sizely)](https://goreportcard.com/report/github.com/gr1m0h/sizely)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Release](https://img.shields.io/github/release/gr1m0h/sizely.svg)](https://github.com/gr1m0h/sizely/releases)

A command-line tool for SRE teams to calculate sprint capacity using T-shirt size estimation and find optimal task combinations for target sprint points.

## 🎯 Features

- **Calculate Sprint Points**: Convert T-shirt size estimates (XS, S, M, L) to points
- **Reverse Calculation**: Find all possible task combinations for target points
- **Capacity Assessment**: Get recommendations based on SRE best practices
- **JSON Support**: Accept input from files or command-line JSON strings
- **Multiple Output Formats**: Human-readable tables and JSON for automation

## 📦 Installation

### Using Go Install

```bash
go install github.com/gr1m0h/sizely/cmd/sizely@latest
```

### Download Binary

Download the latest binary from [releases page](https://github.com/gr1m0h/sizely/releases).

### Using Homebrew (macOS/Linux)

```bash
brew install gr1m0h/tap/sizely
```

### Building from Source

```bash
git clone https://github.com/gr1m0h/sizely.git
cd sizely
make build
```

## 🚀 Quick Start

### Calculate Points from Tasks

```bash
# From JSON file
sizely -calc -input examples/basic/tasks.json

# From JSON string
sizely -calc -json '{"xs":3,"s":2,"m":1,"l":1}'
```

### Find Task Combinations

```bash
# Find all combinations for 33 points
sizely -reverse -points 33

# Limit to maximum 10 tasks
sizely -reverse -points 33 -max 10
```

## 📊 T-shirt Size Points

| Size | Points | Time Estimate |
| ---- | ------ | ------------- |
| XS   | 1      | 30min - 4hrs  |
| S    | 3      | 4hrs - 1 day  |
| M    | 5      | 2-3 days      |
| L    | 10     | 1 week        |

## 📋 Usage Examples

### Basic Calculation

```bash
$ sizely -calc -json '{"xs":3,"s":2,"m":1,"l":1}'

📊 Sprint Capacity Calculation
═══════════════════════════════
XS (1pt):   3 tasks =  3 points
S  (3pt):   2 tasks =  6 points
M  (5pt):   1 tasks =  5 points
L (10pt):   1 tasks = 10 points
───────────────────────────────
Total:      7 tasks = 24 points

🎯 Capacity Assessment
⚠️  CONSERVATIVE: 24 points is below optimal range
```

### Reverse Calculation

```bash
$ sizely -reverse -points 33

🔍 Finding combinations for 33 points (max 15 tasks)
═══════════════════════════════════════════════════
Found 12 combination(s):

 1. L×3 + XS×3 = 33 points (6 tasks)
    ✅ Good mix of large and small tasks

 2. L×2 + M×2 + XS×3 = 33 points (7 tasks)
    ✅ Good mix of large and small tasks
```

## 🔧 JSON Input Format

```json
{
  "xs": 2,
  "s": 3,
  "m": 1,
  "l": 2
}
```

## 📚 Documentation

- [Installation Guide](docs/installation.md)
- [Usage Guide](docs/usage.md)
- [API Documentation](docs/api.md)
- [Contributing Guide](docs/contributing.md)

## 🤝 Contributing

We welcome contributions! Please see our [Contributing Guide](CONTRIBUTING.md) for details.

### Development Setup

```bash
# Clone the repository
git clone https://github.com/gr1m0h/sizely.git
cd sizely

# Install dependencies
go mod download

# Run tests
make test

# Build
make build
```

## 📝 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

- Built for teams practicing ScrumBan methodology

## 📈 Roadmap

- [ ] Web UI interface
- [ ] Integration with JIRA/Linear
- [ ] Team velocity tracking
- [ ] Historical data analysis
- [ ] Custom point configurations
