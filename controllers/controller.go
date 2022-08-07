package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pedroppd/api-go-rest.git/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home Page")
}

func GetAllPersonalities(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.Personalitys)
}
