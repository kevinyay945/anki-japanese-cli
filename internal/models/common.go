package models

import "fmt"

// CardType 卡片類型介面
type CardType interface {
	GetCardType() string
	Validate() error
}

// ValidationError 驗證錯誤類型
type ValidationError struct {
	Field   string
	Message string
}

// Error 實作錯誤介面
func (e *ValidationError) Error() string {
	return fmt.Sprintf("欄位 '%s' %s", e.Field, e.Message)
}

// NewValidationError 建立新的驗證錯誤
func NewValidationError(field, message string) *ValidationError {
	return &ValidationError{
		Field:   field,
		Message: message,
	}
}

// CardData 通用卡片資料介面
type CardData interface {
	ToMap() map[string]interface{}
	FromMap(data map[string]interface{}) error
}
