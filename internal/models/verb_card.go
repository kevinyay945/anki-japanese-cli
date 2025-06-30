package models

import "encoding/json"

// VerbCard 動詞卡片類型
type VerbCard struct {
	CoreWord        string `json:"核心單字"`
	WordType        string `json:"詞性分類"`
	CoreMeaning     string `json:"核心意義"`
	Pronunciation   string `json:"發音"`
	Accent          string `json:"重音"`
	Conjugations    string `json:"常用變化"`
	ContextSentence string `json:"情境例句"`
	Translation     string `json:"例句翻譯"`
	ImageHint       string `json:"圖片提示,omitempty"`
}

// GetCardType 返回卡片類型
func (v *VerbCard) GetCardType() string {
	return "verb"
}

// Validate 驗證卡片資料
func (v *VerbCard) Validate() error {
	if v.CoreWord == "" {
		return NewValidationError("核心單字", "不能為空")
	}
	if v.CoreMeaning == "" {
		return NewValidationError("核心意義", "不能為空")
	}
	if v.Pronunciation == "" {
		return NewValidationError("發音", "不能為空")
	}
	if v.ContextSentence == "" {
		return NewValidationError("情境例句", "不能為空")
	}
	if v.Translation == "" {
		return NewValidationError("例句翻譯", "不能為空")
	}
	return nil
}

// ToMap 轉換為 map 格式
func (v *VerbCard) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"核心單字": v.CoreWord,
		"詞性分類": v.WordType,
		"核心意義": v.CoreMeaning,
		"發音":   v.Pronunciation,
		"重音":   v.Accent,
		"常用變化": v.Conjugations,
		"情境例句": v.ContextSentence,
		"例句翻譯": v.Translation,
		"圖片提示": v.ImageHint,
	}
}

// FromMap 從 map 載入資料
func (v *VerbCard) FromMap(data map[string]interface{}) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return json.Unmarshal(jsonData, v)
}
