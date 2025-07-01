package cmd

import (
	"bytes"
	"os"
	"strings"
	"testing"

	"anki-japanese-cli/internal/anki"
	"anki-japanese-cli/internal/config"
)

// TestInitCommand tests the init command
func TestInitCommand(t *testing.T) {
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
			name:    "Invalid card type",
			args:    []string{"init", "invalid"},
			wantErr: true,
			contains: []string{
				"不支援的卡片類型",
			},
		},
		{
			name:    "Valid verb card type",
			args:    []string{"init", "verb"},
			wantErr: false,
			contains: []string{
				"成功連線到 Anki",
				"牌組 '日文動詞' 已就緒",
				"初始化完成",
			},
		},
		{
			name:    "Valid adjective card type",
			args:    []string{"init", "adjective"},
			wantErr: false,
			contains: []string{
				"成功連線到 Anki",
				"牌組 '日文形容詞' 已就緒",
				"初始化完成",
			},
		},
		{
			name:    "Valid normal card type",
			args:    []string{"init", "normal"},
			wantErr: false,
			contains: []string{
				"成功連線到 Anki",
				"牌組 '日文單字' 已就緒",
				"初始化完成",
			},
		},
		{
			name:    "Valid grammar card type",
			args:    []string{"init", "grammar"},
			wantErr: false,
			contains: []string{
				"成功連線到 Anki",
				"牌組 '日文文法' 已就緒",
				"初始化完成",
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
