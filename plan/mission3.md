# Mission3: 卡片類型模型和模板系統

## 目標
實作四種卡片類型的資料模型和 HTML 模板系統

## 任務細節

### 3.1 定義卡片類型資料結構
建立每種卡片類型的 Go struct:

#### 動詞卡片 (Verb)
```go
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
```

#### 形容詞卡片 (Adjective)
```go
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
```

#### 一般單字卡片 (Normal Words)
#### 文法卡片 (Grammar)

### 3.2 建立 HTML 模板系統
- 建立每種卡片類型的 HTML 模板檔案
- 實作模板渲染功能
- 包含 CSS 樣式的整合

### 3.3 模板管理器
- 建立模板載入和管理系統
- 支援模板的動態選擇
- 模板驗證功能

### 3.4 卡片工廠模式
- 實作卡片建立的工廠模式
- 支援從 JSON 輸入建立不同類型的卡片
- 卡片資料驗證功能

## 預期產出
- 完整的四種卡片類型資料模型
- HTML 模板檔案和渲染系統
- 卡片建立和驗證機制
