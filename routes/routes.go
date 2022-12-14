package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	controller "github.com/pedroppd/api-go-rest.git/controllers"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/", controller.Home).Methods("GET")
	r.HandleFunc("/api/personalities", controller.GetAllPersonalities).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", r))
}
