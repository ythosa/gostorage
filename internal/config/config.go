package config

import (
	"github.com/BurntSushi/toml"
	"log"
	"os"
	"sync"
)

type Config struct {
	Server
	Logger
	Files
}

func newConfig() *Config {
	return &Config{
		Server: Server{
			BindHost:     "localhost",
			BindAddr:     ":8080",
			WriteTimeout: defaultWriteTimeout,
			ReadTimeout:  defaultReadTimeout,
			IdleTimeout:  defaultIdleTimeout,
		},
		Logger: newLogger("debug"),
		Files:  newFiles("/static/"),
	}
}

func Get() *Config {
	var (
		cfg  *Config
		once sync.Once
	)

	once.Do(func() {
		cfg = newConfig()

		cfgPath := os.Getenv("CONFIG_PATH")
		if cfgPath == "" {
			cfgPath = "../../configs/prod.toml"
		}

		_, err := toml.DecodeFile(cfgPath, cfg)
		if err != nil {
			log.Fatal(err)
		}
	})

	return cfg
}

const (
	defaultWriteTimeout = 15
	defaultReadTimeout  = 15
	defaultIdleTimeout  = 60
)

type Server struct {
	BindHost     string `toml:"bind_host"`
	BindAddr     string `toml:"bind_addr"`
	WriteTimeout int    `toml:"write_timeout"`
	ReadTimeout  int    `toml:"read_timeout"`
	IdleTimeout  int    `toml:"idle_timeout"`
}

type Logger struct {
	LogLevel string `toml:"log_level"`
}

func newLogger(logLevel string) Logger {
	return Logger{LogLevel: logLevel}
}

type Files struct {
	Path string `toml:"file_path"`
}

func newFiles(path string) Files {
	return Files{Path: path}
}
