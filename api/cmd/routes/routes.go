package routes

import (
	"api/cmd/controllers"
	"api/internal/middleware"
	services "api/internal/service"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func HandleRequest(characterService services.CharacterService) {
	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleware)

	characterController := controllers.NewCharacterController(characterService)

	r.HandleFunc("/api/v1/personagens/", characterController.GetAllCharacters).Methods("GET")
	r.HandleFunc("/api/v1/personagens/specific", characterController.GetCharacterByName).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
}
