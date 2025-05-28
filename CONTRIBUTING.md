# Contributing to Sprint Capacity Calculator

Thank you for your interest in contributing to Sprint Capacity Calculator! We welcome contributions from everyone.

## 🚀 Getting Started

### Prerequisites

- Go 1.20 or later
- Git
- Make (optional, but recommended)

### Development Setup

1. **Fork the repository** on GitHub
2. **Clone your fork**:

   ```bash
   git clone https://github.com/yourusername/sizely.git
   cd sizely
   ```

3. **Add the upstream remote**:

   ```bash
   git remote add upstream https://github.com/originaluser/sizely.git
   ```

4. **Install dependencies**:

   ```bash
   make deps
   ```

5. **Run tests** to ensure everything works:

   ```bash
   make test
   ```

## 🛠️ Development Workflow

### Making Changes

1. **Create a feature branch**:

   ```bash
   git checkout -b feature/your-feature-name
   ```

2. **Make your changes** following our coding standards

3. **Run tests and linting**:

   ```bash
   make check
   ```

4. **Build the project**:

   ```bash
   make build
   ```

5. **Test your changes**:

   ```bash
   make run-example
   ```

### Commit Message Guidelines

We follow the [Conventional Commits](https://conventionalcommits.org/) specification:

- `feat:` New features
- `fix:` Bug fixes
- `docs:` Documentation changes
- `style:` Code style changes (formatting, etc.)
- `refactor:` Code refactoring
- `perf:` Performance improvements
- `test:` Adding or updating tests
- `chore:` Maintenance tasks

Examples:

```
feat: add support for custom point configurations
fix: handle edge case in combination calculation
docs: update installation instructions
test: add benchmarks for calculator package
```

### Code Style

- Follow standard Go formatting (use `gofmt`)
- Run `make fmt` before committing
- Use meaningful variable and function names
- Add comments for exported functions and complex logic
- Keep functions focused and small

## 🧪 Testing

### Running Tests

```bash
# Run all tests
make test

# Run tests with coverage
make coverage

# Run benchmarks
go test -bench=. ./...
```

### Writing Tests

- Write unit tests for all new functionality
- Use table-driven tests when appropriate
- Include edge cases and error conditions
- Use `testify/assert` for assertions
- Add benchmarks for performance-critical code

Example test structure:

```go
func TestNewFeature(t *testing.T) {
    tests := []struct {
        name     string
        input    InputType
        expected ExpectedType
        wantErr  bool
    }{
        {
            name:     "valid input",
            input:    validInput,
            expected: expectedOutput,
            wantErr:  false,
        },
        // more test cases...
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result, err := NewFeature(tt.input)
            if tt.wantErr {
                assert.Error(t, err)
                return
            }
            assert.NoError(t, err)
            assert.Equal(t, tt.expected, result)
        })
    }
}
```

## 📋 Pull Request Process

1. **Update documentation** if needed
2. **Ensure all tests pass**:

   ```bash
   make check
   ```

3. **Update CHANGELOG.md** with your changes
4. **Create a pull request** with:
   - Clear title describing the change
   - Detailed description of what was changed and why
   - Link to any related issues
   - Screenshots (if applicable)

### Pull Request Template

```markdown
## Description

Brief description of changes made.

## Type of Change

- [ ] Bug fix (non-breaking change which fixes an issue)
- [ ] New feature (non-breaking change which adds functionality)
- [ ] Breaking change (fix or feature that would cause existing functionality to not work as expected)
- [ ] Documentation update

## Testing

- [ ] Tests pass locally
- [ ] New tests added for new functionality
- [ ] Manual testing completed

## Checklist

- [ ] Code follows project style guidelines
- [ ] Self-review completed
- [ ] Documentation updated
- [ ] CHANGELOG.md updated
```

## 🐛 Reporting Issues

### Bug Reports

When reporting bugs, please include:

1. **Environment information**:

   - Go version
   - Operating system
   - Binary version (if using pre-built)

2. **Steps to reproduce**:

   - Clear, numbered steps
   - Minimal example that reproduces the issue

3. **Expected vs actual behavior**

4. **Additional context**:
   - Error messages
   - Logs
   - Screenshots (if applicable)

### Feature Requests

When requesting features:

1. **Describe the problem** you're trying to solve
2. **Explain the proposed solution**
3. **Provide use cases** where this would be helpful
4. **Consider alternatives** you've explored

## 📚 Documentation

### Code Documentation

- Add comments for all exported functions
- Use Go doc format: `// FunctionName does X and returns Y`
- Include examples in comments when helpful

### User Documentation

- Update README.md for user-facing changes
- Add examples in the `examples/` directory
- Update CLI help text if commands change

## 🏗️ Project Structure

```
sizely/
├── cmd/capacity-calc/     # Main application entry point
├── internal/              # Private application code
│   ├── calculator/        # Core calculation logic
│   ├── cli/              # CLI interface and output
│   └── models/           # Data models and types
├── pkg/                  # Public library code (if any)
├── examples/             # Usage examples
├── docs/                 # Documentation
└── scripts/              # Build and utility scripts
```

### Package Guidelines

- `internal/` packages are private to this project
- `pkg/` packages could be imported by other projects
- Keep packages focused and cohesive
- Avoid circular dependencies

## 🔧 Advanced Development

### Adding New Features

1. **Design first**: Consider the API and user experience
2. **Start with tests**: Write failing tests for the new feature
3. **Implement incrementally**: Small, focused commits
4. **Update documentation**: Both code and user docs

### Performance Considerations

- Profile code for performance-critical paths
- Use benchmarks to validate improvements
- Consider memory allocation patterns
- Test with realistic data sizes

### Dependencies

- Minimize external dependencies
- Use standard library when possible
- For new dependencies, consider:
  - Maintenance status
  - License compatibility
  - Security track record

## 🤝 Community Guidelines

### Code of Conduct

Please read and follow our [Code of Conduct](CODE_OF_CONDUCT.md).

### Getting Help

- Check existing issues and discussions
- Ask questions in GitHub Discussions
- Tag maintainers if needed (but be patient)

### Recognition

Contributors will be:

- Listed in CHANGELOG.md for their contributions
- Mentioned in release notes for significant features
- Added to the contributors list

## 📄 License

By contributing, you agree that your contributions will be licensed under the MIT License.

## ❓ Questions?

If you have questions about contributing, please:

1. Check this document first
2. Search existing issues and discussions
3. Open a new discussion or issue
4. Contact the maintainers

Thank you for contributing! 🎉
