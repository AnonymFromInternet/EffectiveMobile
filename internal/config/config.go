package config

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Mode           string `yaml:"mode" env-default:"local" env-required:"true"`
	HTTPServer     string `yaml:"http_server"`
	DataSourceName string `yaml:"data_source_name" env-required:"true"`
}

type HTTPServer struct {
	Address     string        `yalm:"address" env-default:"localhost:8082"`
	Timeout     time.Duration `yaml:"timeout"`
	IdleTimeout time.Duration `yaml:"idle_timeout"`
}

func MustCreate() *Config {

	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		// TODO: Можно ли уже тут логгировать приложение по разным уровням, хотя логгер еще не загружен?
		// Хотя наверное нет смысла
		log.Fatal("package config.MustCreate: cannot get CONFIG_PATH")
	}

	if _, e := os.Stat(configPath); os.IsNotExist(e) {
		log.Fatalf("package config.MustCreate: config path { %s } does not exist", configPath)
	}

	var config Config
	e := cleanenv.ReadConfig(configPath, &config)
	if e != nil {
		log.Fatal("package config.MustCreate: cannot read config. Getted error :", e)
	}

	return &config
}