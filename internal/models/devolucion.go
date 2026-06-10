package models

import "time"

type MotivoDevolucion string

const (
	MotivoDefectoFabrica    MotivoDevolucion = "DEFECTO_FABRICA"
	MotivoMalFuncionamiento MotivoDevolucion = "MAL_FUNCIONAMIENTO"
	MotivoCambioMente       MotivoDevolucion = "CAMBIO_MENTE"
	MotivoDanoTransporte    MotivoDevolucion = "DAÑO_TRANSPORTE"
	MotivoOtro              MotivoDevolucion = "OTRO"
)

type EstadoDevolucion string

const (
	EstadoPendiente  EstadoDevolucion = "PENDIENTE"
	EstadoEnRevision EstadoDevolucion = "EN_REVISION"
	EstadoAprobada   EstadoDevolucion = "APROBADA"
	EstadoRechazada  EstadoDevolucion = "RECHAZADA"
	EstadoResuelta   EstadoDevolucion = "RESUELTA"
)

type Devolucion struct {
	ID              string           `json:"id" gorm:"primaryKey"`
	OrdenID         string           `json:"orden_id"`
	ProductoID      string           `json:"producto_id"`
	ClienteNombre   string           `json:"cliente_nombre"`
	Motivo          MotivoDevolucion `json:"motivo"`
	Descripcion     string           `json:"descripcion"`
	Estado          EstadoDevolucion `json:"estado"`
	FechaSolicitud  time.Time        `json:"fecha_solicitud"`
	FechaResolucion *time.Time       `json:"fecha_resolucion,omitempty"`
}
