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
		&models.Cliente{},
		&models.Devolucion{},
		&models.RegistroMantenimiento{},
	); err != nil {
		t.Fatalf("migrate: %v", err)
	}
	return db
}

func crearClienteSQLite(t *testing.T, almacen *AlmacenSQLite) models.Cliente {
	t.Helper()
	c, err := almacen.CrearCliente(models.Cliente{
		Nombre: "Cliente Test",
		Cedula: "0999999999",
	})
	if err != nil {
		t.Fatalf("crear cliente: %v", err)
	}
	return c
}
