package anki

import (
	"fmt"
	"net/url"
	"strings"
	"time"
)

// ConnectionStatus represents the status of the connection to Anki
type ConnectionStatus struct {
	Connected      bool
	Version        string
	Error          string
	ResponseTimeMs int64
	URL            string
}

// CheckConnection checks if the connection to Anki is working
func (c *Client) CheckConnection() ConnectionStatus {
	status := ConnectionStatus{
		Connected: false,
		URL:       c.config.ConnectURL,
	}

	startTime := time.Now()
	result, err := c.Call("version", nil)
	endTime := time.Now()

	status.ResponseTimeMs = endTime.Sub(startTime).Milliseconds()

	if err != nil {
		status.Error = err.Error()
		return status
	}

	// Convert the result to a string
	version, ok := result.(float64)
	if !ok {
		status.Error = fmt.Sprintf("unexpected result type: %T", result)
		return status
	}

	status.Connected = true
	status.Version = fmt.Sprintf("%.1f", version)

	return status
}

// DiagnoseConnection provides diagnostic information about the connection
func (c *Client) DiagnoseConnection() string {
	status := c.CheckConnection()

	var sb strings.Builder

	sb.WriteString("Anki Connect 連線診斷\n")
	sb.WriteString("====================\n")
	sb.WriteString(fmt.Sprintf("連線 URL: %s\n", status.URL))
	sb.WriteString(fmt.Sprintf("連線狀態: %s\n", getStatusText(status.Connected)))
	sb.WriteString(fmt.Sprintf("回應時間: %d ms\n", status.ResponseTimeMs))

	if status.Connected {
		sb.WriteString(fmt.Sprintf("Anki Connect 版本: %s\n", status.Version))
	} else {
		sb.WriteString("\n錯誤診斷:\n")
		sb.WriteString(fmt.Sprintf("錯誤訊息: %s\n", status.Error))
		sb.WriteString("\n可能的解決方案:\n")
		sb.WriteString(getErrorSuggestions(status.Error, status.URL))
	}

	return sb.String()
}

// getStatusText returns a human-readable status text
func getStatusText(connected bool) string {
	if connected {
		return "已連線 ✓"
	}
	return "未連線 ✗"
}

// getErrorSuggestions returns suggestions for fixing the error
func getErrorSuggestions(errMsg, connectURL string) string {
	var sb strings.Builder

	// Check for common error patterns
	if strings.Contains(errMsg, "connection refused") {
		sb.WriteString("1. 確認 Anki 是否已啟動\n")
		sb.WriteString("2. 確認 AnkiConnect 插件是否已安裝\n")
		sb.WriteString("3. 重新啟動 Anki\n")
	} else if strings.Contains(errMsg, "no such host") {
		sb.WriteString("1. 檢查連線 URL 是否正確\n")
		sb.WriteString("2. 檢查網路連線\n")
	} else if strings.Contains(errMsg, "timeout") {
		sb.WriteString("1. Anki 可能正在處理其他請求，請稍後再試\n")
		sb.WriteString("2. 檢查 Anki 是否響應\n")
	} else if strings.Contains(errMsg, "API error") {
		sb.WriteString("1. 檢查 AnkiConnect 插件版本是否最新\n")
		sb.WriteString("2. 檢查請求參數是否正確\n")
	} else {
		sb.WriteString("1. 確認 Anki 是否已啟動\n")
		sb.WriteString("2. 確認 AnkiConnect 插件是否已安裝\n")
		sb.WriteString("3. 檢查連線 URL 是否正確\n")
	}

	// Add URL validation suggestion if the URL seems invalid
	_, err := url.Parse(connectURL)
	if err != nil {
		sb.WriteString("\n連線 URL 格式不正確，請檢查設定檔中的 anki.connect_url 值\n")
		sb.WriteString("預設值應為: http://localhost:8765\n")
	}

	// Add installation instructions for AnkiConnect
	sb.WriteString("\n安裝 AnkiConnect 插件:\n")
	sb.WriteString("1. 在 Anki 中，點擊「工具」>「附加元件」\n")
	sb.WriteString("2. 點擊「取得附加元件」\n")
	sb.WriteString("3. 輸入代碼: 2055492159\n")
	sb.WriteString("4. 重新啟動 Anki\n")

	return sb.String()
}

// FormatError formats an error message in a user-friendly way
func FormatError(err error) string {
	if err == nil {
		return ""
	}

	errMsg := err.Error()

	// Make the error message more user-friendly
	if strings.Contains(errMsg, "connection refused") {
		return "無法連線到 Anki。請確認 Anki 已啟動且 AnkiConnect 插件已安裝。"
	} else if strings.Contains(errMsg, "no such host") {
		return "找不到 Anki 伺服器。請檢查連線 URL 是否正確。"
	} else if strings.Contains(errMsg, "timeout") {
		return "連線到 Anki 逾時。請確認 Anki 是否正常運作。"
	} else if strings.Contains(errMsg, "API error") {
		// Extract the actual API error message
		parts := strings.Split(errMsg, "API error: ")
		if len(parts) > 1 {
			return fmt.Sprintf("Anki API 錯誤: %s", parts[1])
		}
		return "Anki API 發生錯誤。"
	}

	return fmt.Sprintf("錯誤: %s", errMsg)
}
