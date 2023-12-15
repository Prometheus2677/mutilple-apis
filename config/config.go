package config

type DatabaseConfig struct {
	DBHost string
	DBPort string
	DBUser string
	DBPass string
	DBName string
}

type ServerConfig struct {
	Port string
}
