package models

import "time"

// Usuario es el empleado de Portotech que inicia sesión en el sistema (auth JWT).
// No confundir con Cliente: el cliente es quien compra o deja equipos en el taller.
type Usuario struct {
	ID           int       `json:"id" gorm:"primaryKey"`
	Email        string    `json:"email" gorm:"not null;uniqueIndex"`
	PasswordHash string    `json:"-" gorm:"not null"`
	CreadoEn     time.Time `json:"creado_en" gorm:"autoCreateTime"`
}
