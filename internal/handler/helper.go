package handler

import (
	"encoding/json"
	"log"
	"net/http"
)

func (h *Handler) handleError(w http.ResponseWriter, statusCode int, err interface{}) {
	log.Println(err, "handleError")

	errorResponse := map[string]interface{}{
		"error": err,
	}

	h.writeJSONResponse(w, statusCode, errorResponse)
}

func (h *Handler) writeJSONResponse(w http.ResponseWriter, statusCode int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	err := json.NewEncoder(w).Encode(payload)
	if err != nil {
		h.handleError(w, http.StatusInternalServerError, "Failed to encode JSON response")
		return
	}
}