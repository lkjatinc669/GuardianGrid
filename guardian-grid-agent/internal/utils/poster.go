package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func SendData(url string, payload interface{}) (string, error) {
	// Convert payload to JSON
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return "", fmt.Errorf("failed to marshal json: %w", err)
	}

	// Create HTTP client with timeout
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	// Create request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")

	// Send request
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	// Read response
	var responseBody bytes.Buffer
	_, err = responseBody.ReadFrom(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response: %w", err)
	}

	return responseBody.String(), nil
}
