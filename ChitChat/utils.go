package main

import (
	"log"
	"os"
)

func version() string {
	return "0.1"
}

func loadConfig() {
	file, err := os.Open("config.json")
	if err != nil {
		log.Fatalln("Cannot open log file", err)

	}
}
