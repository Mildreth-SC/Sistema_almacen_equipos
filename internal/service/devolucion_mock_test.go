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

type mockClienteRepoDevolucion struct {
	id string
}

func (m *mockClienteRepoDevolucion) ListarClientes() []models.Cliente { return nil }
func (m *mockClienteRepoDevolucion) BuscarClientePorID(id string) (models.Cliente, bool) {
	if id == m.id {
		return models.Cliente{ID: id, Nombre: "Cliente test"}, true
	}
	return models.Cliente{}, false
}
func (m *mockClienteRepoDevolucion) BuscarClientePorCedula(string) (models.Cliente, bool) {
	return models.Cliente{}, false
}
func (m *mockClienteRepoDevolucion) CrearCliente(models.Cliente) (models.Cliente, error) {
	return models.Cliente{}, nil
}
func (m *mockClienteRepoDevolucion) ActualizarCliente(string, models.Cliente) (models.Cliente, bool) {
	return models.Cliente{}, false
}
func (m *mockClienteRepoDevolucion) BorrarCliente(string) bool { return false }

func TestDevolucionService_Mock_MotivoInvalido_NoLlamaRepositorio(t *testing.T) {
	devMock := &mockDevolucionRepo{}
	piezaMock := &mockPiezaRepoDevolucion{id: "pieza-123"}
	clienteMock := &mockClienteRepoDevolucion{id: "cliente-123"}
	svc := NewDevolucionService(devMock, piezaMock, clienteMock)

	_, err := svc.Crear(models.Devolucion{
		PiezaID:       "pieza-123",
		ClienteID:     "cliente-123",
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
