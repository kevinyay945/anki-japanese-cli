package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "新增日文詞彙卡片到 Anki",
	Long: `新增日文詞彙卡片到 Anki 套牌。

這個指令支援多種新增方式：
- 單一詞彙新增
- 批次匯入
- 從檔案讀取

範例:
  anki-japanese-cli add --word "こんにちは" --meaning "你好"
  anki-japanese-cli add --file vocabulary.txt`,
	Run: func(cmd *cobra.Command, args []string) {
		word, _ := cmd.Flags().GetString("word")
		meaning, _ := cmd.Flags().GetString("meaning")
		file, _ := cmd.Flags().GetString("file")

		if file != "" {
			fmt.Printf("從檔案匯入詞彙: %s\n", file)
			// TODO: 實作檔案匯入功能
			fmt.Println("檔案匯入功能將在後續版本實作")
			return
		}

		if word == "" {
			fmt.Println("錯誤：請提供要新增的詞彙")
			cmd.Help()
			return
		}

		if meaning == "" {
			fmt.Println("錯誤：請提供詞彙的意思")
			cmd.Help()
			return
		}

		fmt.Printf("正在新增詞彙卡片...\n")
		fmt.Printf("詞彙: %s\n", word)
		fmt.Printf("意思: %s\n", meaning)

		// TODO: 實作與 Anki Connect 的整合
		fmt.Println("Anki Connect 整合功能將在後續版本實作")
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// 定義 flags
	addCmd.Flags().StringP("word", "w", "", "要新增的日文詞彙")
	addCmd.Flags().StringP("meaning", "m", "", "詞彙的意思")
	addCmd.Flags().StringP("file", "f", "", "包含詞彙清單的檔案路徑")
}
