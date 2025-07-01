package anki

import (
	"testing"
	"time"

	"anki-japanese-cli/internal/config"
)

// TestAnkiConnection tests the connection to Anki
func TestAnkiConnection(t *testing.T) {
	// Create a new client with default configuration
	cfg := &config.AnkiConfig{
		ConnectURL: "http://localhost:8765",
		DeckName:   "日文學習",
	}
	client := NewClient(cfg)

	// Test connection
	err := client.Ping()
	if err != nil {
		t.Fatalf("Failed to connect to Anki: %v", err)
	}

	// Test connection status
	status := client.CheckConnection()
	if !status.Connected {
		t.Fatalf("Connection status check failed: %s", status.Error)
	}

	t.Logf("Successfully connected to Anki Connect. Version: %s", status.Version)
}

// TestDeckOperations tests deck operations
func TestDeckOperations(t *testing.T) {
	// Create a new client with default configuration
	cfg := &config.AnkiConfig{
		ConnectURL: "http://localhost:8765",
		DeckName:   "日文學習",
	}
	client := NewClient(cfg)

	// Test getting deck names
	deckNames, err := client.DeckNames()
	if err != nil {
		t.Fatalf("Failed to get deck names: %v", err)
	}
	t.Logf("Available decks: %v", deckNames)

	// Test if the deck exists or create it
	exists, err := client.DeckExists("日文學習")
	if err != nil {
		t.Fatalf("Failed to check if deck exists: %v", err)
	}
	if !exists {
		// Create the deck if it doesn't exist
		deckID, err := client.CreateDeck("日文學習")
		if err != nil {
			t.Fatalf("Failed to create deck '日文學習': %v", err)
		}
		t.Logf("Created deck '日文學習' with ID: %d", deckID)
	} else {
		t.Logf("Deck '日文學習' exists")
	}

	// Test creating a test deck
	testDeckName := "日文學習_test_" + time.Now().Format("20060102150405")
	deckID, err := client.CreateDeck(testDeckName)
	if err != nil {
		t.Fatalf("Failed to create test deck: %v", err)
	}
	t.Logf("Created test deck '%s' with ID: %d", testDeckName, deckID)

	// Verify the test deck was created
	exists, err = client.DeckExists(testDeckName)
	if err != nil {
		t.Fatalf("Failed to check if test deck exists: %v", err)
	}
	if !exists {
		t.Fatalf("Test deck '%s' was not created", testDeckName)
	}
	t.Logf("Test deck '%s' exists", testDeckName)
}

// TestModelOperations tests model operations
func TestModelOperations(t *testing.T) {
	// Create a new client with default configuration
	cfg := &config.AnkiConfig{
		ConnectURL: "http://localhost:8765",
		DeckName:   "日文學習",
	}
	client := NewClient(cfg)

	// Test getting model names
	modelNames, err := client.ModelNames()
	if err != nil {
		t.Fatalf("Failed to get model names: %v", err)
	}
	t.Logf("Available models: %v", modelNames)

	// Test creating a test model
	testModelName := "TestModel_" + time.Now().Format("20060102150405")
	modelConfig := ModelConfig{
		ModelName: testModelName,
		InOrderFields: []string{
			"Front",
			"Back",
		},
		CSS: ".card { font-family: arial; font-size: 20px; text-align: center; color: black; background-color: white; }",
		CardTemplates: []CardTemplateConfig{
			{
				Name:  "Card 1",
				Front: "{{Front}}",
				Back:  "{{Front}}<hr>{{Back}}",
			},
		},
	}

	err = client.CreateModel(modelConfig)
	if err != nil {
		t.Fatalf("Failed to create test model: %v", err)
	}
	t.Logf("Created test model '%s'", testModelName)

	// Verify the test model was created
	exists, err := client.ModelExists(testModelName)
	if err != nil {
		t.Fatalf("Failed to check if test model exists: %v", err)
	}
	if !exists {
		t.Fatalf("Test model '%s' was not created", testModelName)
	}
	t.Logf("Test model '%s' exists", testModelName)

	// Test getting model field names
	fieldNames, err := client.ModelFieldNames(testModelName)
	if err != nil {
		t.Fatalf("Failed to get model field names: %v", err)
	}
	t.Logf("Model '%s' fields: %v", testModelName, fieldNames)
}

// TestNoteOperations tests note operations
func TestNoteOperations(t *testing.T) {
	// Create a new client with default configuration
	cfg := &config.AnkiConfig{
		ConnectURL: "http://localhost:8765",
		DeckName:   "日文學習",
	}
	client := NewClient(cfg)

	// Create a test deck for notes
	testDeckName := "日文學習_test_notes_" + time.Now().Format("20060102150405")
	_, err := client.CreateDeck(testDeckName)
	if err != nil {
		t.Fatalf("Failed to create test deck: %v", err)
	}
	t.Logf("Created test deck '%s' for notes", testDeckName)

	// Use the Basic model which should exist in any Anki installation
	modelName := "Basic"
	exists, err := client.ModelExists(modelName)
	if err != nil {
		t.Fatalf("Failed to check if model exists: %v", err)
	}
	if !exists {
		t.Fatalf("Model '%s' does not exist", modelName)
	}

	// Get the field names for the Basic model
	fieldNames, err := client.ModelFieldNames(modelName)
	if err != nil {
		t.Fatalf("Failed to get field names for model '%s': %v", modelName, err)
	}
	t.Logf("Model '%s' fields: %v", modelName, fieldNames)

	// Ensure we have the expected fields
	if len(fieldNames) < 2 {
		t.Fatalf("Model '%s' does not have enough fields: %v", modelName, fieldNames)
	}

	// Create a map for the fields using the actual field names
	// Add a timestamp to make the note content unique for each test run
	timestamp := time.Now().Format("20060102150405")
	fields := make(map[string]string)
	for i, fieldName := range fieldNames {
		if i == 0 {
			fields[fieldName] = "テスト単語 " + timestamp
		} else if i == 1 {
			fields[fieldName] = "Test Word " + timestamp
		} else {
			// Set empty string for any additional fields
			fields[fieldName] = ""
		}
	}

	// Test adding a single note
	note := NoteInfo{
		DeckName:  testDeckName,
		ModelName: modelName,
		Fields:    fields,
		Tags:      []string{"test", "japanese"},
	}

	noteID, err := client.AddNote(note)
	if err != nil {
		t.Fatalf("Failed to add note: %v", err)
	}
	t.Logf("Added note with ID: %d", noteID)

	// Test adding multiple notes
	// Create fields for the first note using the same timestamp to make content unique
	fields1 := make(map[string]string)
	for i, fieldName := range fieldNames {
		if i == 0 {
			fields1[fieldName] = "こんにちは " + timestamp
		} else if i == 1 {
			fields1[fieldName] = "Hello " + timestamp
		} else {
			// Set empty string for any additional fields
			fields1[fieldName] = ""
		}
	}

	// Create fields for the second note using the same timestamp
	fields2 := make(map[string]string)
	for i, fieldName := range fieldNames {
		if i == 0 {
			fields2[fieldName] = "さようなら " + timestamp
		} else if i == 1 {
			fields2[fieldName] = "Goodbye " + timestamp
		} else {
			// Set empty string for any additional fields
			fields2[fieldName] = ""
		}
	}

	notes := []NoteInfo{
		{
			DeckName:  testDeckName,
			ModelName: modelName,
			Fields:    fields1,
			Tags:      []string{"test", "japanese", "greeting"},
		},
		{
			DeckName:  testDeckName,
			ModelName: modelName,
			Fields:    fields2,
			Tags:      []string{"test", "japanese", "greeting"},
		},
	}

	noteIDs, err := client.AddNotes(notes)
	if err != nil {
		t.Fatalf("Failed to add notes: %v", err)
	}
	t.Logf("Added %d notes with IDs: %v", len(noteIDs), noteIDs)
}

// TestConverter tests the converter functionality
func TestConverter(t *testing.T) {
	// Create a new client with default configuration
	cfg := &config.AnkiConfig{
		ConnectURL: "http://localhost:8765",
		DeckName:   "日文學習",
	}
	client := NewClient(cfg)
	converter := NewConverter(client)

	// Create a test deck for converter
	testDeckName := "日文學習_test_converter_" + time.Now().Format("20060102150405")
	_, err := client.CreateDeck(testDeckName)
	if err != nil {
		t.Fatalf("Failed to create test deck: %v", err)
	}
	t.Logf("Created test deck '%s' for converter", testDeckName)

	// Define a test model with only the fields we need
	testModelName := "TestConverterModel_" + time.Now().Format("20060102150405")
	fields := []string{"Word", "Reading", "Meaning"}
	css := ".card { font-family: arial; font-size: 20px; text-align: center; color: black; background-color: white; }"
	templates := []CardTemplateConfig{
		{
			Name:  "Card 1",
			Front: "{{Word}}",
			Back:  "{{Word}}<hr>{{Reading}}<br>{{Meaning}}",
		},
	}

	// Test ensuring model exists
	err = converter.EnsureModelExists(testModelName, fields, css, templates)
	if err != nil {
		t.Fatalf("Failed to ensure model exists: %v", err)
	}
	t.Logf("Ensured model '%s' exists", testModelName)

	// Define a test struct that matches the model
	type TestWord struct {
		Word    string `json:"Word"`
		Reading string `json:"Reading"`
		Meaning string `json:"Meaning"`
	}

	// Test converting a struct to a note
	testWord := TestWord{
		Word:    "食べる",
		Reading: "たべる",
		Meaning: "to eat",
	}

	note, err := converter.ConvertToNote(testWord, testDeckName, []string{"test", "converter"})
	if err != nil {
		t.Fatalf("Failed to convert struct to note: %v", err)
	}

	// Override the model name to use the one we created
	note.ModelName = testModelName

	t.Logf("Converted struct to note: %+v", note)

	// Test creating a card from a model by directly adding the note
	cardID, err := client.AddNote(note)
	if err != nil {
		t.Fatalf("Failed to create card from model: %v", err)
	}
	t.Logf("Created card with ID: %d", cardID)

	// Test converting multiple structs to notes
	testWords := []TestWord{
		{
			Word:    "飲む",
			Reading: "のむ",
			Meaning: "to drink",
		},
		{
			Word:    "走る",
			Reading: "はしる",
			Meaning: "to run",
		},
	}

	// Manually convert the structs to notes and set the model name
	var notes []NoteInfo
	for _, word := range testWords {
		note, err := converter.ConvertToNote(word, testDeckName, []string{"test", "converter", "batch"})
		if err != nil {
			t.Fatalf("Failed to convert struct to note: %v", err)
		}
		// Override the model name to use the one we created
		note.ModelName = testModelName
		notes = append(notes, note)
	}
	t.Logf("Converted %d structs to notes", len(notes))

	// Test creating cards by directly adding the notes
	var cardIDs []int64
	for _, note := range notes {
		cardID, err := client.AddNote(note)
		if err != nil {
			t.Fatalf("Failed to add note: %v", err)
		}
		cardIDs = append(cardIDs, cardID)
	}
	t.Logf("Created %d cards with IDs: %v", len(cardIDs), cardIDs)
}

// TestAll runs all the tests
// Note: This test is disabled because it causes duplicate model and note errors
// when run after the individual tests. The individual tests should be run instead.
func TestAll(t *testing.T) {
	t.Skip("Skipping TestAll to avoid duplicate model and note errors. Run individual tests instead.")

	t.Run("TestAnkiConnection", TestAnkiConnection)
	t.Run("TestDeckOperations", TestDeckOperations)
	t.Run("TestModelOperations", TestModelOperations)
	t.Run("TestNoteOperations", TestNoteOperations)
	t.Run("TestConverter", TestConverter)
}
