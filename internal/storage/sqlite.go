package storage

import (
	"errors"
	"strings"
	"time"

	"github.com/Mildreth-SC/Sistema_almacen_equipos/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AlmacenSQLite struct {
	db *gorm.DB
}

func NewAlmacenSQLite(db *gorm.DB) *AlmacenSQLite {
	return &AlmacenSQLite{db: db}
}

func esDuplicado(err error) bool {
	if err == nil {
		return false
	}
	msg := strings.ToLower(err.Error())
	return errors.Is(err, gorm.ErrDuplicatedKey) ||
		strings.Contains(msg, "unique") ||
		strings.Contains(msg, "duplicate")
}

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

func (a *AlmacenSQLite) CrearPieza(p models.Pieza) (models.Pieza, error) {
	now := time.Now()
	p.ID = uuid.New().String()
	if p.FechaIngreso.IsZero() {
		p.FechaIngreso = now
	}
	p.CreatedAt = now
	p.UpdatedAt = now
	if err := a.db.Create(&p).Error; err != nil {
		if esDuplicado(err) {
			return models.Pieza{}, ErrDuplicado
		}
		return models.Pieza{}, err
	}
	return p, nil
}

func (a *AlmacenSQLite) ActualizarPieza(id string, datos models.Pieza) (models.Pieza, bool) {
	var existente models.Pieza
	if err := a.db.First(&existente, "id = ?", id).Error; err != nil {
		return models.Pieza{}, false
	}
	datos.ID = id
	datos.UpdatedAt = time.Now()
	if err := a.db.Save(&datos).Error; err != nil {
		return models.Pieza{}, false
	}
	return datos, true
}

func (a *AlmacenSQLite) BorrarPieza(id string) bool {
	if err := a.db.Delete(&models.Pieza{}, "id = ?", id).Error; err != nil {
		return false
	}
	return true
}

// --- Devoluciones — Ivanna Zamora ---

func (a *AlmacenSQLite) ListarDevoluciones() []models.Devolucion {
	var lista []models.Devolucion
	a.db.Preload("Pieza").Find(&lista)
	return lista
}

func (a *AlmacenSQLite) BuscarDevolucionPorID(id string) (models.Devolucion, bool) {
	var d models.Devolucion
	if err := a.db.Preload("Pieza").First(&d, "id = ?", id).Error; err != nil {
		return models.Devolucion{}, false
	}
	return d, true
}

func (a *AlmacenSQLite) CrearDevolucion(d models.Devolucion) models.Devolucion {
	d.ID = uuid.New().String()
	d.Pieza = models.Pieza{}
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
	datos.Pieza = models.Pieza{}
	if err := a.db.Save(&datos).Error; err != nil {
		return models.Devolucion{}, false
	}
	return datos, true
}

func (a *AlmacenSQLite) BorrarDevolucion(id string) bool {
	if err := a.db.Delete(&models.Devolucion{}, "id = ?", id).Error; err != nil {
		return false
	}
	return true
}

// --- Mantenimientos — José Mieles ---

func (a *AlmacenSQLite) ListarMantenimientos() []models.RegistroMantenimiento {
	var lista []models.RegistroMantenimiento
	a.db.Preload("Pieza").Find(&lista)
	return lista
}

func (a *AlmacenSQLite) BuscarMantenimientoPorID(id string) (models.RegistroMantenimiento, bool) {
	var m models.RegistroMantenimiento
	if err := a.db.Preload("Pieza").First(&m, "id = ?", id).Error; err != nil {
		return models.RegistroMantenimiento{}, false
	}
	return m, true
}

func (a *AlmacenSQLite) CrearMantenimiento(m models.RegistroMantenimiento) models.RegistroMantenimiento {
	m.ID = uuid.New().String()
	m.Pieza = models.Pieza{}
	if m.FechaIngreso.IsZero() {
		m.FechaIngreso = time.Now()
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
	datos.Pieza = models.Pieza{}
	if err := a.db.Save(&datos).Error; err != nil {
		return models.RegistroMantenimiento{}, false
	}
	return datos, true
}

func (a *AlmacenSQLite) BorrarMantenimiento(id string) bool {
	if err := a.db.Delete(&models.RegistroMantenimiento{}, "id = ?", id).Error; err != nil {
		return false
	}
	return true
}

func (a *AlmacenSQLite) Sembrarvacio() {
	now := time.Now()
	var countPiezas int64
	a.db.Model(&models.Pieza{}).Count(&countPiezas)
	if countPiezas == 0 {
		p1 := models.Pieza{
			ID:           uuid.New().String(),
			NumeroSerial: "SN-SAM-LCD-001",
			CodigoBarras: "BAR-001",
			Nombre:       "Pantalla LCD 15.6",
			Categoria:    "Display",
			Marca:        "Samsung",
			Modelo:       "NP300E5A",
			Garantia:     12,
			Stock:        12,
			StockMinimo:  3,
			PrecioCompra: 65.00,
			PrecioVenta:  89.99,
			Proveedor:    "TechParts SA",
			Ubicacion:    "Estante A1",
			Estado:       models.Disponible,
			FechaIngreso: now,
			CreatedAt:    now,
			UpdatedAt:    now,
		}
		p2 := models.Pieza{
			ID:           uuid.New().String(),
			NumeroSerial: "SN-LOG-KBD-002",
			CodigoBarras: "BAR-002",
			Nombre:       "Teclado laptop",
			Categoria:    "Perifericos",
			Marca:        "Logitech",
			Modelo:       "Universal",
			Garantia:     6,
			Stock:        25,
			StockMinimo:  5,
			PrecioCompra: 18.00,
			PrecioVenta:  24.50,
			Proveedor:    "Repuestos EC",
			Ubicacion:    "Estante B2",
			Estado:       models.Disponible,
			FechaIngreso: now,
			CreatedAt:    now,
			UpdatedAt:    now,
		}
		a.db.Create(&[]models.Pieza{p1, p2})

		devoluciones := []models.Devolucion{
			{
				ID:              uuid.New().String(),
				PiezaID:         p1.ID,
				ClienteNombre:   "María López",
				ClienteTelefono: "0991234567",
				NumeroFactura:   "FAC-2024-001",
				Motivo:          models.MotivoDefectuoso,
				Descripcion:     "Laptop no enciende",
				Estado:          models.EstadoPendiente,
				AtendidoPor:     "Mildreth G.",
				FechaSolicitud:  now,
			},
			{
				ID:              uuid.New().String(),
				PiezaID:         p1.ID,
				ClienteNombre:   "Carlos Ruiz",
				ClienteTelefono: "0987654321",
				NumeroFactura:   "FAC-2024-002",
				Motivo:          models.MotivoGarantia,
				Descripcion:     "Pantalla con pixeles muertos",
				Estado:          models.EstadoAprobada,
				Resolucion:      "cambio",
				AtendidoPor:     "Ivanna Z.",
				FechaSolicitud:  now,
			},
		}
		a.db.Create(&devoluciones)

		mantenimientos := []models.RegistroMantenimiento{
			{
				ID:                uuid.New().String(),
				PiezaID:           p2.ID,
				ClienteNombre:     "Pedro Sánchez",
				ClienteTelefono:   "0991112233",
				EquipoDescripcion: "Laptop HP 15, negro",
				NumeroSerial:      "HP-CLIENTE-9988",
				FallaReportada:    "Teclado no responde",
				DiagnosticoPrevio: "Teclas dañadas por líquido",
				Tipo:              models.TipoCorrectivo,
				Tecnico:           "Juan Pérez",
				Costo:             45.00,
				Anticipo:          20.00,
				Estado:            models.MantenimientoListo,
				Observaciones:     "Reemplazo de teclado completado",
				FechaIngreso:      now,
			},
			{
				ID:                uuid.New().String(),
				ClienteNombre:     "Ana Torres",
				ClienteTelefono:   "0976543210",
				EquipoDescripcion: "Desktop Dell Optiplex",
				NumeroSerial:      "DELL-5544",
				FallaReportada:    "Lentitud general",
				DiagnosticoPrevio: "Polvo acumulado, pasta térmica seca",
				Tipo:              models.TipoPreventivo,
				Tecnico:           "José Mieles",
				Costo:             20.00,
				Anticipo:          10.00,
				Estado:            models.MantenimientoPendiente,
				FechaIngreso:      now,
			},
		}
		a.db.Create(&mantenimientos)
	}
}
