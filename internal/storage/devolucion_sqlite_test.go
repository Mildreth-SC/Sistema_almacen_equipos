// Test repositorio GORM :memory: — Ivanna Zamora (devoluciones)

package storage

import (
	"testing"

	"github.com/Mildreth-SC/Sistema_almacen_equipos/internal/models"
)

func TestDevolucionSQLite_CrearYBuscar(t *testing.T) {
	db := abrirSQLiteMemoria(t)
	almacen := NewAlmacenSQLite(db)

	pieza, err := almacen.CrearPieza(models.Pieza{
		NumeroSerial: "SN-DEV-GORM",
		CodigoBarras: "BAR-DEV-GORM",
		Nombre:       "Pantalla 15",
		Stock:        2,
	})
	if err != nil {
		t.Fatalf("crear pieza: %v", err)
	}

	creada := almacen.CrearDevolucion(models.Devolucion{
		PiezaID:       pieza.ID,
		ClienteID:     crearClienteSQLite(t, almacen).ID,
		NumeroFactura: "FAC-GORM-001",
		Motivo:        models.MotivoGarantia,
		Descripcion:   "Pixeles muertos",
		Estado:        models.EstadoPendiente,
	})
	if creada.ID == "" {
		t.Fatal("id debe generarse")
	}

	encontrada, ok := almacen.BuscarDevolucionPorID(creada.ID)
	if !ok {
		t.Fatal("devolución no encontrada")
	}
	if encontrada.NumeroFactura != "FAC-GORM-001" {
		t.Fatalf("factura: %s", encontrada.NumeroFactura)
	}
	if encontrada.Pieza.ID != pieza.ID {
		t.Fatalf("preload pieza: %s", encontrada.Pieza.ID)
	}
}

func TestDevolucionSQLite_ListarReflejaCreadas(t *testing.T) {
	db := abrirSQLiteMemoria(t)
	almacen := NewAlmacenSQLite(db)

	pieza, _ := almacen.CrearPieza(models.Pieza{
		NumeroSerial: "SN-DEV-LIST",
		CodigoBarras: "BAR-DEV-LIST",
		Nombre:       "Cable HDMI",
		Stock:        1,
	})

	almacen.CrearDevolucion(models.Devolucion{
		PiezaID: pieza.ID, ClienteID: crearClienteSQLite(t, almacen).ID,
		NumeroFactura: "FAC-002", Motivo: models.MotivoEquivocado,
	})

	lista := almacen.ListarDevoluciones()
	if len(lista) != 1 {
		t.Fatalf("esperaba 1 devolución, obtuvo %d", len(lista))
	}
}
