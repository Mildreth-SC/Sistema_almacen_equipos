// Test repositorio GORM :memory: — José Mieles (mantenimientos)

package storage

import (
	"testing"

	"github.com/Mildreth-SC/Sistema_almacen_equipos/internal/models"
)

func TestMantenimientoSQLite_CrearYBuscar(t *testing.T) {
	db := abrirSQLiteMemoria(t)
	almacen := NewAlmacenSQLite(db)

	creado := almacen.CrearMantenimiento(models.RegistroMantenimiento{
		ClienteNombre:     "Luis Méndez",
		EquipoDescripcion: "Desktop Dell",
		FallaReportada:    "Pantalla azul",
		Tipo:              models.TipoCorrectivo,
		Tecnico:           "José Mieles",
		Costo:             50,
		Anticipo:          15,
		Estado:            models.MantenimientoPendiente,
	})
	if creado.ID == "" {
		t.Fatal("id debe generarse")
	}

	encontrado, ok := almacen.BuscarMantenimientoPorID(creado.ID)
	if !ok {
		t.Fatal("mantenimiento no encontrado")
	}
	if encontrado.Tecnico != "José Mieles" {
		t.Fatalf("tecnico: %s", encontrado.Tecnico)
	}
}

func TestMantenimientoSQLite_ListarReflejaCreados(t *testing.T) {
	db := abrirSQLiteMemoria(t)
	almacen := NewAlmacenSQLite(db)

	almacen.CrearMantenimiento(models.RegistroMantenimiento{
		ClienteNombre: "Cliente", EquipoDescripcion: "Laptop",
		FallaReportada: "Lento", Tipo: models.TipoPreventivo,
		Tecnico: "Ana", Costo: 20,
	})

	lista := almacen.ListarMantenimientos()
	if len(lista) != 1 {
		t.Fatalf("esperaba 1 mantenimiento, obtuvo %d", len(lista))
	}
}
