package templates

import (
	"strings"
	"testing"
)

func TestNewTemplateManager(t *testing.T) {
	// Create a new template manager
	manager, err := NewTemplateManager()
	
	// Check if the manager was created successfully
	if err != nil {
		t.Fatalf("Failed to create template manager: %v", err)
	}
	
	// Check if the manager is not nil
	if manager == nil {
		t.Fatal("Template manager is nil")
	}
	
	// Check if templates were loaded
	if len(manager.templates) == 0 {
		t.Fatal("No templates were loaded")
	}
	
	// Check if all expected templates are loaded
	expectedTemplates := []string{"verb", "adjective", "normal", "grammar"}
	for _, templateName := range expectedTemplates {
		template, exists := manager.templates[templateName]
		if !exists {
			t.Errorf("Template '%s' was not loaded", templateName)
			continue
		}
		
		if template.Front == nil {
			t.Errorf("Template '%s' front is nil", templateName)
		}
		
		if template.Back == nil {
			t.Errorf("Template '%s' back is nil", templateName)
		}
	}
}

func TestTemplateManager_RenderCardFront(t *testing.T) {
	// Create a new template manager
	manager, err := NewTemplateManager()
	if err != nil {
		t.Fatalf("Failed to create template manager: %v", err)
	}
	
	// Test data for different card types
	testCases := []struct {
		name     string
		cardType string
		data     map[string]interface{}
		wantErr  bool
		contains []string
	}{
		{
			name:     "Verb card front",
			cardType: "verb",
			data: map[string]interface{}{
				"情境例句": "水を飲む",
				"核心意義": "喝",
				"圖片提示": "http://example.com/image.jpg",
			},
			wantErr:  false,
			contains: []string{"水を飲む", "喝", "http://example.com/image.jpg"},
		},
		{
			name:     "Adjective card front",
			cardType: "adjective",
			data: map[string]interface{}{
				"情境例句": "美しい花",
				"核心意義": "美麗的",
			},
			wantErr:  false,
			contains: []string{"美しい花", "美麗的"},
		},
		{
			name:     "Normal card front",
			cardType: "normal",
			data: map[string]interface{}{
				"情境例句": "猫が屋根の上にいる",
				"核心意義": "貓",
				"圖片提示": "http://example.com/cat.jpg",
			},
			wantErr:  false,
			contains: []string{"猫が屋根の上にいる", "貓", "http://example.com/cat.jpg"},
		},
		{
			name:     "Grammar card front",
			cardType: "grammar",
			data: map[string]interface{}{
				"情境課題": "即使下雨，也要去。",
				"文法點":  "〜ても",
			},
			wantErr:  false,
			contains: []string{"即使下雨，也要去。", "〜ても"},
		},
		{
			name:     "Invalid card type",
			cardType: "invalid",
			data:     map[string]interface{}{},
			wantErr:  true,
			contains: nil,
		},
	}
	
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Render the card front
			result, err := manager.RenderCardFront(tc.cardType, tc.data)
			
			// Check if the error matches the expectation
			if (err != nil) != tc.wantErr {
				t.Errorf("RenderCardFront() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
			
			// If no error is expected, check the result
			if !tc.wantErr {
				// Check if the result contains the expected strings
				for _, s := range tc.contains {
					if !strings.Contains(result, s) {
						t.Errorf("RenderCardFront() result does not contain '%s'", s)
					}
				}
			}
		})
	}
}

func TestTemplateManager_RenderCardBack(t *testing.T) {
	// Create a new template manager
	manager, err := NewTemplateManager()
	if err != nil {
		t.Fatalf("Failed to create template manager: %v", err)
	}
	
	// Test data for different card types
	testCases := []struct {
		name     string
		cardType string
		data     map[string]interface{}
		wantErr  bool
		contains []string
	}{
		{
			name:     "Verb card back",
			cardType: "verb",
			data: map[string]interface{}{
				"情境例句": "水を飲む",
				"核心單字": "飲む",
				"發音":   "のむ",
				"重音":   "1",
				"詞性分類": "五段動詞",
				"例句翻譯": "喝水",
				"常用變化": "飲みます、飲んで",
				"圖片提示": "http://example.com/image.jpg",
			},
			wantErr: false,
			contains: []string{
				"水を飲む", "飲む", "のむ", "1", "五段動詞", "喝水", "飲みます、飲んで", "http://example.com/image.jpg",
			},
		},
		{
			name:     "Adjective card back",
			cardType: "adjective",
			data: map[string]interface{}{
				"情境例句": "美しい花",
				"核心單字": "美しい",
				"發音":   "うつくしい",
				"重音":   "4",
				"詞性分類": "い形容詞",
				"例句翻譯": "美麗的花",
				"主要變化": "美しくない、美しかった",
				"相關詞彙": "綺麗、素敵",
			},
			wantErr: false,
			contains: []string{
				"美しい花", "美しい", "うつくしい", "4", "い形容詞", "美麗的花", "美しくない、美しかった", "綺麗、素敵",
			},
		},
		{
			name:     "Normal card back",
			cardType: "normal",
			data: map[string]interface{}{
				"情境例句": "猫が屋根の上にいる",
				"核心單字": "猫",
				"發音":   "ねこ",
				"重音":   "2",
				"詞性分類": "名詞",
				"例句翻譯": "貓在屋頂上",
				"同義詞":  "ニャンコ",
				"反義詞":  "犬",
				"圖片提示": "http://example.com/cat.jpg",
			},
			wantErr: false,
			contains: []string{
				"猫が屋根の上にいる", "猫", "ねこ", "2", "名詞", "貓在屋頂上", "ニャンコ", "犬", "http://example.com/cat.jpg",
			},
		},
		{
			name:     "Grammar card back",
			cardType: "grammar",
			data: map[string]interface{}{
				"情境課題": "即使下雨，也要去。",
				"解答範例": "雨が降っても、行きます。",
				"文法要點": "〜ても",
				"意義說明": "即使...也...",
				"結構形式": "動詞て形 + も",
				"使用時機": "表示讓步",
				"相關文法": "〜ながら、〜のに",
			},
			wantErr: false,
			contains: []string{
				"即使下雨，也要去。", "雨が降っても、行きます。", "〜ても", "即使...也...", "動詞て形 + も", "表示讓步", "〜ながら、〜のに",
			},
		},
		{
			name:     "Invalid card type",
			cardType: "invalid",
			data:     map[string]interface{}{},
			wantErr:  true,
			contains: nil,
		},
	}
	
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Render the card back
			result, err := manager.RenderCardBack(tc.cardType, tc.data)
			
			// Check if the error matches the expectation
			if (err != nil) != tc.wantErr {
				t.Errorf("RenderCardBack() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
			
			// If no error is expected, check the result
			if !tc.wantErr {
				// Check if the result contains the expected strings
				for _, s := range tc.contains {
					if !strings.Contains(result, s) {
						t.Errorf("RenderCardBack() result does not contain '%s'", s)
					}
				}
			}
		})
	}
}

func TestTemplateManager_ValidateTemplate(t *testing.T) {
	// Create a new template manager
	manager, err := NewTemplateManager()
	if err != nil {
		t.Fatalf("Failed to create template manager: %v", err)
	}
	
	// Test cases
	testCases := []struct {
		name     string
		cardType string
		wantErr  bool
	}{
		{
			name:     "Valid verb template",
			cardType: "verb",
			wantErr:  false,
		},
		{
			name:     "Valid adjective template",
			cardType: "adjective",
			wantErr:  false,
		},
		{
			name:     "Valid normal template",
			cardType: "normal",
			wantErr:  false,
		},
		{
			name:     "Valid grammar template",
			cardType: "grammar",
			wantErr:  false,
		},
		{
			name:     "Invalid template",
			cardType: "invalid",
			wantErr:  true,
		},
	}
	
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Validate the template
			err := manager.ValidateTemplate(tc.cardType)
			
			// Check if the error matches the expectation
			if (err != nil) != tc.wantErr {
				t.Errorf("ValidateTemplate() error = %v, wantErr %v", err, tc.wantErr)
			}
		})
	}
}

func TestTemplateManager_GetAvailableTemplates(t *testing.T) {
	// Create a new template manager
	manager, err := NewTemplateManager()
	if err != nil {
		t.Fatalf("Failed to create template manager: %v", err)
	}
	
	// Get available templates
	templates := manager.GetAvailableTemplates()
	
	// Check if the result is not nil
	if templates == nil {
		t.Fatal("GetAvailableTemplates() returned nil")
	}
	
	// Check if all expected templates are included
	expectedTemplates := []string{"verb", "adjective", "normal", "grammar"}
	for _, expected := range expectedTemplates {
		found := false
		for _, actual := range templates {
			if actual == expected {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("GetAvailableTemplates() did not include '%s'", expected)
		}
	}
}

func TestTemplateManager_GetRawTemplate(t *testing.T) {
	// Create a new template manager
	manager, err := NewTemplateManager()
	if err != nil {
		t.Fatalf("Failed to create template manager: %v", err)
	}
	
	// Test cases
	testCases := []struct {
		name     string
		cardType string
		side     string
		wantErr  bool
	}{
		{
			name:     "Verb front template",
			cardType: "verb",
			side:     "front",
			wantErr:  false,
		},
		{
			name:     "Verb back template",
			cardType: "verb",
			side:     "back",
			wantErr:  false,
		},
		{
			name:     "Adjective front template",
			cardType: "adjective",
			side:     "front",
			wantErr:  false,
		},
		{
			name:     "Adjective back template",
			cardType: "adjective",
			side:     "back",
			wantErr:  false,
		},
		{
			name:     "Normal front template",
			cardType: "normal",
			side:     "front",
			wantErr:  false,
		},
		{
			name:     "Normal back template",
			cardType: "normal",
			side:     "back",
			wantErr:  false,
		},
		{
			name:     "Grammar front template",
			cardType: "grammar",
			side:     "front",
			wantErr:  false,
		},
		{
			name:     "Grammar back template",
			cardType: "grammar",
			side:     "back",
			wantErr:  false,
		},
		{
			name:     "Invalid card type",
			cardType: "invalid",
			side:     "front",
			wantErr:  true,
		},
		{
			name:     "Invalid side",
			cardType: "verb",
			side:     "invalid",
			wantErr:  true,
		},
	}
	
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Get the raw template
			content, err := manager.GetRawTemplate(tc.cardType, tc.side)
			
			// Check if the error matches the expectation
			if (err != nil) != tc.wantErr {
				t.Errorf("GetRawTemplate() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
			
			// If no error is expected, check the result
			if !tc.wantErr {
				// Check if the content is not empty
				if content == "" {
					t.Error("GetRawTemplate() returned empty content")
				}
			}
		})
	}
}