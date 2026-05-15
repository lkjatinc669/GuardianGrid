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

func (s *AService) ProcessTelemetry(agentID string, payload map[string]interface{}) error {
	for teleType, data := range payload {
		err := s.repo.SaveTelemetry(agentID, teleType, data)
		if err != nil {
			return err
		}
	}
	return nil
}
