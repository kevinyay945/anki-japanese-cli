package models

import (
	"testing"
)

func TestCardFactory_NewCardFactory(t *testing.T) {
	factory := NewCardFactory()
	if factory == nil {
		t.Error("NewCardFactory() returned nil")
	}
}

func TestCardFactory_GetSupportedCardTypes(t *testing.T) {
	factory := NewCardFactory()
	types := factory.GetSupportedCardTypes()

	expected := []string{"verb", "adjective", "normal", "grammar"}
	if len(types) != len(expected) {
		t.Errorf("GetSupportedCardTypes() returned %d types, expected %d", len(types), len(expected))
	}

	// Check that all expected types are present
	for _, expectedType := range expected {
		found := false
		for _, actualType := range types {
			if actualType == expectedType {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("GetSupportedCardTypes() did not include '%s'", expectedType)
		}
	}
}

func TestCardFactory_ValidateCardType(t *testing.T) {
	factory := NewCardFactory()

	// Test valid card types
	validTypes := []string{"verb", "adjective", "normal", "grammar"}
	for _, cardType := range validTypes {
		err := factory.ValidateCardType(cardType)
		if err != nil {
			t.Errorf("ValidateCardType(%s) returned error: %v", cardType, err)
		}
	}

	// Test invalid card type
	err := factory.ValidateCardType("invalid")
	if err == nil {
		t.Error("ValidateCardType(invalid) did not return error")
	}
}

func TestCardFactory_CreateCard(t *testing.T) {
	factory := NewCardFactory()

	// Test creating a verb card
	verbData := map[string]interface{}{
		"核心單字": "飲む",
		"詞性分類": "五段動詞",
		"核心意義": "喝",
		"發音":   "のむ",
		"重音":   "1",
		"常用變化": "飲みます、飲んで",
		"情境例句": "水を飲む",
		"例句翻譯": "喝水",
	}

	verbCard, err := factory.CreateCard("verb", verbData)
	if err != nil {
		t.Errorf("CreateCard(verb) returned error: %v", err)
	}
	if verbCard == nil {
		t.Error("CreateCard(verb) returned nil")
	} else if verbCard.GetCardType() != "verb" {
		t.Errorf("CreateCard(verb) returned card of type '%s'", verbCard.GetCardType())
	}

	// Test creating an adjective card
	adjectiveData := map[string]interface{}{
		"核心單字": "美しい",
		"詞性分類": "い形容詞",
		"核心意義": "美麗的",
		"發音":   "うつくしい",
		"重音":   "4",
		"主要變化": "美しくない、美しかった",
		"情境例句": "美しい花",
		"例句翻譯": "美麗的花",
	}

	adjectiveCard, err := factory.CreateCard("adjective", adjectiveData)
	if err != nil {
		t.Errorf("CreateCard(adjective) returned error: %v", err)
	}
	if adjectiveCard == nil {
		t.Error("CreateCard(adjective) returned nil")
	} else if adjectiveCard.GetCardType() != "adjective" {
		t.Errorf("CreateCard(adjective) returned card of type '%s'", adjectiveCard.GetCardType())
	}

	// Test creating a normal card
	normalData := map[string]interface{}{
		"核心單字": "猫",
		"詞性分類": "名詞",
		"核心意義": "貓",
		"發音":   "ねこ",
		"重音":   "2",
		"使用方式": "動物を表す名詞",
		"情境例句": "猫が屋根の上にいる",
		"例句翻譯": "貓在屋頂上",
	}

	normalCard, err := factory.CreateCard("normal", normalData)
	if err != nil {
		t.Errorf("CreateCard(normal) returned error: %v", err)
	}
	if normalCard == nil {
		t.Error("CreateCard(normal) returned nil")
	} else if normalCard.GetCardType() != "normal" {
		t.Errorf("CreateCard(normal) returned card of type '%s'", normalCard.GetCardType())
	}

	// Test creating a grammar card
	grammarData := map[string]interface{}{
		"文法要點": "〜ても",
		"結構形式": "動詞て形 + も",
		"意義說明": "即使...也...",
		"例句示範": "雨が降っても、行きます。",
		"例句翻譯": "即使下雨，也要去。",
	}

	grammarCard, err := factory.CreateCard("grammar", grammarData)
	if err != nil {
		t.Errorf("CreateCard(grammar) returned error: %v", err)
	}
	if grammarCard == nil {
		t.Error("CreateCard(grammar) returned nil")
	} else if grammarCard.GetCardType() != "grammar" {
		t.Errorf("CreateCard(grammar) returned card of type '%s'", grammarCard.GetCardType())
	}

	// Test creating a card with invalid type
	_, err = factory.CreateCard("invalid", map[string]interface{}{})
	if err == nil {
		t.Error("CreateCard(invalid) did not return error")
	}

	// Test creating a card with invalid data
	_, err = factory.CreateCard("verb", map[string]interface{}{})
	if err == nil {
		t.Error("CreateCard(verb, {}) did not return error")
	}
}

func TestCardFactory_CreateCardFromJSON(t *testing.T) {
	factory := NewCardFactory()

	// Test creating a verb card from JSON
	verbJSON := `{
		"核心單字": "飲む",
		"詞性分類": "五段動詞",
		"核心意義": "喝",
		"發音": "のむ",
		"重音": "1",
		"常用變化": "飲みます、飲んで",
		"情境例句": "水を飲む",
		"例句翻譯": "喝水"
	}`

	verbCard, err := factory.CreateCardFromJSON("verb", []byte(verbJSON))
	if err != nil {
		t.Errorf("CreateCardFromJSON(verb) returned error: %v", err)
	}
	if verbCard == nil {
		t.Error("CreateCardFromJSON(verb) returned nil")
	} else if verbCard.GetCardType() != "verb" {
		t.Errorf("CreateCardFromJSON(verb) returned card of type '%s'", verbCard.GetCardType())
	}

	// Test creating a card from invalid JSON
	_, err = factory.CreateCardFromJSON("verb", []byte("{invalid json}"))
	if err == nil {
		t.Error("CreateCardFromJSON(verb, invalid) did not return error")
	}

	// Test creating a card with invalid type
	_, err = factory.CreateCardFromJSON("invalid", []byte("{}"))
	if err == nil {
		t.Error("CreateCardFromJSON(invalid) did not return error")
	}
}

func TestCardFactory_CreateCardWithInvalidData(t *testing.T) {
	factory := NewCardFactory()

	// Test cases for invalid data
	testCases := []struct {
		name     string
		cardType string
		data     map[string]interface{}
	}{
		{
			name:     "Verb missing required fields",
			cardType: "verb",
			data: map[string]interface{}{
				"核心單字": "飲む",
				// Missing other required fields
			},
		},
		{
			name:     "Adjective missing required fields",
			cardType: "adjective",
			data: map[string]interface{}{
				"核心單字": "美しい",
				// Missing other required fields
			},
		},
		{
			name:     "Normal missing required fields",
			cardType: "normal",
			data: map[string]interface{}{
				"核心單字": "猫",
				// Missing other required fields
			},
		},
		{
			name:     "Grammar missing required fields",
			cardType: "grammar",
			data: map[string]interface{}{
				"文法要點": "〜ても",
				// Missing other required fields
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := factory.CreateCard(tc.cardType, tc.data)
			if err == nil {
				t.Errorf("CreateCard(%s) with invalid data did not return error", tc.cardType)
			}
		})
	}
}

func TestCardFactory_CreateCardWithValidData(t *testing.T) {
	factory := NewCardFactory()

	// Test cases for valid data
	testCases := []struct {
		name     string
		cardType string
		data     map[string]interface{}
	}{
		{
			name:     "Valid verb card",
			cardType: "verb",
			data: map[string]interface{}{
				"核心單字": "飲む",
				"詞性分類": "五段動詞",
				"核心意義": "喝",
				"發音":   "のむ",
				"重音":   "1",
				"常用變化": "飲みます、飲んで",
				"情境例句": "水を飲む",
				"例句翻譯": "喝水",
			},
		},
		{
			name:     "Valid adjective card",
			cardType: "adjective",
			data: map[string]interface{}{
				"核心單字": "美しい",
				"詞性分類": "い形容詞",
				"核心意義": "美麗的",
				"發音":   "うつくしい",
				"重音":   "4",
				"主要變化": "美しくない、美しかった",
				"情境例句": "美しい花",
				"例句翻譯": "美麗的花",
			},
		},
		{
			name:     "Valid normal card",
			cardType: "normal",
			data: map[string]interface{}{
				"核心單字": "猫",
				"詞性分類": "名詞",
				"核心意義": "貓",
				"發音":   "ねこ",
				"重音":   "2",
				"使用方式": "動物を表す名詞",
				"情境例句": "猫が屋根の上にいる",
				"例句翻譯": "貓在屋頂上",
			},
		},
		{
			name:     "Valid grammar card",
			cardType: "grammar",
			data: map[string]interface{}{
				"文法要點": "〜ても",
				"結構形式": "動詞て形 + も",
				"意義說明": "即使...也...",
				"例句示範": "雨が降っても、行きます。",
				"例句翻譯": "即使下雨，也要去。",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			card, err := factory.CreateCard(tc.cardType, tc.data)
			if err != nil {
				t.Errorf("CreateCard(%s) with valid data returned error: %v", tc.cardType, err)
			}
			if card == nil {
				t.Errorf("CreateCard(%s) with valid data returned nil", tc.cardType)
			} else if card.GetCardType() != tc.cardType {
				t.Errorf("CreateCard(%s) returned card of type '%s'", tc.cardType, card.GetCardType())
			}
		})
	}
}
