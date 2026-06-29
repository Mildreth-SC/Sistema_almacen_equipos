// MODULO REALIZADO POR MILDRETH GUANOLUISA — Inventario de piezas

package models

import "time"

type EstadoPieza string

const (
	Disponible EstadoPieza = "DISPONIBLE"
	Agotado    EstadoPieza = "AGOTADO"
	Reservado  EstadoPieza = "RESERVADO"
)

type Pieza struct {
	ID           string      `json:"id"             gorm:"primaryKey"`
	NumeroSerial   string      `json:"numero_serial"  gorm:"uniqueIndex;not null"`
	CodigoBarras   string      `json:"codigo_barras"  gorm:"uniqueIndex;not null"`
	Nombre         string      `json:"nombre"         gorm:"not null"`
	Categoria      string      `json:"categoria"`
	Marca          string      `json:"marca"`
	Modelo         string      `json:"modelo"`
	Garantia       int         `json:"garantia_meses"`
	Stock          int         `json:"stock"`
	StockMinimo    int         `json:"stock_minimo"`
	PrecioCompra   float64     `json:"precio_compra"`
	PrecioVenta    float64     `json:"precio_venta"`
	Proveedor      string      `json:"proveedor"`
	Ubicacion      string      `json:"ubicacion"`
	Estado         EstadoPieza `json:"estado"`
	FechaIngreso   time.Time   `json:"fecha_ingreso"`
	CreatedAt      time.Time   `json:"created_at"`
	UpdatedAt      time.Time   `json:"updated_at"`
}
