package alerts

import (
	"context"
	"time"

	"guardian-grid-api/internal/platform/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	collection *mongo.Collection
}

func NewRepository() *Repository {
	db := database.GetMongoDatabase()
	return &Repository{
		collection: db.Collection("alerts"),
	}
}

func (r *Repository) SaveAlert(alert Alert) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	alert.CreatedAt = time.Now()
	_, err := r.collection.InsertOne(ctx, alert)
	return err
}

func (r *Repository) GetAlertsByAgent(agentID string) ([]Alert, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := r.collection.Find(ctx, bson.M{"agent_id": agentID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var alerts []Alert
	if err = cursor.All(ctx, &alerts); err != nil {
		return nil, err
	}
	return alerts, nil
}

func (r *Repository) GetAllAlerts() ([]Alert, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var alerts []Alert
	if err = cursor.All(ctx, &alerts); err != nil {
		return nil, err
	}
	return alerts, nil
}

func (r *Repository) AlertExists(agentID, cveID, programName, version string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	count, err := r.collection.CountDocuments(ctx, bson.M{
		"agent_id":     agentID,
		"cve_id":       cveID,
		"program_name": programName,
		"version":      version,
	})
	return count > 0, err
}
