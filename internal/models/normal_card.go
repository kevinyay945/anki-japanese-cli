package models

import "encoding/json"

// NormalWordCard 一般單字卡片類型
type NormalWordCard struct {
	CoreWord        string `json:"核心單字"`
	WordType        string `json:"詞性分類"`
	CoreMeaning     string `json:"核心意義"`
	Pronunciation   string `json:"發音"`
	Accent          string `json:"重音"`
	Usage           string `json:"使用方式"`
	ContextSentence string `json:"情境例句"`
	Translation     string `json:"例句翻譯"`
	Synonyms        string `json:"同義詞,omitempty"`
	Antonyms        string `json:"反義詞,omitempty"`
	ImageHint       string `json:"圖片提示,omitempty"`
}

// GetCardType 返回卡片類型
func (n *NormalWordCard) GetCardType() string {
	return "normal"
}

// Validate 驗證卡片資料
func (n *NormalWordCard) Validate() error {
	if n.CoreWord == "" {
		return NewValidationError("核心單字", "不能為空")
	}
	if n.CoreMeaning == "" {
		return NewValidationError("核心意義", "不能為空")
	}
	if n.Pronunciation == "" {
		return NewValidationError("發音", "不能為空")
	}
	if n.ContextSentence == "" {
		return NewValidationError("情境例句", "不能為空")
	}
	if n.Translation == "" {
		return NewValidationError("例句翻譯", "不能為空")
	}
	return nil
}

// ToMap 轉換為 map 格式
func (n *NormalWordCard) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"核心單字": n.CoreWord,
		"詞性分類": n.WordType,
		"核心意義": n.CoreMeaning,
		"發音":   n.Pronunciation,
		"重音":   n.Accent,
		"使用方式": n.Usage,
		"情境例句": n.ContextSentence,
		"例句翻譯": n.Translation,
		"同義詞":  n.Synonyms,
		"反義詞":  n.Antonyms,
		"圖片提示": n.ImageHint,
	}
}

// FromMap 從 map 載入資料
func (n *NormalWordCard) FromMap(data map[string]interface{}) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return json.Unmarshal(jsonData, n)
}
