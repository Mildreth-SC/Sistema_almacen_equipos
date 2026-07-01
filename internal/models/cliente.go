package models

import "time"

// Cliente representa a una persona que realiza órdenes y/o solicita devoluciones.
type Cliente struct {
	ID            string    `json:"id" gorm:"primaryKey"`
	Nombre        string    `json:"nombre"`
	Cedula        string    `json:"cedula" gorm:"uniqueIndex"`
	Telefono      string    `json:"telefono"`
	Email         string    `json:"email"`
	Direccion     string    `json:"direccion"`
	FechaRegistro time.Time `json:"fecha_registro"`
}
