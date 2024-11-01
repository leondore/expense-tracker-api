package config

import (
	"log"
	"time"

	"github.com/joeshaw/envdecode"
	"github.com/joho/godotenv"
)

type Conf struct {
	Server ConfServer
	DB     ConfDB
}

type ConfServer struct {
	Port         int           `env:"SERVER_PORT,required"`
	TimeoutRead  time.Duration `env:"SERVER_TIMEOUT_READ,required"`
	TimeoutWrite time.Duration `env:"SERVER_TIMEOUT_WRITE,required"`
	TimeoutIdle  time.Duration `env:"SERVER_TIMEOUT_IDLE,required"`
	Debug        bool          `env:"SERVER_DEBUG,required"`
}

type ConfDB struct {
	Host     string `env:"DB_HOST"`
	Port     int    `env:"DB_PORT"`
	Username string `env:"DB_USER"`
	Password string `env:"DB_PASS"`
	DBName   string `env:"DB_NAME"`
	Debug    bool   `env:"DB_DEBUG"`
}

func LoadEnv() error {
	return godotenv.Load()
}

func New() *Conf {
	var c Conf
	if err := envdecode.StrictDecode(&c); err != nil {
		log.Fatalf("Failed to decode: %v", err)
	}

	return &c
}

func NewDB() *ConfDB {
	var c ConfDB
	if err := envdecode.StrictDecode(&c); err != nil {
		log.Fatalf("Failed to decode: %v", err)
	}

	return &c
}
