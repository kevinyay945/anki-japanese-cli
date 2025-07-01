package models

import (
	"encoding/json"
	"testing"
)

func TestAdjectiveCard_GetCardType(t *testing.T) {
	card := &AdjectiveCard{}
	if card.GetCardType() != "adjective" {
		t.Errorf("Expected card type to be 'adjective', got '%s'", card.GetCardType())
	}
}

func TestAdjectiveCard_Validate(t *testing.T) {
	tests := []struct {
		name     string
		card     AdjectiveCard
		wantErr  bool
		errField string
	}{
		{
			name: "Valid card",
			card: AdjectiveCard{
				CoreWord:        "美しい",
				WordType:        "い形容詞",
				CoreMeaning:     "美麗的",
				Pronunciation:   "うつくしい",
				Accent:          "4",
				MainChanges:     "美しくない、美しかった",
				ContextSentence: "美しい花",
				Translation:     "美麗的花",
				RelatedWords:    "綺麗、素敵",
			},
			wantErr: false,
		},
		{
			name: "Missing CoreWord",
			card: AdjectiveCard{
				CoreMeaning:     "美麗的",
				Pronunciation:   "うつくしい",
				ContextSentence: "美しい花",
				Translation:     "美麗的花",
			},
			wantErr:  true,
			errField: "核心單字",
		},
		{
			name: "Missing CoreMeaning",
			card: AdjectiveCard{
				CoreWord:        "美しい",
				Pronunciation:   "うつくしい",
				ContextSentence: "美しい花",
				Translation:     "美麗的花",
			},
			wantErr:  true,
			errField: "核心意義",
		},
		{
			name: "Missing Pronunciation",
			card: AdjectiveCard{
				CoreWord:        "美しい",
				CoreMeaning:     "美麗的",
				ContextSentence: "美しい花",
				Translation:     "美麗的花",
			},
			wantErr:  true,
			errField: "發音",
		},
		{
			name: "Missing ContextSentence",
			card: AdjectiveCard{
				CoreWord:      "美しい",
				CoreMeaning:   "美麗的",
				Pronunciation: "うつくしい",
				Translation:   "美麗的花",
			},
			wantErr:  true,
			errField: "情境例句",
		},
		{
			name: "Missing Translation",
			card: AdjectiveCard{
				CoreWord:        "美しい",
				CoreMeaning:     "美麗的",
				Pronunciation:   "うつくしい",
				ContextSentence: "美しい花",
			},
			wantErr:  true,
			errField: "例句翻譯",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.card.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("AdjectiveCard.Validate() error = %v, wantErr %v", err, tt.wantErr)
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

func TestAdjectiveCard_ToMap(t *testing.T) {
	card := AdjectiveCard{
		CoreWord:        "美しい",
		WordType:        "い形容詞",
		CoreMeaning:     "美麗的",
		Pronunciation:   "うつくしい",
		Accent:          "4",
		MainChanges:     "美しくない、美しかった",
		ContextSentence: "美しい花",
		Translation:     "美麗的花",
		RelatedWords:    "綺麗、素敵",
	}

	expected := map[string]interface{}{
		"核心單字": "美しい",
		"詞性分類": "い形容詞",
		"核心意義": "美麗的",
		"發音":   "うつくしい",
		"重音":   "4",
		"主要變化": "美しくない、美しかった",
		"情境例句": "美しい花",
		"例句翻譯": "美麗的花",
		"相關詞彙": "綺麗、素敵",
	}

	result := card.ToMap()
	for key, expectedValue := range expected {
		if result[key] != expectedValue {
			t.Errorf("ToMap() for key '%s', expected '%v', got '%v'", key, expectedValue, result[key])
		}
	}
}

func TestAdjectiveCard_FromMap(t *testing.T) {
	data := map[string]interface{}{
		"核心單字": "美しい",
		"詞性分類": "い形容詞",
		"核心意義": "美麗的",
		"發音":   "うつくしい",
		"重音":   "4",
		"主要變化": "美しくない、美しかった",
		"情境例句": "美しい花",
		"例句翻譯": "美麗的花",
		"相關詞彙": "綺麗、素敵",
	}

	var card AdjectiveCard
	err := card.FromMap(data)
	if err != nil {
		t.Errorf("FromMap() error = %v", err)
		return
	}

	expected := AdjectiveCard{
		CoreWord:        "美しい",
		WordType:        "い形容詞",
		CoreMeaning:     "美麗的",
		Pronunciation:   "うつくしい",
		Accent:          "4",
		MainChanges:     "美しくない、美しかった",
		ContextSentence: "美しい花",
		Translation:     "美麗的花",
		RelatedWords:    "綺麗、素敵",
	}

	if card.CoreWord != expected.CoreWord ||
		card.WordType != expected.WordType ||
		card.CoreMeaning != expected.CoreMeaning ||
		card.Pronunciation != expected.Pronunciation ||
		card.Accent != expected.Accent ||
		card.MainChanges != expected.MainChanges ||
		card.ContextSentence != expected.ContextSentence ||
		card.Translation != expected.Translation ||
		card.RelatedWords != expected.RelatedWords {
		t.Errorf("FromMap() = %v, want %v", card, expected)
	}
}

func TestAdjectiveCard_JSONSerialization(t *testing.T) {
	original := AdjectiveCard{
		CoreWord:        "美しい",
		WordType:        "い形容詞",
		CoreMeaning:     "美麗的",
		Pronunciation:   "うつくしい",
		Accent:          "4",
		MainChanges:     "美しくない、美しかった",
		ContextSentence: "美しい花",
		Translation:     "美麗的花",
		RelatedWords:    "綺麗、素敵",
	}

	// Serialize to JSON
	jsonData, err := json.Marshal(original)
	if err != nil {
		t.Errorf("json.Marshal() error = %v", err)
		return
	}

	// Deserialize from JSON
	var deserialized AdjectiveCard
	err = json.Unmarshal(jsonData, &deserialized)
	if err != nil {
		t.Errorf("json.Unmarshal() error = %v", err)
		return
	}

	// Compare original and deserialized
	if original.CoreWord != deserialized.CoreWord ||
		original.WordType != deserialized.WordType ||
		original.CoreMeaning != deserialized.CoreMeaning ||
		original.Pronunciation != deserialized.Pronunciation ||
		original.Accent != deserialized.Accent ||
		original.MainChanges != deserialized.MainChanges ||
		original.ContextSentence != deserialized.ContextSentence ||
		original.Translation != deserialized.Translation ||
		original.RelatedWords != deserialized.RelatedWords {
		t.Errorf("JSON serialization/deserialization failed. Original: %v, Deserialized: %v", original, deserialized)
	}
}
