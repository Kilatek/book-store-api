package api

import (
	"encoding/json"
	"net/http"

	"bookstore.com/domain/service"
	"bookstore.com/port/payload"
)

type userHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) UserHandler {
	return &userHandler{
		userService: userService,
	}
}

func (h *userHandler) Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	user := &payload.RegisterRequest{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		responseErr(w, err)
		return
	}

	err = h.userService.Register(r.Context(), user)
	if err != nil {
		responseErr(w, err)
		return
	}
	response(w, http.StatusOK)
}

func (h *userHandler) Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	user := &payload.LoginRequest{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		responseErr(w, err)
		return
	}

	token, err := h.userService.Login(r.Context(), user)
	if err != nil {
		responseErr(w, err)
		return
	}

	responseJSON(w, http.StatusOK, token)
}
