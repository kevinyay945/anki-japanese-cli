package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"anki-japanese-cli/internal/anki"
	"anki-japanese-cli/internal/config"
	"anki-japanese-cli/internal/models"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add [card-type]",
	Short: "新增日文卡片到 Anki",
	Long: `新增日文卡片到 Anki 牌組。

這個指令支援四種卡片類型：
- verb: 動詞卡片
- adjective: 形容詞卡片
- normal: 一般單字卡片
- grammar: 文法卡片

支援多種新增方式：
- 從 JSON 字串新增
- 從 JSON 檔案讀取
- 批次處理模式

範例:
  anki-japanese-cli add verb --deckName="日文動詞" --json='{"核心單字":"飲む", "詞性分類":"五段動詞", "核心意義":"喝"}'
  anki-japanese-cli add normal --deckName="日文單字" --file=words.json
  anki-japanese-cli add grammar --deckName="日文文法" --batch --file=grammar_batch.json`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		cardType := strings.ToLower(args[0])

		// 驗證卡片類型
		factory := models.NewCardFactory()
		if err := factory.ValidateCardType(cardType); err != nil {
			cmd.PrintErrf("錯誤: %v\n", err)
			return err
		}

		// 取得選項
		deckName, _ := cmd.Flags().GetString("deckName")
		jsonStr, _ := cmd.Flags().GetString("json")
		filePath, _ := cmd.Flags().GetString("file")
		batchMode, _ := cmd.Flags().GetBool("batch")

		// 檢查必要參數
		if deckName == "" {
			cmd.Println("錯誤: 請指定目標牌組名稱 (--deckName)")
			cmd.Help()
			return fmt.Errorf("請指定目標牌組名稱")
		}

		// 載入設定
		cfg, err := config.LoadConfig()
		if err != nil {
			cmd.PrintErrf("錯誤: 無法載入設定: %v\n", err)
			return fmt.Errorf("無法載入設定: %w", err)
		}

		// 建立 Anki 客戶端
		client := anki.NewClient(&cfg.Anki)

		// 檢查 Anki Connect 連線狀態
		cmd.Println("檢查 Anki Connect 連線狀態...")
		if err := client.Ping(); err != nil {
			fmt.Printf("錯誤: 無法連線到 Anki: %v\n", err)
			fmt.Println("請確認 Anki 已啟動且已安裝 AnkiConnect 插件。")
			return fmt.Errorf("無法連線到 Anki: %w", err)
		}
		fmt.Println("✓ 成功連線到 Anki")

		// 確保牌組存在
		fmt.Printf("確保牌組 '%s' 存在...\n", deckName)
		if err := client.EnsureDeckExists(deckName); err != nil {
			fmt.Printf("錯誤: 無法確保牌組存在: %v\n", err)
			return fmt.Errorf("無法確保牌組存在: %w", err)
		}
		fmt.Printf("✓ 牌組 '%s' 已就緒\n", deckName)

		// 檢查模型是否存在
		modelName := ""
		switch cardType {
		case "verb":
			modelName = "Japanese Verb"
		case "adjective":
			modelName = "Japanese Adjective"
		case "normal":
			modelName = "Japanese Normal Word"
		case "grammar":
			modelName = "Japanese Grammar"
		}

		exists, err := client.ModelExists(modelName)
		if err != nil {
			fmt.Printf("錯誤: 檢查模型時發生錯誤: %v\n", err)
			return fmt.Errorf("檢查模型時發生錯誤: %w", err)
		}

		if !exists {
			fmt.Printf("錯誤: 模型 '%s' 不存在。請先執行 'init %s' 指令建立模型。\n", modelName, cardType)
			return fmt.Errorf("模型 '%s' 不存在", modelName)
		}

		// 處理卡片資料
		var cardData []map[string]interface{}

		// 從檔案讀取
		if filePath != "" {
			fmt.Printf("從檔案 '%s' 讀取卡片資料...\n", filePath)
			fileContent, err := os.ReadFile(filePath)
			if err != nil {
				fmt.Printf("錯誤: 無法讀取檔案: %v\n", err)
				return fmt.Errorf("無法讀取檔案: %w", err)
			}

			// 批次模式
			if batchMode {
				if err := json.Unmarshal(fileContent, &cardData); err != nil {
					fmt.Printf("錯誤: JSON 解析失敗: %v\n", err)
					return fmt.Errorf("JSON 解析失敗: %w", err)
				}
			} else {
				// 單一卡片模式
				var singleCard map[string]interface{}
				if err := json.Unmarshal(fileContent, &singleCard); err != nil {
					fmt.Printf("錯誤: JSON 解析失敗: %v\n", err)
					return fmt.Errorf("JSON 解析失敗: %w", err)
				}
				cardData = append(cardData, singleCard)
			}
		} else if jsonStr != "" {
			// 從 JSON 字串讀取
			var singleCard map[string]interface{}
			if err := json.Unmarshal([]byte(jsonStr), &singleCard); err != nil {
				fmt.Printf("錯誤: JSON 解析失敗: %v\n", err)
				return fmt.Errorf("JSON 解析失敗: %w", err)
			}
			cardData = append(cardData, singleCard)
		} else {
			fmt.Println("錯誤: 請提供卡片資料 (--json 或 --file)")
			cmd.Help()
			return fmt.Errorf("請提供卡片資料")
		}

		// 驗證卡片資料
		if len(cardData) == 0 {
			fmt.Println("錯誤: 沒有有效的卡片資料")
			return fmt.Errorf("沒有有效的卡片資料")
		}

		// 驗證所有卡片資料
		fmt.Printf("驗證 %d 張卡片資料...\n", len(cardData))

		// 處理每張卡片
		var notes []anki.NoteInfo
		for i, data := range cardData {
			// 驗證卡片資料
			_, err := factory.CreateCard(cardType, data)
			if err != nil {
				fmt.Printf("錯誤: 卡片 #%d 驗證失敗: %v\n", i+1, err)
				return fmt.Errorf("卡片 #%d 驗證失敗: %w", i+1, err)
			}

			// 建立 Anki 筆記
			note := anki.NoteInfo{
				DeckName:  deckName,
				ModelName: modelName,
				Fields:    make(map[string]string),
				Tags:      []string{"anki-japanese-cli", cardType},
			}

			// 轉換欄位
			for key, value := range data {
				if strValue, ok := value.(string); ok {
					note.Fields[key] = strValue
				} else {
					// 將非字串值轉換為字串
					jsonValue, _ := json.Marshal(value)
					note.Fields[key] = string(jsonValue)
				}
			}

			notes = append(notes, note)
		}

		// 新增卡片到 Anki
		if len(notes) == 1 {
			// 單一卡片模式
			fmt.Println("正在新增卡片到 Anki...")
			noteID, err := client.AddNote(notes[0])
			if err != nil {
				fmt.Printf("錯誤: 無法新增卡片: %v\n", err)
				return fmt.Errorf("無法新增卡片: %w", err)
			}
			fmt.Printf("✓ 成功新增卡片 (ID: %d)\n", noteID)
		} else {
			// 批次模式
			fmt.Printf("正在批次新增 %d 張卡片到 Anki...\n", len(notes))
			noteIDs, err := client.AddNotes(notes)
			if err != nil {
				fmt.Printf("錯誤: 無法批次新增卡片: %v\n", err)
				return fmt.Errorf("無法批次新增卡片: %w", err)
			}

			// 計算成功和失敗的數量
			successCount := 0
			for _, id := range noteIDs {
				if id != 0 {
					successCount++
				}
			}

			fmt.Printf("✓ 成功新增 %d/%d 張卡片\n", successCount, len(notes))
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// 定義 flags
	addCmd.Flags().String("deckName", "", "目標牌組名稱")
	addCmd.Flags().String("json", "", "JSON 格式的卡片資料")
	addCmd.Flags().StringP("file", "f", "", "包含卡片資料的 JSON 檔案路徑")
	addCmd.Flags().BoolP("batch", "b", false, "批次處理模式 (從檔案讀取多張卡片)")
}
