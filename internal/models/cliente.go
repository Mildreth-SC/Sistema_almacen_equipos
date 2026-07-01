package models

import "time"

// Cliente representa a una persona que compra piezas, solicita devoluciones o deja equipos en el taller.
type Cliente struct {
	ID            string    `json:"id" gorm:"primaryKey"`
	Nombre        string    `json:"nombre" gorm:"not null"`
	Cedula        string    `json:"cedula" gorm:"uniqueIndex;not null"`
	Telefono      string    `json:"telefono"`
	Email         string    `json:"email"`
	Direccion     string    `json:"direccion"`
	FechaRegistro time.Time `json:"fecha_registro"`
}
