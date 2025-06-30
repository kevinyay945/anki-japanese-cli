# Mission6: 測試和文件

## 目標
建立完整的測試套件和使用文件

## 任務細節

### 6.1 單元測試
建立測試檔案:
- `internal/models/*_test.go`: 卡片模型測試
- `internal/anki/*_test.go`: Anki Connect 客戶端測試
- `internal/templates/*_test.go`: 模板系統測試
- `cmd/*_test.go`: CLI 指令測試

測試覆蓋範圍:
- 資料結構驗證
- JSON 序列化/反序列化
- HTTP 客戶端模擬測試
- 模板渲染測試
- CLI 參數解析測試

### 6.2 整合測試
- Anki Connect API 整合測試
- 端到端 CLI 工作流程測試
- 錯誤情境測試
- 效能測試

### 6.3 測試工具和模擬
- HTTP 客戶端模擬器
- Anki Connect API 模擬伺服器
- 測試資料產生器
- 測試輔助函式

### 6.4 文件撰寫
建立以下文件:
- `README.md`: 專案說明和安裝指南
- `USAGE.md`: 詳細使用說明
- `API.md`: Anki Connect 整合說明
- `CONTRIBUTING.md`: 開發貢獻指南
- `CHANGELOG.md`: 版本更新記錄

### 6.5 範例檔案
建立範例檔案:
- `examples/verb_cards.json`: 動詞卡片範例
- `examples/adjective_cards.json`: 形容詞卡片範例
- `examples/normal_cards.json`: 一般單字卡片範例
- `examples/grammar_cards.json`: 文法卡片範例
- `examples/batch_import.json`: 批次匯入範例

### 6.6 程式碼品質
- 實作 linting 規則
- 程式碼格式化設定
- 持續整合設定 (GitHub Actions)
- 程式碼覆蓋率報告

## 預期產出
- 完整的測試套件 (單元測試 + 整合測試)
- 詳細的使用文件和 API 文件
- 豐富的範例檔案
- 程式碼品質保證機制
