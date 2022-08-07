package routes

import (
	"log"
	"net/http"

	controller "github.com/pedroppd/api-go-rest.git/controllers"
)

func HandleRequest() {
	http.HandleFunc("/", controller.Home)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
