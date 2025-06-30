package models

import (
	"anki-japanese-cli/internal/templates"
	"fmt"
)

// CardService 卡片服務
type CardService struct {
	factory         *CardFactory
	templateManager *templates.TemplateManager
}

// NewCardService 建立新的卡片服務
func NewCardService() (*CardService, error) {
	templateManager, err := templates.NewTemplateManager()
	if err != nil {
		return nil, fmt.Errorf("初始化模板管理器失敗: %w", err)
	}

	return &CardService{
		factory:         NewCardFactory(),
		templateManager: templateManager,
	}, nil
}

// CreateAndRenderCard 建立並渲染卡片
func (cs *CardService) CreateAndRenderCard(cardType string, data map[string]interface{}) (string, error) {
	// 驗證卡片類型
	err := cs.factory.ValidateCardType(cardType)
	if err != nil {
		return "", err
	}

	// 驗證模板存在
	err = cs.templateManager.ValidateTemplate(cardType)
	if err != nil {
		return "", err
	}

	// 建立卡片
	card, err := cs.factory.CreateCard(cardType, data)
	if err != nil {
		return "", err
	}

	// 獲取卡片資料
	var cardData interface{}
	if cardDataProvider, ok := card.(CardData); ok {
		cardData = cardDataProvider.ToMap()
	} else {
		cardData = card
	}

	// 渲染卡片背面 (向後相容)
	html, err := cs.templateManager.RenderCardBack(cardType, cardData)
	if err != nil {
		return "", fmt.Errorf("渲染卡片失敗: %w", err)
	}

	return html, nil
}

// CreateAndRenderCardFront 建立並渲染卡片正面
func (cs *CardService) CreateAndRenderCardFront(cardType string, data map[string]interface{}) (string, error) {
	// 驗證卡片類型
	err := cs.factory.ValidateCardType(cardType)
	if err != nil {
		return "", err
	}

	// 驗證模板存在
	err = cs.templateManager.ValidateTemplate(cardType)
	if err != nil {
		return "", err
	}

	// 建立卡片
	card, err := cs.factory.CreateCard(cardType, data)
	if err != nil {
		return "", err
	}

	// 獲取卡片資料
	var cardData interface{}
	if cardDataProvider, ok := card.(CardData); ok {
		cardData = cardDataProvider.ToMap()
	} else {
		cardData = card
	}

	// 渲染卡片正面
	html, err := cs.templateManager.RenderCardFront(cardType, cardData)
	if err != nil {
		return "", fmt.Errorf("渲染卡片正面失敗: %w", err)
	}

	return html, nil
}

// CreateAndRenderCardBack 建立並渲染卡片背面
func (cs *CardService) CreateAndRenderCardBack(cardType string, data map[string]interface{}) (string, error) {
	// 驗證卡片類型
	err := cs.factory.ValidateCardType(cardType)
	if err != nil {
		return "", err
	}

	// 驗證模板存在
	err = cs.templateManager.ValidateTemplate(cardType)
	if err != nil {
		return "", err
	}

	// 建立卡片
	card, err := cs.factory.CreateCard(cardType, data)
	if err != nil {
		return "", err
	}

	// 獲取卡片資料
	var cardData interface{}
	if cardDataProvider, ok := card.(CardData); ok {
		cardData = cardDataProvider.ToMap()
	} else {
		cardData = card
	}

	// 渲染卡片背面
	html, err := cs.templateManager.RenderCardBack(cardType, cardData)
	if err != nil {
		return "", fmt.Errorf("渲染卡片背面失敗: %w", err)
	}

	return html, nil
}

// CreateCardFromJSON 從 JSON 建立並渲染卡片
func (cs *CardService) CreateCardFromJSON(cardType string, jsonData []byte) (string, error) {
	card, err := cs.factory.CreateCardFromJSON(cardType, jsonData)
	if err != nil {
		return "", err
	}

	// 獲取卡片資料
	var cardData interface{}
	if cardDataProvider, ok := card.(CardData); ok {
		cardData = cardDataProvider.ToMap()
	} else {
		cardData = card
	}

	// 渲染卡片
	html, err := cs.templateManager.RenderCard(cardType, cardData)
	if err != nil {
		return "", fmt.Errorf("渲染卡片失敗: %w", err)
	}

	return html, nil
}

// GetSupportedCardTypes 獲取支援的卡片類型
func (cs *CardService) GetSupportedCardTypes() []string {
	return cs.factory.GetSupportedCardTypes()
}

// GetAvailableTemplates 獲取可用的模板
func (cs *CardService) GetAvailableTemplates() []string {
	return cs.templateManager.GetAvailableTemplates()
}
