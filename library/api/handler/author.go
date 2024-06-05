package handler

import (
	"authors/models"
	"authors/storage"
	"encoding/json"
	"net/http"
	"strconv"
)

type Handler struct {
	Storage *storage.Author
}

func NewHandeler(s *storage.Author) *Handler {
	return &Handler{Storage: s}
}

func (h *Handler) CreateAuthor(w http.ResponseWriter, r *http.Request) {

	var body models.AuthorRequest

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp, err := h.Storage.CreateNewAuthor(body)
	if err != nil {
		http.Error(w, "failed creating new author", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func (h *Handler) GetAuthorByID(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	res, err := h.Storage.GetAuthorByID(intId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) GetAuthorByName(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	res, err := h.Storage.GetAuthorByName(name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
