package main

import "digit-caster/internal/app"

const (
	configPath = "config/config"
)

func main() {
	app.Run(configPath)
}
