// Tests httptest — José Mieles (mantenimientos)

package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestListarMantenimientos_SinToken_401(t *testing.T) {
	router, _ := nuevoRouterTest(t)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/mantenimientos", nil)
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusUnauthorized {
		t.Fatalf("esperaba 401, obtuvo %d body=%s", rec.Code, rec.Body.String())
	}
}

func TestCrearMantenimiento_ConToken_201(t *testing.T) {
	router, token := nuevoRouterTest(t)

	body := `{
		"cliente_nombre": "Carlos Ruiz",
		"cliente_telefono": "0987654321",
		"equipo_descripcion": "Laptop HP 15, negro",
		"numero_serial": "HP-HTTP-9988",
		"falla_reportada": "No enciende",
		"tipo": "CORRECTIVO",
		"tecnico": "Juan Pérez",
		"costo": 45.00,
		"anticipo": 20.00
	}`
	req := httptest.NewRequest(http.MethodPost, "/api/v1/mantenimientos", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusCreated {
		t.Fatalf("esperaba 201, obtuvo %d body=%s", rec.Code, rec.Body.String())
	}
	if !strings.Contains(rec.Body.String(), "HP-HTTP-9988") {
		t.Fatalf("respuesta debe incluir numero_serial del equipo: %s", rec.Body.String())
	}
}
