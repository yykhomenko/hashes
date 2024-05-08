package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync/atomic"

	"github.com/gorilla/mux"
)

type response struct {
	Value    string `json:"value,omitempty"`
	ErrorID  byte   `json:"errorID,omitempty"`
	ErrorMsg string `json:"errorMsg,omitempty"`
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

	if !validateMsisdnLen(msisdn) {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(response{ErrorID: 2, ErrorMsg: "Not supported msisdn format: " + msisdn})
		return
	}

	if cc, ok := validateCC(msisdn); !ok {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(response{ErrorID: 3, ErrorMsg: "Not supported cc: " + cc})
		return
	}

	if ndc, ok := validateNDC(msisdn, s.config.NDCS); !ok {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(response{ErrorID: 4, ErrorMsg: "Not supported ndc: " + ndc})
		return
	}

	_ = json.NewEncoder(w).Encode(response{Value: s.store.Hash(msisdn[3:])})
}
