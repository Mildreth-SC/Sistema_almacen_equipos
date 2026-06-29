// Test con MOCK — Ivanna Zamora (devoluciones)

package service

import (
	"errors"
	"testing"

	"github.com/Mildreth-SC/Sistema_almacen_equipos/internal/models"
)

type mockDevolucionRepo struct {
	crearLlamado bool
}

func (m *mockDevolucionRepo) ListarDevoluciones() []models.Devolucion { return nil }
func (m *mockDevolucionRepo) BuscarDevolucionPorID(string) (models.Devolucion, bool) {
	return models.Devolucion{}, false
}
func (m *mockDevolucionRepo) ActualizarDevolucion(string, models.Devolucion) (models.Devolucion, bool) {
	return models.Devolucion{}, false
}
func (m *mockDevolucionRepo) BorrarDevolucion(string) bool { return false }
func (m *mockDevolucionRepo) CrearDevolucion(models.Devolucion) models.Devolucion {
	m.crearLlamado = true
	return models.Devolucion{}
}

type mockPiezaRepoDevolucion struct {
	id string
}

func (m *mockPiezaRepoDevolucion) ListarPiezas() []models.Pieza { return nil }
func (m *mockPiezaRepoDevolucion) BuscarPiezaPorID(id string) (models.Pieza, bool) {
	if id == m.id {
		return models.Pieza{ID: id, Nombre: "Pieza test"}, true
	}
	return models.Pieza{}, false
}
func (m *mockPiezaRepoDevolucion) CrearPieza(models.Pieza) (models.Pieza, error) {
	return models.Pieza{}, nil
}
func (m *mockPiezaRepoDevolucion) ActualizarPieza(string, models.Pieza) (models.Pieza, bool) {
	return models.Pieza{}, false
}
func (m *mockPiezaRepoDevolucion) BorrarPieza(string) bool { return false }

func TestDevolucionService_Mock_MotivoInvalido_NoLlamaRepositorio(t *testing.T) {
	devMock := &mockDevolucionRepo{}
	piezaMock := &mockPiezaRepoDevolucion{id: "pieza-123"}
	svc := NewDevolucionService(devMock, piezaMock)

	_, err := svc.Crear(models.Devolucion{
		PiezaID:       "pieza-123",
		ClienteNombre: "María López",
		NumeroFactura: "FAC-001",
		Motivo:        "MAL_FUNCIONAMIENTO", // motivo viejo, inválido
	})
	if !errors.Is(err, ErrMotivoInvalido) {
		t.Fatalf("esperaba ErrMotivoInvalido, obtuvo %v", err)
	}
	if devMock.crearLlamado {
		t.Fatal("CrearDevolucion NO debe llamarse cuando el motivo es inválido")
	}
}
