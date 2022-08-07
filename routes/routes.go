package routes

import (
	"log"
	"net/http"

	controller "github.com/pedroppd/api-go-rest.git/controllers"
)

func HandleRequest() {
	http.HandleFunc("/", controller.Home)
	http.HandleFunc("/api/personalities", controller.GetAllPersonalities)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
