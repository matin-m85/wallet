package config

import "os"

type Config struct {
	WalletURL string
	GiftURL   string
	Port      string
}

func LoadConfig() *Config {
	return &Config{
		WalletURL: getEnv("WALLET_URL", "http://localhost:8081"),
		GiftURL:   getEnv("GIFT_URL", "http://localhost:8082"),
		Port:      getEnv("PORT", "8080"),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
