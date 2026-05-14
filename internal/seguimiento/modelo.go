package seguimiento

import "time"

type EstadoOrden string

//valores posibles
const (
	EstadoRecibido     EstadoOrden = "RECIBIDO"
	EstadoDiagnostico  EstadoOrden = "DIAGNOSTICO"
	EstadoEnReparacion EstadoOrden = "EN_REPARACION"
	EstadoListo        EstadoOrden = "LISTO"
	EstadoEntregado    EstadoOrden = "ENTREGADO"
)

type OrdenSoporte struct {
	ID            string      `json:"id"`
	ClienteNombre string      `json:"cliente_nombre"`
	Equipo        string      `json:"equipo"`
	Problema      string      `json:"problema"`
	Estado        EstadoOrden `json:"estado"`
	Tecnico       string      `json:"tecnico"`
	NumOrden      string      `json:"num_orden"`
	Fecha         time.Time   `json:"fecha"`
}
