package routes

import (
	"api/src/controllers"
	"api/src/middleware"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleware)
	r.HandleFunc("/api/personagens", controllers.TodosPersonagens).Methods("Get")
	r.HandleFunc("/api/", controllers.PersonageSpecific).Methods("Get")
	r.HandleFunc("/api/personagens/mongo", controllers.TodosPersonagensMongo).Methods("Get")
	r.HandleFunc("/api/personagens/mongo/specific", controllers.PersonageSpecificMongo).Methods("Get")
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
}
