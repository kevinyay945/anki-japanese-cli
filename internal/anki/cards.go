package anki

import (
	"fmt"
)

// NoteInfo represents the information for a note to be added
type NoteInfo struct {
	DeckName  string                 `json:"deckName"`
	ModelName string                 `json:"modelName"`
	Fields    map[string]string      `json:"fields"`
	Tags      []string               `json:"tags,omitempty"`
	Options   map[string]interface{} `json:"options,omitempty"`
}

// AddNote adds a single note to Anki
func (c *Client) AddNote(note NoteInfo) (int64, error) {
	// Create the note map without options first
	noteMap := map[string]interface{}{
		"deckName":  note.DeckName,
		"modelName": note.ModelName,
		"fields":    note.Fields,
	}

	// Add tags if present
	if note.Tags != nil && len(note.Tags) > 0 {
		noteMap["tags"] = note.Tags
	}

	// Add options if present
	if note.Options != nil && len(note.Options) > 0 {
		noteMap["options"] = note.Options
	}

	params := map[string]interface{}{
		"note": noteMap,
	}

	result, err := c.Call("addNote", params)
	if err != nil {
		return 0, fmt.Errorf("failed to add note: %w", err)
	}

	// Convert the result to an int64 (note ID)
	noteID, ok := result.(float64)
	if !ok {
		return 0, fmt.Errorf("unexpected result type: %T", result)
	}

	return int64(noteID), nil
}

// AddNotes adds multiple notes to Anki
func (c *Client) AddNotes(notes []NoteInfo) ([]int64, error) {
	// Convert notes to the format expected by the API
	apiNotes := make([]map[string]interface{}, len(notes))
	for i, note := range notes {
		// Create the note map without options first
		noteMap := map[string]interface{}{
			"deckName":  note.DeckName,
			"modelName": note.ModelName,
			"fields":    note.Fields,
		}

		// Add tags if present
		if note.Tags != nil && len(note.Tags) > 0 {
			noteMap["tags"] = note.Tags
		}

		// Add options if present
		if note.Options != nil && len(note.Options) > 0 {
			noteMap["options"] = note.Options
		}

		apiNotes[i] = noteMap
	}

	params := map[string]interface{}{
		"notes": apiNotes,
	}

	result, err := c.Call("addNotes", params)
	if err != nil {
		return nil, fmt.Errorf("failed to add notes: %w", err)
	}

	// Convert the result to an int64 slice (note IDs)
	noteIDs, ok := result.([]interface{})
	if !ok {
		return nil, fmt.Errorf("unexpected result type: %T", result)
	}

	var ids []int64
	for _, id := range noteIDs {
		// Note IDs can be null if the note wasn't added
		if id == nil {
			ids = append(ids, 0)
			continue
		}

		floatID, ok := id.(float64)
		if !ok {
			return nil, fmt.Errorf("unexpected ID type: %T", id)
		}
		ids = append(ids, int64(floatID))
	}

	return ids, nil
}

// DeckNames returns a list of all deck names
func (c *Client) DeckNames() ([]string, error) {
	result, err := c.Call("deckNames", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get deck names: %w", err)
	}

	// Convert the result to a string slice
	decks, ok := result.([]interface{})
	if !ok {
		return nil, fmt.Errorf("unexpected result type: %T", result)
	}

	var deckNames []string
	for _, deck := range decks {
		strDeck, ok := deck.(string)
		if !ok {
			return nil, fmt.Errorf("unexpected deck type: %T", deck)
		}
		deckNames = append(deckNames, strDeck)
	}

	return deckNames, nil
}

// CreateDeck creates a new deck
func (c *Client) CreateDeck(deckName string) (int64, error) {
	params := map[string]interface{}{
		"deck": deckName,
	}

	result, err := c.Call("createDeck", params)
	if err != nil {
		return 0, fmt.Errorf("failed to create deck: %w", err)
	}

	// Convert the result to an int64 (deck ID)
	deckID, ok := result.(float64)
	if !ok {
		return 0, fmt.Errorf("unexpected result type: %T", result)
	}

	return int64(deckID), nil
}

// DeckExists checks if a deck with the given name exists
func (c *Client) DeckExists(deckName string) (bool, error) {
	decks, err := c.DeckNames()
	if err != nil {
		return false, err
	}

	for _, deck := range decks {
		if deck == deckName {
			return true, nil
		}
	}

	return false, nil
}

// EnsureDeckExists creates the deck if it doesn't exist
func (c *Client) EnsureDeckExists(deckName string) error {
	exists, err := c.DeckExists(deckName)
	if err != nil {
		return err
	}

	if !exists {
		_, err = c.CreateDeck(deckName)
		if err != nil {
			return err
		}
	}

	return nil
}
