package server

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yykhomenko/hashes/internal/config"
	"github.com/yykhomenko/hashes/internal/store"
)

type counter struct {
	hashes  uint64
	msisdns uint64
}

type Server struct {
	config  *config.Config
	store   *store.Store
	counter *counter
}

func New(c *config.Config, s *store.Store) *Server {
	return &Server{
		store:   s,
		counter: &counter{},
	}
}

func (s *Server) Start() {
	log.Println("http-Server listening...")

	r := mux.NewRouter()
	r.HandleFunc("/metrics", s.getMetrics).Methods("GET")
	r.HandleFunc("/hashes/{msisdn}", s.getHash).Methods("GET")
	r.HandleFunc("/msisdns/{hash}", s.getMsisdn).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", r))
}
