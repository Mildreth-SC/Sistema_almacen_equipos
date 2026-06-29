package service

import (
	"errors"
	"testing"

	"github.com/Mildreth-SC/Sistema_almacen_equipos/internal/models"
	"github.com/Mildreth-SC/Sistema_almacen_equipos/internal/storage"
)

func mantenimientoValido() models.RegistroMantenimiento {
	return models.RegistroMantenimiento{
		ClienteNombre:     "Carlos Ruiz",
		ClienteTelefono:   "0987654321",
		EquipoDescripcion: "Laptop HP 15, negro",
		NumeroSerial:      "HP-9988",
		FallaReportada:    "No enciende",
		Tipo:              models.TipoCorrectivo,
		Tecnico:           "Juan Pérez",
		Costo:             45.00,
		Anticipo:          20.00,
	}
}

func TestMantenimientoService_AnticipoInvalido(t *testing.T) {
	svc := NewMantenimientoService(storage.NewAlmacenMemoria(), storage.NewAlmacenMemoria())

	m := mantenimientoValido()
	m.Anticipo = 60

	_, err := svc.Crear(m)
	if !errors.Is(err, ErrAnticipoInvalido) {
		t.Fatalf("esperaba ErrAnticipoInvalido, obtuvo %v", err)
	}
}

func TestMantenimientoService_TipoInvalido(t *testing.T) {
	svc := NewMantenimientoService(storage.NewAlmacenMemoria(), storage.NewAlmacenMemoria())

	m := mantenimientoValido()
	m.Tipo = "REVISION"

	_, err := svc.Crear(m)
	if !errors.Is(err, ErrTipoInvalido) {
		t.Fatalf("esperaba ErrTipoInvalido, obtuvo %v", err)
	}
}

func TestMantenimientoService_SinPiezaID_Ok(t *testing.T) {
	svc := NewMantenimientoService(storage.NewAlmacenMemoria(), storage.NewAlmacenMemoria())

	creado, err := svc.Crear(mantenimientoValido())
	if err != nil {
		t.Fatalf("crear sin pieza_id: %v", err)
	}
	if creado.Estado != models.MantenimientoPendiente {
		t.Fatalf("estado por defecto: %s", creado.Estado)
	}
}

func TestMantenimientoService_CambiarEstado_FlujoCompleto(t *testing.T) {
	svc := NewMantenimientoService(storage.NewAlmacenMemoria(), storage.NewAlmacenMemoria())

	creado, err := svc.Crear(mantenimientoValido())
	if err != nil {
		t.Fatalf("crear: %v", err)
	}

	enProceso, err := svc.CambiarEstado(creado.ID, models.MantenimientoEnProceso)
	if err != nil {
		t.Fatalf("EN_PROCESO: %v", err)
	}
	if enProceso.Estado != models.MantenimientoEnProceso {
		t.Fatalf("estado: %s", enProceso.Estado)
	}

	listo, err := svc.CambiarEstado(creado.ID, models.MantenimientoListo)
	if err != nil {
		t.Fatalf("LISTO: %v", err)
	}
	if listo.Estado != models.MantenimientoListo {
		t.Fatalf("estado: %s", listo.Estado)
	}

	entregado, err := svc.CambiarEstado(creado.ID, models.MantenimientoEntregado)
	if err != nil {
		t.Fatalf("ENTREGADO: %v", err)
	}
	if entregado.FechaEntrega == nil {
		t.Fatal("fecha_entrega debe estar seteada")
	}
}

func TestMantenimientoService_CambiarEstado_TransicionInvalida(t *testing.T) {
	svc := NewMantenimientoService(storage.NewAlmacenMemoria(), storage.NewAlmacenMemoria())

	creado, err := svc.Crear(mantenimientoValido())
	if err != nil {
		t.Fatalf("crear: %v", err)
	}

	// Saltar de PENDIENTE directo a ENTREGADO
	_, err = svc.CambiarEstado(creado.ID, models.MantenimientoEntregado)
	if !errors.Is(err, ErrTransicionEstadoInvalida) {
		t.Fatalf("esperaba ErrTransicionEstadoInvalida, obtuvo %v", err)
	}
}

func TestMantenimientoService_ConPiezaDelStock(t *testing.T) {
	repo := storage.NewAlmacenMemoria()
	piezaSvc := NewPiezaService(repo)
	manSvc := NewMantenimientoService(repo, repo)

	pieza, err := piezaSvc.Crear(piezaValida())
	if err != nil {
		t.Fatalf("crear pieza: %v", err)
	}

	m := mantenimientoValido()
	m.PiezaID = pieza.ID

	creado, err := manSvc.Crear(m)
	if err != nil {
		t.Fatalf("crear mantenimiento con pieza: %v", err)
	}
	if creado.PiezaID != pieza.ID {
		t.Fatalf("pieza_id: %s", creado.PiezaID)
	}
}

func TestMantenimientoService_ListarPorEstado(t *testing.T) {
	svc := NewMantenimientoService(storage.NewAlmacenMemoria(), storage.NewAlmacenMemoria())

	if _, err := svc.Crear(mantenimientoValido()); err != nil {
		t.Fatalf("crear: %v", err)
	}

	pendientes := svc.Listar("PENDIENTE")
	if len(pendientes) != 1 {
		t.Fatalf("esperaba 1 pendiente, obtuvo %d", len(pendientes))
	}
}
