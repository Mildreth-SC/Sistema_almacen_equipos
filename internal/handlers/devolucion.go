package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Mildreth-SC/Sistema_almacen_equipos/internal/models"
	"github.com/Mildreth-SC/Sistema_almacen_equipos/internal/storage"
	"github.com/go-chi/chi/v5"
)

type DevolucionHandler struct {
	almacen storage.Almacen
}

func NewDevolucionHandler(a storage.Almacen) *DevolucionHandler {
	return &DevolucionHandler{almacen: a}
}

func validarDevolucion(d models.Devolucion) string {
	if d.ClienteNombre == "" {
		return "el nombre del cliente es requerido"
	}
	if d.OrdenID == "" {
		return "el orden_id es requerido"
	}
	if d.Motivo == "" {
		return "el motivo es requerido"
	}
	return ""
}

func (h *DevolucionHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, h.almacen.ListarDevoluciones())
}

func (h *DevolucionHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	d, ok := h.almacen.BuscarDevolucionPorID(id)
	if !ok {
		writeError(w, http.StatusNotFound, "no encontrado")
		return
	}
	writeJSON(w, http.StatusOK, d)
}

func (h *DevolucionHandler) Create(w http.ResponseWriter, r *http.Request) {
	var d models.Devolucion
	if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
		writeError(w, http.StatusBadRequest, "json invalido")
		return
	}
	if msg := validarDevolucion(d); msg != "" {
		writeError(w, http.StatusBadRequest, msg)
		return
	}
	if d.Estado == "" {
		d.Estado = models.EstadoPendiente
	}
	creada := h.almacen.CrearDevolucion(d)
	writeJSON(w, http.StatusCreated, creada)
}

func (h *DevolucionHandler) Update(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var d models.Devolucion
	if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
		writeError(w, http.StatusBadRequest, "json invalido")
		return
	}
	if msg := validarDevolucion(d); msg != "" {
		writeError(w, http.StatusBadRequest, msg)
		return
	}
	actualizada, ok := h.almacen.ActualizarDevolucion(id, d)
	if !ok {
		writeError(w, http.StatusNotFound, "no encontrado")
		return
	}
	writeJSON(w, http.StatusOK, actualizada)
}

func (h *DevolucionHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if !h.almacen.BorrarDevolucion(id) {
		writeError(w, http.StatusNotFound, "no encontrado")
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{"mensaje": "eliminado"})
}
