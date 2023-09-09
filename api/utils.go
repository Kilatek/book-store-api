package api

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"

	portError "bookstore.com/port/error"
	"bookstore.com/port/payload"
)

func decodeBody(r *http.Request, v interface{}) error {
	err := json.NewDecoder(r.Body).Decode(&v)
	if err != nil {
		if errors.Is(err, io.EOF) || errors.Is(err, io.ErrUnexpectedEOF) {
			return portError.NewBadRequestError(err.Error(), nil)
		}
		return err
	}

	return nil
}

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
