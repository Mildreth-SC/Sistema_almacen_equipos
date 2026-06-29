package storage

import "github.com/Mildreth-SC/Sistema_almacen_equipos/internal/models"

// PiezaRepository — inventario de piezas (Mildreth).
type PiezaRepository interface {
	ListarPiezas() []models.Pieza
	BuscarPiezaPorID(id string) (models.Pieza, bool)
	CrearPieza(p models.Pieza) (models.Pieza, error)
	ActualizarPieza(id string, datos models.Pieza) (models.Pieza, bool)
	BorrarPieza(id string) bool
}

// DevolucionRepository — devoluciones (Ivanna).
type DevolucionRepository interface {
	ListarDevoluciones() []models.Devolucion
	BuscarDevolucionPorID(id string) (models.Devolucion, bool)
	CrearDevolucion(d models.Devolucion) models.Devolucion
	ActualizarDevolucion(id string, datos models.Devolucion) (models.Devolucion, bool)
	BorrarDevolucion(id string) bool
}

// MantenimientoRepository — mantenimientos (José).
type MantenimientoRepository interface {
	ListarMantenimientos() []models.RegistroMantenimiento
	BuscarMantenimientoPorID(id string) (models.RegistroMantenimiento, bool)
	CrearMantenimiento(m models.RegistroMantenimiento) models.RegistroMantenimiento
	ActualizarMantenimiento(id string, datos models.RegistroMantenimiento) (models.RegistroMantenimiento, bool)
	BorrarMantenimiento(id string) bool
}

// Almacen agrupa los 3 repositorios.
type Almacen interface {
	PiezaRepository
	DevolucionRepository
	MantenimientoRepository
}

var _ Almacen = (*AlmacenSQLite)(nil)
var _ Almacen = (*AlmacenMemoria)(nil)
