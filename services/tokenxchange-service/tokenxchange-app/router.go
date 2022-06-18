package main

import "github.com/gorilla/mux"

func RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/oauth/google/redirect",RedirectToGoogle).Methods("POST")
	
}
