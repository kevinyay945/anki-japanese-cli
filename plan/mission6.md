# Mission6: 測試和文件

## 目標

建立完整的測試套件和使用文件

## 任務細節

### 6.1 單元測試

建立測試檔案:

- [x] `internal/models/*_test.go`: 卡片模型測試
- [x] `internal/anki/*_test.go`: Anki Connect 客戶端測試
- [x] `internal/templates/*_test.go`: 模板系統測試
- [x] `cmd/*_test.go`: CLI 指令測試

測試覆蓋範圍:

- [x] 資料結構驗證
- [x] JSON 序列化/反序列化
- [x] HTTP 客戶端模擬測試
- [x] 模板渲染測試
- [x] CLI 參數解析測試

### 6.2 整合測試

- [x] Anki Connect API 整合測試
- [x] 端到端 CLI 工作流程測試
- [x] 錯誤情境測試

### 6.3 測試工具和模擬

- [x] HTTP 客戶端模擬器
- [x] Anki Connect API 模擬伺服器
- [x] 測試資料產生器
- [x] 測試輔助函式

### 6.4 文件撰寫

建立以下文件:

- [x] `README.md`: 專案說明和安裝指南
- [x] `USAGE.md`: 詳細使用說明
- [x] `API.md`: Anki Connect 整合說明
- [x] `CONTRIBUTING.md`: 開發貢獻指南
- [x] `CHANGELOG.md`: 版本更新記錄

### 6.5 範例檔案

建立範例檔案:

- [x] `examples/verb_cards.json`: 動詞卡片範例
- [x] `examples/adjective_cards.json`: 形容詞卡片範例
- [x] `examples/normal_cards.json`: 一般單字卡片範例
- [x] `examples/grammar_cards.json`: 文法卡片範例
- [x] `examples/batch_import.json`: 批次匯入範例

## 預期產出

- [x] 完整的測試套件 (單元測試 + 整合測試)
- [x] 詳細的使用文件和 API 文件
- [x] 豐富的範例檔案
