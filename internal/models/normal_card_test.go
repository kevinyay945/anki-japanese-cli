package models

import (
	"encoding/json"
	"testing"
)

func TestNormalWordCard_GetCardType(t *testing.T) {
	card := &NormalWordCard{}
	if card.GetCardType() != "normal" {
		t.Errorf("Expected card type to be 'normal', got '%s'", card.GetCardType())
	}
}

func TestNormalWordCard_Validate(t *testing.T) {
	tests := []struct {
		name     string
		card     NormalWordCard
		wantErr  bool
		errField string
	}{
		{
			name: "Valid card",
			card: NormalWordCard{
				CoreWord:        "猫",
				WordType:        "名詞",
				CoreMeaning:     "貓",
				Pronunciation:   "ねこ",
				Accent:          "2",
				Usage:           "動物を表す名詞",
				ContextSentence: "猫が屋根の上にいる",
				Translation:     "貓在屋頂上",
				Synonyms:        "ニャンコ",
				Antonyms:        "犬",
				ImageHint:       "http://example.com/cat.jpg",
			},
			wantErr: false,
		},
		{
			name: "Missing CoreWord",
			card: NormalWordCard{
				CoreMeaning:     "貓",
				Pronunciation:   "ねこ",
				ContextSentence: "猫が屋根の上にいる",
				Translation:     "貓在屋頂上",
			},
			wantErr:  true,
			errField: "核心單字",
		},
		{
			name: "Missing CoreMeaning",
			card: NormalWordCard{
				CoreWord:        "猫",
				Pronunciation:   "ねこ",
				ContextSentence: "猫が屋根の上にいる",
				Translation:     "貓在屋頂上",
			},
			wantErr:  true,
			errField: "核心意義",
		},
		{
			name: "Missing Pronunciation",
			card: NormalWordCard{
				CoreWord:        "猫",
				CoreMeaning:     "貓",
				ContextSentence: "猫が屋根の上にいる",
				Translation:     "貓在屋頂上",
			},
			wantErr:  true,
			errField: "發音",
		},
		{
			name: "Missing ContextSentence",
			card: NormalWordCard{
				CoreWord:      "猫",
				CoreMeaning:   "貓",
				Pronunciation: "ねこ",
				Translation:   "貓在屋頂上",
			},
			wantErr:  true,
			errField: "情境例句",
		},
		{
			name: "Missing Translation",
			card: NormalWordCard{
				CoreWord:        "猫",
				CoreMeaning:     "貓",
				Pronunciation:   "ねこ",
				ContextSentence: "猫が屋根の上にいる",
			},
			wantErr:  true,
			errField: "例句翻譯",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.card.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("NormalWordCard.Validate() error = %v, wantErr %v", err, tt.wantErr)
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

func TestNormalWordCard_ToMap(t *testing.T) {
	card := NormalWordCard{
		CoreWord:        "猫",
		WordType:        "名詞",
		CoreMeaning:     "貓",
		Pronunciation:   "ねこ",
		Accent:          "2",
		Usage:           "動物を表す名詞",
		ContextSentence: "猫が屋根の上にいる",
		Translation:     "貓在屋頂上",
		Synonyms:        "ニャンコ",
		Antonyms:        "犬",
		ImageHint:       "http://example.com/cat.jpg",
	}

	expected := map[string]interface{}{
		"核心單字": "猫",
		"詞性分類": "名詞",
		"核心意義": "貓",
		"發音":   "ねこ",
		"重音":   "2",
		"使用方式": "動物を表す名詞",
		"情境例句": "猫が屋根の上にいる",
		"例句翻譯": "貓在屋頂上",
		"同義詞":  "ニャンコ",
		"反義詞":  "犬",
		"圖片提示": "http://example.com/cat.jpg",
	}

	result := card.ToMap()
	for key, expectedValue := range expected {
		if result[key] != expectedValue {
			t.Errorf("ToMap() for key '%s', expected '%v', got '%v'", key, expectedValue, result[key])
		}
	}
}

func TestNormalWordCard_FromMap(t *testing.T) {
	data := map[string]interface{}{
		"核心單字": "猫",
		"詞性分類": "名詞",
		"核心意義": "貓",
		"發音":   "ねこ",
		"重音":   "2",
		"使用方式": "動物を表す名詞",
		"情境例句": "猫が屋根の上にいる",
		"例句翻譯": "貓在屋頂上",
		"同義詞":  "ニャンコ",
		"反義詞":  "犬",
		"圖片提示": "http://example.com/cat.jpg",
	}

	var card NormalWordCard
	err := card.FromMap(data)
	if err != nil {
		t.Errorf("FromMap() error = %v", err)
		return
	}

	expected := NormalWordCard{
		CoreWord:        "猫",
		WordType:        "名詞",
		CoreMeaning:     "貓",
		Pronunciation:   "ねこ",
		Accent:          "2",
		Usage:           "動物を表す名詞",
		ContextSentence: "猫が屋根の上にいる",
		Translation:     "貓在屋頂上",
		Synonyms:        "ニャンコ",
		Antonyms:        "犬",
		ImageHint:       "http://example.com/cat.jpg",
	}

	if card.CoreWord != expected.CoreWord ||
		card.WordType != expected.WordType ||
		card.CoreMeaning != expected.CoreMeaning ||
		card.Pronunciation != expected.Pronunciation ||
		card.Accent != expected.Accent ||
		card.Usage != expected.Usage ||
		card.ContextSentence != expected.ContextSentence ||
		card.Translation != expected.Translation ||
		card.Synonyms != expected.Synonyms ||
		card.Antonyms != expected.Antonyms ||
		card.ImageHint != expected.ImageHint {
		t.Errorf("FromMap() = %v, want %v", card, expected)
	}
}

func TestNormalWordCard_JSONSerialization(t *testing.T) {
	original := NormalWordCard{
		CoreWord:        "猫",
		WordType:        "名詞",
		CoreMeaning:     "貓",
		Pronunciation:   "ねこ",
		Accent:          "2",
		Usage:           "動物を表す名詞",
		ContextSentence: "猫が屋根の上にいる",
		Translation:     "貓在屋頂上",
		Synonyms:        "ニャンコ",
		Antonyms:        "犬",
		ImageHint:       "http://example.com/cat.jpg",
	}

	// Serialize to JSON
	jsonData, err := json.Marshal(original)
	if err != nil {
		t.Errorf("json.Marshal() error = %v", err)
		return
	}

	// Deserialize from JSON
	var deserialized NormalWordCard
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
		original.Usage != deserialized.Usage ||
		original.ContextSentence != deserialized.ContextSentence ||
		original.Translation != deserialized.Translation ||
		original.Synonyms != deserialized.Synonyms ||
		original.Antonyms != deserialized.Antonyms ||
		original.ImageHint != deserialized.ImageHint {
		t.Errorf("JSON serialization/deserialization failed. Original: %v, Deserialized: %v", original, deserialized)
	}
}
