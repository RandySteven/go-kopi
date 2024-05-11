package main

import (
	"github.com/RandySteven/go-kopi/pkg/config"
	"github.com/gorilla/mux"
	"log"
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

	//repositories, err := db.NewRepositories(config)
	//if err != nil {
	//	log.Fatal(err)
	//	return
	//}
	//
	//services, err := apps.NewServices(context.Background())
	//if err != nil {
	//	log.Fatal(err)
	//	return
	//}
	//
	//handlers := apps.NewHandlers(repositories, services)
	//
	r := mux.NewRouter()
	//
	//
	//apps.RegisterMiddleware(r)
	//
	//handlers.InitRouter(r)
	config.Run(r)
}
