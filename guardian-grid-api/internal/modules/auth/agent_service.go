package auth

import (
	"crypto/rand"
	"encoding/hex"

	"github.com/google/uuid"
)

type AService struct {
	repo *ARepository
}

func NewAgentService(repo *ARepository) *AService {
	return &AService{repo: repo}
}

func generateAgentToken() string {
	b := make([]byte, 32)
	rand.Read(b)
	return hex.EncodeToString(b)
}

func (s *AService) RegisterAgent(hostname string) (string, string, error) {
	id := uuid.New().String()
	token := generateAgentToken()

	err := s.repo.CreateAgent(id, hostname, token)
	if err != nil {
		return "", "", err
	}

	return id, token, nil
}

func (s *AService) ProcessTelemetry(agentID string, payload map[string]interface{}) (map[string]interface{}, error) {
	latestData := make(map[string]interface{})

	for teleType, data := range payload {
		// The agent sends data in batches (slices of maps)
		if dataList, ok := data.([]interface{}); ok && len(dataList) > 0 {
			// Save each telemetry point individually
			for _, item := range dataList {
				err := s.repo.SaveTelemetry(agentID, teleType, item)
				if err != nil {
					return nil, err
				}
			}
			// Keep the last (latest) one for broadcasting to the dashboard
			latestData[teleType] = dataList[len(dataList)-1]
		} else {
			// Fallback if it's a single item
			err := s.repo.SaveTelemetry(agentID, teleType, data)
			if err != nil {
				return nil, err
			}
			latestData[teleType] = data
		}
	}
	return latestData, nil
}
