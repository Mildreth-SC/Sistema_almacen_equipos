package service

import (
	"errors"
	"testing"

	"github.com/Mildreth-SC/Sistema_almacen_equipos/internal/models"
	"github.com/Mildreth-SC/Sistema_almacen_equipos/internal/storage"
)

func setupDevolucion(t *testing.T) (*DevolucionService, models.Pieza) {
	t.Helper()
	repo := storage.NewAlmacenMemoria()
	piezaSvc := NewPiezaService(repo)
	devSvc := NewDevolucionService(repo, repo)

	creada, err := piezaSvc.Crear(piezaValida())
	if err != nil {
		t.Fatalf("crear pieza: %v", err)
	}
	return devSvc, creada
}

func devolucionValida(piezaID string) models.Devolucion {
	return models.Devolucion{
		PiezaID:         piezaID,
		ClienteNombre:   "María López",
		ClienteTelefono: "0991234567",
		NumeroFactura:   "FAC-2024-001",
		Motivo:          models.MotivoDefectuoso,
		Descripcion:     "Pieza no funciona",
	}
}

func TestDevolucionService_Crear_RequiereFactura(t *testing.T) {
	devSvc, pieza := setupDevolucion(t)
	d := devolucionValida(pieza.ID)
	d.NumeroFactura = ""

	_, err := devSvc.Crear(d)
	if !errors.Is(err, ErrNumeroFacturaVacio) {
		t.Fatalf("esperaba ErrNumeroFacturaVacio, obtuvo %v", err)
	}
}

func TestDevolucionService_Crear_MotivoInvalido(t *testing.T) {
	devSvc, pieza := setupDevolucion(t)
	d := devolucionValida(pieza.ID)
	d.Motivo = "OTRO"

	_, err := devSvc.Crear(d)
	if !errors.Is(err, ErrMotivoInvalido) {
		t.Fatalf("esperaba ErrMotivoInvalido, obtuvo %v", err)
	}
}

func TestDevolucionService_RequierePiezaExistente(t *testing.T) {
	devSvc, _ := setupDevolucion(t)

	_, err := devSvc.Crear(models.Devolucion{
		PiezaID:       "id-inexistente",
		ClienteNombre: "Cliente Test",
		NumeroFactura: "FAC-002",
		Motivo:        models.MotivoDefectuoso,
	})
	if !errors.Is(err, ErrNoEncontrado) {
		t.Fatalf("esperaba ErrNoEncontrado, obtuvo %v", err)
	}
}

func TestDevolucionService_CambiarEstado_Aprobar(t *testing.T) {
	devSvc, pieza := setupDevolucion(t)

	creada, err := devSvc.Crear(devolucionValida(pieza.ID))
	if err != nil {
		t.Fatalf("crear: %v", err)
	}

	resuelta, err := devSvc.CambiarEstado(creada.ID, models.EstadoAprobada, "cambio", "Ivanna Z.")
	if err != nil {
		t.Fatalf("cambiar estado: %v", err)
	}
	if resuelta.Estado != models.EstadoAprobada {
		t.Fatalf("estado: %s", resuelta.Estado)
	}
	if resuelta.FechaResolucion == nil {
		t.Fatal("fecha_resolucion debe estar seteada")
	}
}

func TestDevolucionService_CambiarEstado_YaResuelta(t *testing.T) {
	devSvc, pieza := setupDevolucion(t)

	creada, err := devSvc.Crear(devolucionValida(pieza.ID))
	if err != nil {
		t.Fatalf("crear: %v", err)
	}

	_, err = devSvc.CambiarEstado(creada.ID, models.EstadoAprobada, "reembolso", "Ivanna Z.")
	if err != nil {
		t.Fatalf("primera resolucion: %v", err)
	}

	_, err = devSvc.CambiarEstado(creada.ID, models.EstadoRechazada, "rechazo", "Ivanna Z.")
	if !errors.Is(err, ErrDevolucionYaResuelta) {
		t.Fatalf("esperaba ErrDevolucionYaResuelta, obtuvo %v", err)
	}
}

func TestDevolucionService_ListarPorEstado(t *testing.T) {
	devSvc, pieza := setupDevolucion(t)

	if _, err := devSvc.Crear(devolucionValida(pieza.ID)); err != nil {
		t.Fatalf("crear: %v", err)
	}

	pendientes := devSvc.Listar("PENDIENTE")
	if len(pendientes) != 1 {
		t.Fatalf("esperaba 1 pendiente, obtuvo %d", len(pendientes))
	}

	aprobados := devSvc.Listar("APROBADA")
	if len(aprobados) != 0 {
		t.Fatalf("esperaba 0 aprobados, obtuvo %d", len(aprobados))
	}
}
