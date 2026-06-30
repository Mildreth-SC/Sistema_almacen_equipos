package repository

import (
	"github.com/Mildreth-SC/Sistema_almacen_equipos/internal/models"
	"gorm.io/gorm"
)

type DevolucionRepository interface {
	Crear(d models.Devolucion) (models.Devolucion, error)
	Listar() ([]models.Devolucion, error)
	BuscarPorID(id string) (models.Devolucion, error)
}

type devolucionRepository struct {
	db *gorm.DB
}

func NewDevolucionRepository(db *gorm.DB) DevolucionRepository {
	return &devolucionRepository{
		db: db,
	}
}

func (r *devolucionRepository) Crear(d models.Devolucion) (models.Devolucion, error) {

	err := r.db.Create(&d).Error

	return d, err
}

func (r *devolucionRepository) Listar() ([]models.Devolucion, error) {

	var devoluciones []models.Devolucion

	err := r.db.Find(&devoluciones).Error

	return devoluciones, err
}

func (r *devolucionRepository) BuscarPorID(id string) (models.Devolucion, error) {

	var devolucion models.Devolucion

	err := r.db.First(&devolucion, "id = ?", id).Error

	return devolucion, err
}
