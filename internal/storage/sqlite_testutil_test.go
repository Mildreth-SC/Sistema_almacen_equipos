package storage

import (
	"testing"

	"github.com/Mildreth-SC/Sistema_almacen_equipos/internal/models"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func abrirSQLiteMemoria(t *testing.T) *gorm.DB {
	t.Helper()
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("abrir sqlite :memory:: %v", err)
	}
	if err := db.AutoMigrate(
		&models.Pieza{},
		&models.Devolucion{},
		&models.RegistroMantenimiento{},
	); err != nil {
		t.Fatalf("migrate: %v", err)
	}
	return db
}
