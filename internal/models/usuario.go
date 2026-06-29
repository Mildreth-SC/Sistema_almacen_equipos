package models

import "time"

type Usuario struct {
	ID           int       `json:"id" gorm:"primaryKey"`
	Email        string    `json:"email" gorm:"not null;uniqueIndex"`
	PasswordHash string    `json:"-" gorm:"not null"`
	CreadoEn     time.Time `json:"creado_en" gorm:"autoCreateTime"`
}
