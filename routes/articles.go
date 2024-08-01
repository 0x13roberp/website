package routes

import (
	"net/http"
	"web/handlers"

	"github.com/gorilla/mux"
)

func RegisterArticlesRouter(r *mux.Router) {
	// definimos la "raiz" de nuestros url de articulos
	articleRouter := r.PathPrefix("/articles").Subrouter()

	articleRouter.HandleFunc("/", handlers.ListAllArticles).Methods(http.MethodGet)
	articleRouter.HandleFunc("", handlers.ListAllArticles).Methods(http.MethodGet)
	articleRouter.HandleFunc("/create", handlers.CreateArticle).Methods(http.MethodPost)
	articleRouter.HandleFunc("/update/{id}", handlers.UpdateArticle).Methods(http.MethodPut)
	articleRouter.HandleFunc("/delete/{id}", handlers.DeleteArticle).Methods(http.MethodDelete)
}
