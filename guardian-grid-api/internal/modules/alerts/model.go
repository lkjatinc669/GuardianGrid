package alerts

import (
	"time"
)

type Alert struct {
	ID          string    `json:"id" bson:"_id,omitempty"`
	AgentID     string    `json:"agent_id" bson:"agent_id"`
	ProgramName string    `json:"program_name" bson:"program_name"`
	Version     string    `json:"version" bson:"version"`
	CVEID       string    `json:"cve_id" bson:"cve_id"`
	Severity    string    `json:"severity" bson:"severity"`
	Description string    `json:"description" bson:"description"`
	Score       float64   `json:"score" bson:"score"`
	CreatedAt   time.Time `json:"created_at" bson:"created_at"`
}
