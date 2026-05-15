package dashboard

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetGlobalStatus() ([]map[string]interface{}, error) {
	return s.repo.GetAllAgentsLatest()
}

func (s *Service) GetAgentDetails(agentID string) (map[string]interface{}, error) {
	return s.repo.GetLatestTelemetry(agentID)
}
