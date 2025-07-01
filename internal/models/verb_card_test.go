package models

import (
	"encoding/json"
	"testing"
)

func TestVerbCard_GetCardType(t *testing.T) {
	card := &VerbCard{}
	if card.GetCardType() != "verb" {
		t.Errorf("Expected card type to be 'verb', got '%s'", card.GetCardType())
	}
}

func TestVerbCard_Validate(t *testing.T) {
	tests := []struct {
		name     string
		card     VerbCard
		wantErr  bool
		errField string
	}{
		{
			name: "Valid card",
			card: VerbCard{
				CoreWord:        "飲む",
				WordType:        "五段動詞",
				CoreMeaning:     "喝",
				Pronunciation:   "のむ",
				Accent:          "1",
				Conjugations:    "飲みます、飲んで",
				ContextSentence: "水を飲む",
				Translation:     "喝水",
			},
			wantErr: false,
		},
		{
			name: "Missing CoreWord",
			card: VerbCard{
				CoreMeaning:     "喝",
				Pronunciation:   "のむ",
				ContextSentence: "水を飲む",
				Translation:     "喝水",
			},
			wantErr:  true,
			errField: "核心單字",
		},
		{
			name: "Missing CoreMeaning",
			card: VerbCard{
				CoreWord:        "飲む",
				Pronunciation:   "のむ",
				ContextSentence: "水を飲む",
				Translation:     "喝水",
			},
			wantErr:  true,
			errField: "核心意義",
		},
		{
			name: "Missing Pronunciation",
			card: VerbCard{
				CoreWord:        "飲む",
				CoreMeaning:     "喝",
				ContextSentence: "水を飲む",
				Translation:     "喝水",
			},
			wantErr:  true,
			errField: "發音",
		},
		{
			name: "Missing ContextSentence",
			card: VerbCard{
				CoreWord:      "飲む",
				CoreMeaning:   "喝",
				Pronunciation: "のむ",
				Translation:   "喝水",
			},
			wantErr:  true,
			errField: "情境例句",
		},
		{
			name: "Missing Translation",
			card: VerbCard{
				CoreWord:        "飲む",
				CoreMeaning:     "喝",
				Pronunciation:   "のむ",
				ContextSentence: "水を飲む",
			},
			wantErr:  true,
			errField: "例句翻譯",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.card.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("VerbCard.Validate() error = %v, wantErr %v", err, tt.wantErr)
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

func TestVerbCard_ToMap(t *testing.T) {
	card := VerbCard{
		CoreWord:        "飲む",
		WordType:        "五段動詞",
		CoreMeaning:     "喝",
		Pronunciation:   "のむ",
		Accent:          "1",
		Conjugations:    "飲みます、飲んで",
		ContextSentence: "水を飲む",
		Translation:     "喝水",
		ImageHint:       "http://example.com/image.jpg",
	}

	expected := map[string]interface{}{
		"核心單字": "飲む",
		"詞性分類": "五段動詞",
		"核心意義": "喝",
		"發音":   "のむ",
		"重音":   "1",
		"常用變化": "飲みます、飲んで",
		"情境例句": "水を飲む",
		"例句翻譯": "喝水",
		"圖片提示": "http://example.com/image.jpg",
	}

	result := card.ToMap()
	for key, expectedValue := range expected {
		if result[key] != expectedValue {
			t.Errorf("ToMap() for key '%s', expected '%v', got '%v'", key, expectedValue, result[key])
		}
	}
}

func TestVerbCard_FromMap(t *testing.T) {
	data := map[string]interface{}{
		"核心單字": "飲む",
		"詞性分類": "五段動詞",
		"核心意義": "喝",
		"發音":   "のむ",
		"重音":   "1",
		"常用變化": "飲みます、飲んで",
		"情境例句": "水を飲む",
		"例句翻譯": "喝水",
		"圖片提示": "http://example.com/image.jpg",
	}

	var card VerbCard
	err := card.FromMap(data)
	if err != nil {
		t.Errorf("FromMap() error = %v", err)
		return
	}

	expected := VerbCard{
		CoreWord:        "飲む",
		WordType:        "五段動詞",
		CoreMeaning:     "喝",
		Pronunciation:   "のむ",
		Accent:          "1",
		Conjugations:    "飲みます、飲んで",
		ContextSentence: "水を飲む",
		Translation:     "喝水",
		ImageHint:       "http://example.com/image.jpg",
	}

	if card.CoreWord != expected.CoreWord ||
		card.WordType != expected.WordType ||
		card.CoreMeaning != expected.CoreMeaning ||
		card.Pronunciation != expected.Pronunciation ||
		card.Accent != expected.Accent ||
		card.Conjugations != expected.Conjugations ||
		card.ContextSentence != expected.ContextSentence ||
		card.Translation != expected.Translation ||
		card.ImageHint != expected.ImageHint {
		t.Errorf("FromMap() = %v, want %v", card, expected)
	}
}

func TestVerbCard_JSONSerialization(t *testing.T) {
	original := VerbCard{
		CoreWord:        "飲む",
		WordType:        "五段動詞",
		CoreMeaning:     "喝",
		Pronunciation:   "のむ",
		Accent:          "1",
		Conjugations:    "飲みます、飲んで",
		ContextSentence: "水を飲む",
		Translation:     "喝水",
		ImageHint:       "http://example.com/image.jpg",
	}

	// Serialize to JSON
	jsonData, err := json.Marshal(original)
	if err != nil {
		t.Errorf("json.Marshal() error = %v", err)
		return
	}

	// Deserialize from JSON
	var deserialized VerbCard
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
		original.Conjugations != deserialized.Conjugations ||
		original.ContextSentence != deserialized.ContextSentence ||
		original.Translation != deserialized.Translation ||
		original.ImageHint != deserialized.ImageHint {
		t.Errorf("JSON serialization/deserialization failed. Original: %v, Deserialized: %v", original, deserialized)
	}
}
