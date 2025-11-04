package configs

import (
	"fmt"
	"user-manage-backend/internal/utils"
)

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}
type Config struct {
	ServerAddress string
	Database      DatabaseConfig
}

func NewConfig() *Config {
	return &Config{
		ServerAddress: "127.0.0.1:8080",
		Database: DatabaseConfig{
			Host:     utils.GetEnv("DB_HOST", "127.0.0.1"),
			Port:     utils.GetEnv("DB_PORT", "5432"),
			User:     utils.GetEnv("DB_USER", "postgres"),
			Password: utils.GetEnv("DB_PASSWORD", "postgres"),
			DBName:   utils.GetEnv("DB_NAME", "user_manage"),
			SSLMode:  utils.GetEnv("DB_SSLMODE", "disable"),
		},
	}
}
func (c *Config) GetDSN() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		c.Database.Host,
		c.Database.Port,
		c.Database.User,
		c.Database.Password,
		c.Database.DBName,
		c.Database.SSLMode,
	)
}
