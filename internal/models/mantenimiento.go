package models

import "time"

type TipoMantenimiento string

const (
	TipoPreventivo TipoMantenimiento = "PREVENTIVO"
	TipoCorrectivo TipoMantenimiento = "CORRECTIVO"
	TipoRevision   TipoMantenimiento = "REVISION"
)

type EstadoMantenimiento string

const (
	EstadoProgramado EstadoMantenimiento = "PROGRAMADO"
	EstadoEnProgreso EstadoMantenimiento = "EN_PROGRESO"
	EstadoCompletado EstadoMantenimiento = "COMPLETADO"
	EstadoCancelado  EstadoMantenimiento = "CANCELADO"
)

type RegistroMantenimiento struct {
	ID          string              `json:"id"`
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
