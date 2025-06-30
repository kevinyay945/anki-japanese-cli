package main

import (
	"anki-japanese-cli/internal/models"
	"fmt"
	"log"
)

func main() {
	// 建立卡片服務
	cardService, err := models.NewCardService()
	if err != nil {
		log.Fatalf("初始化卡片服務失敗: %v", err)
	}

	// 範例1: 動詞卡片 - 正面和背面
	fmt.Println("=== 動詞卡片範例 ===")
	verbData := map[string]interface{}{
		"核心單字": "食べる",
		"詞性分類": "動詞 (一段動詞)",
		"核心意義": "吃",
		"發音":   "たべる",
		"重音":   "②",
		"常用變化": "食べます (敬語形)\n食べた (過去形)\n食べない (否定形)\n食べて (て形)",
		"情境例句": "昨日、美味しいラーメンを食べました。",
		"例句翻譯": "昨天吃了美味的拉麵。",
		"圖片提示": "https://example.com/ramen.jpg",
	}

	frontHTML, err := cardService.CreateAndRenderCardFront("verb", verbData)
	if err != nil {
		log.Fatalf("建立動詞卡片正面失敗: %v", err)
	}

	backHTML, err := cardService.CreateAndRenderCardBack("verb", verbData)
	if err != nil {
		log.Fatalf("建立動詞卡片背面失敗: %v", err)
	}

	fmt.Println("--- 正面 ---")
	fmt.Println(frontHTML)
	fmt.Println("\n--- 背面 ---")
	fmt.Println(backHTML)
	fmt.Println()

	// 範例2: 形容詞卡片
	fmt.Println("=== 形容詞卡片範例 ===")
	adjectiveData := map[string]interface{}{
		"核心單字": "美しい",
		"詞性分類": "い形容詞",
		"核心意義": "美麗的",
		"發音":   "うつくしい",
		"重音":   "④",
		"主要變化": "美しく (副詞形)\n美しくない (否定形)\n美しかった (過去形)",
		"情境例句": "桜の花がとても美しいです。",
		"例句翻譯": "櫻花非常美麗。",
		"相關詞彙": "綺麗 (きれい), 素敵 (すてき)",
	}

	frontHTML, err = cardService.CreateAndRenderCardFront("adjective", adjectiveData)
	if err != nil {
		log.Fatalf("建立形容詞卡片正面失敗: %v", err)
	}

	backHTML, err = cardService.CreateAndRenderCardBack("adjective", adjectiveData)
	if err != nil {
		log.Fatalf("建立形容詞卡片背面失敗: %v", err)
	}

	fmt.Println("--- 正面 ---")
	fmt.Println(frontHTML)
	fmt.Println("\n--- 背面 ---")
	fmt.Println(backHTML)
	fmt.Println()

	// 範例3: 文法卡片 - 按照 README 的格式
	fmt.Println("=== 文法卡片範例 ===")
	grammarData := map[string]interface{}{
		"文法要點": "～ている",
		"結構形式": "動詞て形 + いる",
		"意義說明": "表示動作的進行或狀態的持續",
		"使用時機": "描述正在進行的動作或持續的狀態",
		"例句示範": "今、本を読んでいます。\n彼は結婚しています。",
		"例句翻譯": "現在正在讀書。\n他已經結婚了。",
		"情境課題": "請用「～ている」來表達「我現在正在學習日語」",
		"解答範例": "今、日本語を勉強しています。",
		"難度等級": "N4",
		"相關文法": "～ていた (過去進行), ～ていく/てくる (方向性)",
		"常見錯誤": "忘記動詞變て形, 混淆進行和狀態的用法",
		"記憶技巧": "想像動作「正在持續」的畫面",
	}

	frontHTML, err = cardService.CreateAndRenderCardFront("grammar", grammarData)
	if err != nil {
		log.Fatalf("建立文法卡片正面失敗: %v", err)
	}

	backHTML, err = cardService.CreateAndRenderCardBack("grammar", grammarData)
	if err != nil {
		log.Fatalf("建立文法卡片背面失敗: %v", err)
	}

	fmt.Println("--- 正面 ---")
	fmt.Println(frontHTML)
	fmt.Println("\n--- 背面 ---")
	fmt.Println(backHTML)
	fmt.Println()

	// 範例4: 一般單字卡片
	fmt.Println("=== 一般單字卡片範例 ===")
	normalData := map[string]interface{}{
		"核心單字": "学校",
		"詞性分類": "名詞",
		"核心意義": "學校",
		"發音":   "がっこう",
		"重音":   "①",
		"使用方式": "教育機構的通稱",
		"情境例句": "毎日学校に行きます。",
		"例句翻譯": "每天去學校。",
		"同義詞":  "学園 (がくえん)",
		"圖片提示": "https://example.com/school.jpg",
	}

	frontHTML, err = cardService.CreateAndRenderCardFront("normal", normalData)
	if err != nil {
		log.Fatalf("建立一般單字卡片正面失敗: %v", err)
	}

	backHTML, err = cardService.CreateAndRenderCardBack("normal", normalData)
	if err != nil {
		log.Fatalf("建立一般單字卡片背面失敗: %v", err)
	}

	fmt.Println("--- 正面 ---")
	fmt.Println(frontHTML)
	fmt.Println("\n--- 背面 ---")
	fmt.Println(backHTML)
	fmt.Println()

	// 顯示支援的卡片類型
	fmt.Println("=== 支援的卡片類型 ===")
	types := cardService.GetSupportedCardTypes()
	for _, t := range types {
		fmt.Printf("- %s\n", t)
	}

	// 顯示可用的模板
	fmt.Println("\n=== 可用的模板 ===")
	templates := cardService.GetAvailableTemplates()
	for _, t := range templates {
		fmt.Printf("- %s (正面與背面)\n", t)
	}
}
