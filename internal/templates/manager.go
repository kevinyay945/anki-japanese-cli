package templates

import (
	"bytes"
	"embed"
	"fmt"
	"html/template"
	"path/filepath"
	"strings"
)

//go:embed *.html
var templateFS embed.FS

// CardTemplate 卡片模板結構
type CardTemplate struct {
	Front *template.Template
	Back  *template.Template
}

// TemplateManager 模板管理器
type TemplateManager struct {
	templates map[string]*CardTemplate
}

// NewTemplateManager 建立新的模板管理器
func NewTemplateManager() (*TemplateManager, error) {
	tm := &TemplateManager{
		templates: make(map[string]*CardTemplate),
	}

	err := tm.loadTemplates()
	if err != nil {
		return nil, fmt.Errorf("載入模板失敗: %w", err)
	}

	return tm, nil
}

// loadTemplates 載入所有模板檔案
func (tm *TemplateManager) loadTemplates() error {
	templateFiles := []string{
		"verb_front.html", "verb_back.html",
		"adjective_front.html", "adjective_back.html",
		"normal_front.html", "normal_back.html",
		"grammar_front.html", "grammar_back.html",
	}

	// 按卡片類型分組
	cardTypes := make(map[string]*CardTemplate)

	for _, file := range templateFiles {
		content, err := templateFS.ReadFile(file)
		if err != nil {
			return fmt.Errorf("讀取模板檔案 %s 失敗: %w", file, err)
		}

		cardType, side := parseFilename(file)
		if cardTypes[cardType] == nil {
			cardTypes[cardType] = &CardTemplate{}
		}

		tmpl, err := template.New(file).Parse(string(content))
		if err != nil {
			return fmt.Errorf("解析模板 %s 失敗: %w", file, err)
		}

		if side == "front" {
			cardTypes[cardType].Front = tmpl
		} else {
			cardTypes[cardType].Back = tmpl
		}
	}

	tm.templates = cardTypes
	return nil
}

// parseFilename 解析檔案名稱獲取卡片類型和面
func parseFilename(filename string) (cardType, side string) {
	name := filepath.Base(filename)
	ext := filepath.Ext(name)
	nameWithoutExt := name[:len(name)-len(ext)]

	parts := strings.Split(nameWithoutExt, "_")
	if len(parts) >= 2 {
		cardType = parts[0]
		side = parts[1]
	}
	return
}

// RenderCardFront 渲染卡片正面
func (tm *TemplateManager) RenderCardFront(cardType string, data interface{}) (string, error) {
	cardTemplate, exists := tm.templates[cardType]
	if !exists || cardTemplate.Front == nil {
		return "", fmt.Errorf("找不到卡片類型 '%s' 的正面模板", cardType)
	}

	var buf bytes.Buffer
	err := cardTemplate.Front.Execute(&buf, data)
	if err != nil {
		return "", fmt.Errorf("渲染正面模板失敗: %w", err)
	}

	return buf.String(), nil
}

// RenderCardBack 渲染卡片背面
func (tm *TemplateManager) RenderCardBack(cardType string, data interface{}) (string, error) {
	cardTemplate, exists := tm.templates[cardType]
	if !exists || cardTemplate.Back == nil {
		return "", fmt.Errorf("找不到卡片類型 '%s' 的背面模板", cardType)
	}

	var buf bytes.Buffer
	err := cardTemplate.Back.Execute(&buf, data)
	if err != nil {
		return "", fmt.Errorf("渲染背面模板失敗: %w", err)
	}

	return buf.String(), nil
}

// RenderCard 渲染完整卡片 (舊方法，向後相容)
func (tm *TemplateManager) RenderCard(cardType string, data interface{}) (string, error) {
	return tm.RenderCardBack(cardType, data)
}

// ValidateTemplate 驗證模板是否有效
func (tm *TemplateManager) ValidateTemplate(cardType string) error {
	cardTemplate, exists := tm.templates[cardType]
	if !exists {
		return fmt.Errorf("模板 '%s' 不存在", cardType)
	}
	if cardTemplate.Front == nil {
		return fmt.Errorf("模板 '%s' 缺少正面模板", cardType)
	}
	if cardTemplate.Back == nil {
		return fmt.Errorf("模板 '%s' 缺少背面模板", cardType)
	}
	return nil
}

// GetAvailableTemplates 獲取可用的模板列表
func (tm *TemplateManager) GetAvailableTemplates() []string {
	var templates []string
	for name := range tm.templates {
		templates = append(templates, name)
	}
	return templates
}

// ReloadTemplates 重新載入模板
func (tm *TemplateManager) ReloadTemplates() error {
	tm.templates = make(map[string]*CardTemplate)
	return tm.loadTemplates()
}

// GetRawTemplate 獲取原始模板內容
func (tm *TemplateManager) GetRawTemplate(cardType string, side string) (string, error) {
	filename := fmt.Sprintf("%s_%s.html", cardType, side)
	content, err := templateFS.ReadFile(filename)
	if err != nil {
		return "", fmt.Errorf("讀取模板檔案 %s 失敗: %w", filename, err)
	}
	return string(content), nil
}
