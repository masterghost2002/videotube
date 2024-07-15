package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PublicHost string //ip address where if host the server
	DBPort     string
	DBUser     string
	DBPassword string
	DBAddress  string
	DBName     string
	JWTSecret  string
	CryptSalt  string
	SSLMODE    string
	CRT_PATH   string
}

var ENVS = initConfig()

func initConfig() Config {
	godotenv.Load("/home/rakesh/Downloads/videotube/.env")
	return Config{
		PublicHost: getEnv("PUBLIC_HOST", "http://localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBUser:     getEnv("DB_USER", "root"),
		DBPassword: getEnv("DB_PASSWORD", "password"),
		DBAddress:  getEnv("DB_HOST", "localhost"),
		DBName:     getEnv("DB_NAME", "test"),
		JWTSecret:  getEnv("JWT_SECRET", "secret-key"),
		CryptSalt:  getEnv("CRYPT_SALT", "salt"),
		SSLMODE:    getEnv("SSLMODE", "disable"),
		CRT_PATH:   getEnv("CRT_PATH", "./ca.crt"),
	}
}
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
