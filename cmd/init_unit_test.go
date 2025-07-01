package cmd

import (
	"bytes"
	"errors"
	"strings"
	"testing"

	"anki-japanese-cli/internal/anki"
	"anki-japanese-cli/internal/config"
)

// Ensure imports are used
var (
	_ = anki.NoteInfo{}
	_ = config.AnkiConfig{}
)

// TestInitCommandUnit tests the init command with mock Anki client
func TestInitCommandUnit(t *testing.T) {
	// Save the original GetAnkiClient function and restore it after the test
	originalGetAnkiClient := GetAnkiClient
	defer func() {
		GetAnkiClient = originalGetAnkiClient
	}()

	// Test cases
	testCases := []struct {
		name        string
		args        []string
		mockClient  *MockAnkiClient
		wantErr     bool
		contains    []string
		notContains []string
	}{
		{
			name:       "Missing card type",
			args:       []string{"init"},
			mockClient: NewMockAnkiClient(),
			wantErr:    true,
			contains:   []string{},
		},
		{
			name:       "Invalid card type",
			args:       []string{"init", "invalid"},
			mockClient: NewMockAnkiClient(),
			wantErr:    true,
			contains: []string{
				"不支援的卡片類型",
			},
		},
		{
			name:       "Anki connection error",
			args:       []string{"init", "verb"},
			mockClient: NewMockAnkiClientWithError(errors.New("connection error")),
			wantErr:    true,
			contains: []string{
				"無法連線到 Anki",
			},
		},
		{
			name: "Model already exists",
			args: []string{"init", "verb"},
			mockClient: NewMockAnkiClientWithCustomResponses(
				nil,       // pingErr
				true, nil, // deckExists, deckExistsErr
				0, nil, // createDeckID, createDeckErr
				true, nil, // modelExists, modelExistsErr
				nil,    // createModelErr
				0, nil, // addNoteID, addNoteErr
				nil, nil, // addNotesIDs, addNotesErr
				nil, // ensureDeckExistsErr
			),
			wantErr: false,
			contains: []string{
				"成功連線到 Anki",
				"模型 'Japanese Verb' 已存在",
				"牌組 '日文動詞' 已就緒",
				"初始化完成",
			},
		},
		{
			name: "Create model error",
			args: []string{"init", "verb"},
			mockClient: NewMockAnkiClientWithCustomResponses(
				nil,       // pingErr
				true, nil, // deckExists, deckExistsErr
				0, nil, // createDeckID, createDeckErr
				false, nil, // modelExists, modelExistsErr
				errors.New("create model error"), // createModelErr
				0, nil,                           // addNoteID, addNoteErr
				nil, nil, // addNotesIDs, addNotesErr
				nil, // ensureDeckExistsErr
			),
			wantErr: true,
			contains: []string{
				"無法建立模型",
			},
		},
		{
			name: "Create deck error",
			args: []string{"init", "verb"},
			mockClient: NewMockAnkiClientWithCustomResponses(
				nil,        // pingErr
				false, nil, // deckExists, deckExistsErr
				0, errors.New("create deck error"), // createDeckID, createDeckErr
				false, nil, // modelExists, modelExistsErr
				nil,    // createModelErr
				0, nil, // addNoteID, addNoteErr
				nil, nil, // addNotesIDs, addNotesErr
				errors.New("create deck error"), // ensureDeckExistsErr
			),
			wantErr: true,
			contains: []string{
				"無法建立牌組",
			},
		},
		{
			name: "Valid verb card type",
			args: []string{"init", "verb"},
			mockClient: NewMockAnkiClientWithCustomResponses(
				nil,       // pingErr
				true, nil, // deckExists, deckExistsErr
				0, nil, // createDeckID, createDeckErr
				false, nil, // modelExists, modelExistsErr
				nil,    // createModelErr
				0, nil, // addNoteID, addNoteErr
				nil, nil, // addNotesIDs, addNotesErr
				nil, // ensureDeckExistsErr
			),
			wantErr: false,
			contains: []string{
				"成功連線到 Anki",
				"成功建立模型 'Japanese Verb'",
				"牌組 '日文動詞' 已就緒",
				"初始化完成",
			},
		},
		{
			name: "Valid adjective card type",
			args: []string{"init", "adjective"},
			mockClient: NewMockAnkiClientWithCustomResponses(
				nil,       // pingErr
				true, nil, // deckExists, deckExistsErr
				0, nil, // createDeckID, createDeckErr
				false, nil, // modelExists, modelExistsErr
				nil,    // createModelErr
				0, nil, // addNoteID, addNoteErr
				nil, nil, // addNotesIDs, addNotesErr
				nil, // ensureDeckExistsErr
			),
			wantErr: false,
			contains: []string{
				"成功連線到 Anki",
				"成功建立模型 'Japanese Adjective'",
				"牌組 '日文形容詞' 已就緒",
				"初始化完成",
			},
		},
		{
			name: "Valid normal card type",
			args: []string{"init", "normal"},
			mockClient: NewMockAnkiClientWithCustomResponses(
				nil,       // pingErr
				true, nil, // deckExists, deckExistsErr
				0, nil, // createDeckID, createDeckErr
				false, nil, // modelExists, modelExistsErr
				nil,    // createModelErr
				0, nil, // addNoteID, addNoteErr
				nil, nil, // addNotesIDs, addNotesErr
				nil, // ensureDeckExistsErr
			),
			wantErr: false,
			contains: []string{
				"成功連線到 Anki",
				"成功建立模型 'Japanese Normal Word'",
				"牌組 '日文單字' 已就緒",
				"初始化完成",
			},
		},
		{
			name: "Valid grammar card type",
			args: []string{"init", "grammar"},
			mockClient: NewMockAnkiClientWithCustomResponses(
				nil,       // pingErr
				true, nil, // deckExists, deckExistsErr
				0, nil, // createDeckID, createDeckErr
				false, nil, // modelExists, modelExistsErr
				nil,    // createModelErr
				0, nil, // addNoteID, addNoteErr
				nil, nil, // addNotesIDs, addNotesErr
				nil, // ensureDeckExistsErr
			),
			wantErr: false,
			contains: []string{
				"成功連線到 Anki",
				"成功建立模型 'Japanese Grammar'",
				"牌組 '日文文法' 已就緒",
				"初始化完成",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Set the mock client
			SetMockAnkiClient(tc.mockClient)

			// Capture output
			buf := new(bytes.Buffer)
			rootCmd.SetOut(buf)
			rootCmd.SetErr(buf)

			// Set args
			rootCmd.SetArgs(tc.args)

			// Execute command
			err := rootCmd.Execute()

			// Check error
			if (err != nil) != tc.wantErr {
				t.Errorf("Execute() error = %v, wantErr %v", err, tc.wantErr)
				return
			}

			// Check output
			output := buf.String()
			for _, s := range tc.contains {
				if !strings.Contains(output, s) {
					t.Errorf("Output does not contain %q\nOutput: %s", s, output)
				}
			}
			for _, s := range tc.notContains {
				if strings.Contains(output, s) {
					t.Errorf("Output contains %q but should not\nOutput: %s", s, output)
				}
			}
		})
	}
}
