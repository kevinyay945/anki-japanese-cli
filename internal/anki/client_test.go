package anki

import (
	"errors"
	"net/http"
	"testing"

	"anki-japanese-cli/internal/config"
)

func TestClient_Ping(t *testing.T) {
	tests := []struct {
		name        string
		mockStatus  int
		mockBody    string
		mockErr     error
		expectError bool
	}{
		{
			name:        "Successful ping",
			mockStatus:  http.StatusOK,
			mockBody:    `{"result": 6, "error": null}`,
			mockErr:     nil,
			expectError: false,
		},
		{
			name:        "HTTP error",
			mockStatus:  http.StatusOK,
			mockBody:    ``,
			mockErr:     errors.New("connection error"),
			expectError: true,
		},
		{
			name:        "API error",
			mockStatus:  http.StatusOK,
			mockBody:    `{"result": null, "error": "some error"}`,
			mockErr:     nil,
			expectError: true,
		},
		{
			name:        "Non-200 status code",
			mockStatus:  http.StatusInternalServerError,
			mockBody:    ``,
			mockErr:     nil,
			expectError: true,
		},
		{
			name:        "Invalid JSON response",
			mockStatus:  http.StatusOK,
			mockBody:    `{invalid json}`,
			mockErr:     nil,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a mock HTTP client
			mockClient := NewMockHTTPClient(tt.mockStatus, tt.mockBody, tt.mockErr)

			// Create an Anki client with the mock HTTP client
			cfg := &config.AnkiConfig{
				ConnectURL: "http://localhost:8765",
				DeckName:   "test",
			}
			client := NewClientWithHTTPClient(cfg, mockClient)

			// Call the Ping method
			err := client.Ping()

			// Check if the error matches the expectation
			if (err != nil) != tt.expectError {
				t.Errorf("Ping() error = %v, expectError %v", err, tt.expectError)
			}
		})
	}
}

func TestClient_Call(t *testing.T) {
	tests := []struct {
		name        string
		action      string
		params      interface{}
		mockStatus  int
		mockBody    string
		mockErr     error
		expectError bool
		expectResult interface{}
	}{
		{
			name:        "Successful call",
			action:      "version",
			params:      nil,
			mockStatus:  http.StatusOK,
			mockBody:    `{"result": 6, "error": null}`,
			mockErr:     nil,
			expectError: false,
			expectResult: float64(6), // JSON numbers are parsed as float64
		},
		{
			name:        "HTTP error",
			action:      "version",
			params:      nil,
			mockStatus:  http.StatusOK,
			mockBody:    ``,
			mockErr:     errors.New("connection error"),
			expectError: true,
			expectResult: nil,
		},
		{
			name:        "API error",
			action:      "version",
			params:      nil,
			mockStatus:  http.StatusOK,
			mockBody:    `{"result": null, "error": "some error"}`,
			mockErr:     nil,
			expectError: true,
			expectResult: nil,
		},
		{
			name:        "Non-200 status code",
			action:      "version",
			params:      nil,
			mockStatus:  http.StatusInternalServerError,
			mockBody:    ``,
			mockErr:     nil,
			expectError: true,
			expectResult: nil,
		},
		{
			name:        "Invalid JSON response",
			action:      "version",
			params:      nil,
			mockStatus:  http.StatusOK,
			mockBody:    `{invalid json}`,
			mockErr:     nil,
			expectError: true,
			expectResult: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a mock HTTP client
			mockClient := NewMockHTTPClient(tt.mockStatus, tt.mockBody, tt.mockErr)

			// Create an Anki client with the mock HTTP client
			cfg := &config.AnkiConfig{
				ConnectURL: "http://localhost:8765",
				DeckName:   "test",
			}
			client := NewClientWithHTTPClient(cfg, mockClient)

			// Call the Call method
			result, err := client.Call(tt.action, tt.params)

			// Check if the error matches the expectation
			if (err != nil) != tt.expectError {
				t.Errorf("Call() error = %v, expectError %v", err, tt.expectError)
			}

			// Check if the result matches the expectation
			if !tt.expectError && result != tt.expectResult {
				t.Errorf("Call() result = %v, expectResult %v", result, tt.expectResult)
			}
		})
	}
}

func TestClient_DeckExists(t *testing.T) {
	tests := []struct {
		name        string
		deckName    string
		mockStatus  int
		mockBody    string
		mockErr     error
		expectError bool
		expectExists bool
	}{
		{
			name:        "Deck exists",
			deckName:    "test",
			mockStatus:  http.StatusOK,
			mockBody:    `{"result": ["test", "other"], "error": null}`,
			mockErr:     nil,
			expectError: false,
			expectExists: true,
		},
		{
			name:        "Deck does not exist",
			deckName:    "nonexistent",
			mockStatus:  http.StatusOK,
			mockBody:    `{"result": ["test", "other"], "error": null}`,
			mockErr:     nil,
			expectError: false,
			expectExists: false,
		},
		{
			name:        "API error",
			deckName:    "test",
			mockStatus:  http.StatusOK,
			mockBody:    `{"result": null, "error": "some error"}`,
			mockErr:     nil,
			expectError: true,
			expectExists: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a mock HTTP client
			mockClient := NewMockHTTPClient(tt.mockStatus, tt.mockBody, tt.mockErr)

			// Create an Anki client with the mock HTTP client
			cfg := &config.AnkiConfig{
				ConnectURL: "http://localhost:8765",
				DeckName:   "test",
			}
			client := NewClientWithHTTPClient(cfg, mockClient)

			// Call the DeckExists method
			exists, err := client.DeckExists(tt.deckName)

			// Check if the error matches the expectation
			if (err != nil) != tt.expectError {
				t.Errorf("DeckExists() error = %v, expectError %v", err, tt.expectError)
			}

			// Check if the result matches the expectation
			if !tt.expectError && exists != tt.expectExists {
				t.Errorf("DeckExists() exists = %v, expectExists %v", exists, tt.expectExists)
			}
		})
	}
}

func TestClient_CreateDeck(t *testing.T) {
	tests := []struct {
		name        string
		deckName    string
		mockStatus  int
		mockBody    string
		mockErr     error
		expectError bool
	}{
		{
			name:        "Create deck success",
			deckName:    "test",
			mockStatus:  http.StatusOK,
			mockBody:    `{"result": 1234, "error": null}`,
			mockErr:     nil,
			expectError: false,
		},
		{
			name:        "API error",
			deckName:    "test",
			mockStatus:  http.StatusOK,
			mockBody:    `{"result": null, "error": "some error"}`,
			mockErr:     nil,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a mock HTTP client
			mockClient := NewMockHTTPClient(tt.mockStatus, tt.mockBody, tt.mockErr)

			// Create an Anki client with the mock HTTP client
			cfg := &config.AnkiConfig{
				ConnectURL: "http://localhost:8765",
				DeckName:   "test",
			}
			client := NewClientWithHTTPClient(cfg, mockClient)

			// Call the CreateDeck method
			deckID, err := client.CreateDeck(tt.deckName)

			// Check if the error matches the expectation
			if (err != nil) != tt.expectError {
				t.Errorf("CreateDeck() error = %v, expectError %v", err, tt.expectError)
			}

			// Check if the result is valid
			if !tt.expectError && deckID <= 0 {
				t.Errorf("CreateDeck() deckID = %v, expected > 0", deckID)
			}
		})
	}
}

func TestClient_AddNote(t *testing.T) {
	tests := []struct {
		name        string
		note        NoteInfo
		mockStatus  int
		mockBody    string
		mockErr     error
		expectError bool
	}{
		{
			name: "Add note success",
			note: NoteInfo{
				DeckName:  "test",
				ModelName: "Basic",
				Fields: map[string]string{
					"Front": "Test front",
					"Back":  "Test back",
				},
				Tags: []string{"test"},
			},
			mockStatus:  http.StatusOK,
			mockBody:    `{"result": 1234, "error": null}`,
			mockErr:     nil,
			expectError: false,
		},
		{
			name: "API error",
			note: NoteInfo{
				DeckName:  "test",
				ModelName: "Basic",
				Fields: map[string]string{
					"Front": "Test front",
					"Back":  "Test back",
				},
				Tags: []string{"test"},
			},
			mockStatus:  http.StatusOK,
			mockBody:    `{"result": null, "error": "some error"}`,
			mockErr:     nil,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a mock HTTP client
			mockClient := NewMockHTTPClient(tt.mockStatus, tt.mockBody, tt.mockErr)

			// Create an Anki client with the mock HTTP client
			cfg := &config.AnkiConfig{
				ConnectURL: "http://localhost:8765",
				DeckName:   "test",
			}
			client := NewClientWithHTTPClient(cfg, mockClient)

			// Call the AddNote method
			noteID, err := client.AddNote(tt.note)

			// Check if the error matches the expectation
			if (err != nil) != tt.expectError {
				t.Errorf("AddNote() error = %v, expectError %v", err, tt.expectError)
			}

			// Check if the result is valid
			if !tt.expectError && noteID <= 0 {
				t.Errorf("AddNote() noteID = %v, expected > 0", noteID)
			}
		})
	}
}

func TestClient_AddNotes(t *testing.T) {
	tests := []struct {
		name        string
		notes       []NoteInfo
		mockStatus  int
		mockBody    string
		mockErr     error
		expectError bool
	}{
		{
			name: "Add notes success",
			notes: []NoteInfo{
				{
					DeckName:  "test",
					ModelName: "Basic",
					Fields: map[string]string{
						"Front": "Test front 1",
						"Back":  "Test back 1",
					},
					Tags: []string{"test"},
				},
				{
					DeckName:  "test",
					ModelName: "Basic",
					Fields: map[string]string{
						"Front": "Test front 2",
						"Back":  "Test back 2",
					},
					Tags: []string{"test"},
				},
			},
			mockStatus:  http.StatusOK,
			mockBody:    `{"result": [1234, 5678], "error": null}`,
			mockErr:     nil,
			expectError: false,
		},
		{
			name: "API error",
			notes: []NoteInfo{
				{
					DeckName:  "test",
					ModelName: "Basic",
					Fields: map[string]string{
						"Front": "Test front",
						"Back":  "Test back",
					},
					Tags: []string{"test"},
				},
			},
			mockStatus:  http.StatusOK,
			mockBody:    `{"result": null, "error": "some error"}`,
			mockErr:     nil,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a mock HTTP client
			mockClient := NewMockHTTPClient(tt.mockStatus, tt.mockBody, tt.mockErr)

			// Create an Anki client with the mock HTTP client
			cfg := &config.AnkiConfig{
				ConnectURL: "http://localhost:8765",
				DeckName:   "test",
			}
			client := NewClientWithHTTPClient(cfg, mockClient)

			// Call the AddNotes method
			noteIDs, err := client.AddNotes(tt.notes)

			// Check if the error matches the expectation
			if (err != nil) != tt.expectError {
				t.Errorf("AddNotes() error = %v, expectError %v", err, tt.expectError)
			}

			// Check if the result is valid
			if !tt.expectError {
				if len(noteIDs) != len(tt.notes) {
					t.Errorf("AddNotes() returned %d IDs, expected %d", len(noteIDs), len(tt.notes))
				}
				for i, id := range noteIDs {
					if id <= 0 {
						t.Errorf("AddNotes() noteID[%d] = %v, expected > 0", i, id)
					}
				}
			}
		})
	}
}

func TestClient_ModelExists(t *testing.T) {
	tests := []struct {
		name        string
		modelName   string
		mockStatus  int
		mockBody    string
		mockErr     error
		expectError bool
		expectExists bool
	}{
		{
			name:        "Model exists",
			modelName:   "Basic",
			mockStatus:  http.StatusOK,
			mockBody:    `{"result": ["Basic", "Other"], "error": null}`,
			mockErr:     nil,
			expectError: false,
			expectExists: true,
		},
		{
			name:        "Model does not exist",
			modelName:   "NonExistent",
			mockStatus:  http.StatusOK,
			mockBody:    `{"result": ["Basic", "Other"], "error": null}`,
			mockErr:     nil,
			expectError: false,
			expectExists: false,
		},
		{
			name:        "API error",
			modelName:   "Basic",
			mockStatus:  http.StatusOK,
			mockBody:    `{"result": null, "error": "some error"}`,
			mockErr:     nil,
			expectError: true,
			expectExists: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a mock HTTP client
			mockClient := NewMockHTTPClient(tt.mockStatus, tt.mockBody, tt.mockErr)

			// Create an Anki client with the mock HTTP client
			cfg := &config.AnkiConfig{
				ConnectURL: "http://localhost:8765",
				DeckName:   "test",
			}
			client := NewClientWithHTTPClient(cfg, mockClient)

			// Call the ModelExists method
			exists, err := client.ModelExists(tt.modelName)

			// Check if the error matches the expectation
			if (err != nil) != tt.expectError {
				t.Errorf("ModelExists() error = %v, expectError %v", err, tt.expectError)
			}

			// Check if the result matches the expectation
			if !tt.expectError && exists != tt.expectExists {
				t.Errorf("ModelExists() exists = %v, expectExists %v", exists, tt.expectExists)
			}
		})
	}
}

func TestClient_CreateModel(t *testing.T) {
	tests := []struct {
		name        string
		model       ModelConfig
		mockStatus  int
		mockBody    string
		mockErr     error
		expectError bool
	}{
		{
			name: "Create model success",
			model: ModelConfig{
				ModelName: "TestModel",
				InOrderFields: []string{"Front", "Back"},
				CSS: ".card { font-family: arial; }",
				CardTemplates: []CardTemplateConfig{
					{
						Name: "Card 1",
						Front: "{{Front}}",
						Back: "{{Front}}<hr>{{Back}}",
					},
				},
			},
			mockStatus:  http.StatusOK,
			mockBody:    `{"result": null, "error": null}`,
			mockErr:     nil,
			expectError: false,
		},
		{
			name: "API error",
			model: ModelConfig{
				ModelName: "TestModel",
				InOrderFields: []string{"Front", "Back"},
				CSS: ".card { font-family: arial; }",
				CardTemplates: []CardTemplateConfig{
					{
						Name: "Card 1",
						Front: "{{Front}}",
						Back: "{{Front}}<hr>{{Back}}",
					},
				},
			},
			mockStatus:  http.StatusOK,
			mockBody:    `{"result": null, "error": "some error"}`,
			mockErr:     nil,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a mock HTTP client
			mockClient := NewMockHTTPClient(tt.mockStatus, tt.mockBody, tt.mockErr)

			// Create an Anki client with the mock HTTP client
			cfg := &config.AnkiConfig{
				ConnectURL: "http://localhost:8765",
				DeckName:   "test",
			}
			client := NewClientWithHTTPClient(cfg, mockClient)

			// Call the CreateModel method
			err := client.CreateModel(tt.model)

			// Check if the error matches the expectation
			if (err != nil) != tt.expectError {
				t.Errorf("CreateModel() error = %v, expectError %v", err, tt.expectError)
			}
		})
	}
}