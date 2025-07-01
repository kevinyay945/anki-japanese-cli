package anki

import (
	"fmt"
	"reflect"
	"strings"
)

// Converter handles conversion between Go structs and Anki Connect JSON format
type Converter struct {
	client *Client
}

// NewConverter creates a new converter
func NewConverter(client *Client) *Converter {
	return &Converter{
		client: client,
	}
}

// ConvertToNote converts a model to an Anki note
func (c *Converter) ConvertToNote(model interface{}, deckName string, tags []string) (NoteInfo, error) {
	// Get the type and value of the model
	modelType := reflect.TypeOf(model)
	modelValue := reflect.ValueOf(model)

	// If the model is a pointer, get the element it points to
	if modelType.Kind() == reflect.Ptr {
		modelType = modelType.Elem()
		modelValue = modelValue.Elem()
	}

	// Check if the model is a struct
	if modelType.Kind() != reflect.Struct {
		return NoteInfo{}, fmt.Errorf("model must be a struct, got %s", modelType.Kind())
	}

	// Get the model name from the struct name
	modelName := modelType.Name()

	// Create a map for the fields
	fields := make(map[string]string)

	// Iterate over the struct fields
	for i := 0; i < modelType.NumField(); i++ {
		field := modelType.Field(i)
		fieldValue := modelValue.Field(i)

		// Skip unexported fields
		if !field.IsExported() {
			continue
		}

		// Get the field name from the json tag or use the field name
		fieldName := field.Name
		if jsonTag := field.Tag.Get("json"); jsonTag != "" {
			parts := strings.Split(jsonTag, ",")
			if parts[0] != "" && parts[0] != "-" {
				fieldName = parts[0]
			}
		}

		// Convert the field value to string
		var fieldStr string
		switch fieldValue.Kind() {
		case reflect.String:
			fieldStr = fieldValue.String()
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			fieldStr = fmt.Sprintf("%d", fieldValue.Int())
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			fieldStr = fmt.Sprintf("%d", fieldValue.Uint())
		case reflect.Float32, reflect.Float64:
			fieldStr = fmt.Sprintf("%f", fieldValue.Float())
		case reflect.Bool:
			fieldStr = fmt.Sprintf("%t", fieldValue.Bool())
		case reflect.Slice, reflect.Array:
			// Handle slices of strings
			if fieldValue.Type().Elem().Kind() == reflect.String {
				var strs []string
				for j := 0; j < fieldValue.Len(); j++ {
					strs = append(strs, fieldValue.Index(j).String())
				}
				fieldStr = strings.Join(strs, ", ")
			} else {
				fieldStr = fmt.Sprintf("%v", fieldValue.Interface())
			}
		default:
			fieldStr = fmt.Sprintf("%v", fieldValue.Interface())
		}

		fields[fieldName] = fieldStr
	}

	// Create the note
	note := NoteInfo{
		DeckName:  deckName,
		ModelName: modelName,
		Fields:    fields,
		Tags:      tags,
	}

	return note, nil
}

// ConvertToNotes converts a slice of models to Anki notes
func (c *Converter) ConvertToNotes(models interface{}, deckName string, tags []string) ([]NoteInfo, error) {
	// Get the type and value of the models
	modelsType := reflect.TypeOf(models)
	modelsValue := reflect.ValueOf(models)

	// If the models is a pointer, get the element it points to
	if modelsType.Kind() == reflect.Ptr {
		modelsType = modelsType.Elem()
		modelsValue = modelsValue.Elem()
	}

	// Check if the models is a slice
	if modelsType.Kind() != reflect.Slice {
		return nil, fmt.Errorf("models must be a slice, got %s", modelsType.Kind())
	}

	// Get the element type
	elemType := modelsType.Elem()

	// Check if the element is a struct or a pointer to a struct
	isPtr := false
	if elemType.Kind() == reflect.Ptr {
		elemType = elemType.Elem()
		isPtr = true
	}

	if elemType.Kind() != reflect.Struct {
		return nil, fmt.Errorf("models must be a slice of structs, got slice of %s", elemType.Kind())
	}

	// Create a slice for the notes
	var notes []NoteInfo

	// Iterate over the slice
	for i := 0; i < modelsValue.Len(); i++ {
		elemValue := modelsValue.Index(i)
		if isPtr {
			// Skip nil pointers
			if elemValue.IsNil() {
				continue
			}
			elemValue = elemValue.Elem()
		}

		// Convert the element to a note
		note, err := c.ConvertToNote(elemValue.Interface(), deckName, tags)
		if err != nil {
			return nil, err
		}

		notes = append(notes, note)
	}

	return notes, nil
}

// EnsureModelExists creates the model if it doesn't exist
func (c *Converter) EnsureModelExists(modelName string, fields []string, css string, templates []CardTemplateConfig) error {
	exists, err := c.client.ModelExists(modelName)
	if err != nil {
		return err
	}

	if !exists {
		config := ModelConfig{
			ModelName:     modelName,
			InOrderFields: fields,
			CSS:           css,
			CardTemplates: templates,
		}

		err = c.client.CreateModel(config)
		if err != nil {
			return err
		}
	}

	return nil
}

// CreateCardFromModel creates an Anki card from a model
func (c *Converter) CreateCardFromModel(model interface{}, deckName string, tags []string) (int64, error) {
	// Ensure the deck exists
	if err := c.client.EnsureDeckExists(deckName); err != nil {
		return 0, err
	}

	// Convert the model to a note
	note, err := c.ConvertToNote(model, deckName, tags)
	if err != nil {
		return 0, err
	}

	// Add the note
	return c.client.AddNote(note)
}

// CreateCardsFromModels creates Anki cards from models
func (c *Converter) CreateCardsFromModels(models interface{}, deckName string, tags []string) ([]int64, error) {
	// Ensure the deck exists
	if err := c.client.EnsureDeckExists(deckName); err != nil {
		return nil, err
	}

	// Convert the models to notes
	notes, err := c.ConvertToNotes(models, deckName, tags)
	if err != nil {
		return nil, err
	}

	// Add the notes
	return c.client.AddNotes(notes)
}
