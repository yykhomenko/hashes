package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync/atomic"

	"github.com/gorilla/mux"

	"github.com/yykhomenko/hashes/internal/hashes/store"
)

type response struct {
	Value    string `json:"value,omitempty"`
	ErrorID  byte   `json:"errorID,omitempty"`
	ErrorMsg string `json:"errorMsg,omitempty"`
}

type Server struct {
	counter *Counter
	store   *store.Store
}

func New(s *store.Store) *Server {
	return &Server{
		store:   s,
		counter: &Counter{},
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

func (s *Server) getMetrics(w http.ResponseWriter, _ *http.Request) {
	_, _ = fmt.Fprint(w,
		"hashes_total ", s.counter.hashes, "\n",
		"msisdns_total ", s.counter.msisdns)
}

func (s *Server) getMsisdn(w http.ResponseWriter, r *http.Request) {
	atomic.AddUint64(&s.counter.msisdns, 1)
	w.Header().Set("Content-Type", "application/json")

	hash := mux.Vars(r)["hash"]
	if msisdn, ok := s.store.Msisdn(hash); !ok {
		w.WriteHeader(http.StatusNotFound)
		_ = json.NewEncoder(w).Encode(response{ErrorID: 1, ErrorMsg: "Not found"})
	} else {
		_ = json.NewEncoder(w).Encode(response{Value: "380" + msisdn})
	}
}

func (s *Server) getHash(w http.ResponseWriter, r *http.Request) {
	atomic.AddUint64(&s.counter.hashes, 1)
	w.Header().Set("Content-Type", "application/json")

	msisdn := mux.Vars(r)["msisdn"]

	if !s.store.ValidateMsisdnLen(msisdn) {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(response{ErrorID: 2, ErrorMsg: "Not supported msisdn format: " + msisdn})
		return
	}

	if cc, ok := s.store.ValidateCC(msisdn); !ok {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(response{ErrorID: 3, ErrorMsg: "Not supported cc: " + cc})
		return
	}

	if ndc, ok := s.store.ValidateNDC(msisdn); !ok {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(response{ErrorID: 4, ErrorMsg: "Not supported ndc: " + ndc})
		return
	}

	_ = json.NewEncoder(w).Encode(response{Value: s.store.Hash(msisdn[3:])})
}
