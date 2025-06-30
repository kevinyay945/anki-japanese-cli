package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

// Config 代表應用程式的設定結構
type Config struct {
	Anki     AnkiConfig     `mapstructure:"anki"`
	Template TemplateConfig `mapstructure:"template"`
}

// AnkiConfig 包含 Anki Connect 相關設定
type AnkiConfig struct {
	ConnectURL string `mapstructure:"connect_url"`
	DeckName   string `mapstructure:"deck_name"`
}

// TemplateConfig 包含卡片模板相關設定
type TemplateConfig struct {
	NoteType string   `mapstructure:"note_type"`
	Tags     []string `mapstructure:"tags"`
}

// LoadConfig 載入設定檔案
func LoadConfig() (*Config, error) {
	var config Config

	// 設定預設值
	viper.SetDefault("anki.connect_url", "http://localhost:8765")
	viper.SetDefault("anki.deck_name", "日文學習")
	viper.SetDefault("template.note_type", "Basic")
	viper.SetDefault("template.tags", []string{"japanese", "vocabulary"})

	// 尋找並讀取設定檔案
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("無法取得使用者家目錄: %w", err)
	}

	viper.AddConfigPath(home)
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")
	viper.SetConfigName(".anki-japanese-cli")

	// 自動讀取環境變數
	viper.AutomaticEnv()

	// 讀取設定檔案
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// 設定檔案不存在，使用預設值
			fmt.Println("未找到設定檔案，使用預設設定")
		} else {
			return nil, fmt.Errorf("讀取設定檔案時發生錯誤: %w", err)
		}
	}

	// 將設定資料映射到結構體
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("解析設定檔案時發生錯誤: %w", err)
	}

	return &config, nil
}

// SaveConfig 儲存設定到檔案
func SaveConfig(config *Config) error {
	home, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("無法取得使用者家目錄: %w", err)
	}

	viper.Set("anki", config.Anki)
	viper.Set("template", config.Template)

	configPath := fmt.Sprintf("%s/.anki-japanese-cli.yaml", home)
	return viper.WriteConfigAs(configPath)
}
