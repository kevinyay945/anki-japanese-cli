# Contributing to Anki Japanese CLI

Thank you for your interest in contributing to the Anki Japanese CLI project! This document provides guidelines and instructions for contributing.

## Table of Contents

- [Code of Conduct](#code-of-conduct)
- [Getting Started](#getting-started)
- [Development Environment](#development-environment)
- [Coding Standards](#coding-standards)
- [Testing](#testing)
- [Pull Request Process](#pull-request-process)
- [Project Structure](#project-structure)

## Code of Conduct

Please be respectful and considerate of others when contributing to this project. We aim to foster an inclusive and welcoming community.

## Getting Started

1. Fork the repository on GitHub
2. Clone your fork locally
3. Add the original repository as a remote named "upstream"
4. Create a new branch for your feature or bug fix
5. Make your changes
6. Push your branch to your fork
7. Submit a pull request

## Development Environment

### Prerequisites

- Go 1.23 or higher
- Anki with AnkiConnect plugin (for testing)

### Setup

```bash
# Clone the repository
git clone https://github.com/yourusername/anki-japanese-cli.git

# Navigate to the project directory
cd anki-japanese-cli

# Install dependencies
go mod download

# Build the project
go build -o anki-japanese-cli
```

## Coding Standards

We follow standard Go coding conventions:

1. Use `gofmt` to format your code
2. Follow the [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)
3. Document all exported functions, types, and constants
4. Write meaningful commit messages
5. Keep functions small and focused on a single responsibility
6. Use meaningful variable and function names

## Testing

All new features and bug fixes should include tests. We use Go's standard testing package.

### Running Tests

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run tests for a specific package
go test ./internal/anki
```

### Integration Tests

For integration tests that interact with Anki:

1. Make sure Anki is running with the AnkiConnect plugin installed
2. Set the environment variable `ANKI_INTEGRATION_TEST=true`
3. Run the tests:

```bash
ANKI_INTEGRATION_TEST=true go test ./...
```

## Pull Request Process

1. Ensure your code follows the coding standards
2. Add or update tests as necessary
3. Update documentation to reflect any changes
4. Ensure all tests pass
5. Rebase your branch on the latest upstream master
6. Submit your pull request with a clear description of the changes

## Project Structure

The project is organized as follows:

- `cmd/`: Command-line interface code
  - `root.go`: Root command
  - `init.go`: Initialize card types
  - `add.go`: Add cards
- `internal/`: Internal packages
  - `anki/`: Anki Connect client
  - `models/`: Card models
  - `templates/`: HTML templates for cards
  - `config/`: Configuration
- `examples/`: Example files
- `docs/`: Documentation

### Key Components

#### Anki Client

The `internal/anki/client.go` file contains the client for interacting with the Anki Connect API. It handles:

- Creating decks
- Creating note models
- Adding notes
- Error handling and retries

#### Card Models

The `internal/models/` directory contains the card models:

- `verb_card.go`: Verb card model
- `adjective_card.go`: Adjective card model
- `normal_card.go`: Normal word card model
- `grammar_card.go`: Grammar card model

#### Templates

The `internal/templates/` directory contains the HTML templates for the cards:

- `verb_front.html`, `verb_back.html`: Verb card templates
- `adjective_front.html`, `adjective_back.html`: Adjective card templates
- `normal_front.html`, `normal_back.html`: Normal word card templates
- `grammar_front.html`, `grammar_back.html`: Grammar card templates

## Thank You

Your contributions are greatly appreciated! Together, we can make this tool even better for Japanese language learners.