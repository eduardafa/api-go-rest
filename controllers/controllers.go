package controllers

import (
	"api-go-rest/database"
	"api-go-rest/models"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home page")
}

func GetAllPersonalities(w http.ResponseWriter, r *http.Request) {
	var personalities []models.Personality
	database.DB.Find(&personalities)
	errEncode := json.NewEncoder(w).Encode(personalities)
	if errEncode != nil {
		fmt.Println("Error: ", errEncode)
	}
}

func GetPersonalityById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var personality models.Personality
	database.DB.First(&personality, id)
	errEncode := json.NewEncoder(w).Encode(personality)
	if errEncode != nil {
		fmt.Println("Error: ", errEncode)
	}
}

func CreatePersonality(w http.ResponseWriter, r *http.Request) {
	var personality models.Personality
	errDecode := json.NewDecoder(r.Body).Decode(&personality)
	if errDecode != nil {
		fmt.Println("Error: ", errDecode)
	}
	database.DB.Create(&personality)
	errEncode := json.NewEncoder(w).Encode(personality)
	if errEncode != nil {
		fmt.Println("Error: ", errEncode)
	}
}

func DeletePersonalityById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var personality models.Personality
	database.DB.Delete(&personality, id)
	errEncode := json.NewEncoder(w).Encode(personality)
	if errEncode != nil {
		fmt.Println("Error: ", errEncode)
	}
}

func UpdatePersonalityById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var personality models.Personality
	database.DB.First(&personality, id)
	errDecode := json.NewDecoder(r.Body).Decode(&personality)
	if errDecode != nil {
		fmt.Println("Error: ", errDecode)
	}
	database.DB.Save(&personality)
	errEncode := json.NewEncoder(w).Encode(personality)
	if errEncode != nil {
		fmt.Println("Error: ", errEncode)
	}
}
