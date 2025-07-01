package cmd

import (
	"anki-japanese-cli/internal/anki"
	"anki-japanese-cli/internal/config"
)

// MockAnkiClient is a mock implementation of the Anki client for testing
type MockAnkiClient struct {
	// PingFunc will be executed when Ping is called
	PingFunc func() error

	// DeckExistsFunc will be executed when DeckExists is called
	DeckExistsFunc func(deckName string) (bool, error)

	// CreateDeckFunc will be executed when CreateDeck is called
	CreateDeckFunc func(deckName string) (int64, error)

	// ModelExistsFunc will be executed when ModelExists is called
	ModelExistsFunc func(modelName string) (bool, error)

	// CreateModelFunc will be executed when CreateModel is called
	CreateModelFunc func(model anki.ModelConfig) error

	// AddNoteFunc will be executed when AddNote is called
	AddNoteFunc func(note anki.NoteInfo) (int64, error)

	// AddNotesFunc will be executed when AddNotes is called
	AddNotesFunc func(notes []anki.NoteInfo) ([]int64, error)

	// EnsureDeckExistsFunc will be executed when EnsureDeckExists is called
	EnsureDeckExistsFunc func(deckName string) error
}

// Ping implements the Ping method of the Anki client
func (m *MockAnkiClient) Ping() error {
	if m.PingFunc != nil {
		return m.PingFunc()
	}
	return nil
}

// DeckExists implements the DeckExists method of the Anki client
func (m *MockAnkiClient) DeckExists(deckName string) (bool, error) {
	if m.DeckExistsFunc != nil {
		return m.DeckExistsFunc(deckName)
	}
	return true, nil
}

// CreateDeck implements the CreateDeck method of the Anki client
func (m *MockAnkiClient) CreateDeck(deckName string) (int64, error) {
	if m.CreateDeckFunc != nil {
		return m.CreateDeckFunc(deckName)
	}
	return 1234, nil
}

// ModelExists implements the ModelExists method of the Anki client
func (m *MockAnkiClient) ModelExists(modelName string) (bool, error) {
	if m.ModelExistsFunc != nil {
		return m.ModelExistsFunc(modelName)
	}
	return true, nil
}

// CreateModel implements the CreateModel method of the Anki client
func (m *MockAnkiClient) CreateModel(model anki.ModelConfig) error {
	if m.CreateModelFunc != nil {
		return m.CreateModelFunc(model)
	}
	return nil
}

// AddNote implements the AddNote method of the Anki client
func (m *MockAnkiClient) AddNote(note anki.NoteInfo) (int64, error) {
	if m.AddNoteFunc != nil {
		return m.AddNoteFunc(note)
	}
	return 1234, nil
}

// AddNotes implements the AddNotes method of the Anki client
func (m *MockAnkiClient) AddNotes(notes []anki.NoteInfo) ([]int64, error) {
	if m.AddNotesFunc != nil {
		return m.AddNotesFunc(notes)
	}
	ids := make([]int64, len(notes))
	for i := range notes {
		ids[i] = int64(i + 1)
	}
	return ids, nil
}

// EnsureDeckExists implements the EnsureDeckExists method of the Anki client
func (m *MockAnkiClient) EnsureDeckExists(deckName string) error {
	if m.EnsureDeckExistsFunc != nil {
		return m.EnsureDeckExistsFunc(deckName)
	}
	return nil
}

// NewMockAnkiClient creates a new mock Anki client with default success responses
func NewMockAnkiClient() *MockAnkiClient {
	return &MockAnkiClient{}
}

// NewMockAnkiClientWithError creates a new mock Anki client that returns an error for all methods
func NewMockAnkiClientWithError(err error) *MockAnkiClient {
	return &MockAnkiClient{
		PingFunc: func() error {
			return err
		},
		DeckExistsFunc: func(deckName string) (bool, error) {
			return false, err
		},
		CreateDeckFunc: func(deckName string) (int64, error) {
			return 0, err
		},
		ModelExistsFunc: func(modelName string) (bool, error) {
			return false, err
		},
		CreateModelFunc: func(model anki.ModelConfig) error {
			return err
		},
		AddNoteFunc: func(note anki.NoteInfo) (int64, error) {
			return 0, err
		},
		AddNotesFunc: func(notes []anki.NoteInfo) ([]int64, error) {
			return nil, err
		},
		EnsureDeckExistsFunc: func(deckName string) error {
			return err
		},
	}
}

// NewMockAnkiClientWithCustomResponses creates a new mock Anki client with custom responses
func NewMockAnkiClientWithCustomResponses(
	pingErr error,
	deckExists bool, deckExistsErr error,
	createDeckID int64, createDeckErr error,
	modelExists bool, modelExistsErr error,
	createModelErr error,
	addNoteID int64, addNoteErr error,
	addNotesIDs []int64, addNotesErr error,
	ensureDeckExistsErr error,
) *MockAnkiClient {
	return &MockAnkiClient{
		PingFunc: func() error {
			return pingErr
		},
		DeckExistsFunc: func(deckName string) (bool, error) {
			return deckExists, deckExistsErr
		},
		CreateDeckFunc: func(deckName string) (int64, error) {
			return createDeckID, createDeckErr
		},
		ModelExistsFunc: func(modelName string) (bool, error) {
			return modelExists, modelExistsErr
		},
		CreateModelFunc: func(model anki.ModelConfig) error {
			return createModelErr
		},
		AddNoteFunc: func(note anki.NoteInfo) (int64, error) {
			return addNoteID, addNoteErr
		},
		AddNotesFunc: func(notes []anki.NoteInfo) ([]int64, error) {
			return addNotesIDs, addNotesErr
		},
		EnsureDeckExistsFunc: func(deckName string) error {
			return ensureDeckExistsErr
		},
	}
}

// GetAnkiClient returns the Anki client to use for testing
// This function can be monkey patched in tests to return a mock client
var GetAnkiClient = func(cfg *config.AnkiConfig) interface{} {
	return anki.NewClient(cfg)
}

// ResetGetAnkiClient resets the GetAnkiClient function to its default implementation
func ResetGetAnkiClient() {
	GetAnkiClient = func(cfg *config.AnkiConfig) interface{} {
		return anki.NewClient(cfg)
	}
}

// SetMockAnkiClient sets the GetAnkiClient function to return the given mock client
func SetMockAnkiClient(mockClient *MockAnkiClient) {
	GetAnkiClient = func(cfg *config.AnkiConfig) interface{} {
		return mockClient
	}
}
