package models

import "encoding/json"

// AdjectiveCard 形容詞卡片類型
type AdjectiveCard struct {
	CoreWord        string `json:"核心單字"`
	WordType        string `json:"詞性分類"`
	CoreMeaning     string `json:"核心意義"`
	Pronunciation   string `json:"發音"`
	Accent          string `json:"重音"`
	MainChanges     string `json:"主要變化"`
	ContextSentence string `json:"情境例句"`
	Translation     string `json:"例句翻譯"`
	RelatedWords    string `json:"相關詞彙,omitempty"`
}

// GetCardType 返回卡片類型
func (a *AdjectiveCard) GetCardType() string {
	return "adjective"
}

// Validate 驗證卡片資料
func (a *AdjectiveCard) Validate() error {
	if a.CoreWord == "" {
		return NewValidationError("核心單字", "不能為空")
	}
	if a.CoreMeaning == "" {
		return NewValidationError("核心意義", "不能為空")
	}
	if a.Pronunciation == "" {
		return NewValidationError("發音", "不能為空")
	}
	if a.ContextSentence == "" {
		return NewValidationError("情境例句", "不能為空")
	}
	if a.Translation == "" {
		return NewValidationError("例句翻譯", "不能為空")
	}
	return nil
}

// ToMap 轉換為 map 格式
func (a *AdjectiveCard) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"核心單字": a.CoreWord,
		"詞性分類": a.WordType,
		"核心意義": a.CoreMeaning,
		"發音":   a.Pronunciation,
		"重音":   a.Accent,
		"主要變化": a.MainChanges,
		"情境例句": a.ContextSentence,
		"例句翻譯": a.Translation,
		"相關詞彙": a.RelatedWords,
	}
}

// FromMap 從 map 載入資料
func (a *AdjectiveCard) FromMap(data map[string]interface{}) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return json.Unmarshal(jsonData, a)
}
