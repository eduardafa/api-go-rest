package routes

import (
	"api-go-rest/controllers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/allpersonalities", controllers.AllPersonalities).Methods("Get")
	r.HandleFunc("/api/personality/{id}", controllers.GetPersonality).Methods("Get")
	log.Fatal(http.ListenAndServe(":8080", r))
}
