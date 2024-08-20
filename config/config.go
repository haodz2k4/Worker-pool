package config

type Config struct {
	WorkerCount int
}

func LoadConfig() *Config {
	return &Config{
		WorkerCount: 4, //default worker
	}
}
