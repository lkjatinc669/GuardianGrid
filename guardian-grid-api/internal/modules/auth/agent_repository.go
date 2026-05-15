package auth

import (
	"context"
	"database/sql"
	"time"

	"guardian-grid-api/internal/platform/database"
	"go.mongodb.org/mongo-driver/mongo"
)

type ARepository struct {
	sqlite *sql.DB
	mongo  *mongo.Database
}

func NewAgentRepository() *ARepository {
	return &ARepository{
		sqlite: database.GetSQLite(),
		mongo:  database.GetMongoDatabase(),
	}
}

// Metadata Operations (SQLite)
func (r *ARepository) CreateAgent(id, hostname, token string) error {
	_, err := r.sqlite.Exec(
		"INSERT INTO agents (id, hostname, token) VALUES (?, ?, ?)",
		id, hostname, token,
	)
	return err
}

func (r *ARepository) GetAgentByToken(token string) (string, error) {
	var id string
	err := r.sqlite.QueryRow(
		"SELECT id FROM agents WHERE token = ?",
		token,
	).Scan(&id)

	return id, err
}

// Telemetry Operations (MongoDB - Multi-collection)
func (r *ARepository) SaveTelemetry(agentID, teleType string, data interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Map teleType to specific collection names
	collectionName := "telemetry_misc"
	switch teleType {
	case "dnscache":
		collectionName = "telemetry_dns"
	case "processes":
		collectionName = "telemetry_processes"
	case "programs":
		collectionName = "telemetry_programs"
	case "openports":
		collectionName = "telemetry_ports"
	case "network":
		collectionName = "telemetry_network"
	case "persistence":
		collectionName = "telemetry_persistence"
	case "liveactivity":
		collectionName = "telemetry_activity"
	case "activeusers":
		collectionName = "telemetry_users"
	case "pcdata":
		collectionName = "telemetry_host_info"
	case "uptime":
		collectionName = "telemetry_uptime"
	}

	collection := r.mongo.Collection(collectionName)

	document := map[string]interface{}{
		"agent_id":  agentID,
		"type":      teleType,
		"data":      data,
		"timestamp": time.Now(),
	}

	_, err := r.mongo.Collection(collectionName).InsertOne(ctx, document)
	_ = collection // avoiding unused warning if needed, but we used it above
	return err
}
