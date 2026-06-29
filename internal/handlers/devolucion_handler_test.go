// Tests httptest — Ivanna Zamora (devoluciones)

package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestListarDevoluciones_SinToken_401(t *testing.T) {
	router, _ := nuevoRouterTest(t)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/devoluciones", nil)
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusUnauthorized {
		t.Fatalf("esperaba 401, obtuvo %d body=%s", rec.Code, rec.Body.String())
	}
}

func TestCrearDevolucion_ConToken_201(t *testing.T) {
	router, token := nuevoRouterTest(t)

	// Primero crear pieza (devolución requiere pieza_id válido)
	piezaBody := `{
		"numero_serial": "SN-DEV-001",
		"codigo_barras": "BAR-DEV-001",
		"nombre": "RAM 8GB",
		"stock": 3,
		"precio_compra": 20,
		"precio_venta": 35
	}`
	reqPieza := httptest.NewRequest(http.MethodPost, "/api/v1/inventario-piezas", strings.NewReader(piezaBody))
	reqPieza.Header.Set("Content-Type", "application/json")
	reqPieza.Header.Set("Authorization", "Bearer "+token)
	recPieza := httptest.NewRecorder()
	router.ServeHTTP(recPieza, reqPieza)
	if recPieza.Code != http.StatusCreated {
		t.Fatalf("crear pieza: esperaba 201, obtuvo %d %s", recPieza.Code, recPieza.Body.String())
	}

	// Extraer id de la pieza del JSON (buscar "id":")
	piezaJSON := recPieza.Body.String()
	idStart := strings.Index(piezaJSON, `"id":"`) + 6
	idEnd := strings.Index(piezaJSON[idStart:], `"`) + idStart
	piezaID := piezaJSON[idStart:idEnd]

	devBody := `{
		"pieza_id": "` + piezaID + `",
		"cliente_nombre": "María López",
		"cliente_telefono": "0991234567",
		"numero_factura": "FAC-HTTP-001",
		"motivo": "DEFECTUOSO",
		"descripcion": "No funciona"
	}`
	req := httptest.NewRequest(http.MethodPost, "/api/v1/devoluciones", strings.NewReader(devBody))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusCreated {
		t.Fatalf("esperaba 201, obtuvo %d body=%s", rec.Code, rec.Body.String())
	}
	if !strings.Contains(rec.Body.String(), "FAC-HTTP-001") {
		t.Fatalf("respuesta debe incluir numero_factura: %s", rec.Body.String())
	}
}
