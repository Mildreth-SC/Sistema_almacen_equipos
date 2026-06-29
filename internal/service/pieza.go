// MODULO REALIZADO POR MILDRETH GUANOLUISA — Lógica de negocio de inventario

package service

import (
	"errors"
	"strings"

	"github.com/Mildreth-SC/Sistema_almacen_equipos/internal/models"
	"github.com/Mildreth-SC/Sistema_almacen_equipos/internal/storage"
)

type PiezaService struct {
	repo storage.PiezaRepository
}

func NewPiezaService(repo storage.PiezaRepository) *PiezaService {
	return &PiezaService{repo: repo}
}

func (s *PiezaService) Listar() []models.Pieza {
	return s.repo.ListarPiezas()
}

func (s *PiezaService) Obtener(id string) (models.Pieza, error) {
	p, ok := s.repo.BuscarPiezaPorID(id)
	if !ok {
		return models.Pieza{}, ErrNoEncontrado
	}
	return p, nil
}

func (s *PiezaService) Crear(p models.Pieza) (models.Pieza, error) {
	if err := validarPieza(p); err != nil {
		return models.Pieza{}, err
	}
	aplicarEstadoStock(&p)
	creada, err := s.repo.CrearPieza(p)
	if err != nil {
		if errors.Is(err, storage.ErrDuplicado) {
			return models.Pieza{}, ErrRegistroDuplicado
		}
		return models.Pieza{}, err
	}
	return creada, nil
}

func (s *PiezaService) Actualizar(id string, p models.Pieza) (models.Pieza, error) {
	if err := validarPieza(p); err != nil {
		return models.Pieza{}, err
	}
	aplicarEstadoStock(&p)
	actualizada, ok := s.repo.ActualizarPieza(id, p)
	if !ok {
		return models.Pieza{}, ErrNoEncontrado
	}
	return actualizada, nil
}

func (s *PiezaService) Borrar(id string) error {
	if !s.repo.BorrarPieza(id) {
		return ErrNoEncontrado
	}
	return nil
}

func (s *PiezaService) AjustarStock(id string, delta int) (models.Pieza, error) {
	p, ok := s.repo.BuscarPiezaPorID(id)
	if !ok {
		return models.Pieza{}, ErrNoEncontrado
	}
	nuevoStock := p.Stock + delta
	if nuevoStock < 0 {
		return models.Pieza{}, ErrStockInsuficiente
	}
	p.Stock = nuevoStock
	aplicarEstadoStock(&p)
	actualizada, ok := s.repo.ActualizarPieza(id, p)
	if !ok {
		return models.Pieza{}, ErrNoEncontrado
	}
	return actualizada, nil
}

func aplicarEstadoStock(p *models.Pieza) {
	if p.Estado == models.Reservado {
		return
	}
	if p.Stock == 0 {
		p.Estado = models.Agotado
	} else if p.Estado == "" || p.Estado == models.Agotado {
		p.Estado = models.Disponible
	}
}

func validarPieza(p models.Pieza) error {
	if strings.TrimSpace(p.Nombre) == "" {
		return ErrNombreVacio
	}
	if strings.TrimSpace(p.NumeroSerial) == "" {
		return ErrNumeroSerialVacio
	}
	if strings.TrimSpace(p.CodigoBarras) == "" {
		return ErrCodigoBarrasVacio
	}
	if p.Stock < 0 {
		return ErrStockNegativo
	}
	if p.StockMinimo < 0 {
		return ErrStockMinimoNegativo
	}
	if p.PrecioCompra < 0 || p.PrecioVenta < 0 {
		return ErrPrecioNegativo
	}
	if p.Garantia < 0 {
		return ErrGarantiaNegativa
	}
	return nil
}
