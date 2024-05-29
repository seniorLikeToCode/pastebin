package service

import (
	"encoding/json"
	"log"
	"net/http"

	"database/sql"

	"github.com/gorilla/mux"
	utils "github.com/seniorLikeToCode/pastebin"
	"github.com/seniorLikeToCode/pastebin/generator"
)

type Handler struct {
	store *Store
}

func NewHandler(db *sql.DB) *Handler {
	return &Handler{
		store: NewStore(db),
	}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/", h.createContentHandler).Methods("POST")
	router.HandleFunc("/{id}", h.getContentHandler).Methods("GET")
}

func (h *Handler) getContentHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uid, ok := vars["id"]
	if !ok {
		http.Error(w, "ID not found", http.StatusBadRequest)
		return
	}

	log.Println(uid)

	data, err := h.store.GetDatabyUid(uid)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Content not found", http.StatusNotFound)
		} else {
			log.Printf("Error getting content by Uid: %v", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
		return
	}

	if err := utils.WriteJSON(w, http.StatusOK, data); err != nil {
		log.Printf("Error writing JSON response: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}

func (h *Handler) createContentHandler(w http.ResponseWriter, r *http.Request) {
	var d Data
	if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
		log.Printf("Error decoding JSON request body: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	log.Println(d.Content)
	uid := generator.GenerateTag(d.Content)
	d.Uid = uid

	if err := h.store.Create(d); err != nil {
		log.Printf("Error creating content: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	if err := utils.WriteJSON(w, http.StatusOK, map[string]string{"id": uid}); err != nil {
		log.Printf("Error writing JSON response: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}
