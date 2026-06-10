package models

type EstadoPieza string

const (
	Disponible EstadoPieza = "DISPONIBLE"
	Agotado    EstadoPieza = "AGOTADO"
	Reservado  EstadoPieza = "RESERVADO"
)

type Pieza struct {
	ID          string      `json:"id" gorm:"primaryKey"`
	Nombre      string      `json:"nombre"`
	Categoria   string      `json:"categoria"`
	Marca       string      `json:"marca"`
	ModeloComp  string      `json:"modelo_compatible"`
	Stock       int         `json:"stock"`
	StockMinimo int         `json:"stock_minimo"`
	PrecioUnit  float64     `json:"precio_unitario"`
	Proveedor   string      `json:"proveedor"`
	Ubicacion   string      `json:"ubicacion"`
	Estado      EstadoPieza `json:"estado"`
}
