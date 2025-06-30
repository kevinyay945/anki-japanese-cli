package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "初始化 anki-japanese-cli 設定",
	Long: `初始化 anki-japanese-cli 的設定檔案。

這個指令會建立一個預設的設定檔案，包含：
- Anki Connect 連線設定
- 預設的卡片模板設定
- 其他必要的組態資訊`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("正在初始化 anki-japanese-cli 設定...")

		home, err := os.UserHomeDir()
		if err != nil {
			fmt.Printf("錯誤：無法取得使用者家目錄: %v\n", err)
			os.Exit(1)
		}

		configPath := filepath.Join(home, ".anki-japanese-cli.yaml")

		// 檢查設定檔案是否已存在
		if _, err := os.Stat(configPath); err == nil {
			fmt.Printf("設定檔案已存在於: %s\n", configPath)
			return
		}

		// 建立預設設定
		defaultConfig := map[string]interface{}{
			"anki": map[string]interface{}{
				"connect_url": "http://localhost:8765",
				"deck_name":   "日文學習",
			},
			"template": map[string]interface{}{
				"note_type": "Basic",
				"tags":      []string{"japanese", "vocabulary"},
			},
		}

		// 設定 viper
		for key, value := range defaultConfig {
			viper.Set(key, value)
		}

		// 寫入設定檔案
		err = viper.WriteConfigAs(configPath)
		if err != nil {
			fmt.Printf("錯誤：無法建立設定檔案: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("設定檔案已建立於: %s\n", configPath)
		fmt.Println("您可以編輯此檔案來自訂設定。")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
