package storage

import (
	"testing"

	"github.com/Mildreth-SC/Sistema_almacen_equipos/internal/models"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func TestGORM_CrearYBuscarDevolucionEnMemoria(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		t.Fatalf("Error al conectar a la BD en memoria: %v", err)
	}

	err = db.AutoMigrate(&models.Devolucion{})
	if err != nil {
		t.Fatalf("Error en AutoMigrate: %v", err)
	}

	nuevaDevolucion := models.Devolucion{}

	if err := db.Create(&nuevaDevolucion).Error; err != nil {
		t.Fatalf("Error al insertar registro: %v", err)
	}

	var devResultado models.Devolucion
	if err := db.First(&devResultado, nuevaDevolucion.ID).Error; err != nil {
		t.Fatalf("Error al buscar el registro: %v", err)
	}
}
