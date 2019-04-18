package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

// HomeHandler handles request for /
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

// ProductsHandler handles request for /products
func ProductsHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

// ArticlesHandler handles request for /articles
func ArticlesHandler(w http.ResponseWriter, r *http.Request) {

}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/products", ProductsHandler)
	r.HandleFunc("/articles", ArticlesHandler)
	// http.Handle("/", r)
	http.ListenAndServe("", r)
}
