package main

import (
	"log"
	"net/http"
	"web/handlers"
)

func main() {
    handlers.ServeStaticFiles() // traer la funcion para manejar archivos estaticos

    log.Println("Server running on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Error while serving the server: %v", err)
	}
}
