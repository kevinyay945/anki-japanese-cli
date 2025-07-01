# Anki Connect API Integration Guide

This document explains how the Anki Japanese CLI tool integrates with Anki through the AnkiConnect API.

## What is AnkiConnect?

[AnkiConnect](https://foosoft.net/projects/anki-connect/) is an Anki add-on that provides a REST API for interacting with Anki. It allows external applications to perform operations such as adding notes, creating decks, and querying the Anki database.

## Installation

1. Open Anki
2. Go to Tools > Add-ons > Get Add-ons...
3. Enter the code: `2055492159`
4. Restart Anki

## API Basics

AnkiConnect exposes a REST API on `http://localhost:8765` by default. All requests are POST requests with a JSON payload.

The basic structure of a request is:

```json
{
  "action": "actionName",
  "version": 6,
  "params": {
    "param1": "value1",
    "param2": "value2"
  }
}
```

The response structure is:

```json
{
  "result": "some result value",
  "error": null
}
```

## Key API Actions Used by Anki Japanese CLI

### 1. Testing Connection

```json
{
  "action": "version",
  "version": 6
}
```

This is used to check if AnkiConnect is available and running.

### 2. Creating Decks

```json
{
  "action": "createDeck",
  "version": 6,
  "params": {
    "deck": "deckName"
  }
}
```

This creates a new deck if it doesn't already exist.

### 3. Creating Models (Note Types)

```json
{
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
    "css": "/* CSS styling */",
    "cardTemplates": [
      {
        "Name": "Card 1",
        "Front": "<!-- Front template HTML -->",
        "Back": "<!-- Back template HTML -->"
      }
    ]
  }
}
```

This creates a new note type with the specified fields and templates.

### 4. Adding Notes

```json
{
  "action": "addNote",
  "version": 6,
  "params": {
    "note": {
      "deckName": "Japanese Verbs",
      "modelName": "Japanese Verb",
      "fields": {
        "核心單字": "飲む",
        "詞性分類": "五段動詞",
        "核心意義": "喝",
        "發音": "のむ",
        "重音": "1",
        "常用變化": "飲みます、飲んで",
        "情境例句": "水を飲む",
        "例句翻譯": "喝水"
      },
      "tags": ["verb", "godan"]
    }
  }
}
```

This adds a new note to the specified deck using the specified note type.

### 5. Adding Multiple Notes

```json
{
  "action": "addNotes",
  "version": 6,
  "params": {
    "notes": [
      {
        "deckName": "Japanese Verbs",
        "modelName": "Japanese Verb",
        "fields": {
          "核心單字": "飲む",
          "詞性分類": "五段動詞",
          "核心意義": "喝",
          "發音": "のむ",
          "重音": "1",
          "常用變化": "飲みます、飲んで",
          "情境例句": "水を飲む",
          "例句翻譯": "喝水"
        },
        "tags": ["verb", "godan"]
      },
      {
        "deckName": "Japanese Verbs",
        "modelName": "Japanese Verb",
        "fields": {
          "核心單字": "食べる",
          "詞性分類": "一段動詞",
          "核心意義": "吃",
          "發音": "たべる",
          "重音": "2",
          "常用變化": "食べます、食べて",
          "情境例句": "ご飯を食べる",
          "例句翻譯": "吃飯"
        },
        "tags": ["verb", "ichidan"]
      }
    ]
  }
}
```

This adds multiple notes in a single request.

## Implementation in Anki Japanese CLI

The Anki Japanese CLI tool implements these API calls in the `internal/anki/client.go` file. The main client struct is:

```
// Example of the Client struct in internal/anki/client.go
// Client represents an Anki Connect API client
type Client struct {
    config     *config.AnkiConfig
    httpClient HTTPClient
    retries    int
    retryDelay time.Duration
}
```

Key methods include:

- `Ping()`: Tests the connection to AnkiConnect
- `CreateDeck(deckName string)`: Creates a new deck
- `ModelExists(modelName string)`: Checks if a model exists
- `CreateModel(model ModelConfig)`: Creates a new model
- `AddNote(note NoteInfo)`: Adds a single note
- `AddNotes(notes []NoteInfo)`: Adds multiple notes

## Error Handling

The client implements retry logic for failed requests:

```
// Example of retry logic in internal/anki/client.go
// Retry logic
for attempt := 0; attempt <= c.retries; attempt++ {
    if attempt > 0 {
        // Wait before retrying
        time.Sleep(c.retryDelay)
    }

    result, lastErr = c.doRequest(req)
    if lastErr == nil {
        return result, nil
    }
}
```

By default, it will retry 3 times with a 1-second delay between retries.

## Configuration

The AnkiConnect client is configured through the `config.AnkiConfig` struct:

```
// Example of the AnkiConfig struct in internal/config/config.go
type AnkiConfig struct {
    ConnectURL string `mapstructure:"connect_url"`
    DeckName   string `mapstructure:"deck_name"`
}
```

The default configuration uses:
- ConnectURL: "http://localhost:8765"
- DeckName: "日文學習" (Japanese Learning)

## Testing

For testing purposes, the CLI includes a mock client that can be used to simulate AnkiConnect responses without actually connecting to Anki. This is implemented in `internal/anki/mock_client.go`.

## References

- [AnkiConnect GitHub Repository](https://github.com/FooSoft/anki-connect)
- [AnkiConnect API Documentation](https://foosoft.net/projects/anki-connect/)
- [Anki Manual](https://docs.ankiweb.net/)
