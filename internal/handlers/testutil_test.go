package handlers

import (
	"testing"

	"github.com/Mildreth-SC/Sistema_almacen_equipos/internal/service"
	"github.com/Mildreth-SC/Sistema_almacen_equipos/internal/storage"
	"github.com/go-chi/chi/v5"
)

// nuevoRouterTest arma el servidor con AlmacenMemoria (FAKE): misma interfaz que SQLite,
// responde crear/buscar/listar en RAM — el handler no nota la diferencia en httptest.
func nuevoRouterTest(t *testing.T) (chi.Router, string) {
	t.Helper()

	mem := storage.NewAlmacenMemoria()
	users := storage.NewUsuarioMemoria()
	auth := service.NewAuthService(users)

	if _, err := auth.Registrar("test@portotech.com", "secret123"); err != nil {
		t.Fatalf("registrar usuario test: %v", err)
	}
	token, err := auth.Login("test@portotech.com", "secret123")
	if err != nil {
		t.Fatalf("login test: %v", err)
	}

	srv := NewServer(
		service.NewPiezaService(mem),
		service.NewClienteService(mem, mem, mem),
		service.NewDevolucionService(mem, mem, mem),
		service.NewMantenimientoService(mem, mem, mem),
		auth,
	)
	r := chi.NewRouter()
	srv.RegisterRoutes(r)
	return r, token
}
