package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func WriteJson(w http.ResponseWriter, status int, v any) error {
	w.Header().Set("Content-Type", "application/json");
	w.WriteHeader(status);

	return json.NewEncoder(w).Encode(v);
}

type ApiError struct {
	Error string;
}

type apiFunc func(http.ResponseWriter, *http.Request) error;

func makeHTTPHandleFunc (f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJson(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
}

type APIServer struct {
	listenAddr string
}

func NewApiServer(listenAddr string) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
	}
}

func (s *APIServer) Run() {
	router := mux.NewRouter();

	router.HandleFunc("/account", makeHTTPHandleFunc(s.handleAccount))
	router.HandleFunc("/account/{id}", makeHTTPHandleFunc(s.handleGetAccount))

	log.Printf("Running server on port: %s\n", s.listenAddr)

	http.ListenAndServe(s.listenAddr, router)
}

func (s *APIServer) handleAccount(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return s.handleGetAccount(w, r);
	}

	if r.Method == "POST" {
		return s.handleCreateAccount(w, r);
	}

	if r.Method == "DELETE" {
		return s.handleDeleteAccount(w, r)
	}

	return fmt.Errorf("Method not allowed %s", r.Method);
}

func (s *APIServer) handleGetAccount(w http.ResponseWriter, r *http.Request) error {
	id := mux.Vars(r)["id"]

	return WriteJson(w, http.StatusOK, id);
}

func (s *APIServer) handleCreateAccount(w http.ResponseWriter, r *http.Request) error {
	account := NewAccout("Paweł", "Kromołowski")
	
	return WriteJson(w, http.StatusOK, account);
}

func (s *APIServer) handleDeleteAccount(w http.ResponseWriter, r *http.Request) error {
	return nil;
}

func (s *APIServer) handleTransfer(w http.ResponseWriter, r *http.Request) error {
	return nil;
}