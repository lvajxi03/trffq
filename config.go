package main

import (
	"os"
	"path/filepath"
)

const (
	CONFIG_FILENAME = ".traffiqrc"
)

type Config struct {
	Configfile string
	Lastnick   string
}

func NewConfig() *Config {
	c := &Config{}
	return c
}

func ReadConfig() *Config {
	dir, err := os.UserHomeDir()
	if err != nil {
		c := &Config{
			Configfile: filepath.Join(dir, CONFIG_FILENAME),
		}
		return c
	}
	return nil

}

func (config *Config) SaveConfig() {

}
