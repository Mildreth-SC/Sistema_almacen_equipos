// Tests httptest — Mildreth Guanoluisa (inventario piezas)

package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestListarPiezas_SinToken_401(t *testing.T) {
	router, _ := nuevoRouterTest(t)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/inventario-piezas", nil)
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusUnauthorized {
		t.Fatalf("esperaba 401, obtuvo %d body=%s", rec.Code, rec.Body.String())
	}
}

func TestCrearPieza_ConToken_201(t *testing.T) {
	router, token := nuevoRouterTest(t)

	body := `{
		"numero_serial": "SN-HTTP-001",
		"codigo_barras": "BAR-HTTP-001",
		"nombre": "SSD 256GB",
		"stock": 5,
		"stock_minimo": 1,
		"precio_compra": 40,
		"precio_venta": 65,
		"garantia_meses": 12
	}`
	req := httptest.NewRequest(http.MethodPost, "/api/v1/inventario-piezas", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusCreated {
		t.Fatalf("esperaba 201, obtuvo %d body=%s", rec.Code, rec.Body.String())
	}
	if !strings.Contains(rec.Body.String(), "SN-HTTP-001") {
		t.Fatalf("respuesta debe incluir numero_serial: %s", rec.Body.String())
	}
}
