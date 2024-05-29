package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq" // assuming you are using PostgreSQL
	"github.com/seniorLikeToCode/pastebin/service"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{addr: addr, db: db}
}

func (s *APIServer) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	// Initialize handler and register routes
	handler := service.NewHandler(s.db)
	handler.RegisterRoutes(subrouter)

	// Serve static files
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./templates/"))))

	// Serve index page on all unhandled routes
	router.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./templates/index.html")
	})

	log.Printf("Listening on %s", s.addr)
	return http.ListenAndServe(s.addr, router)
}
