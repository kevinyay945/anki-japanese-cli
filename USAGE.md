# Anki Japanese CLI Usage Guide

This document provides detailed instructions on how to use the Anki Japanese CLI tool for creating and managing Japanese flashcards in Anki.

## Prerequisites

Before using this tool, make sure you have:

1. [Anki](https://apps.ankiweb.net/) installed on your computer
2. [AnkiConnect](https://ankiweb.net/shared/info/2055492159) plugin installed in Anki
3. Anki running in the background while using this CLI tool

## Installation

```bash
# Clone the repository
git clone https://github.com/yourusername/anki-japanese-cli.git

# Navigate to the project directory
cd anki-japanese-cli

# Build the project
go build -o anki-japanese-cli

# Make it executable (Linux/macOS)
chmod +x anki-japanese-cli
```

## Basic Commands

### Initialize Card Types

Before adding cards, you need to initialize the card types in Anki:

```bash
./anki-japanese-cli init <card-type>
```

Where `<card-type>` can be one of:
- `verb` - For Japanese verbs
- `adjective` - For Japanese adjectives
- `normal` - For general Japanese vocabulary
- `grammar` - For Japanese grammar points

This command will:
1. Create the necessary card model in Anki
2. Create a default deck for the card type
3. Set up the card templates with proper styling

Example:
```bash
./anki-japanese-cli init verb
```

### Add Cards

To add a new card:

```bash
./anki-japanese-cli add <card-type> --deckName='<deck-name>' --json='<json-data>'
```

Parameters:
- `<card-type>`: The type of card (verb, adjective, normal, grammar)
- `--deckName`: The name of the Anki deck to add the card to
- `--json`: JSON string containing the card data

Example:
```bash
./anki-japanese-cli add verb --deckName='Japanese Verbs' --json='{"核心單字":"飲む", "詞性分類":"五段動詞", "核心意義":"喝", "發音":"のむ", "重音":"1", "常用變化":"飲みます、飲んで", "情境例句":"水を飲む", "例句翻譯":"喝水"}'
```

### Add Cards from File

You can also add cards from a JSON file:

```bash
./anki-japanese-cli add <card-type> --deckName='<deck-name>' --file='<file-path>'
```

Parameters:
- `<card-type>`: The type of card (verb, adjective, normal, grammar)
- `--deckName`: The name of the Anki deck to add the card to
- `--file`: Path to a JSON file containing the card data

Example:
```bash
./anki-japanese-cli add verb --deckName='Japanese Verbs' --file='examples/verb_cards.json'
```

### Batch Import

For importing multiple cards at once:

```bash
./anki-japanese-cli add <card-type> --deckName='<deck-name>' --file='<file-path>' --batch
```

The `--batch` flag indicates that the JSON file contains an array of card data.

Example:
```bash
./anki-japanese-cli add verb --deckName='Japanese Verbs' --file='examples/batch_import.json' --batch
```

## Card Type Details

### Verb Cards

Required fields:
- `核心單字`: The verb in dictionary form
- `詞性分類`: Verb type (五段動詞, 一段動詞, etc.)
- `核心意義`: Core meaning in Chinese
- `發音`: Pronunciation in hiragana
- `情境例句`: Example sentence
- `例句翻譯`: Translation of the example sentence

Optional fields:
- `重音`: Pitch accent
- `常用變化`: Common conjugations
- `圖片提示`: URL to an image

### Adjective Cards

Required fields:
- `核心單字`: The adjective in dictionary form
- `詞性分類`: Adjective type (い形容詞, な形容詞)
- `核心意義`: Core meaning in Chinese
- `發音`: Pronunciation in hiragana
- `情境例句`: Example sentence
- `例句翻譯`: Translation of the example sentence

Optional fields:
- `重音`: Pitch accent
- `主要變化`: Main conjugations
- `相關詞彙`: Related words

### Normal Word Cards

Required fields:
- `核心單字`: The word
- `詞性`: Part of speech
- `核心意義`: Core meaning in Chinese
- `發音`: Pronunciation in hiragana
- `情境例句`: Example sentence
- `例句翻譯`: Translation of the example sentence

Optional fields:
- `重音`: Pitch accent
- `相關詞彙`: Related words
- `圖片提示`: URL to an image

### Grammar Cards

Required fields:
- `文法要點`: Grammar point
- `結構形式`: Structure
- `意義說明`: Meaning explanation
- `例句示範`: Example sentence
- `例句翻譯`: Translation of the example sentence

Optional fields:
- `使用時機`: Usage context
- `情境課題`: Challenge scenario
- `解答範例`: Example answer
- `相關文法`: Related grammar points

## Examples

See the `examples` directory for sample JSON files for each card type.