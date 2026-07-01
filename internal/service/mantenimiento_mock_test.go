// Test con MOCK — José Mieles (mantenimientos)

package service

import (
	"errors"
	"testing"

	"github.com/Mildreth-SC/Sistema_almacen_equipos/internal/models"
)

type mockMantenimientoRepo struct {
	crearLlamado bool
}

func (m *mockMantenimientoRepo) ListarMantenimientos() []models.RegistroMantenimiento { return nil }
func (m *mockMantenimientoRepo) BuscarMantenimientoPorID(string) (models.RegistroMantenimiento, bool) {
	return models.RegistroMantenimiento{}, false
}
func (m *mockMantenimientoRepo) ActualizarMantenimiento(string, models.RegistroMantenimiento) (models.RegistroMantenimiento, bool) {
	return models.RegistroMantenimiento{}, false
}
func (m *mockMantenimientoRepo) BorrarMantenimiento(string) bool { return false }
func (m *mockMantenimientoRepo) CrearMantenimiento(models.RegistroMantenimiento) models.RegistroMantenimiento {
	m.crearLlamado = true
	return models.RegistroMantenimiento{}
}

func TestMantenimientoService_Mock_TipoInvalido_NoLlamaRepositorio(t *testing.T) {
	mock := &mockMantenimientoRepo{}
	clienteMock := &mockClienteRepoDevolucion{id: "cliente-123"}
	svc := NewMantenimientoService(mock, &mockPiezaRepoDevolucion{}, clienteMock)

	_, err := svc.Crear(models.RegistroMantenimiento{
		ClienteID:         "cliente-123",
		EquipoDescripcion: "Laptop HP 15",
		FallaReportada:    "No enciende",
		Tecnico:           "Juan Pérez",
		Tipo:              "REVISION", // inválido: solo PREVENTIVO o CORRECTIVO
		Costo:             45,
	})
	if !errors.Is(err, ErrTipoInvalido) {
		t.Fatalf("esperaba ErrTipoInvalido, obtuvo %v", err)
	}
	if mock.crearLlamado {
		t.Fatal("CrearMantenimiento NO debe llamarse cuando el tipo es inválido")
	}
}
