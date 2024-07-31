package handlers

import "net/http"

func ServeStaticFiles() {
	// le decimos donde buscar los archivos estaticos
	staticFiles := http.FileServer(http.Dir("static"))

	// StripPrefix se usa para modificar la ruta del URL antes de que lo pase al handler. se usa para facilitar la interaccion con archivos estaticos
	http.Handle("/", http.StripPrefix("/", staticFiles))
}
