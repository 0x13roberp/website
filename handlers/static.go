package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
)

func ServeStaticFiles(r *mux.Router) {
	// le decimos donde buscar los archivos estaticos
	staticFiles := http.FileServer(http.Dir("static"))

	// StripPrefix se usa para modificar la ruta del URL antes de que lo pase al handler. se usa para facilitar la interaccion con archivos estaticos
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", staticFiles))

	// ruta para servir el index.html
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/index.html")
	})
}
