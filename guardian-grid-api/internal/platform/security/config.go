package security

import (
	"os"

	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
}

func GetJWTSecret() []byte {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return []byte("default_fallback_secret_for_dev_only")
	}
	return []byte(secret)
}

func GetDBPath() string {
	path := os.Getenv("DB_PATH")
	if path == "" {
		return "static/guardian.db"
	}
	return path
}

func GetMongoURI() string {
	uri := os.Getenv("MONGO_URI")
	if uri == "" {
		return "mongodb://localhost:27017"
	}
	return uri
}

func GetMongoDBName() string {
	name := os.Getenv("MONGO_DB_NAME")
	if name == "" {
		return "guardian_grid_telemetry"
	}
	return name
}

func GetPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		return "8080"
	}
	return port
}
