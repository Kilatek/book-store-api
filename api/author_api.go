package api

import (
	"net/http"

	"bookstore.com/domain/service"
	"bookstore.com/port/payload"
	"github.com/go-chi/chi"
)

type authorHandler struct {
	authorService service.AuthorService
}

func NewAuthorHandler(authorService service.AuthorService) AuthorHandler {
	return &authorHandler{
		authorService: authorService,
	}
}

func (h *authorHandler) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := chi.URLParam(r, "id")
	author, err := h.authorService.Find(r.Context(), id)
	if err != nil {
		responseErr(w, err)
		return
	}

	responseJSON(w, http.StatusOK, author)
}

func (h *authorHandler) Post(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	author := &payload.AuthorRequest{}
	if err := decodeBody(r, author); err != nil {
		responseErr(w, err)
		return
	}

	err := h.authorService.Store(r.Context(), author)
	if err != nil {
		responseErr(w, err)
		return
	}

	responseJSON(w, http.StatusOK, author)
}

func (h *authorHandler) Put(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := chi.URLParam(r, "id")

	author := &payload.AuthorRequest{}
	if err := decodeBody(r, author); err != nil {
		responseErr(w, err)
		return
	}

	err := h.authorService.Update(r.Context(), id, author)
	if err != nil {
		responseErr(w, err)
		return
	}

	responseJSON(w, http.StatusOK, author)
}

func (h *authorHandler) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := chi.URLParam(r, "id")

	err := h.authorService.Delete(r.Context(), id)
	if err != nil {
		responseErr(w, err)
		return
	}

	responseJSON(w, http.StatusOK, payload.MessageResponse{
		Message: "Deleted author successfully!",
	})
}
func (h *authorHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	authors, err := h.authorService.FindAll(r.Context())
	if err != nil {
		responseErr(w, err)
		return
	}

	responseJSON(w, http.StatusOK, authors)
}
