package env

import (
	"os"
	"sync"

	"github.com/joho/godotenv"
)

var lock = &sync.Mutex{}

type Config struct {
	Port          string
	DSN           string
	DB            string
	BaseURL       string
	ServiceID     string
	FoodURL       string
	WinefamilyURL string
}

func newConfig() *Config {
	godotenv.Load(".env")

	config := &Config{}

	// ----- SET Values -----
	port := ":8080"
	if val, ok := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT"); ok {
		port = ":" + val
	}

	config.Port = port
	config.DSN = os.Getenv("MONGO_CONN_STRING")
	config.DB = os.Getenv("MONGO_COLLECTION")

	config.BaseURL = "/api"

	return config
}

var singleton *Config

func GetInstance() *Config {
	if singleton == nil {
		lock.Lock()
		defer lock.Unlock()

		if singleton == nil {
			singleton = newConfig()
		}
	}

	return singleton
}
