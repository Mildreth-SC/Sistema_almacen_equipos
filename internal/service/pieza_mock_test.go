// Test con MOCK — Mildreth Guanoluisa (inventario piezas)
// El mock NO guarda; solo registra si se llamó al repositorio.

package service

import (
	"errors"
	"testing"

	"github.com/Mildreth-SC/Sistema_almacen_equipos/internal/models"
)

type mockPiezaRepo struct {
	crearLlamado bool
}

func (m *mockPiezaRepo) ListarPiezas() []models.Pieza                       { return nil }
func (m *mockPiezaRepo) BuscarPiezaPorID(string) (models.Pieza, bool)      { return models.Pieza{}, false }
func (m *mockPiezaRepo) ActualizarPieza(string, models.Pieza) (models.Pieza, bool) {
	return models.Pieza{}, false
}
func (m *mockPiezaRepo) BorrarPieza(string) bool { return false }
func (m *mockPiezaRepo) CrearPieza(models.Pieza) (models.Pieza, error) {
	m.crearLlamado = true
	return models.Pieza{}, nil
}

func TestPiezaService_Mock_NombreVacio_NoLlamaRepositorio(t *testing.T) {
	mock := &mockPiezaRepo{}
	svc := NewPiezaService(mock)

	p := piezaValida()
	p.Nombre = ""

	_, err := svc.Crear(p)
	if !errors.Is(err, ErrNombreVacio) {
		t.Fatalf("esperaba ErrNombreVacio, obtuvo %v", err)
	}
	if mock.crearLlamado {
		t.Fatal("CrearPieza NO debe llamarse cuando el nombre está vacío")
	}
}
