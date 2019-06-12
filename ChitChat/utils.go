package main

import (
	"encoding/json"
	"log"
	"os"
)

type Configuration struct {
	Address      string
	ReadTimeout  int64
	WriteTimeout int64
	Static       string
}

var config Configuration
var logger *log.Logger

func version() string {
	return "0.1"
}

func loadConfig() {
	// 打开配置文件
	file, err := os.Open("config.json")
	if err != nil {
		log.Fatalln("Cannot open log file", err)
	}

	// 读取 JSON
	decoder := json.NewDecoder(file)
	config = Configuration{}
	err = decoder.Decode(&config)
	if err != nil {
		log.Fatalln("Cannot parse configuration from file", err)
	}
}
