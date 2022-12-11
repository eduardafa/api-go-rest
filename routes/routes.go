package routes

import (
	"api-go-rest/controllers"
	"api-go-rest/middleware"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.Use(middleware.SetContentTypeMiddleware)
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/allpersonalities", controllers.GetAllPersonalities).Methods("Get")
	r.HandleFunc("/api/personality/{id}", controllers.GetPersonalityById).Methods("Get")
	r.HandleFunc("/api/personality", controllers.CreatePersonality).Methods("Post")
	r.HandleFunc("/api/personality/{id}", controllers.DeletePersonalityById).Methods("Delete")
	r.HandleFunc("/api/personality/{id}", controllers.UpdatePersonalityById).Methods("Put")
	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
}
