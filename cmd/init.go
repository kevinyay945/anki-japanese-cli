package cmd

import (
	"fmt"
	"strings"

	"anki-japanese-cli/internal/anki"
	"anki-japanese-cli/internal/config"
	"anki-japanese-cli/internal/models"
	"anki-japanese-cli/internal/templates"

	"github.com/spf13/cobra"
)

// CSS 樣式
const defaultCSS = `
/* 共通樣式 */
.card {
  font-family: "Hiragino Sans", "Yu Gothic", "Meiryo", sans-serif;
  background: linear-gradient(135deg, #f5f7fa 0%, #c3cfe2 100%);
  padding: 20px;
  border-radius: 10px;
  box-shadow: 0 4px 15px rgba(0,0,0,0.1);
}

.card-front, .card-back {
  text-align: center;
  max-width: 500px;
  margin: 0 auto;
}

/* 正面樣式 */
.context-sentence {
  font-size: 1.4em;
  line-height: 1.6;
  margin-bottom: 15px;
  color: #2c3e50;
  font-weight: 500;
}

.meaning-hint, .grammar-point-hint {
  font-size: 1em;
  color: #7f8c8d;
  margin-bottom: 10px;
  font-style: italic;
}

/* 背面樣式 */
.core-word, .grammar-point {
  font-size: 2em;
  color: #e74c3c;
  font-weight: bold;
  margin-bottom: 10px;
}

.word-info {
  display: flex;
  justify-content: center;
  gap: 15px;
  margin-bottom: 15px;
  flex-wrap: wrap;
}

.pronunciation, .accent, .word-type {
  background: #3498db;
  color: white;
  padding: 5px 10px;
  border-radius: 5px;
  font-size: 0.9em;
}

.translation, .answer {
  font-size: 1.2em;
  color: #27ae60;
  margin-bottom: 15px;
  font-weight: 500;
}

.conjugations, .related-words, .confusing-grammar {
  background: #ecf0f1;
  padding: 10px;
  border-radius: 5px;
  margin-bottom: 10px;
  font-size: 0.9em;
  color: #34495e;
}

.grammar-info {
  text-align: left;
  background: #f8f9fa;
  padding: 15px;
  border-radius: 8px;
  margin-bottom: 10px;
}

.connection-rules, .usage-notes {
  margin: 8px 0;
  font-size: 0.95em;
  line-height: 1.4;
}

.image-hint img, .image img {
  max-width: 200px;
  border-radius: 8px;
  box-shadow: 0 2px 8px rgba(0,0,0,0.2);
}
`

// 卡片模型定義
var cardModels = map[string]struct {
	Name   string
	Fields []string
	Deck   string
}{
	"verb": {
		Name: "Japanese Verb",
		Fields: []string{
			"核心單字", "詞性分類", "核心意義", "發音", "重音",
			"常用變化", "情境例句", "例句翻譯", "圖片提示",
		},
		Deck: "日文動詞",
	},
	"adjective": {
		Name: "Japanese Adjective",
		Fields: []string{
			"核心單字", "詞性分類", "核心意義", "發音", "重音",
			"主要變化", "情境例句", "例句翻譯", "相關詞彙",
		},
		Deck: "日文形容詞",
	},
	"normal": {
		Name: "Japanese Normal Word",
		Fields: []string{
			"核心單字", "詞性", "核心意義", "發音", "重音",
			"情境例句", "例句翻譯", "相關詞彙", "圖片提示",
		},
		Deck: "日文單字",
	},
	"grammar": {
		Name: "Japanese Grammar",
		Fields: []string{
			"文法點", "核心意義", "接續規則", "語感說明",
			"情境課題", "解答範例", "易混淆文法",
		},
		Deck: "日文文法",
	},
}

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init [card-type]",
	Short: "初始化 Anki 卡片模型",
	Long: `初始化 Anki 卡片模型和牌組。

這個指令會在 Anki 中建立指定類型的卡片模型和必要的牌組。
支援的卡片類型:
- verb: 動詞卡片
- adjective: 形容詞卡片
- normal: 一般單字卡片
- grammar: 文法卡片`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		cardType := strings.ToLower(args[0])

		// 驗證卡片類型
		factory := models.NewCardFactory()
		if err := factory.ValidateCardType(cardType); err != nil {
			cmd.PrintErrf("錯誤: %v\n", err)
			return err
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
			cmd.PrintErrf("錯誤: 無法連線到 Anki: %v\n", err)
			cmd.Println("請確認 Anki 已啟動且已安裝 AnkiConnect 插件。")
			return fmt.Errorf("無法連線到 Anki: %w", err)
		}
		cmd.Println("✓ 成功連線到 Anki")

		// 取得模型定義
		modelDef, exists := cardModels[cardType]
		if !exists {
			cmd.PrintErrf("錯誤: 找不到卡片類型 '%s' 的定義\n", cardType)
			return fmt.Errorf("找不到卡片類型 '%s' 的定義", cardType)
		}

		// 檢查模型是否已存在
		exists, err = client.ModelExists(modelDef.Name)
		if err != nil {
			cmd.PrintErrf("錯誤: 檢查模型時發生錯誤: %v\n", err)
			return fmt.Errorf("檢查模型時發生錯誤: %w", err)
		}

		if exists {
			cmd.Printf("模型 '%s' 已存在\n", modelDef.Name)
		} else {
			// 建立模型
			cmd.Printf("正在建立模型 '%s'...\n", modelDef.Name)

			// 載入模板
			templateManager, err := templates.NewTemplateManager()
			if err != nil {
				cmd.PrintErrf("錯誤: 無法初始化模板管理器: %v\n", err)
				return fmt.Errorf("無法初始化模板管理器: %w", err)
			}

			// 驗證模板
			if err := templateManager.ValidateTemplate(cardType); err != nil {
				cmd.PrintErrf("錯誤: 模板驗證失敗: %v\n", err)
				return fmt.Errorf("模板驗證失敗: %w", err)
			}

			// 使用簡化的 Anki 模板
			frontTemplate := `
<div class="card-front">
  <div class="context-sentence">{{情境例句}}</div>
  <div class="meaning-hint">{{核心意義}}</div>
  {{#圖片提示}}<div class="image-hint"><img src="{{圖片提示}}"></div>{{/圖片提示}}
</div>
`
			backTemplate := `
<div class="card-back">
  <div class="context-sentence">{{情境例句}}</div>
  <div class="core-word">{{核心單字}}</div>
  <div class="word-info">
    <div class="pronunciation">{{發音}}</div>
    <div class="accent">{{重音}}</div>
    <div class="word-type">{{詞性分類}}</div>
  </div>
  <div class="translation">{{例句翻譯}}</div>
  <div class="conjugations">{{常用變化}}</div>
  {{#圖片提示}}<div class="image"><img src="{{圖片提示}}"></div>{{/圖片提示}}
</div>
`
			// 根據卡片類型調整模板
			if cardType == "adjective" {
				frontTemplate = `
<div class="card-front">
  <div class="context-sentence">{{情境例句}}</div>
  <div class="meaning-hint">{{核心意義}}</div>
</div>
`
				backTemplate = `
<div class="card-back">
  <div class="context-sentence">{{情境例句}}</div>
  <div class="core-word">{{核心單字}}</div>
  <div class="word-info">
    <div class="pronunciation">{{發音}}</div>
    <div class="accent">{{重音}}</div>
    <div class="word-type">{{詞性分類}}</div>
  </div>
  <div class="translation">{{例句翻譯}}</div>
  <div class="conjugations">{{主要變化}}</div>
  {{#相關詞彙}}<div class="related-words">{{相關詞彙}}</div>{{/相關詞彙}}
</div>
`
			} else if cardType == "normal" {
				frontTemplate = `
<div class="card-front">
  <div class="context-sentence">{{情境例句}}</div>
  <div class="meaning-hint">{{核心意義}}</div>
  {{#圖片提示}}<div class="image-hint"><img src="{{圖片提示}}"></div>{{/圖片提示}}
</div>
`
				backTemplate = `
<div class="card-back">
  <div class="context-sentence">{{情境例句}}</div>
  <div class="core-word">{{核心單字}}</div>
  <div class="word-info">
    <div class="pronunciation">{{發音}}</div>
    <div class="accent">{{重音}}</div>
    <div class="word-type">{{詞性}}</div>
  </div>
  <div class="translation">{{例句翻譯}}</div>
  {{#相關詞彙}}<div class="related-words">{{相關詞彙}}</div>{{/相關詞彙}}
  {{#圖片提示}}<div class="image"><img src="{{圖片提示}}"></div>{{/圖片提示}}
</div>
`
			} else if cardType == "grammar" {
				frontTemplate = `
<div class="card-front">
  <div class="grammar-challenge">{{情境課題}}</div>
  <div class="grammar-point-hint">使用「{{文法點}}」</div>
</div>
`
				backTemplate = `
<div class="card-back">
  <div class="challenge">{{情境課題}}</div>
  <div class="answer">{{解答範例}}</div>
  <div class="grammar-info">
    <div class="grammar-point">{{文法點}}</div>
    <div class="meaning">{{核心意義}}</div>
    <div class="connection-rules">{{接續規則}}</div>
    <div class="usage-notes">{{語感說明}}</div>
  </div>
  {{#易混淆文法}}<div class="confusing-grammar">{{易混淆文法}}</div>{{/易混淆文法}}
</div>
`
			}

			// 建立模型設定
			modelConfig := anki.ModelConfig{
				ModelName:     modelDef.Name,
				InOrderFields: modelDef.Fields,
				CSS:           defaultCSS,
				CardTemplates: []anki.CardTemplateConfig{
					{
						Name:  modelDef.Name,
						Front: frontTemplate,
						Back:  backTemplate,
					},
				},
			}

			// 建立模型
			if err := client.CreateModel(modelConfig); err != nil {
				cmd.PrintErrf("錯誤: 無法建立模型: %v\n", err)
				return fmt.Errorf("無法建立模型: %w", err)
			}

			cmd.Printf("✓ 成功建立模型 '%s'\n", modelDef.Name)
		}

		// 確保牌組存在
		cmd.Printf("正在確保牌組 '%s' 存在...\n", modelDef.Deck)
		if err := client.EnsureDeckExists(modelDef.Deck); err != nil {
			cmd.PrintErrf("錯誤: 無法建立牌組: %v\n", err)
			return fmt.Errorf("無法建立牌組: %w", err)
		}
		cmd.Printf("✓ 牌組 '%s' 已就緒\n", modelDef.Deck)

		cmd.Println("\n初始化完成！您現在可以使用以下指令新增卡片:")
		cmd.Printf("./anki-japanese-cli add %s --deckName='%s' --json='{...}'\n", cardType, modelDef.Deck)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
