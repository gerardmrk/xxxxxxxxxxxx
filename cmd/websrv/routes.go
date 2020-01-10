package main

import (
	"database/sql"
	"fmt"
	"net/http"
)

type WebService struct {
	db *sql.DB
}

func (sv *WebService) IndexEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/plain")
	fmt.Fprint(w, "INDEX ROUTE")
}

func (sv *WebService) Login(w http.ResponseWriter, r *http.Request) {
	http.Error(w, http.StatusText(http.StatusNotImplemented), http.StatusNotImplemented)
}

func (sv *WebService) Logout(w http.ResponseWriter, r *http.Request) {
	http.Error(w, http.StatusText(http.StatusNotImplemented), http.StatusNotImplemented)
}
