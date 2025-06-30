package models

import (
	"encoding/json"
	"fmt"
)

// CardFactory 卡片工廠
type CardFactory struct{}

// NewCardFactory 建立新的卡片工廠
func NewCardFactory() *CardFactory {
	return &CardFactory{}
}

// CreateCard 根據類型和資料建立卡片
func (cf *CardFactory) CreateCard(cardType string, data map[string]interface{}) (CardType, error) {
	switch cardType {
	case "verb":
		return cf.createVerbCard(data)
	case "adjective":
		return cf.createAdjectiveCard(data)
	case "normal":
		return cf.createNormalCard(data)
	case "grammar":
		return cf.createGrammarCard(data)
	default:
		return nil, fmt.Errorf("不支援的卡片類型: %s", cardType)
	}
}

// CreateCardFromJSON 從 JSON 資料建立卡片
func (cf *CardFactory) CreateCardFromJSON(cardType string, jsonData []byte) (CardType, error) {
	var data map[string]interface{}
	err := json.Unmarshal(jsonData, &data)
	if err != nil {
		return nil, fmt.Errorf("JSON 解析失敗: %w", err)
	}

	return cf.CreateCard(cardType, data)
}

// createVerbCard 建立動詞卡片
func (cf *CardFactory) createVerbCard(data map[string]interface{}) (*VerbCard, error) {
	card := &VerbCard{}
	err := card.FromMap(data)
	if err != nil {
		return nil, fmt.Errorf("建立動詞卡片失敗: %w", err)
	}

	err = card.Validate()
	if err != nil {
		return nil, fmt.Errorf("動詞卡片驗證失敗: %w", err)
	}

	return card, nil
}

// createAdjectiveCard 建立形容詞卡片
func (cf *CardFactory) createAdjectiveCard(data map[string]interface{}) (*AdjectiveCard, error) {
	card := &AdjectiveCard{}
	err := card.FromMap(data)
	if err != nil {
		return nil, fmt.Errorf("建立形容詞卡片失敗: %w", err)
	}

	err = card.Validate()
	if err != nil {
		return nil, fmt.Errorf("形容詞卡片驗證失敗: %w", err)
	}

	return card, nil
}

// createNormalCard 建立一般單字卡片
func (cf *CardFactory) createNormalCard(data map[string]interface{}) (*NormalWordCard, error) {
	card := &NormalWordCard{}
	err := card.FromMap(data)
	if err != nil {
		return nil, fmt.Errorf("建立一般單字卡片失敗: %w", err)
	}

	err = card.Validate()
	if err != nil {
		return nil, fmt.Errorf("一般單字卡片驗證失敗: %w", err)
	}

	return card, nil
}

// createGrammarCard 建立文法卡片
func (cf *CardFactory) createGrammarCard(data map[string]interface{}) (*GrammarCard, error) {
	card := &GrammarCard{}
	err := card.FromMap(data)
	if err != nil {
		return nil, fmt.Errorf("建立文法卡片失敗: %w", err)
	}

	err = card.Validate()
	if err != nil {
		return nil, fmt.Errorf("文法卡片驗證失敗: %w", err)
	}

	return card, nil
}

// GetSupportedCardTypes 獲取支援的卡片類型
func (cf *CardFactory) GetSupportedCardTypes() []string {
	return []string{"verb", "adjective", "normal", "grammar"}
}

// ValidateCardType 驗證卡片類型是否支援
func (cf *CardFactory) ValidateCardType(cardType string) error {
	supportedTypes := cf.GetSupportedCardTypes()
	for _, t := range supportedTypes {
		if t == cardType {
			return nil
		}
	}
	return fmt.Errorf("不支援的卡片類型: %s。支援的類型: %v", cardType, supportedTypes)
}
