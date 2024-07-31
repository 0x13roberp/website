package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"
	"web/models"

	"github.com/gorilla/mux"
)

var articles = map[int]models.Article{} // creamos un array para guardar articulos
var nextID = 1                          // la variable que asignara un id a cada articulo

// manejar las solicitudes para crear un articulo
func CreateArticle(w http.ResponseWriter, r *http.Request) {
	var article models.Article

	// NewDecoder sirve para leer datos json de la peticion http y deserializarlo en una estructura de go
	err := json.NewDecoder(r.Body).Decode(&article)
	if err != nil {
		http.Error(w, "Invalid input!", http.StatusBadRequest)
		return
	}

	// comprobar que no pasan campos vacios
	if article.Title == "" || article.Author == "" || article.Content == "" {
		http.Error(w, "Fill the fields!", http.StatusBadRequest)
		return
	}

	article.ID = nextID            // le asignamos el id guardado en nextID
	nextID++                       // aumentamos la variable anterior para que no hayan id repetidos
	article.CreatedAt = time.Now() // asignamos los valores de tiempo a ahora
	article.UpdatedAt = time.Now()
	articles[article.ID] = article // guardamos el articulo creado en el array de articulos y con el id de la misma

	// retornar informacion para especificar que se creo sin problema
	w.WriteHeader(http.StatusCreated)

	// NewEncoder sirve para escribir datos json a la peticion http
	json.NewEncoder(w).Encode(article)
}

func DeleteArticle(w http.ResponseWriter, r *http.Request) {
	// mux vars agarra el id de la ruta pasada
	idStr := mux.Vars(r)["id"] // id del articulo que queremos eliminar, por defecto viene en forma string con lo cual tenemos que castearlo a int
	id, err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(w, "Invalid ID!", http.StatusBadRequest)
		return
	}

	// comprobar si existe el articulo
	if _, exists := articles[id]; !exists {
		http.Error(w, "Article not found!", http.StatusNotFound)
		return
	}

	delete(articles, id)                // eliminar el articulo
	w.WriteHeader(http.StatusNoContent) // decirle que ya no existe ese articulo
}

func UpdateArticle(w http.ResponseWriter, r *http.Request)  {
    idStr := mux.Vars(r)["id"]
    id, err := strconv.Atoi(idStr)

    if err != nil{
        http.Error(w, "Invalid ID!", http.StatusBadRequest)
        return
    }

    if _, exists := articles[id]; !exists {
        http.Error(w, "Article not found!", http.StatusNotFound)
        return
    }


}

func ListAllArticles(w http.ResponseWriter, r *http.Request) {
}
