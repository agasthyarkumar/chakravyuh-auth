package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	// =========================
	// SERVER
	// =========================
	Port string

	// =========================
	// DATABASE
	// =========================
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string

	// =========================
	// AUTH
	// =========================
	JWTSecret string

	AccessTokenExpiryMinutes string
	RefreshTokenExpiryDays   string

	// =========================
	// REDIS
	// =========================
	RedisURL string

	// =========================
	// SUPERADMIN
	// =========================
	SuperAdminUsername string
	SuperAdminPassword string

	// =========================
	// FRONTEND
	// =========================
	FrontendURL string

	// =========================
	// RATE LIMITING
	// =========================
	RateLimitRequests     string
	RateLimitWindowSeconds string
}

var AppConfig Config

func LoadEnv() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Failed to load .env file")
	}

	AppConfig = Config{
		// =========================
		// SERVER
		// =========================
		Port: os.Getenv("PORT"),

		// =========================
		// DATABASE
		// =========================
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),

		// =========================
		// AUTH
		// =========================
		JWTSecret: os.Getenv("JWT_SECRET"),

		AccessTokenExpiryMinutes: os.Getenv("ACCESS_TOKEN_EXPIRY_MINUTES"),
		RefreshTokenExpiryDays:   os.Getenv("REFRESH_TOKEN_EXPIRY_DAYS"),

		// =========================
		// REDIS
		// =========================
		RedisURL: os.Getenv("REDIS_URL"),

		// =========================
		// SUPERADMIN
		// =========================
		SuperAdminUsername: os.Getenv("SUPERADMIN_USERNAME"),
		SuperAdminPassword: os.Getenv("SUPERADMIN_PASSWORD"),

		// =========================
		// FRONTEND
		// =========================
		FrontendURL: os.Getenv("FRONTEND_URL"),

		// =========================
		// RATE LIMITING
		// =========================
		RateLimitRequests:      os.Getenv("RATE_LIMIT_REQUESTS"),
		RateLimitWindowSeconds: os.Getenv("RATE_LIMIT_WINDOW_SECONDS"),
	}
}