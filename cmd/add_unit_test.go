package cmd

import (
	"bytes"
	"errors"
	"os"
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

// TestAddCommandUnit tests the add command with mock Anki client
func TestAddCommandUnit(t *testing.T) {
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
			args:       []string{"add"},
			mockClient: NewMockAnkiClient(),
			wantErr:    true,
			contains:   []string{},
		},
		{
			name:       "Invalid card type",
			args:       []string{"add", "invalid", "--deckName=test"},
			mockClient: NewMockAnkiClient(),
			wantErr:    true,
			contains: []string{
				"不支援的卡片類型",
			},
		},
		{
			name:       "Missing deck name",
			args:       []string{"add", "verb"},
			mockClient: NewMockAnkiClient(),
			wantErr:    true,
			contains: []string{
				"請指定目標牌組名稱",
			},
		},
		{
			name:       "Missing JSON data",
			args:       []string{"add", "verb", "--deckName=test"},
			mockClient: NewMockAnkiClient(),
			wantErr:    true,
			contains: []string{
				"請提供卡片資料",
			},
		},
		{
			name:       "Invalid JSON data",
			args:       []string{"add", "verb", "--deckName=test", "--json={invalid}"},
			mockClient: NewMockAnkiClient(),
			wantErr:    true,
			contains: []string{
				"JSON 解析失敗",
			},
		},
		{
			name: "Anki connection error",
			args: []string{
				"add", "verb",
				"--deckName=test",
				"--json={\"核心單字\":\"飲む\", \"詞性分類\":\"五段動詞\", \"核心意義\":\"喝\", \"發音\":\"のむ\", \"重音\":\"1\", \"常用變化\":\"飲みます、飲んで\", \"情境例句\":\"水を飲む\", \"例句翻譯\":\"喝水\"}",
			},
			mockClient: NewMockAnkiClientWithError(errors.New("connection error")),
			wantErr:    true,
			contains: []string{
				"無法連線到 Anki",
			},
		},
		{
			name: "Model does not exist",
			args: []string{
				"add", "verb",
				"--deckName=test",
				"--json={\"核心單字\":\"飲む\", \"詞性分類\":\"五段動詞\", \"核心意義\":\"喝\", \"發音\":\"のむ\", \"重音\":\"1\", \"常用變化\":\"飲みます、飲んで\", \"情境例句\":\"水を飲む\", \"例句翻譯\":\"喝水\"}",
			},
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
			wantErr: true,
			contains: []string{
				"模型 'Japanese Verb' 不存在",
			},
		},
		{
			name: "Valid verb card",
			args: []string{
				"add", "verb",
				"--deckName=test",
				"--json={\"核心單字\":\"飲む\", \"詞性分類\":\"五段動詞\", \"核心意義\":\"喝\", \"發音\":\"のむ\", \"重音\":\"1\", \"常用變化\":\"飲みます、飲んで\", \"情境例句\":\"水を飲む\", \"例句翻譯\":\"喝水\"}",
			},
			mockClient: NewMockAnkiClient(),
			wantErr:    false,
			contains: []string{
				"成功連線到 Anki",
				"牌組 'test' 已就緒",
				"驗證 1 張卡片資料",
				"成功新增卡片",
			},
		},
		{
			name: "Valid adjective card",
			args: []string{
				"add", "adjective",
				"--deckName=test",
				"--json={\"核心單字\":\"美しい\", \"詞性分類\":\"い形容詞\", \"核心意義\":\"美麗的\", \"發音\":\"うつくしい\", \"重音\":\"4\", \"主要變化\":\"美しくない、美しかった\", \"情境例句\":\"美しい花\", \"例句翻譯\":\"美麗的花\"}",
			},
			mockClient: NewMockAnkiClient(),
			wantErr:    false,
			contains: []string{
				"成功連線到 Anki",
				"牌組 'test' 已就緒",
				"驗證 1 張卡片資料",
				"成功新增卡片",
			},
		},
		{
			name: "Valid normal card",
			args: []string{
				"add", "normal",
				"--deckName=test",
				"--json={\"核心單字\":\"猫\", \"詞性\":\"名詞\", \"核心意義\":\"貓\", \"發音\":\"ねこ\", \"重音\":\"2\", \"情境例句\":\"猫が屋根の上にいる\", \"例句翻譯\":\"貓在屋頂上\"}",
			},
			mockClient: NewMockAnkiClient(),
			wantErr:    false,
			contains: []string{
				"成功連線到 Anki",
				"牌組 'test' 已就緒",
				"驗證 1 張卡片資料",
				"成功新增卡片",
			},
		},
		{
			name: "Valid grammar card",
			args: []string{
				"add", "grammar",
				"--deckName=test",
				"--json={\"文法要點\":\"〜ても\", \"結構形式\":\"動詞て形 + も\", \"意義說明\":\"即使...也...\", \"使用時機\":\"表示讓步\", \"例句示範\":\"雨が降っても、行きます。\", \"例句翻譯\":\"即使下雨，也要去。\"}",
			},
			mockClient: NewMockAnkiClient(),
			wantErr:    false,
			contains: []string{
				"成功連線到 Anki",
				"牌組 'test' 已就緒",
				"驗證 1 張卡片資料",
				"成功新增卡片",
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

// TestAddCommandWithFileUnit tests the add command with file input
func TestAddCommandWithFileUnit(t *testing.T) {
	// Save the original GetAnkiClient function and restore it after the test
	originalGetAnkiClient := GetAnkiClient
	defer func() {
		GetAnkiClient = originalGetAnkiClient
	}()

	// Create a temporary file with JSON data
	tempFile, err := os.CreateTemp("", "test-*.json")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tempFile.Name())

	// Write JSON data to the file
	jsonData := `{"核心單字":"飲む", "詞性分類":"五段動詞", "核心意義":"喝", "發音":"のむ", "重音":"1", "常用變化":"飲みます、飲んで", "情境例句":"水を飲む", "例句翻譯":"喝水"}`
	if _, err := tempFile.Write([]byte(jsonData)); err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}
	if err := tempFile.Close(); err != nil {
		t.Fatalf("Failed to close temp file: %v", err)
	}

	// Create a temporary file with batch JSON data
	batchTempFile, err := os.CreateTemp("", "test-batch-*.json")
	if err != nil {
		t.Fatalf("Failed to create batch temp file: %v", err)
	}
	defer os.Remove(batchTempFile.Name())

	// Write batch JSON data to the file
	batchJsonData := `[
		{"核心單字":"飲む", "詞性分類":"五段動詞", "核心意義":"喝", "發音":"のむ", "重音":"1", "常用變化":"飲みます、飲んで", "情境例句":"水を飲む", "例句翻譯":"喝水"},
		{"核心單字":"食べる", "詞性分類":"一段動詞", "核心意義":"吃", "發音":"たべる", "重音":"2", "常用變化":"食べます、食べて", "情境例句":"ご飯を食べる", "例句翻譯":"吃飯"}
	]`
	if _, err := batchTempFile.Write([]byte(batchJsonData)); err != nil {
		t.Fatalf("Failed to write to batch temp file: %v", err)
	}
	if err := batchTempFile.Close(); err != nil {
		t.Fatalf("Failed to close batch temp file: %v", err)
	}

	// Test cases
	testCases := []struct {
		name       string
		args       []string
		mockClient *MockAnkiClient
		wantErr    bool
		contains   []string
	}{
		{
			name: "Add from file",
			args: []string{
				"add", "verb",
				"--deckName=test",
				"--file=" + tempFile.Name(),
			},
			mockClient: NewMockAnkiClient(),
			wantErr:    false,
			contains: []string{
				"成功連線到 Anki",
				"牌組 'test' 已就緒",
				"從檔案",
				"驗證 1 張卡片資料",
				"成功新增卡片",
			},
		},
		{
			name: "Add batch from file",
			args: []string{
				"add", "verb",
				"--deckName=test",
				"--file=" + batchTempFile.Name(),
				"--batch",
			},
			mockClient: NewMockAnkiClient(),
			wantErr:    false,
			contains: []string{
				"成功連線到 Anki",
				"牌組 'test' 已就緒",
				"從檔案",
				"驗證 2 張卡片資料",
				"正在批次新增 2 張卡片",
				"成功新增 2/2 張卡片",
			},
		},
		{
			name: "File not found",
			args: []string{
				"add", "verb",
				"--deckName=test",
				"--file=nonexistent.json",
			},
			mockClient: NewMockAnkiClient(),
			wantErr:    true,
			contains: []string{
				"無法讀取檔案",
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
		})
	}
}
