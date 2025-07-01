package models

import (
	"encoding/json"
	"testing"
)

func TestGrammarCard_GetCardType(t *testing.T) {
	card := &GrammarCard{}
	if card.GetCardType() != "grammar" {
		t.Errorf("Expected card type to be 'grammar', got '%s'", card.GetCardType())
	}
}

func TestGrammarCard_Validate(t *testing.T) {
	tests := []struct {
		name     string
		card     GrammarCard
		wantErr  bool
		errField string
	}{
		{
			name: "Valid card",
			card: GrammarCard{
				GrammarPoint:   "〜ても",
				Structure:      "動詞て形 + も",
				Meaning:        "即使...也...",
				Usage:          "表示讓步",
				Examples:       "雨が降っても、行きます。",
				Translation:    "即使下雨，也要去。",
				Challenge:      "即使很忙，也要學習日文。",
				Answer:         "忙しくても、日本語を勉強します。",
				Level:          "N4",
				RelatedGrammar: "〜ながら、〜のに",
				CommonMistakes: "混淆「〜ても」和「〜でも」",
				Tips:           "記住「即使...也...」的讓步關係",
			},
			wantErr: false,
		},
		{
			name: "Missing GrammarPoint",
			card: GrammarCard{
				Structure:   "動詞て形 + も",
				Meaning:     "即使...也...",
				Examples:    "雨が降っても、行きます。",
				Translation: "即使下雨，也要去。",
			},
			wantErr:  true,
			errField: "文法要點",
		},
		{
			name: "Missing Structure",
			card: GrammarCard{
				GrammarPoint: "〜ても",
				Meaning:      "即使...也...",
				Examples:     "雨が降っても、行きます。",
				Translation:  "即使下雨，也要去。",
			},
			wantErr:  true,
			errField: "結構形式",
		},
		{
			name: "Missing Meaning",
			card: GrammarCard{
				GrammarPoint: "〜ても",
				Structure:    "動詞て形 + も",
				Examples:     "雨が降っても、行きます。",
				Translation:  "即使下雨，也要去。",
			},
			wantErr:  true,
			errField: "意義說明",
		},
		{
			name: "Missing Examples",
			card: GrammarCard{
				GrammarPoint: "〜ても",
				Structure:    "動詞て形 + も",
				Meaning:      "即使...也...",
				Translation:  "即使下雨，也要去。",
			},
			wantErr:  true,
			errField: "例句示範",
		},
		{
			name: "Missing Translation",
			card: GrammarCard{
				GrammarPoint: "〜ても",
				Structure:    "動詞て形 + も",
				Meaning:      "即使...也...",
				Examples:     "雨が降っても、行きます。",
			},
			wantErr:  true,
			errField: "例句翻譯",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.card.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("GrammarCard.Validate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil {
				validationErr, ok := err.(*ValidationError)
				if !ok {
					t.Errorf("Expected ValidationError, got %T", err)
					return
				}
				if validationErr.Field != tt.errField {
					t.Errorf("Expected error field to be '%s', got '%s'", tt.errField, validationErr.Field)
				}
			}
		})
	}
}

func TestGrammarCard_ToMap(t *testing.T) {
	card := GrammarCard{
		GrammarPoint:   "〜ても",
		Structure:      "動詞て形 + も",
		Meaning:        "即使...也...",
		Usage:          "表示讓步",
		Examples:       "雨が降っても、行きます。",
		Translation:    "即使下雨，也要去。",
		Challenge:      "即使很忙，也要學習日文。",
		Answer:         "忙しくても、日本語を勉強します。",
		Level:          "N4",
		RelatedGrammar: "〜ながら、〜のに",
		CommonMistakes: "混淆「〜ても」和「〜でも」",
		Tips:           "記住「即使...也...」的讓步關係",
	}

	expected := map[string]interface{}{
		"文法要點": "〜ても",
		"結構形式": "動詞て形 + も",
		"意義說明": "即使...也...",
		"使用時機": "表示讓步",
		"例句示範": "雨が降っても、行きます。",
		"例句翻譯": "即使下雨，也要去。",
		"情境課題": "即使很忙，也要學習日文。",
		"解答範例": "忙しくても、日本語を勉強します。",
		"難度等級": "N4",
		"相關文法": "〜ながら、〜のに",
		"常見錯誤": "混淆「〜ても」和「〜でも」",
		"記憶技巧": "記住「即使...也...」的讓步關係",
	}

	result := card.ToMap()
	for key, expectedValue := range expected {
		if result[key] != expectedValue {
			t.Errorf("ToMap() for key '%s', expected '%v', got '%v'", key, expectedValue, result[key])
		}
	}
}

func TestGrammarCard_FromMap(t *testing.T) {
	data := map[string]interface{}{
		"文法要點": "〜ても",
		"結構形式": "動詞て形 + も",
		"意義說明": "即使...也...",
		"使用時機": "表示讓步",
		"例句示範": "雨が降っても、行きます。",
		"例句翻譯": "即使下雨，也要去。",
		"情境課題": "即使很忙，也要學習日文。",
		"解答範例": "忙しくても、日本語を勉強します。",
		"難度等級": "N4",
		"相關文法": "〜ながら、〜のに",
		"常見錯誤": "混淆「〜ても」和「〜でも」",
		"記憶技巧": "記住「即使...也...」的讓步關係",
	}

	var card GrammarCard
	err := card.FromMap(data)
	if err != nil {
		t.Errorf("FromMap() error = %v", err)
		return
	}

	expected := GrammarCard{
		GrammarPoint:   "〜ても",
		Structure:      "動詞て形 + も",
		Meaning:        "即使...也...",
		Usage:          "表示讓步",
		Examples:       "雨が降っても、行きます。",
		Translation:    "即使下雨，也要去。",
		Challenge:      "即使很忙，也要學習日文。",
		Answer:         "忙しくても、日本語を勉強します。",
		Level:          "N4",
		RelatedGrammar: "〜ながら、〜のに",
		CommonMistakes: "混淆「〜ても」和「〜でも」",
		Tips:           "記住「即使...也...」的讓步關係",
	}

	if card.GrammarPoint != expected.GrammarPoint ||
		card.Structure != expected.Structure ||
		card.Meaning != expected.Meaning ||
		card.Usage != expected.Usage ||
		card.Examples != expected.Examples ||
		card.Translation != expected.Translation ||
		card.Challenge != expected.Challenge ||
		card.Answer != expected.Answer ||
		card.Level != expected.Level ||
		card.RelatedGrammar != expected.RelatedGrammar ||
		card.CommonMistakes != expected.CommonMistakes ||
		card.Tips != expected.Tips {
		t.Errorf("FromMap() = %v, want %v", card, expected)
	}
}

func TestGrammarCard_JSONSerialization(t *testing.T) {
	original := GrammarCard{
		GrammarPoint:   "〜ても",
		Structure:      "動詞て形 + も",
		Meaning:        "即使...也...",
		Usage:          "表示讓步",
		Examples:       "雨が降っても、行きます。",
		Translation:    "即使下雨，也要去。",
		Challenge:      "即使很忙，也要學習日文。",
		Answer:         "忙しくても、日本語を勉強します。",
		Level:          "N4",
		RelatedGrammar: "〜ながら、〜のに",
		CommonMistakes: "混淆「〜ても」和「〜でも」",
		Tips:           "記住「即使...也...」的讓步關係",
	}

	// Serialize to JSON
	jsonData, err := json.Marshal(original)
	if err != nil {
		t.Errorf("json.Marshal() error = %v", err)
		return
	}

	// Deserialize from JSON
	var deserialized GrammarCard
	err = json.Unmarshal(jsonData, &deserialized)
	if err != nil {
		t.Errorf("json.Unmarshal() error = %v", err)
		return
	}

	// Compare original and deserialized
	if original.GrammarPoint != deserialized.GrammarPoint ||
		original.Structure != deserialized.Structure ||
		original.Meaning != deserialized.Meaning ||
		original.Usage != deserialized.Usage ||
		original.Examples != deserialized.Examples ||
		original.Translation != deserialized.Translation ||
		original.Challenge != deserialized.Challenge ||
		original.Answer != deserialized.Answer ||
		original.Level != deserialized.Level ||
		original.RelatedGrammar != deserialized.RelatedGrammar ||
		original.CommonMistakes != deserialized.CommonMistakes ||
		original.Tips != deserialized.Tips {
		t.Errorf("JSON serialization/deserialization failed. Original: %v, Deserialized: %v", original, deserialized)
	}
}
