// MODULO REALIZADO POR JOSÉ MIELES — Mantenimiento de equipos

package models

import "time"

type TipoMantenimiento string

const (
	TipoPreventivo TipoMantenimiento = "PREVENTIVO"
	TipoCorrectivo TipoMantenimiento = "CORRECTIVO"
)

type EstadoMantenimiento string

const (
	MantenimientoPendiente EstadoMantenimiento = "PENDIENTE"
	MantenimientoEnProceso EstadoMantenimiento = "EN_PROCESO"
	MantenimientoListo     EstadoMantenimiento = "LISTO"
	MantenimientoEntregado EstadoMantenimiento = "ENTREGADO"
)

// RegistroMantenimiento registra un equipo que entra al taller a reparar.
type RegistroMantenimiento struct {
	ID                string              `json:"id" gorm:"primaryKey"`
	ClienteID         string              `json:"cliente_id" gorm:"not null;index"`
	Cliente           Cliente             `json:"cliente,omitempty" gorm:"foreignKey:ClienteID;references:ID"`
	PiezaID           string              `json:"pieza_id" gorm:"index"`
	Pieza             Pieza               `json:"pieza,omitempty" gorm:"foreignKey:PiezaID;references:ID"`
	EquipoDescripcion string              `json:"equipo_descripcion"`
	NumeroSerial      string              `json:"numero_serial"`
	FallaReportada    string              `json:"falla_reportada"`
	DiagnosticoPrevio string              `json:"diagnostico_previo"`
	Tipo              TipoMantenimiento   `json:"tipo"`
	Tecnico           string              `json:"tecnico"`
	Costo             float64             `json:"costo"`
	Anticipo          float64             `json:"anticipo"`
	Estado            EstadoMantenimiento `json:"estado"`
	Observaciones     string              `json:"observaciones"`
	FechaIngreso      time.Time           `json:"fecha_ingreso"`
	FechaEstimada     *time.Time          `json:"fecha_estimada,omitempty"`
	FechaEntrega      *time.Time          `json:"fecha_entrega,omitempty"`
}

func EsTipoMantenimientoValido(t TipoMantenimiento) bool {
	switch t {
	case TipoPreventivo, TipoCorrectivo:
		return true
	default:
		return false
	}
}

func EsEstadoMantenimientoValido(e EstadoMantenimiento) bool {
	switch e {
	case MantenimientoPendiente, MantenimientoEnProceso, MantenimientoListo, MantenimientoEntregado:
		return true
	default:
		return false
	}
}

func TransicionMantenimientoValida(actual, nuevo EstadoMantenimiento) bool {
	switch actual {
	case MantenimientoPendiente:
		return nuevo == MantenimientoEnProceso
	case MantenimientoEnProceso:
		return nuevo == MantenimientoListo
	case MantenimientoListo:
		return nuevo == MantenimientoEntregado
	default:
		return false
	}
}
