package main

import "tracker/internal/config"

func main() {
	_, err := config.NewConfig()
	if err != nil {
		panic(err)
	}
	// implement all layers
	// add graceful shutdown
}
