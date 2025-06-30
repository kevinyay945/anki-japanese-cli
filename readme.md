# Anki Japanese Cli

This cli can help me to create the card template,and help me to create the card by CLI

There are four types of card template can create.
Verb, Adjective, Normal Words and Grammar.

## How to use it

### create card type
```bash
./anki-japanese-cli init <card-type>
```

### add flashcard
```bash
cat anki-japanese-cli <card-type> --deckName='japanese-2025' --json='{
"xxx": "yyy"
}'
```

## Tech Stack

- Golang 1.23
- Cobra-cli https://github.com/spf13/cobra-cli
- Viper https://github.com/spf13/viper
- net/http
- anki-connection


## Card types

### Verb
#### 卡片前後樣式

**正面 (Front):**
```html
<div class="card-front">
  <div class="context-sentence">{{情境例句}}</div>
  <div class="meaning-hint">{{核心意義}}</div>
  {{#圖片提示}}<div class="image-hint"><img src="{{圖片提示}}"></div>{{/圖片提示}}
</div>
```

**背面 (Back):**
```html
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
```

#### 動詞卡片欄位：
- `核心單字`: 以辭書形（原形）記錄動詞
- `詞性分類`: 標明動詞類別（五段動詞、上一段/下一段動詞、不規則動詞）
- `核心意義`: 動詞最主要、最常用的中文意思
- `發音`: 記錄單字的假名發音
- `重音`: 用數字或高低線條標示重音（Pitch Accent）
- `常用變化`: 條列重要變化（ます形、て形、ない形、た形）
- `情境例句`: 包含此動詞的完整句子
- `例句翻譯`: 對應情境例句的中文翻譯
- `圖片提示`: 視覺化動詞意義的圖片（可選）

### Adjective
#### 卡片前後樣式

**正面 (Front):**
```html
<div class="card-front">
  <div class="context-sentence">{{情境例句}}</div>
  <div class="meaning-hint">{{核心意義}}</div>
</div>
```

**背面 (Back):**
```html
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
```

#### 形容詞卡片欄位：
- `核心單字`: 形容詞的辭書形（原形）
- `詞性分類`: 標明是「い形容詞」還是「な形容詞」
- `核心意義`: 形容詞最主要、最常用的中文意思
- `發音`: 記錄單字的假名發音
- `重音`: 用數字或高低線條標示重音
- `主要變化`: 修飾名詞、否定形、過去形等關鍵變化
- `情境例句`: 包含此形容詞的完整句子
- `例句翻譯`: 對應情境例句的中文翻譯
- `相關詞彙`: 反義詞或近義詞（可選）

### Normal Words
#### 卡片前後樣式

**正面 (Front):**
```html
<div class="card-front">
  <div class="context-sentence">{{情境例句}}</div>
  <div class="meaning-hint">{{核心意義}}</div>
  {{#圖片提示}}<div class="image-hint"><img src="{{圖片提示}}"></div>{{/圖片提示}}
</div>
```

**背面 (Back):**
```html
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
```

#### 其他單字卡片欄位：
- `核心單字`: 記錄單字本身
- `詞性`: 標明詞性（名詞、副詞、感嘆詞等）
- `核心意義`: 單字最主要、最常用的中文意思
- `發音`: 記錄單字的假名發音
- `重音`: 用數字或高低線條標示重音
- `情境例句`: 包含此單字的完整句子
- `例句翻譯`: 對應情境例句的中文翻譯
- `相關詞彙`: 同義、反義或相關概念詞彙（可選）
- `圖片提示`: 視覺化單字意義的圖片（可選）

### Grammar
#### 卡片前後樣式

**正面 (Front):**
```html
<div class="card-front">
  <div class="grammar-challenge">{{情境課題}}</div>
  <div class="grammar-point-hint">使用「{{文法點}}」</div>
</div>
```

**背面 (Back):**
```html
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
```

#### 文法卡片欄位：
- `文法點`: 要學習的句型或文法
- `核心意義`: 文法所表達的核心功能或意思
- `接續規則`: 詳細說明文法前面如何接續不同詞性
- `語感說明`: 解釋文法的語感、使用場合、正式程度
- `情境課題`: 中文句子或情境描述，要求使用此文法產出日文句子
- `解答範例`: 對應情境課題的日文解答
- `易混淆文法`: 與此文法相近易搞混的其他文法點（可選）

## CSS 樣式建議

```css
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
```

## Anki Connect API 範例

### 建立動詞卡片模型

```javascript
const verbModel = {
  "action": "createModel",
  "version": 6,
  "params": {
    "modelName": "Japanese Verb",
    "inOrderFields": [
      "核心單字",
      "詞性分類",
      "核心意義",
      "發音",
      "重音",
      "常用變化",
      "情境例句",
      "例句翻譯",
      "圖片提示"
    ],
    "css": "/* 在此處加入上方建議的 CSS 樣式 */",
    "cardTemplates": [
      {
        "Name": "動詞卡片",
        "Front": "<!-- 正面模板 -->",
        "Back": "<!-- 背面模板 -->"
      }
    ]
  }
}
```

### 新增動詞卡片

```javascript
const addVerbCard = {
  "action": "addNote",
  "version": 6,
  "params": {
    "note": {
      "deckName": "日本語動詞",
      "modelName": "Japanese Verb",
      "fields": {
        "核心單字": "飲む",
        "詞性分類": "五段動詞 (I)",
        "核心意義": "喝",
        "發音": "のむ",
        "重音": "1 (の① HIGH-low)",
        "常用變化": "ます形: 飲みます<br>て形: 飲んで<br>ない形: 飲まない<br>た形: 飲んだ",
        "情境例句": "寝る前に、温かい牛乳を＿＿＿習慣があります。",
        "例句翻譯": "我有睡前喝溫牛奶的習慣。",
        "圖片提示": ""
      },
      "tags": ["動詞", "五段動詞", "日常生活"]
    }
  }
}
```
