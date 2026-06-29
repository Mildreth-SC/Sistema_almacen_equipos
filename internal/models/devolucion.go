// MODULO REALIZADO POR IVANNA ZAMORA — Devoluciones y garantías

package models

import "time"

type MotivoDevolucion string

const (
	MotivoDefectuoso MotivoDevolucion = "DEFECTUOSO"
	MotivoEquivocado MotivoDevolucion = "EQUIVOCADO"
	MotivoGarantia   MotivoDevolucion = "GARANTIA"
)

type EstadoDevolucion string

const (
	EstadoPendiente EstadoDevolucion = "PENDIENTE"
	EstadoAprobada  EstadoDevolucion = "APROBADA"
	EstadoRechazada EstadoDevolucion = "RECHAZADA"
)

// Devolucion registra cuando un cliente devuelve una pieza comprada.
type Devolucion struct {
	ID              string           `json:"id" gorm:"primaryKey"`
	PiezaID         string           `json:"pieza_id" gorm:"not null;index"`
	Pieza           Pieza            `json:"pieza,omitempty" gorm:"foreignKey:PiezaID;references:ID"`
	ClienteNombre   string           `json:"cliente_nombre"`
	ClienteTelefono string           `json:"cliente_telefono"`
	NumeroFactura   string           `json:"numero_factura"`
	Motivo          MotivoDevolucion `json:"motivo"`
	Descripcion     string           `json:"descripcion"`
	Estado          EstadoDevolucion `json:"estado"`
	Resolucion      string           `json:"resolucion"`
	AtendidoPor     string           `json:"atendido_por"`
	FechaSolicitud  time.Time        `json:"fecha_solicitud"`
	FechaResolucion *time.Time       `json:"fecha_resolucion,omitempty"`
}

func EsMotivoDevolucionValido(m MotivoDevolucion) bool {
	switch m {
	case MotivoDefectuoso, MotivoEquivocado, MotivoGarantia:
		return true
	default:
		return false
	}
}

func EsEstadoDevolucionValido(e EstadoDevolucion) bool {
	switch e {
	case EstadoPendiente, EstadoAprobada, EstadoRechazada:
		return true
	default:
		return false
	}
}
