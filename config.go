package main

import (
	"encoding/json"
	"os"
)

type Account struct {
	Username string `json:"username" yaml:"username" toml:"username" ini:"username"`
	Password string `json:"password" yaml:"password" toml:"password" ini:"password"`
}

type Config struct {
	Account Account `json:"account" yaml:"account" toml:"account" ini:"account"`
}

func NewConfig() *Config {
	file, _ := os.Open("config.json")
	defer file.Close()

	decoder := json.NewDecoder(file)
	config := Config{}
	err := decoder.Decode(&config)
	if err != nil {
		panic(err)
	}
	return &config
}
