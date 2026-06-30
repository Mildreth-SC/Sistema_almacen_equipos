package handlers

import (
	"encoding/json"
	"net/http"
)

func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, `{"error":"error al escribir respuesta"}`, http.StatusInternalServerError)
	}
}

func writeError(w http.ResponseWriter, status int, mensaje string) {
	writeJSON(w, status, map[string]string{"error": mensaje})
}
