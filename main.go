package main

import (
	"log"
	"net/http"
	"web/handlers"
	"web/routes"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterArticlesRouter(r)
	handlers.ServeStaticFiles(r) // traer la funcion para manejar archivos estaticos

    // http.Handle("/", r)
	log.Println("Server running on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Error while serving the server: %v", err)
	}
}
