// Test repositorio GORM :memory: — Mildreth Guanoluisa (inventario piezas)

package storage

import (
	"testing"

	"github.com/Mildreth-SC/Sistema_almacen_equipos/internal/models"
)

func TestPiezaSQLite_CrearYBuscar(t *testing.T) {
	db := abrirSQLiteMemoria(t)
	almacen := NewAlmacenSQLite(db)

	creada, err := almacen.CrearPieza(models.Pieza{
		NumeroSerial: "SN-GORM-001",
		CodigoBarras: "BAR-GORM-001",
		Nombre:       "Teclado mecánico",
		Stock:        10,
		PrecioCompra: 15,
		PrecioVenta:  25,
	})
	if err != nil {
		t.Fatalf("crear: %v", err)
	}
	if creada.ID == "" {
		t.Fatal("id debe generarse")
	}

	encontrada, ok := almacen.BuscarPiezaPorID(creada.ID)
	if !ok {
		t.Fatal("pieza no encontrada después de crear")
	}
	if encontrada.Nombre != "Teclado mecánico" {
		t.Fatalf("nombre: %s", encontrada.Nombre)
	}
}

func TestPiezaSQLite_ListarReflejaCreadas(t *testing.T) {
	db := abrirSQLiteMemoria(t)
	almacen := NewAlmacenSQLite(db)

	if _, err := almacen.CrearPieza(models.Pieza{
		NumeroSerial: "SN-GORM-002",
		CodigoBarras: "BAR-GORM-002",
		Nombre:       "Mouse USB",
		Stock:        5,
	}); err != nil {
		t.Fatalf("crear: %v", err)
	}

	lista := almacen.ListarPiezas()
	if len(lista) != 1 {
		t.Fatalf("esperaba 1 pieza, obtuvo %d", len(lista))
	}
}
