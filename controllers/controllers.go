package controllers

import (
	"api-go-rest/models"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home page")
}

func AllPersonalities(w http.ResponseWriter, r *http.Request) {
	errEncode := json.NewEncoder(w).Encode(models.Personalities)
	if errEncode != nil {
		fmt.Println("Error: ", errEncode)
	}
}

func GetPersonality(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	for _, personality := range models.Personalities {
		if strconv.Itoa(personality.Id) == id {
			errEncode := json.NewEncoder(w).Encode(personality)
			if errEncode != nil {
				fmt.Println("Error: ", errEncode)
			}
		}
	}
}
