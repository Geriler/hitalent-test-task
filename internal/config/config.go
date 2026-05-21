package config

import (
	"flag"
	"sync"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	TimeoutStop time.Duration `env:"TIMEOUT_STOP" yaml:"timeout_stop" env-default:"10s"`
	HTTP        Address       `yaml:"http" env-prefix:"HTTP_"`
	Database    Database      `yaml:"database" env-prefix:"DATABASE_"`
}

type Address struct {
	Host string `env:"HOST" yaml:"host" env-required:"true"`
	Port int    `env:"PORT" yaml:"port" env-required:"true"`
}

type Database struct {
	Host     string `env:"HOST" yaml:"host" env-required:"true"`
	Port     int    `env:"PORT" yaml:"port" env-required:"true"`
	User     string `env:"USER" yaml:"user" env-required:"true"`
	Password string `env:"PASSWORD" yaml:"password" env-required:"true"`
	Name     string `env:"NAME" yaml:"name" env-required:"true"`
}

var once sync.Once
var config Config

func MustGet() Config {
	once.Do(func() {
		var path string
		flag.StringVar(&path, "config", "", "config file path")
		flag.Parse()

		if path == "" {
			err := cleanenv.ReadEnv(&config)
			if err != nil {
				panic(err)
			}
		} else {
			err := cleanenv.ReadConfig(path, &config)
			if err != nil {
				panic(err)
			}
		}
	})

	return config
}
