package main

import (
	"context"
	"log"

	"github.com/RandySteven/go-kopi/apps"
	"github.com/RandySteven/go-kopi/pkg/config"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load("./files/env/.env")
	if err != nil {
		log.Fatalln(`failed to load .env `, err)
		return
	}
}

func main() {
	configPath, err := config.ParseFlags()
	if err != nil {
		log.Fatalln(err)
		return
	}

	config, err := config.NewConfig(configPath)
	if err != nil {
		log.Fatalln(err)
		return
	}
	ctx := context.TODO()

	app, err := apps.NewApp(config)
	if err != nil {
		log.Fatalln(`Error starting app `, err)
		return
	}

	if err = app.ExecuteMigration(ctx); err != nil {
		log.Fatalln(`Error executing migration `, err)
		return
	}

}
