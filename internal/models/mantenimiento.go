package models

import "time"

// Tipos de mantenimiento disponibles en el sistema.
type TipoMantenimiento string

const (
	TipoPreventivo TipoMantenimiento = "PREVENTIVO"
	TipoCorrectivo TipoMantenimiento = "CORRECTIVO"
	TipoRevision   TipoMantenimiento = "REVISION"
)

// Estados posibles de un mantenimiento.
type EstadoMantenimiento string

const (
	EstadoProgramado EstadoMantenimiento = "PROGRAMADO"
	EstadoEnProgreso EstadoMantenimiento = "EN_PROGRESO"
	EstadoCompletado EstadoMantenimiento = "COMPLETADO"
	EstadoCancelado  EstadoMantenimiento = "CANCELADO"
)

// RegistroMantenimiento representa un mantenimiento realizado a un equipo.

type RegistroMantenimiento struct {
	ID          string              `json:"id" gorm:"primaryKey"`
	OrdenID     string              `json:"orden_id"`
	ProductoID  string              `json:"producto_id"`
	Tipo        TipoMantenimiento   `json:"tipo"`
	Descripcion string              `json:"descripcion"`
	Tecnico     string              `json:"tecnico"`
	Costo       float64             `json:"costo"`
	Estado      EstadoMantenimiento `json:"estado"`
	FechaInicio time.Time           `json:"fecha_inicio"`
	FechaFin    *time.Time          `json:"fecha_fin,omitempty"`
}
