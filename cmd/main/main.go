package main

import (
	"context"
	"log"

	"github.com/RandySteven/go-kopi/apps"
	api_http "github.com/RandySteven/go-kopi/handlers/https"
	"github.com/RandySteven/go-kopi/pkg/config"
	"github.com/RandySteven/go-kopi/routes"
	"github.com/RandySteven/go-kopi/usecases"
	"github.com/gorilla/mux"
)

func init() {
	log.Println("  ________                  ____  __.            .__ ")
	log.Println(" /  _____/  ____           |    |/ _|____ ______ |__|")
	log.Println("/   \\  ___ /  _ \\   ______ |      < /  _ \\____ \\|  |")
	log.Println("\\    \\_\\  (  <_> ) /_____/ |    |  (  <_> )  |_> >  |")
	log.Println(" \\______  /\\____/          |____|__ \\____/|   __/|__|")
	log.Println("        \\/                         \\/     |__|        ")
}

func main() {
	configPath, err := config.ParseFlags()

	if err != nil {
		log.Fatal(err)
		return
	}

	config, err := config.NewConfig(configPath)
	if err != nil {
		log.Fatal(err)
		return
	}

	repositories, err := db.NewRepositories(config)
	if err != nil {
		log.Fatal(err)
		return
	}

	usecases := usecases.NewUsecases()
	if err != nil {
		log.Fatal(err)
		return
	}

	apiHttp := api_http.NewHTTPs(usecases)

	r := mux.NewRouter()
	apps.RegisterMiddleware(r)
	routes.InitRouter(apiHttp, r)
	config.Run(r)
}
