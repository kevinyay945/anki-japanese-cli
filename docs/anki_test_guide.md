# Anki Connect Test Guide

This guide explains how to run the tests for the Anki Connect client functionality in the anki-japanese-cli tool.

## Prerequisites

Before running the tests, make sure you have:

1. Anki installed and running
2. AnkiConnect plugin installed in Anki
3. A deck named "日文學習" created in Anki

### Installing AnkiConnect

If you don't have AnkiConnect installed:

1. In Anki, go to "Tools" > "Add-ons"
2. Click "Get Add-ons"
3. Enter the code: 2055492159
4. Restart Anki

### Creating the Required Deck

If you don't have a deck named "日文學習":

1. In Anki, click "Create Deck"
2. Name the deck "日文學習"
3. Click "Save"

## Running the Tests

You can run the tests in two ways:

### Option 1: Using the test script

1. Make sure Anki is running with AnkiConnect plugin installed
2. Run the test script:

```bash
go run test_anki.go
```

### Option 2: Using Go test directly

1. Make sure Anki is running with AnkiConnect plugin installed
2. Run the Go test command:

```bash
go test -v ./internal/anki
```

## What the Tests Cover

The tests verify the following functionality:

1. **Connection to Anki**
   - Verifies that the client can connect to Anki
   - Checks the connection status

2. **Deck Operations**
   - Lists available decks
   - Checks if a specific deck exists
   - Creates a new test deck
   - Verifies the test deck was created

3. **Model Operations**
   - Lists available models
   - Creates a new test model
   - Verifies the test model was created
   - Gets the field names for the test model

4. **Note Operations**
   - Creates a test deck for notes
   - Adds a single note
   - Adds multiple notes

5. **Converter Functionality**
   - Creates a test deck for the converter
   - Defines a test model
   - Ensures the model exists
   - Converts a struct to a note
   - Creates a card from a model
   - Converts multiple structs to notes
   - Creates cards from models

## Test Results

When the tests run successfully, you should see output similar to the following:

```
=== RUN   TestAnkiConnection
    anki_test.go:31: Successfully connected to Anki Connect. Version: 6.0
--- PASS: TestAnkiConnection (0.12s)
=== RUN   TestDeckOperations
    anki_test.go:48: Available decks: [日文學習 ...]
    anki_test.go:58: Deck '日文學習' exists
    anki_test.go:66: Created test deck '日文學習_test_20230615123456' with ID: 1234567890
    anki_test.go:76: Test deck '日文學習_test_20230615123456' exists
--- PASS: TestDeckOperations (0.25s)
...
```

## Troubleshooting

If the tests fail, check the following:

1. **Anki is not running**
   - Make sure Anki is open before running the tests

2. **AnkiConnect plugin is not installed**
   - Follow the installation instructions above

3. **The required deck does not exist**
   - Create a deck named "日文學習" in Anki

4. **Connection issues**
   - The tests assume AnkiConnect is running on http://localhost:8765
   - If you've configured a different port, update the test configuration

5. **Permission issues**
   - Some Anki operations may require permissions
   - Check Anki for any permission dialogs that may appear during testing