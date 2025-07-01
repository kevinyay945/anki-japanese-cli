package anki

import (
	"fmt"
)

// ModelConfig represents the configuration for creating a new note model
type ModelConfig struct {
	ModelName     string               `json:"modelName"`
	InOrderFields []string             `json:"inOrderFields"`
	CSS           string               `json:"css"`
	CardTemplates []CardTemplateConfig `json:"cardTemplates"`
	IsCloze       bool                 `json:"isCloze,omitempty"`
}

// CardTemplateConfig represents a card template configuration
type CardTemplateConfig struct {
	Name  string `json:"Name"`
	Front string `json:"Front"`
	Back  string `json:"Back"`
}

// CreateModel creates a new note model
func (c *Client) CreateModel(config ModelConfig) error {
	params := map[string]interface{}{
		"modelName":     config.ModelName,
		"inOrderFields": config.InOrderFields,
		"css":           config.CSS,
		"cardTemplates": config.CardTemplates,
	}

	if config.IsCloze {
		params["isCloze"] = true
	}

	_, err := c.Call("createModel", params)
	if err != nil {
		return fmt.Errorf("failed to create model: %w", err)
	}

	return nil
}

// ModelNames returns a list of all model names
func (c *Client) ModelNames() ([]string, error) {
	result, err := c.Call("modelNames", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get model names: %w", err)
	}

	// Convert the result to a string slice
	names, ok := result.([]interface{})
	if !ok {
		return nil, fmt.Errorf("unexpected result type: %T", result)
	}

	var modelNames []string
	for _, name := range names {
		strName, ok := name.(string)
		if !ok {
			return nil, fmt.Errorf("unexpected name type: %T", name)
		}
		modelNames = append(modelNames, strName)
	}

	return modelNames, nil
}

// ModelFieldNames returns a list of field names for the specified model
func (c *Client) ModelFieldNames(modelName string) ([]string, error) {
	params := map[string]interface{}{
		"modelName": modelName,
	}

	result, err := c.Call("modelFieldNames", params)
	if err != nil {
		return nil, fmt.Errorf("failed to get model field names: %w", err)
	}

	// Convert the result to a string slice
	fields, ok := result.([]interface{})
	if !ok {
		return nil, fmt.Errorf("unexpected result type: %T", result)
	}

	var fieldNames []string
	for _, field := range fields {
		strField, ok := field.(string)
		if !ok {
			return nil, fmt.Errorf("unexpected field type: %T", field)
		}
		fieldNames = append(fieldNames, strField)
	}

	return fieldNames, nil
}

// UpdateModelTemplates updates the templates for the specified model
func (c *Client) UpdateModelTemplates(modelName string, templates map[string]map[string]string) error {
	params := map[string]interface{}{
		"model": map[string]interface{}{
			"name":      modelName,
			"templates": templates,
		},
	}

	_, err := c.Call("updateModelTemplates", params)
	if err != nil {
		return fmt.Errorf("failed to update model templates: %w", err)
	}

	return nil
}

// ModelExists checks if a model with the given name exists
func (c *Client) ModelExists(modelName string) (bool, error) {
	names, err := c.ModelNames()
	if err != nil {
		return false, err
	}

	for _, name := range names {
		if name == modelName {
			return true, nil
		}
	}

	return false, nil
}
