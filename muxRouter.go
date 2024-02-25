//from the book Building RESTful Web Services with Go, Naren Yellavula

package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// ArticleHandler is a afunction handler
func ArticleHandler(w http.ResponseWriter, r *http.Request) {
	// mux.Vars returns alll path parametrs as a map
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Category is: %v\n", vars["category"])
	fmt.Fprintf(w, "ID is: %v\n", vars["id"])
}

func main() {
	// Create a newe router
	r := mux.NewRouter()
	// Atach a path which hendler
	r.HandleFunc("/articles/{category}/{id:[0-9]+}", ArticleHandler)
	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
