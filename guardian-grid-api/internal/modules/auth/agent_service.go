package auth

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"guardian-grid-api/internal/modules/alerts"
	"guardian-grid-api/internal/modules/cve"
	"guardian-grid-api/internal/modules/websocket"

	"github.com/google/uuid"
)

type AService struct {
	repo         *ARepository
	alertService *alerts.Service
	cveService   *cve.CVEService
	hub          *websocket.Hub
}

func NewAgentService(repo *ARepository, alertService *alerts.Service, cveService *cve.CVEService, hub *websocket.Hub) *AService {
	return &AService{
		repo:         repo,
		alertService: alertService,
		cveService:   cveService,
		hub:          hub,
	}
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

				// If telemetry is 'programs', perform CVE scan
				if teleType == "programs" {
					s.performCVEScan(agentID, item)
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

			if teleType == "programs" {
				s.performCVEScan(agentID, data)
			}
		}
	}
	return latestData, nil
}

func (s *AService) performCVEScan(agentID string, data interface{}) {
	progData, ok := data.(map[string]interface{})
	if !ok {
		return
	}

	programs, ok := progData["programs"].([]interface{})
	if !ok {
		return
	}

	for _, p := range programs {
		prog, ok := p.(map[string]interface{})
		if !ok {
			continue
		}

		name, _ := prog["name"].(string)
		version, _ := prog["version"].(string)

		if name == "" || version == "" {
			continue
		}

		findings := s.cveService.Scan(name, version)
		for _, v := range findings {
			alert := alerts.Alert{
				AgentID:     agentID,
				ProgramName: name,
				Version:     version,
				CVEID:       v.CVEID,
				Severity:    v.Severity,
				Description: v.Description,
				Score:       v.Score,
			}

			err := s.alertService.CreateAlert(alert)
			if err == nil {
				fmt.Printf("⚠️ SECURITY ALERT: %s found in %s v%s\n", v.CVEID, name, version)
				if s.hub != nil {
					s.hub.BroadcastAlert(agentID, alert)
				}
			}
		}
	}
}
