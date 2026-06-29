package models

import "testing"

func TestEsMotivoDevolucionValido(t *testing.T) {
	if !EsMotivoDevolucionValido(MotivoDefectuoso) {
		t.Fatal("DEFECTUOSO debe ser valido")
	}
	if EsMotivoDevolucionValido("OTRO") {
		t.Fatal("OTRO no debe ser valido")
	}
}

func TestTransicionMantenimientoValida(t *testing.T) {
	if !TransicionMantenimientoValida(MantenimientoPendiente, MantenimientoEnProceso) {
		t.Fatal("PENDIENTE → EN_PROCESO debe ser valida")
	}
	if TransicionMantenimientoValida(MantenimientoPendiente, MantenimientoEntregado) {
		t.Fatal("PENDIENTE → ENTREGADO no debe ser valida")
	}
}
