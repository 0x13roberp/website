package routes

import "github.com/gorilla/mux"

func RegisterArticlesRouter(r *mux.Router)  {
    // definimos la "raiz" de nuestros url de articulos
    articleRouter := r.PathPrefix("articles").Subrouter()

    articleRouter.HandleFunc("/")

}
