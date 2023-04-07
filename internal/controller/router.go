package controller

import (
	"net/http"
)

func RoutesInit(h *Handler) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/rest/substr/find", h.findStringHandler)
	mux.HandleFunc("/rest/email/check", h.findEmailHandler)
	mux.HandleFunc("/rest/iin/check", h.findIinHandler)
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	})

	return mux
}
