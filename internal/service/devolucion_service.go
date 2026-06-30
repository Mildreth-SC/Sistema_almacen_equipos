package service

import (
	"errors"

	"github.com/Mildreth-SC/Sistema_almacen_equipos/internal/models"
)

// Interfaz del repositorio (para poder hacer Mocking)
type DevolucionRepository interface {
	Create(dev *models.Devolucion) error
	GetByID(id uint) (*models.Devolucion, error)
}

type DevolucionService struct {
	repo DevolucionRepository
}

func NewDevolucionService(repo DevolucionRepository) *DevolucionService {
	return &DevolucionService{repo: repo}
}

// Regla de negocio: Registrar devolución con validación real
func (s *DevolucionService) RegistrarDevolucion(dev *models.Devolucion) error {
	// Regla de negocio: Si la devolución no tiene datos o es nula, arroja error
	if dev == nil {
		return errors.New("regla de negocio rota: la devolución no puede ser nula")
	}
	return s.repo.Create(dev)
}
