package config

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Mode            string `yaml:"mode" env-default:"local" env-required:"true"`
	DataSourceName  string `yaml:"data_source_name" env-required:"true"`
	HTTPServer      `yaml:"http_server"`
	ExternalApiUrl  string          `yaml:"external_api_url"`
	MigrationsPaths MigrationsPaths `yaml:"migrations_paths"`
}

type HTTPServer struct {
	Address     string        `yalm:"address" env-default:"localhost:8082"`
	Timeout     time.Duration `yaml:"timeout"`
	IdleTimeout time.Duration `yaml:"idle_timeout"`
}

type MigrationsPaths struct {
	Up   string `yaml:"up" env-required:"true"`
	Down string `yaml:"down" env-required:"true"`
}

func MustCreate() *Config {
	// TODO: DELETE
	os.Setenv("CONFIG_PATH", "../../config/local.yaml")

	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		log.Fatal("package config.MustCreate: cannot get CONFIG_PATH")
	}

	if _, e := os.Stat(configPath); os.IsNotExist(e) {
		log.Fatalf("package config.MustCreate: config path { %s } does not exist", configPath)
	}

	var config Config
	e := cleanenv.ReadConfig(configPath, &config)
	if e != nil {
		log.Fatal("package config.MustCreate: cannot read config. Getted error: ", e)
	}

	return &config
}
