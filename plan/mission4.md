# Mission4: Anki Connect 整合

## 目標
實作與 Anki Connect API 的整合，支援卡片模型建立和卡片新增功能

## 任務細節

### 4.1 Anki Connect 客戶端
- [x] 建立 HTTP 客戶端來與 Anki Connect 通訊
- [x] 實作基本的 API 呼叫封裝
- [x] 錯誤處理和重試機制

### 4.2 卡片模型管理
實作以下 Anki Connect API 呼叫:
- [x] `createModel`: 建立新的卡片模型
- [x] `modelNames`: 列出現有模型
- [x] `modelFieldNames`: 獲取模型欄位
- [x] `updateModelTemplates`: 更新模型模板

### 4.3 卡片操作
實作以下 Anki Connect API 呼叫:
- [x] `addNote`: 新增單張卡片
- [x] `addNotes`: 批次新增多張卡片
- [x] `deckNames`: 列出可用的牌組
- [x] `createDeck`: 建立新牌組

### 4.4 資料轉換
- [x] 實作 Go struct 到 Anki Connect JSON 格式的轉換
- [x] 支援批次操作的資料格式化
- [x] 建立卡片模型的自動建立機制

### 4.5 連線狀態檢查
- [x] 實作 Anki Connect 連線狀態檢查
- [x] 提供連線診斷功能
- [x] 友善的錯誤訊息顯示

## 預期產出
- [x] 完整的 Anki Connect 客戶端函式庫
- [x] 卡片模型和卡片的 CRUD 操作
- [x] 連線狀態檢查和診斷工具
