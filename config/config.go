package config

type Config struct {
	DBUsername string `mapstructure:"DB_USERNAME"`
	DBName     string `mapstructure:"DB_NAME"`
	DBHost     string `mapstructure:"DB_HOST"`
	DBPort     string `mapstructure:"DB_PORT"`
}
