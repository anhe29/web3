package config

type Config struct {
	JWTSecret string
}

func LoadConfig() *Config {
	return &Config{
		JWTSecret: "SECRET",
	}
}
