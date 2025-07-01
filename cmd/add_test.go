package cmd

import (
	"bytes"
	"os"
	"strings"
	"testing"

	"anki-japanese-cli/internal/anki"
	"anki-japanese-cli/internal/config"
)

// TestAddCommand tests the add command
func TestAddCommand(t *testing.T) {
	// Skip if not running integration tests
	if os.Getenv("ANKI_INTEGRATION_TEST") != "true" {
		t.Skip("Skipping integration test. Set ANKI_INTEGRATION_TEST=true to run")
	}

	// Load config
	cfg, err := config.LoadConfig()
	if err != nil {
		t.Fatalf("Failed to load config: %v", err)
	}

	// Create Anki client
	client := anki.NewClient(&cfg.Anki)

	// Check Anki connection
	if err := client.Ping(); err != nil {
		t.Fatalf("Failed to connect to Anki: %v", err)
	}

	// Test cases
	testCases := []struct {
		name     string
		args     []string
		wantErr  bool
		contains []string
	}{
		{
			name:    "Missing card type",
			args:    []string{"add"},
			wantErr: true,
		},
		{
			name:    "Invalid card type",
			args:    []string{"add", "invalid", "--deckName=test"},
			wantErr: true,
			contains: []string{
				"不支援的卡片類型",
			},
		},
		{
			name:    "Missing deck name",
			args:    []string{"add", "verb"},
			wantErr: true,
			contains: []string{
				"請指定目標牌組名稱",
			},
		},
		{
			name:    "Missing JSON data",
			args:    []string{"add", "verb", "--deckName=test"},
			wantErr: true,
			contains: []string{
				"請提供卡片資料",
			},
		},
		{
			name:    "Invalid JSON data",
			args:    []string{"add", "verb", "--deckName=test", "--json={invalid}"},
			wantErr: true,
			contains: []string{
				"JSON 解析失敗",
			},
		},
		{
			name: "Valid verb card",
			args: []string{
				"add", "verb",
				"--deckName=日文動詞",
				"--json={\"核心單字\":\"飲む\", \"詞性分類\":\"五段動詞\", \"核心意義\":\"喝\", \"發音\":\"のむ\", \"重音\":\"1\", \"常用變化\":\"飲みます、飲んで\", \"情境例句\":\"水を飲む\", \"例句翻譯\":\"喝水\"}",
			},
			wantErr: false,
			contains: []string{
				"成功連線到 Anki",
				"牌組 '日文動詞' 已就緒",
				"驗證 1 張卡片資料",
				"成功新增卡片",
			},
		},
		{
			name: "Valid adjective card",
			args: []string{
				"add", "adjective",
				"--deckName=日文形容詞",
				"--json={\"核心單字\":\"美しい\", \"詞性分類\":\"い形容詞\", \"核心意義\":\"美麗的\", \"發音\":\"うつくしい\", \"重音\":\"4\", \"主要變化\":\"美しくない、美しかった\", \"情境例句\":\"美しい花\", \"例句翻譯\":\"美麗的花\"}",
			},
			wantErr: false,
			contains: []string{
				"成功連線到 Anki",
				"牌組 '日文形容詞' 已就緒",
				"驗證 1 張卡片資料",
				"成功新增卡片",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
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
