package config

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	Database DatabaseConfig
	JWT      JWTConfig
	Server   ServerConfig
	SSO      SSOConfig
}

type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSLMode  string
}

type JWTConfig struct {
	Secret                string
	TokenExpirationHours  int
	RefreshExpirationDays int
}

type ServerConfig struct {
	Port       string
	Mode       string // debug, release
	Origins    []string
	PublicURL  string // URL publique de l'application (ex: https://tools.marocpme.gov.ma)
}

type SSOConfig struct {
	Enabled        bool
	AutoProvision  bool
	DefaultRole    string
	DefaultGroup   string
	GroupMapping   map[string]string // map[AuthentikGroup]AirboardGroup
	AdminGroups    []string          // Groupes Authentik qui ont le rôle admin
}

func LoadConfig() *Config {
	// Charger le fichier .env si présent
	if err := godotenv.Load(); err != nil {
		log.Println("Aucun fichier .env trouvé, utilisation des variables d'environnement")
	}

	// Configuration base de données
	dbPort, err := strconv.Atoi(getEnv("DB_PORT", "5432"))
	if err != nil {
		dbPort = 5432
	}

	// Configuration JWT
	tokenExp, err := strconv.Atoi(getEnv("JWT_TOKEN_EXPIRATION_HOURS", "24"))
	if err != nil {
		tokenExp = 24
	}

	refreshExp, err := strconv.Atoi(getEnv("JWT_REFRESH_EXPIRATION_DAYS", "7"))
	if err != nil {
		refreshExp = 7
	}

	// Configuration SSO
	ssoEnabled := getEnv("SSO_ENABLED", "false") == "true"
	ssoAutoProvision := getEnv("SSO_AUTO_PROVISION", "true") == "true"

	// Parse admin groups (comma-separated)
	adminGroups := []string{}
	if adminGroupsStr := getEnv("SSO_ADMIN_GROUPS", "airboard-admins"); adminGroupsStr != "" {
		for _, group := range splitAndTrim(adminGroupsStr, ",") {
			adminGroups = append(adminGroups, group)
		}
	}

	return &Config{
		Database: DatabaseConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     dbPort,
			User:     getEnv("DB_USER", "airboard"),
			Password: getEnv("DB_PASSWORD", "airboard123"),
			DBName:   getEnv("DB_NAME", "airboard"),
			SSLMode:  getEnv("DB_SSLMODE", "disable"),
		},
		JWT: JWTConfig{
			Secret:                getEnv("JWT_SECRET", "airboard-super-secret-key-2024"),
			TokenExpirationHours:  tokenExp,
			RefreshExpirationDays: refreshExp,
		},
		Server: ServerConfig{
			Port:      getEnv("PORT", "8080"),
			Mode:      getEnv("GIN_MODE", "debug"),
			PublicURL: getEnv("PUBLIC_URL", "http://localhost:5173"),
			Origins: []string{
				getEnv("FRONTEND_URL", "http://localhost:3000"),
				"http://localhost:5173", // Vite dev server
				"http://localhost:8080", // Swagger
			},
		},
		SSO: SSOConfig{
			Enabled:       ssoEnabled,
			AutoProvision: ssoAutoProvision,
			DefaultRole:   getEnv("SSO_DEFAULT_ROLE", "user"),
			DefaultGroup:  getEnv("SSO_DEFAULT_GROUP", "Common"),
			GroupMapping:  make(map[string]string), // Sera peuplé par les groupes Authentik
			AdminGroups:   adminGroups,
		},
	}
}

func (c *Config) GetDSN() string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		c.Database.Host,
		c.Database.Port,
		c.Database.User,
		c.Database.Password,
		c.Database.DBName,
		c.Database.SSLMode,
	)
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func splitAndTrim(s, sep string) []string {
	var result []string
	for _, item := range strings.Split(s, sep) {
		if trimmed := strings.TrimSpace(item); trimmed != "" {
			result = append(result, trimmed)
		}
	}
	return result
}