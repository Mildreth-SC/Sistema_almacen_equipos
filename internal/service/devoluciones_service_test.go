package service

import (
	"testing"

	"github.com/Mildreth-SC/Sistema_almacen_equipos/internal/models"
)

type mockDevolucionRepository struct {
	createCalled bool
}

func (m *mockDevolucionRepository) Create(dev *models.Devolucion) error {
	m.createCalled = true
	return nil
}

func (m *mockDevolucionRepository) GetByID(id uint) (*models.Devolucion, error) {
	return nil, nil
}

func TestRegistrarDevolucion_ReglaNegocioInvalida(t *testing.T) {
	mockRepo := &mockDevolucionRepository{}
	service := NewDevolucionService(mockRepo)

	// Enviamos un valor nulo (nil) para activar a propósito la regla de negocio
	var devInvalida *models.Devolucion = nil

	err := service.RegistrarDevolucion(devInvalida)
	if err == nil {
		t.Error("Se esperaba un error por regla de negocio.")
	}

	if mockRepo.createCalled {
		t.Error("El repositorio fue llamado. Un dato inválido NO debe llegar al repositorio.")
	}
}
