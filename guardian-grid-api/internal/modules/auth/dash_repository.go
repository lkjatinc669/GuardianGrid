package auth

import (
	"database/sql"
	"errors"

	"guardian-grid-api/internal/platform/database"
)

type User struct {
	ID         int
	Username   string
	TOTPSecret string
}

type DRepository struct {
	db *sql.DB
}

func DashRepository() *DRepository {
	return &DRepository{
		db: database.GetSQLite(),
	}
}

func (r *DRepository) GetUserByUsername(username string) (*User, error) {
	row := r.db.QueryRow(
		"SELECT id, username, totp_secret FROM users WHERE username = ?",
		username,
	)

	var user User
	err := row.Scan(&user.ID, &user.Username, &user.TOTPSecret)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return &user, nil
}
