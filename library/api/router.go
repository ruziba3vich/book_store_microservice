package api

import (
	"authors/api/handler"
	"authors/storage"
	"net/http"
)

func Router(st *storage.Author) http.Handler {
	mux := http.NewServeMux()
	handler := handler.NewHandeler(st)
	mux.HandleFunc("POST /author", handler.CreateAuthor)
	mux.HandleFunc("GET /author/id", handler.GetAuthorByID)
	mux.HandleFunc("GET /author/name", handler.GetAuthorByName)

	return mux
}
