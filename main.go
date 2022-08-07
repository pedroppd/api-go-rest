package main

import (
	"github.com/pedroppd/api-go-rest.git/models"
	"github.com/pedroppd/api-go-rest.git/routes"
)

func main() {
	models.Personalitys = []models.Personality{{Id: 1, Name: "teste", History: "2"}, {Id: 1, Name: "teste 2", History: "3"}}
	routes.HandleRequest()
}
