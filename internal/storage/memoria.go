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
	clientes       map[string]models.Cliente
	devoluciones   map[string]models.Devolucion
	mantenimientos map[string]models.RegistroMantenimiento
}

func NewAlmacenMemoria() *AlmacenMemoria {
	return &AlmacenMemoria{
		piezas:         make(map[string]models.Pieza),
		clientes:       make(map[string]models.Cliente),
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

func (a *AlmacenMemoria) ListarClientes() []models.Cliente {
	a.mu.RLock()
	defer a.mu.RUnlock()
	out := make([]models.Cliente, 0, len(a.clientes))
	for _, c := range a.clientes {
		out = append(out, c)
	}
	return out
}

func (a *AlmacenMemoria) BuscarClientePorID(id string) (models.Cliente, bool) {
	a.mu.RLock()
	defer a.mu.RUnlock()
	c, ok := a.clientes[id]
	return c, ok
}

func (a *AlmacenMemoria) BuscarClientePorCedula(cedula string) (models.Cliente, bool) {
	a.mu.RLock()
	defer a.mu.RUnlock()
	for _, c := range a.clientes {
		if c.Cedula == cedula {
			return c, true
		}
	}
	return models.Cliente{}, false
}

func (a *AlmacenMemoria) CrearCliente(c models.Cliente) (models.Cliente, error) {
	a.mu.Lock()
	defer a.mu.Unlock()
	for _, existente := range a.clientes {
		if existente.Cedula == c.Cedula {
			return models.Cliente{}, ErrDuplicado
		}
	}
	now := time.Now()
	c.ID = uuid.New().String()
	if c.FechaRegistro.IsZero() {
		c.FechaRegistro = now
	}
	a.clientes[c.ID] = c
	return c, nil
}

func (a *AlmacenMemoria) ActualizarCliente(id string, datos models.Cliente) (models.Cliente, bool) {
	a.mu.Lock()
	defer a.mu.Unlock()
	existente, ok := a.clientes[id]
	if !ok {
		return models.Cliente{}, false
	}
	datos.ID = id
	if datos.FechaRegistro.IsZero() {
		datos.FechaRegistro = existente.FechaRegistro
	}
	a.clientes[id] = datos
	return datos, true
}

func (a *AlmacenMemoria) BorrarCliente(id string) bool {
	a.mu.Lock()
	defer a.mu.Unlock()
	if _, ok := a.clientes[id]; !ok {
		return false
	}
	delete(a.clientes, id)
	return true
}

func (a *AlmacenMemoria) preloadDevolucion(d models.Devolucion) models.Devolucion {
	if p, ok := a.piezas[d.PiezaID]; ok {
		d.Pieza = p
	}
	if c, ok := a.clientes[d.ClienteID]; ok {
		d.Cliente = c
	}
	return d
}

func (a *AlmacenMemoria) preloadMantenimiento(m models.RegistroMantenimiento) models.RegistroMantenimiento {
	if m.PiezaID != "" {
		if p, ok := a.piezas[m.PiezaID]; ok {
			m.Pieza = p
		}
	}
	if c, ok := a.clientes[m.ClienteID]; ok {
		m.Cliente = c
	}
	return m
}

func (a *AlmacenMemoria) ListarDevoluciones() []models.Devolucion {
	a.mu.RLock()
	defer a.mu.RUnlock()
	out := make([]models.Devolucion, 0, len(a.devoluciones))
	for _, d := range a.devoluciones {
		out = append(out, a.preloadDevolucion(d))
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
	return a.preloadDevolucion(d), ok
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
		out = append(out, a.preloadMantenimiento(m))
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
	return a.preloadMantenimiento(m), ok
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
