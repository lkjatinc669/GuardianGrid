package auth

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"
)

type AuthConfig struct {
	AgentID string `json:"agent_id"`
	Token   string `json:"token"`
}

const configPath = "agent_auth.json"

func GetConfig() (*AuthConfig, error) {
	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	var config AuthConfig
	err = json.Unmarshal(data, &config)
	return &config, err
}

func Register(baseURL string) (*AuthConfig, error) {
	hostname, _ := os.Hostname()

	payload := map[string]string{
		"hostname": hostname,
	}
	jsonData, _ := json.Marshal(payload)

	resp, err := http.Post(baseURL+"/agent/register", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result struct {
		Status string `json:"status"`
		Data   struct {
			AgentID string `json:"agent_id"`
			Token   string `json:"token"`
		} `json:"data"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	config := &AuthConfig{
		AgentID: result.Data.AgentID,
		Token:   result.Data.Token,
	}

	// Save to file
	configData, _ := json.Marshal(config)
	_ = os.WriteFile(configPath, configData, 0644)

	return config, nil
}
