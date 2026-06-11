package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Mildreth-SC/Sistema_almacen_equipos/internal/models"
	"github.com/Mildreth-SC/Sistema_almacen_equipos/internal/storage"
	"github.com/go-chi/chi/v5"
)

type MantenimientoHandler struct {
	almacen storage.Almacen
}

func NewMantenimientoHandler(a storage.Almacen) *MantenimientoHandler {
	return &MantenimientoHandler{almacen: a}
}

// validarMantenimiento verifica que los campos obligatorios sean válidos.
func validarMantenimiento(m models.RegistroMantenimiento) string {
	if m.Descripcion == "" {
		return "la descripcion es requerida"
	}
	if m.Tecnico == "" {
		return "el tecnico es requerido"
	}
	if m.OrdenID == "" {
		return "el orden_id es requerido"
	}
	if m.Costo < 0 {
		return "el costo no puede ser negativo"
	}
	return ""
}

// GetAll devuelve todos los mantenimientos registrados.
func (h *MantenimientoHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, h.almacen.ListarMantenimientos())
}

// GetByID busca un mantenimiento por su identificador.
func (h *MantenimientoHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	m, ok := h.almacen.BuscarMantenimientoPorID(id)
	if !ok {
		writeError(w, http.StatusNotFound, "no encontrado")
		return
	}
	writeJSON(w, http.StatusOK, m)
}

// Create registra un nuevo mantenimiento.
func (h *MantenimientoHandler) Create(w http.ResponseWriter, r *http.Request) {
	var m models.RegistroMantenimiento
	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		writeError(w, http.StatusBadRequest, "json invalido")
		return
	}
	if msg := validarMantenimiento(m); msg != "" {
		writeError(w, http.StatusBadRequest, msg)
		return
	}
	if m.Estado == "" {
		m.Estado = models.EstadoProgramado
	}
	creado := h.almacen.CrearMantenimiento(m)
	writeJSON(w, http.StatusCreated, creado)
}

// Update modifica un mantenimiento existente.
func (h *MantenimientoHandler) Update(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var m models.RegistroMantenimiento
	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		writeError(w, http.StatusBadRequest, "json invalido")
		return
	}
	if msg := validarMantenimiento(m); msg != "" {
		writeError(w, http.StatusBadRequest, msg)
		return
	}
	actualizado, ok := h.almacen.ActualizarMantenimiento(id, m)
	if !ok {
		writeError(w, http.StatusNotFound, "no encontrado")
		return
	}
	writeJSON(w, http.StatusOK, actualizado)
}

// Delete elimina un mantenimiento existente.
func (h *MantenimientoHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if !h.almacen.BorrarMantenimiento(id) {
		writeError(w, http.StatusNotFound, "no encontrado")
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{"mensaje": "eliminado"})
}
