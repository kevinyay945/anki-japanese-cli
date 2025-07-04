# Mission2: 專案初始化和基礎架構

## 目標
建立 Golang CLI 專案的基礎架構，包含 Cobra CLI 框架和 Viper 設定管理

## 任務細節

### 2.1 專案初始化
- [x] 初始化 Go module (`go mod init anki-japanese-cli`)
- [x] 安裝必要的依賴套件:
  - [x] github.com/spf13/cobra
  - [x] github.com/spf13/viper
  - [x] 其他必要的 HTTP 客戶端函式庫

### 2.2 建立基本目錄結構
```
anki-japanese-cli/
├── cmd/
│   ├── root.go
│   ├── init.go
│   └── add.go
├── internal/
│   ├── config/
│   ├── models/
│   ├── templates/
│   └── anki/
├── pkg/
├── .gitignore
└── main.go
```

### 2.3 設定 Cobra CLI 根指令
- [x] 建立主要的 CLI 入口點
- [x] 設定基本的指令結構
- [x] 實作 `--version` 和 `--help` 功能

### 2.4 建立設定檔案系統
- [x] 使用 Viper 建立設定檔案管理
- [x] 定義預設的設定檔案位置
- [x] 支援 YAML/JSON 設定檔案格式

## 預期產出
- [x] 可執行的基本 CLI 工具
- [x] 完整的專案目錄結構
- [x] 基本的指令架構 (`./anki-japanese-cli --help` 可以執行)
