package storage

import (
	"errors"
	"time"

	"github.com/Mildreth-SC/Sistema_almacen_equipos/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

var (
	ErrPiezaNoEncontrada = errors.New("pieza no encontrada")
	ErrStockInsuficiente = errors.New("stock insuficiente")
)

type AlmacenSQLite struct {
	db *gorm.DB
}

func NewAlmacenSQLite(db *gorm.DB) *AlmacenSQLite {
	return &AlmacenSQLite{db: db}
}
// MODULO REALIZO POR MILDRETH GUANOLUISA
// --- Piezas ---

func (a *AlmacenSQLite) ListarPiezas() []models.Pieza {
	var piezas []models.Pieza
	a.db.Find(&piezas)
	return piezas
}

func (a *AlmacenSQLite) BuscarPiezaPorID(id string) (models.Pieza, bool) {
	var pieza models.Pieza
	if err := a.db.First(&pieza, "id = ?", id).Error; err != nil {
		return models.Pieza{}, false
	}
	return pieza, true
}

func (a *AlmacenSQLite) CrearPieza(p models.Pieza) models.Pieza {
	p.ID = uuid.New().String()
	if err := a.db.Create(&p).Error; err != nil {
		return models.Pieza{}
	}
	return p
}

func (a *AlmacenSQLite) ActualizarPieza(id string, datos models.Pieza) (models.Pieza, bool) {
	var existente models.Pieza
	if err := a.db.First(&existente, "id = ?", id).Error; err != nil {
		return models.Pieza{}, false
	}
	datos.ID = id
	a.db.Save(&datos)
	return datos, true
}

func (a *AlmacenSQLite) BorrarPieza(id string) bool {
	if err := a.db.Delete(&models.Pieza{}, "id = ?", id).Error; err != nil {
		return false
	}
	return true
}

func (a *AlmacenSQLite) AjustarStockPieza(id string, delta int) (models.Pieza, error) {
	var pieza models.Pieza
	if err := a.db.First(&pieza, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.Pieza{}, ErrPiezaNoEncontrada
		}
		return models.Pieza{}, err
	}
	nuevoStock := pieza.Stock + delta
	if nuevoStock < 0 {
		return models.Pieza{}, ErrStockInsuficiente
	}
	pieza.Stock = nuevoStock
	a.db.Save(&pieza)
	return pieza, nil
}

// --- Devoluciones ---

func (a *AlmacenSQLite) ListarDevoluciones() []models.Devolucion {
	var lista []models.Devolucion
	a.db.Find(&lista)
	return lista
}

func (a *AlmacenSQLite) BuscarDevolucionPorID(id string) (models.Devolucion, bool) {
	var d models.Devolucion
	if err := a.db.First(&d, "id = ?", id).Error; err != nil {
		return models.Devolucion{}, false
	}
	return d, true
}

func (a *AlmacenSQLite) CrearDevolucion(d models.Devolucion) models.Devolucion {
	d.ID = uuid.New().String()
	if d.FechaSolicitud.IsZero() {
		d.FechaSolicitud = time.Now()
	}
	if err := a.db.Create(&d).Error; err != nil {
		return models.Devolucion{}
	}
	return d
}

func (a *AlmacenSQLite) ActualizarDevolucion(id string, datos models.Devolucion) (models.Devolucion, bool) {
	var existente models.Devolucion
	if err := a.db.First(&existente, "id = ?", id).Error; err != nil {
		return models.Devolucion{}, false
	}
	datos.ID = id
	a.db.Save(&datos)
	return datos, true
}

func (a *AlmacenSQLite) BorrarDevolucion(id string) bool {
	if err := a.db.Delete(&models.Devolucion{}, "id = ?", id).Error; err != nil {
		return false
	}
	return true
}

// --- Mantenimientos ---

func (a *AlmacenSQLite) ListarMantenimientos() []models.RegistroMantenimiento {
	var lista []models.RegistroMantenimiento
	a.db.Find(&lista)
	return lista
}

func (a *AlmacenSQLite) BuscarMantenimientoPorID(id string) (models.RegistroMantenimiento, bool) {
	var m models.RegistroMantenimiento
	if err := a.db.First(&m, "id = ?", id).Error; err != nil {
		return models.RegistroMantenimiento{}, false
	}
	return m, true
}

func (a *AlmacenSQLite) CrearMantenimiento(m models.RegistroMantenimiento) models.RegistroMantenimiento {
	m.ID = uuid.New().String()
	if m.FechaInicio.IsZero() {
		m.FechaInicio = time.Now()
	}
	if err := a.db.Create(&m).Error; err != nil {
		return models.RegistroMantenimiento{}
	}
	return m
}

func (a *AlmacenSQLite) ActualizarMantenimiento(id string, datos models.RegistroMantenimiento) (models.RegistroMantenimiento, bool) {
	var existente models.RegistroMantenimiento
	if err := a.db.First(&existente, "id = ?", id).Error; err != nil {
		return models.RegistroMantenimiento{}, false
	}
	datos.ID = id
	a.db.Save(&datos)
	return datos, true
}

func (a *AlmacenSQLite) BorrarMantenimiento(id string) bool {
	if err := a.db.Delete(&models.RegistroMantenimiento{}, "id = ?", id).Error; err != nil {
		return false
	}
	return true
}

func (a *AlmacenSQLite) Sembrarvacio() {
	var countPiezas int64
	a.db.Model(&models.Pieza{}).Count(&countPiezas)
	if countPiezas == 0 {
		piezas := []models.Pieza{
			{ID: uuid.New().String(), Nombre: "Pantalla LCD 15.6", Categoria: "Display", Marca: "Samsung", ModeloComp: "NP300E5A", Stock: 12, StockMinimo: 3, PrecioUnit: 89.99, Proveedor: "TechParts SA", Ubicacion: "Estante A1", Estado: models.Disponible},
			{ID: uuid.New().String(), Nombre: "Teclado laptop", Categoria: "Perifericos", Marca: "Logitech", ModeloComp: "Universal", Stock: 25, StockMinimo: 5, PrecioUnit: 24.50, Proveedor: "Repuestos EC", Ubicacion: "Estante B2", Estado: models.Disponible},
		}
		a.db.Create(&piezas)
	}

	var countDev int64
	a.db.Model(&models.Devolucion{}).Count(&countDev)
	if countDev == 0 {
		devoluciones := []models.Devolucion{
			{ID: uuid.New().String(), OrdenID: "ORD-001", ProductoID: "PROD-101", ClienteNombre: "María López", Motivo: models.MotivoMalFuncionamiento, Descripcion: "Laptop no enciende", Estado: models.EstadoPendiente, FechaSolicitud: time.Now()},
			{ID: uuid.New().String(), OrdenID: "ORD-002", ProductoID: "PROD-205", ClienteNombre: "Carlos Ruiz", Motivo: models.MotivoDefectoFabrica, Descripcion: "Pantalla con pixeles muertos", Estado: models.EstadoEnRevision, FechaSolicitud: time.Now()},
		}
		a.db.Create(&devoluciones)
	}

	var countMan int64
	a.db.Model(&models.RegistroMantenimiento{}).Count(&countMan)
	if countMan == 0 {
		mantenimientos := []models.RegistroMantenimiento{
			{ID: uuid.New().String(), OrdenID: "ORD-010", ProductoID: "PROD-301", Tipo: models.TipoCorrectivo, Descripcion: "Cambio de disco SSD", Tecnico: "Juan Pérez", Costo: 45.00, Estado: models.EstadoCompletado, FechaInicio: time.Now()},
			{ID: uuid.New().String(), OrdenID: "ORD-011", ProductoID: "PROD-402", Tipo: models.TipoPreventivo, Descripcion: "Limpieza interna y pasta térmica", Tecnico: "Ana Torres", Costo: 20.00, Estado: models.EstadoProgramado, FechaInicio: time.Now()},
		}
		a.db.Create(&mantenimientos)
	}
}

var _ Almacen = (*AlmacenSQLite)(nil)
