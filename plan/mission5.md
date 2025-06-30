# Mission5: CLI 指令實作

## 目標
實作主要的 CLI 指令，包含 `init` 和 `add` 指令

## 任務細節

### 5.1 實作 `init` 指令
```bash
./anki-japanese-cli init <card-type>
```

功能:
- 支援四種卡片類型: `verb`, `adjective`, `normal`, `grammar`
- 在 Anki 中建立對應的卡片模型
- 自動建立必要的牌組
- 提供初始化狀態的回饋

參數驗證:
- 檢查卡片類型是否有效
- 檢查 Anki Connect 連線狀態
- 檢查模型是否已存在

### 5.2 實作 `add` 指令
```bash
./anki-japanese-cli add <card-type> --deckName='japanese-2025' --json='{...}'
```

功能:
- 支援四種卡片類型的新增
- 從 JSON 輸入解析卡片資料
- 驗證必要欄位
- 新增卡片到指定牌組

參數和選項:
- `--deckName`: 指定目標牌組名稱
- `--json`: JSON 格式的卡片資料
- `--file`: 從檔案讀取 JSON 資料 (額外功能)
- `--batch`: 批次處理模式 (額外功能)

### 5.3 輸入驗證和錯誤處理
- JSON 格式驗證
- 必要欄位檢查
- 資料型別驗證
- 友善的錯誤訊息

### 5.4 互動式模式 (額外功能)
- 提供互動式的卡片建立模式
- 欄位提示和驗證
- 預覽功能

## 預期產出
- 完整的 `init` 指令實作
- 完整的 `add` 指令實作
- 輸入驗證和錯誤處理機制
- 使用範例和說明文件
