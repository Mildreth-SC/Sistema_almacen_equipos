package service

import (
	"errors"
	"testing"

	"github.com/Mildreth-SC/Sistema_almacen_equipos/internal/models"
	"github.com/Mildreth-SC/Sistema_almacen_equipos/internal/storage"
)

func piezaValida() models.Pieza {
	return models.Pieza{
		NumeroSerial: "SN-TEST-001",
		CodigoBarras: "BAR-TEST-001",
		Nombre:       "RAM DDR4 8GB",
		Stock:        10,
		StockMinimo:  2,
		PrecioCompra: 25,
		PrecioVenta:  40,
		Garantia:     12,
	}
}

func TestPiezaService_Crear_ValidaNombre(t *testing.T) {
	svc := NewPiezaService(storage.NewAlmacenMemoria())
	p := piezaValida()
	p.Nombre = ""

	_, err := svc.Crear(p)
	if !errors.Is(err, ErrNombreVacio) {
		t.Fatalf("esperaba ErrNombreVacio, obtuvo %v", err)
	}
}

func TestPiezaService_Crear_DuplicadoSerial(t *testing.T) {
	repo := storage.NewAlmacenMemoria()
	svc := NewPiezaService(repo)

	if _, err := svc.Crear(piezaValida()); err != nil {
		t.Fatalf("primer create: %v", err)
	}

	dup := piezaValida()
	dup.CodigoBarras = "BAR-OTRO"
	_, err := svc.Crear(dup)
	if !errors.Is(err, ErrRegistroDuplicado) {
		t.Fatalf("esperaba ErrRegistroDuplicado, obtuvo %v", err)
	}
}

func TestPiezaService_AjustarStock_Insuficiente(t *testing.T) {
	repo := storage.NewAlmacenMemoria()
	svc := NewPiezaService(repo)

	creada, err := svc.Crear(piezaValida())
	if err != nil {
		t.Fatalf("crear: %v", err)
	}

	_, err = svc.AjustarStock(creada.ID, -100)
	if !errors.Is(err, ErrStockInsuficiente) {
		t.Fatalf("esperaba ErrStockInsuficiente, obtuvo %v", err)
	}
}

func TestPiezaService_AjustarStock_Agotado(t *testing.T) {
	repo := storage.NewAlmacenMemoria()
	svc := NewPiezaService(repo)

	p := piezaValida()
	p.Stock = 3
	creada, err := svc.Crear(p)
	if err != nil {
		t.Fatalf("crear: %v", err)
	}

	actualizada, err := svc.AjustarStock(creada.ID, -3)
	if err != nil {
		t.Fatalf("ajustar: %v", err)
	}
	if actualizada.Estado != models.Agotado {
		t.Fatalf("esperaba AGOTADO, obtuvo %s", actualizada.Estado)
	}
}
