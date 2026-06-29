package storage

import (
	"sync"
	"time"

	"github.com/Mildreth-SC/Sistema_almacen_equipos/internal/models"
	"github.com/google/uuid"
)

// AlmacenMemoria implementa Almacen en memoria para pruebas unitarias.
type AlmacenMemoria struct {
	mu             sync.RWMutex
	piezas         map[string]models.Pieza
	devoluciones   map[string]models.Devolucion
	mantenimientos map[string]models.RegistroMantenimiento
}

func NewAlmacenMemoria() *AlmacenMemoria {
	return &AlmacenMemoria{
		piezas:         make(map[string]models.Pieza),
		devoluciones:   make(map[string]models.Devolucion),
		mantenimientos: make(map[string]models.RegistroMantenimiento),
	}
}

func (a *AlmacenMemoria) ListarPiezas() []models.Pieza {
	a.mu.RLock()
	defer a.mu.RUnlock()
	out := make([]models.Pieza, 0, len(a.piezas))
	for _, p := range a.piezas {
		out = append(out, p)
	}
	return out
}

func (a *AlmacenMemoria) BuscarPiezaPorID(id string) (models.Pieza, bool) {
	a.mu.RLock()
	defer a.mu.RUnlock()
	p, ok := a.piezas[id]
	return p, ok
}

func (a *AlmacenMemoria) CrearPieza(p models.Pieza) (models.Pieza, error) {
	a.mu.Lock()
	defer a.mu.Unlock()
	for _, existente := range a.piezas {
		if existente.NumeroSerial == p.NumeroSerial || existente.CodigoBarras == p.CodigoBarras {
			return models.Pieza{}, ErrDuplicado
		}
	}
	now := time.Now()
	p.ID = uuid.New().String()
	if p.FechaIngreso.IsZero() {
		p.FechaIngreso = now
	}
	p.CreatedAt = now
	p.UpdatedAt = now
	a.piezas[p.ID] = p
	return p, nil
}

func (a *AlmacenMemoria) ActualizarPieza(id string, datos models.Pieza) (models.Pieza, bool) {
	a.mu.Lock()
	defer a.mu.Unlock()
	if _, ok := a.piezas[id]; !ok {
		return models.Pieza{}, false
	}
	datos.ID = id
	datos.UpdatedAt = time.Now()
	a.piezas[id] = datos
	return datos, true
}

func (a *AlmacenMemoria) BorrarPieza(id string) bool {
	a.mu.Lock()
	defer a.mu.Unlock()
	if _, ok := a.piezas[id]; !ok {
		return false
	}
	delete(a.piezas, id)
	return true
}

func (a *AlmacenMemoria) ListarDevoluciones() []models.Devolucion {
	a.mu.RLock()
	defer a.mu.RUnlock()
	out := make([]models.Devolucion, 0, len(a.devoluciones))
	for _, d := range a.devoluciones {
		if p, ok := a.piezas[d.PiezaID]; ok {
			d.Pieza = p
		}
		out = append(out, d)
	}
	return out
}

func (a *AlmacenMemoria) BuscarDevolucionPorID(id string) (models.Devolucion, bool) {
	a.mu.RLock()
	defer a.mu.RUnlock()
	d, ok := a.devoluciones[id]
	if !ok {
		return models.Devolucion{}, false
	}
	if p, ok := a.piezas[d.PiezaID]; ok {
		d.Pieza = p
	}
	return d, ok
}

func (a *AlmacenMemoria) CrearDevolucion(d models.Devolucion) models.Devolucion {
	a.mu.Lock()
	defer a.mu.Unlock()
	d.ID = uuid.New().String()
	if d.FechaSolicitud.IsZero() {
		d.FechaSolicitud = time.Now()
	}
	a.devoluciones[d.ID] = d
	return d
}

func (a *AlmacenMemoria) ActualizarDevolucion(id string, datos models.Devolucion) (models.Devolucion, bool) {
	a.mu.Lock()
	defer a.mu.Unlock()
	if _, ok := a.devoluciones[id]; !ok {
		return models.Devolucion{}, false
	}
	datos.ID = id
	a.devoluciones[id] = datos
	return datos, true
}

func (a *AlmacenMemoria) BorrarDevolucion(id string) bool {
	a.mu.Lock()
	defer a.mu.Unlock()
	if _, ok := a.devoluciones[id]; !ok {
		return false
	}
	delete(a.devoluciones, id)
	return true
}

func (a *AlmacenMemoria) ListarMantenimientos() []models.RegistroMantenimiento {
	a.mu.RLock()
	defer a.mu.RUnlock()
	out := make([]models.RegistroMantenimiento, 0, len(a.mantenimientos))
	for _, m := range a.mantenimientos {
		if m.PiezaID != "" {
			if p, ok := a.piezas[m.PiezaID]; ok {
				m.Pieza = p
			}
		}
		out = append(out, m)
	}
	return out
}

func (a *AlmacenMemoria) BuscarMantenimientoPorID(id string) (models.RegistroMantenimiento, bool) {
	a.mu.RLock()
	defer a.mu.RUnlock()
	m, ok := a.mantenimientos[id]
	if !ok {
		return models.RegistroMantenimiento{}, false
	}
	if m.PiezaID != "" {
		if p, ok := a.piezas[m.PiezaID]; ok {
			m.Pieza = p
		}
	}
	return m, ok
}

func (a *AlmacenMemoria) CrearMantenimiento(m models.RegistroMantenimiento) models.RegistroMantenimiento {
	a.mu.Lock()
	defer a.mu.Unlock()
	m.ID = uuid.New().String()
	if m.FechaIngreso.IsZero() {
		m.FechaIngreso = time.Now()
	}
	a.mantenimientos[m.ID] = m
	return m
}

func (a *AlmacenMemoria) ActualizarMantenimiento(id string, datos models.RegistroMantenimiento) (models.RegistroMantenimiento, bool) {
	a.mu.Lock()
	defer a.mu.Unlock()
	if _, ok := a.mantenimientos[id]; !ok {
		return models.RegistroMantenimiento{}, false
	}
	datos.ID = id
	a.mantenimientos[id] = datos
	return datos, true
}

func (a *AlmacenMemoria) BorrarMantenimiento(id string) bool {
	a.mu.Lock()
	defer a.mu.Unlock()
	if _, ok := a.mantenimientos[id]; !ok {
		return false
	}
	delete(a.mantenimientos, id)
	return true
}

func (a *AlmacenMemoria) Sembrarvacio() {}
