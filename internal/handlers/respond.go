package handlers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/Mildreth-SC/Sistema_almacen_equipos/internal/service"
)

func RespondJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if data == nil {
		return
	}
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Printf("error codificando JSON: %v", err)
	}
}

func RespondError(w http.ResponseWriter, status int, mensaje string) {
	RespondJSON(w, status, map[string]string{"error": mensaje})
}

func statusDeError(err error) int {
	switch {
	case errors.Is(err, service.ErrNoEncontrado),
		errors.Is(err, service.ErrUsuarioNoEncontrado):
		return http.StatusNotFound
	case errors.Is(err, service.ErrCredencialesInvalidas):
		return http.StatusUnauthorized
	case errors.Is(err, service.ErrUsuarioYaExiste),
		errors.Is(err, service.ErrRegistroDuplicado),
		errors.Is(err, service.ErrCedulaDuplicada):
		return http.StatusConflict
	case errors.Is(err, service.ErrDevolucionYaResuelta),
		errors.Is(err, service.ErrTransicionEstadoInvalida),
		errors.Is(err, service.ErrClienteEnUso):
		return http.StatusConflict
	case errors.Is(err, service.ErrStockInsuficiente),
		errors.Is(err, service.ErrNombreVacio),
		errors.Is(err, service.ErrNumeroSerialVacio),
		errors.Is(err, service.ErrCodigoBarrasVacio),
		errors.Is(err, service.ErrStockNegativo),
		errors.Is(err, service.ErrStockMinimoNegativo),
		errors.Is(err, service.ErrPrecioNegativo),
		errors.Is(err, service.ErrGarantiaNegativa),
		errors.Is(err, service.ErrNombreClienteVacio),
		errors.Is(err, service.ErrCedulaVacia),
		errors.Is(err, service.ErrClienteIDVacio),
		errors.Is(err, service.ErrPiezaIDVacio),
		errors.Is(err, service.ErrNumeroFacturaVacio),
		errors.Is(err, service.ErrMotivoVacio),
		errors.Is(err, service.ErrMotivoInvalido),
		errors.Is(err, service.ErrEstadoInvalido),
		errors.Is(err, service.ErrEstadoDevolucionPendiente),
		errors.Is(err, service.ErrResolucionVacia),
		errors.Is(err, service.ErrAtendidoPorVacio),
		errors.Is(err, service.ErrEquipoDescripcionVacia),
		errors.Is(err, service.ErrFallaReportadaVacia),
		errors.Is(err, service.ErrTecnicoVacio),
		errors.Is(err, service.ErrTipoVacio),
		errors.Is(err, service.ErrTipoInvalido),
		errors.Is(err, service.ErrCostoNegativo),
		errors.Is(err, service.ErrAnticipoInvalido),
		errors.Is(err, service.ErrEmailOContrasenaVacios):
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}
