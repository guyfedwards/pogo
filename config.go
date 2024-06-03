package main

import (
	"fmt"
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Path          string
	Port          string
	DefaultSearch string            `yaml:"defaultSearch,omitempty"`
	Aliases       map[string]string `yaml:"aliases,omitempty"`
}

func NewConfig(path, port string) *Config {
	return &Config{
		Path:          path,
		Port:          port,
		DefaultSearch: "https://www.google.com/search?q=%s",
	}
}

func (c *Config) Load() error {
	b, err := os.ReadFile(c.Path)
	if err != nil {
		return fmt.Errorf("Config.Load: %w", err)
	}

	err = yaml.Unmarshal(b, &c)
	if err != nil {
		return fmt.Errorf("Config.Load: %w", err)
	}

	return nil
}

func (c *Config) Monitor() chan os.Signal {
	cancel := make(chan os.Signal, 1)

	go func() {
		t := time.NewTicker(10 * time.Second)

		for {
			select {
			case <-cancel:
				return
			case <-t.C:
				err := c.Load()
				if err != nil {
					fmt.Printf("config.Monitor failed to load: %s\n", err)
					panic(1)
				}
			}
		}
	}()

	return cancel
}
