package dashboard

import (
	"context"
	"sync"
	"time"

	"guardian-grid-api/internal/platform/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository struct {
	db *mongo.Database
}

func NewRepository() *Repository {
	return &Repository{
		db: database.GetMongoDatabase(),
	}
}

var collections = []string{
	"telemetry_dns",
	"telemetry_processes",
	"telemetry_programs",
	"telemetry_ports",
	"telemetry_network",
	"telemetry_persistence",
	"telemetry_activity",
	"telemetry_users",
	"telemetry_host_info",
	"telemetry_uptime",
}

func (r *Repository) GetLatestTelemetry(agentID string) (map[string]interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	mergedData := make(map[string]interface{})
	mergedData["agent_id"] = agentID
	
	details := make(map[string]interface{})

	var wg sync.WaitGroup
	var mu sync.Mutex

	for _, collName := range collections {
		wg.Add(1)
		go func(name string) {
			defer wg.Done()
			var result map[string]interface{}
			err := r.db.Collection(name).FindOne(ctx, bson.M{"agent_id": agentID}, options.FindOne().SetSort(bson.M{"timestamp": -1})).Decode(&result)
			if err == nil {
				mu.Lock()
				// Use the 'type' field as the key in the 'data' map
				if t, ok := result["type"].(string); ok {
					details[t] = result["data"]
				}
				mu.Unlock()
			}
		}(collName)
	}

	wg.Wait()
	mergedData["data"] = details
	return mergedData, nil
}

func (r *Repository) GetAllAgentsLatest() ([]map[string]interface{}, error) {
	// This is more complex now. We first need to find all unique agent IDs from SQLite
	// and then call GetLatestTelemetry for each.
	
	// For now, let's keep it simple and query the agents from SQLite
	sqlite := database.GetSQLite()
	rows, err := sqlite.Query("SELECT id FROM agents")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []map[string]interface{}
	for rows.Next() {
		var id string
		if err := rows.Scan(&id); err == nil {
			data, err := r.GetLatestTelemetry(id)
			if err == nil {
				results = append(results, data)
			}
		}
	}

	return results, nil
}
