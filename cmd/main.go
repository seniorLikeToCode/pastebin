package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	utils "github.com/seniorLikeToCode/pastebin"
)

func main() {
	router := mux.NewRouter()

	api := router.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/{id}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, ok := vars["id"]
		if !ok {
			http.Error(w, "ID not found", http.StatusBadRequest)
			return
		}

		log.Printf("Received ID: %s", id)

		if err := utils.WriteJSON(w, http.StatusOK, id); err != nil {
			log.Printf("Error writing JSON response: %v", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
	}).Methods("POST")

	api.HandleFunc("/save", func(w http.ResponseWriter, r *http.Request) {

	}).Methods("POST")

	// Serve static files
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./templates/"))))

	// Serve index page on all unhandled routes
	router.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./templates/index.html")
	})

	log.Println("Listening on port 5000")
	err := http.ListenAndServe(":5000", router)
	if err != nil {
		log.Fatal(err)
	}
}
