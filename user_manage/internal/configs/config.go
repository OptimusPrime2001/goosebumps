package configs

type Config struct {
	ServerAddress string
}

func NewConfig() *Config {
	return &Config{
		ServerAddress: "127.0.0.1:8080",
	}
}
