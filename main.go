package main

import (
	"log"

	"github.com/pedroppd/api-go-rest.git/routes"
)

func main() {
	log.Println("Starting server...")
	routes.HandleRequest()
}
