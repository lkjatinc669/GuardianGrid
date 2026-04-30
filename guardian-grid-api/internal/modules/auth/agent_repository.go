package auth

import (
	"database/sql"

	"guardian-grid-api/internal/platform/database"
)

type ARepository struct {
	db *sql.DB
}

func AgentRepository() *ARepository {
	return &ARepository{
		db: database.GetSQLite(),
	}
}

func (r *ARepository) CreateAgent(id, hostname, token string) error {
	_, err := r.db.Exec(
		"INSERT INTO agents (id, hostname, token) VALUES (?, ?, ?)",
		id, hostname, token,
	)
	return err
}

func (r *ARepository) GetAgentByToken(token string) (string, error) {
	var id string
	err := r.db.QueryRow(
		"SELECT id FROM agents WHERE token = ?",
		token,
	).Scan(&id)

	return id, err
}
