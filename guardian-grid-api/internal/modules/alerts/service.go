package alerts

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateAlert(alert Alert) error {
	exists, err := s.repo.AlertExists(alert.AgentID, alert.CVEID, alert.ProgramName, alert.Version)
	if err != nil {
		return err
	}
	if exists {
		return nil // Avoid duplicate alerts
	}
	return s.repo.SaveAlert(alert)
}

func (s *Service) GetAgentAlerts(agentID string) ([]Alert, error) {
	return s.repo.GetAlertsByAgent(agentID)
}

func (s *Service) GetAllAlerts() ([]Alert, error) {
	return s.repo.GetAllAlerts()
}
