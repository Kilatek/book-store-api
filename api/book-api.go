package api

import (
	"net/http"

	"bookstore.com/domain/service"
	"bookstore.com/port/payload"
	"github.com/go-chi/chi"
)

type bookHandler struct {
	authorService service.BookService
}

func NewBookHandler(authorService service.BookService) BookHandler {
	return &bookHandler{
		authorService: authorService,
	}
}

func (h *bookHandler) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := chi.URLParam(r, "id")
	book, err := h.authorService.Find(r.Context(), id)
	if err != nil {
		responseErr(w, err)
		return
	}

	responseJSON(w, http.StatusOK, book)
}

func (h *bookHandler) Post(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	book := &payload.BookRequest{}
	if err := decodeBody(r, book); err != nil {
		responseErr(w, err)
		return
	}

	err := h.authorService.Store(r.Context(), book)
	if err != nil {
		responseErr(w, err)
		return
	}

	responseJSON(w, http.StatusOK, book)
}

func (h *bookHandler) Put(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := chi.URLParam(r, "id")

	book := &payload.BookRequest{}
	if err := decodeBody(r, book); err != nil {
		responseErr(w, err)
		return
	}

	err := h.authorService.Update(r.Context(), id, book)
	if err != nil {
		responseErr(w, err)
		return
	}

	responseJSON(w, http.StatusOK, book)
}

func (h *bookHandler) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := chi.URLParam(r, "id")

	err := h.authorService.Delete(r.Context(), id)
	if err != nil {
		responseErr(w, err)
		return
	}

	responseJSON(w, http.StatusOK, payload.MessageResponse{
		Message: "Deleted book successfully!",
	})
}
func (h *bookHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	authors, err := h.authorService.FindAll(r.Context())
	if err != nil {
		responseErr(w, err)
		return
	}

	responseJSON(w, http.StatusOK, authors)
}
