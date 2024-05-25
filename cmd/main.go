package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	utils "github.com/seniorLikeToCode/pastebin"
	"github.com/seniorLikeToCode/pastebin/generator"
)

type Data struct {
	Content string `json:"content"`
}

// temp db
var store = make(map[string]string)

func main() {
	router := mux.NewRouter()

	api := router.PathPrefix("/api/v1").Subrouter()

	api.HandleFunc("/{id}", getContentHandler).Methods("POST")
	api.HandleFunc("/", createContentHandler).Methods("POST")

	// Serve static files
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./templates/"))))

	// Serve index page on all unhandled routes
	router.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./templates/index.html")
	})

	log.Println("Listening on port 5000")
	if err := http.ListenAndServe(":5000", router); err != nil {
		log.Fatal(err)
	}
}

func getContentHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		http.Error(w, "ID not found", http.StatusBadRequest)
		return
	}

	// log.Printf("Received ID: %s", id)

	getContent, exists := store[id]
	if !exists {
		http.Error(w, "Content not found", http.StatusNotFound)
		return
	}

	if err := utils.WriteJSON(w, http.StatusOK, map[string]string{"content": getContent}); err != nil {
		log.Printf("Error writing JSON response: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}

func createContentHandler(w http.ResponseWriter, r *http.Request) {
	var d Data
	if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
		log.Printf("Error decoding JSON request body: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// log.Printf("Received content: %s", d.Content)

	tag := generator.GenerateTag(d.Content)
	store[tag] = d.Content

	if err := utils.WriteJSON(w, http.StatusOK, map[string]string{"id": tag}); err != nil {
		log.Printf("Error writing JSON response: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}
