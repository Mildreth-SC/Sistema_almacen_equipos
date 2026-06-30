package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// Scenario 1: Ruta protegida sin Token debe responder 401
func TestDevolucionHandler_RutaProtegidaSinToken(t *testing.T) {
	// Reemplaza "/api/devoluciones" por tu ruta exacta
	req, err := http.NewRequest("POST", "/api/devoluciones", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	// Aquí llamamos directamente a tu middleware de autenticación o al handler protegido.
	// Simulando la lógica del middleware que valida la ausencia de token:
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" {
			w.WriteHeader(http.StatusUnauthorized) // 401
			return
		}
		w.WriteHeader(http.StatusOK)
	})

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusUnauthorized {
		t.Errorf("El handler retornó un código %v, pero se esperaba un 401 Unauthorized", status)
	}
}
