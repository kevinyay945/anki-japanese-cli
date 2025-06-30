package models

import "encoding/json"

// GrammarCard 文法卡片類型
type GrammarCard struct {
	GrammarPoint   string `json:"文法要點"`
	Structure      string `json:"結構形式"`
	Meaning        string `json:"意義說明"`
	Usage          string `json:"使用時機"`
	Examples       string `json:"例句示範"`
	Translation    string `json:"例句翻譯"`
	Challenge      string `json:"情境課題"`
	Answer         string `json:"解答範例"`
	Level          string `json:"難度等級"`
	RelatedGrammar string `json:"相關文法,omitempty"`
	CommonMistakes string `json:"常見錯誤,omitempty"`
	Tips           string `json:"記憶技巧,omitempty"`
}

// GetCardType 返回卡片類型
func (g *GrammarCard) GetCardType() string {
	return "grammar"
}

// Validate 驗證卡片資料
func (g *GrammarCard) Validate() error {
	if g.GrammarPoint == "" {
		return NewValidationError("文法要點", "不能為空")
	}
	if g.Structure == "" {
		return NewValidationError("結構形式", "不能為空")
	}
	if g.Meaning == "" {
		return NewValidationError("意義說明", "不能為空")
	}
	if g.Examples == "" {
		return NewValidationError("例句示範", "不能為空")
	}
	if g.Translation == "" {
		return NewValidationError("例句翻譯", "不能為空")
	}
	return nil
}

// ToMap 轉換為 map 格式
func (g *GrammarCard) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"文法要點": g.GrammarPoint,
		"結構形式": g.Structure,
		"意義說明": g.Meaning,
		"使用時機": g.Usage,
		"例句示範": g.Examples,
		"例句翻譯": g.Translation,
		"情境課題": g.Challenge,
		"解答範例": g.Answer,
		"難度等級": g.Level,
		"相關文法": g.RelatedGrammar,
		"常見錯誤": g.CommonMistakes,
		"記憶技巧": g.Tips,
	}
}

// FromMap 從 map 載入資料
func (g *GrammarCard) FromMap(data map[string]interface{}) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return json.Unmarshal(jsonData, g)
}
