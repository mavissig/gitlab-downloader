package loader

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"log"
)

type Config struct {
	Url   string `envconfig:"URL" required:"true"`
	Token string `envconfig:"TOKEN" required:"true"`
}

func LoadConfig() *Config {
	for _, fileName := range []string{".env.local", ".env"} {
		err := godotenv.Load(fileName)
		if err != nil {
			log.Println("[CONFIG] ERROR: ", err)
		}
	}

	cfg := &Config{}

	if err := envconfig.Process("", cfg); err != nil {
		log.Fatalln(err)
	}

	return cfg
}
