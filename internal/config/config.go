package config

// Config holds application configuration
type Config struct {
	ServerPort string
	DBPath     string
}

// LoadConfig loads the configuration
func LoadConfig() *Config {
	return &Config{
		ServerPort: ":3000",
		DBPath:     "users.db",
	}
}
