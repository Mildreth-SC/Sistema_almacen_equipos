package storage

import (
	"errors"
	"sync"

	"github.com/Mildreth-SC/Sistema_almacen_equipos/internal/models"
	"github.com/google/uuid"
)

var (
	ErrPiezaNoEncontrada = errors.New("pieza no encontrada")
	ErrStockInsuficiente = errors.New("stock insuficiente")
)

type InventarioStorage struct {
	mu     sync.RWMutex
	piezas map[string]models.Pieza
}

func NewInventarioStorage() *InventarioStorage {
	return &InventarioStorage{
		piezas: make(map[string]models.Pieza),
	}
}

func (s *InventarioStorage) GetAll() []models.Pieza {
	s.mu.RLock()
	defer s.mu.RUnlock()

	resultado := make([]models.Pieza, 0, len(s.piezas))
	for _, p := range s.piezas {
		resultado = append(resultado, p)
	}
	return resultado
}

func (s *InventarioStorage) GetByID(id string) (models.Pieza, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	p, ok := s.piezas[id]
	return p, ok
}

func (s *InventarioStorage) Create(p models.Pieza) models.Pieza {
	s.mu.Lock()
	defer s.mu.Unlock()

	p.ID = uuid.New().String()
	s.piezas[p.ID] = p
	return p
}

func (s *InventarioStorage) Update(id string, p models.Pieza) (models.Pieza, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.piezas[id]; !ok {
		return models.Pieza{}, false
	}

	p.ID = id
	s.piezas[id] = p
	return p, true
}

func (s *InventarioStorage) Delete(id string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.piezas[id]; !ok {
		return false
	}

	delete(s.piezas, id)
	return true
}

func (s *InventarioStorage) AjustarStock(id string, delta int) (models.Pieza, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	p, ok := s.piezas[id]
	if !ok {
		return models.Pieza{}, ErrPiezaNoEncontrada
	}

	nuevoStock := p.Stock + delta
	if nuevoStock < 0 {
		return models.Pieza{}, ErrStockInsuficiente
	}

	p.Stock = nuevoStock
	s.piezas[id] = p
	return p, nil
}
