package storage

import "github.com/Mildreth-SC/Sistema_almacen_equipos/internal/models"

type Almacen interface {
	ListarPiezas() []models.Pieza
	BuscarPiezaPorID(id string) (models.Pieza, bool)
	CrearPieza(p models.Pieza) models.Pieza
	ActualizarPieza(id string, datos models.Pieza) (models.Pieza, bool)
	BorrarPieza(id string) bool
	AjustarStockPieza(id string, delta int) (models.Pieza, error)

	ListarClientes() []models.Cliente
	BuscarClientePorID(id string) (models.Cliente, bool)
	CrearCliente(c models.Cliente) models.Cliente
	ActualizarCliente(id string, datos models.Cliente) (models.Cliente, bool)
	BorrarCliente(id string) bool

	ListarDevoluciones() []models.Devolucion
	BuscarDevolucionPorID(id string) (models.Devolucion, bool)
	CrearDevolucion(d models.Devolucion) (models.Devolucion, error)
	ActualizarDevolucion(id string, datos models.Devolucion) (models.Devolucion, bool, error)
	BorrarDevolucion(id string) bool
	ListarDevolucionesPorCliente(clienteID string) []models.Devolucion

	ListarMantenimientos() []models.RegistroMantenimiento
	BuscarMantenimientoPorID(id string) (models.RegistroMantenimiento, bool)
	CrearMantenimiento(m models.RegistroMantenimiento) models.RegistroMantenimiento
	ActualizarMantenimiento(id string, datos models.RegistroMantenimiento) (models.RegistroMantenimiento, bool)
	BorrarMantenimiento(id string) bool
}
