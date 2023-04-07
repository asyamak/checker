package controller

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"restapi/internal/entity"
	"restapi/internal/usecase"
)

type Handler struct {
	u *usecase.Usecase
}

func NewHandler(u *usecase.Usecase) *Handler {
	return &Handler{
		u,
	}
}

func (h *Handler) findStringHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "not supported type", http.StatusUnsupportedMediaType)
		return
	}

	var req entity.StringRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	if len(req.Text) == 0 {
		http.Error(w, "invalid string length", http.StatusBadRequest)
	}

	maxSubstring := h.u.StringFinderUsecase.FindMaxSubstring(req.Text)
	req.Text = maxSubstring

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(req); err != nil {
		log.Printf("error to encode: %v\n", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	log.Printf("%v - %v\n", r.RequestURI, r.Method)
}

func (h *Handler) findEmailHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "not supported type", http.StatusUnsupportedMediaType)
		return
	}

	var req entity.EmailRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	if len(req.Emails) == 0 {
		http.Error(w, "invalid email length", http.StatusBadRequest)
	}

	emails := h.u.EmailCheckerUsecase.CheckEmail(req.Emails)

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(entity.EmailRequest{
		Emails: emails,
	}); err != nil {
		log.Printf("error to encode: %v\n", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	log.Printf("%v - %v\n", r.RequestURI, r.Method)
}

func (h *Handler) findIinHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "not supported type", http.StatusUnsupportedMediaType)
		return
	}

	var req entity.IinRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	IIN, err := h.u.IinGetterUsecase.GetIINs(req.IIN)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if len(IIN) == 0 {
		http.Error(w, errors.New("invalid input data").Error(), http.StatusForbidden)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(IIN); err != nil {
		log.Printf("error to encode: %v\n", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	log.Printf("%v - %v\n", r.RequestURI, r.Method)
}
