package config

type Config struct {
	App *AppConfig
}

type AppConfig struct {
	CurrentDir string
}
