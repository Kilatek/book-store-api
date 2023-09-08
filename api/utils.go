package api

import (
	"encoding/json"
	"log"
	"net/http"

	portError "bookstore.com/port/error"
	"bookstore.com/port/payload"
)

func responseErr(w http.ResponseWriter, err error) {
	apiErr, ok := err.(*portError.ApiError)
	if ok {
		log.Println(err)
		responseJSON(w, apiErr.Status, &payload.MessageResponse{
			Message: apiErr.Message,
		})
		return
	}

	log.Println(err)
	responseJSON(w, http.StatusInternalServerError, &payload.MessageResponse{
		Message: "Some thing wrong with the server",
	})
}

func responseJSON(w http.ResponseWriter, status int, v interface{}) {
	raw, err := json.Marshal(v)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(status)
	w.Write(raw)
}
