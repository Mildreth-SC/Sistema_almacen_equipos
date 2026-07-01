package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Mildreth-SC/Sistema_almacen_equipos/internal/models"
	"github.com/Mildreth-SC/Sistema_almacen_equipos/internal/storage"
	"github.com/go-chi/chi/v5"
)

type ClienteHandler struct {
	almacen storage.Almacen
}

func NewClienteHandler(a storage.Almacen) *ClienteHandler {
	return &ClienteHandler{almacen: a}
}

func validarCliente(c models.Cliente) string {
	if c.Nombre == "" {
		return "el nombre del cliente es requerido"
	}
	if c.Cedula == "" {
		return "la cedula del cliente es requerida"
	}
	return ""
}

func (h *ClienteHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, h.almacen.ListarClientes())
}

func (h *ClienteHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	c, ok := h.almacen.BuscarClientePorID(id)
	if !ok {
		writeError(w, http.StatusNotFound, "cliente no encontrado")
		return
	}
	writeJSON(w, http.StatusOK, c)
}

func (h *ClienteHandler) Create(w http.ResponseWriter, r *http.Request) {
	var c models.Cliente
	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		writeError(w, http.StatusBadRequest, "json invalido")
		return
	}
	if msg := validarCliente(c); msg != "" {
		writeError(w, http.StatusBadRequest, msg)
		return
	}
	creado := h.almacen.CrearCliente(c)
	writeJSON(w, http.StatusCreated, creado)
}

func (h *ClienteHandler) Update(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var c models.Cliente
	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		writeError(w, http.StatusBadRequest, "json invalido")
		return
	}
	if msg := validarCliente(c); msg != "" {
		writeError(w, http.StatusBadRequest, msg)
		return
	}
	actualizado, ok := h.almacen.ActualizarCliente(id, c)
	if !ok {
		writeError(w, http.StatusNotFound, "cliente no encontrado")
		return
	}
	writeJSON(w, http.StatusOK, actualizado)
}

func (h *ClienteHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if !h.almacen.BorrarCliente(id) {
		writeError(w, http.StatusNotFound, "cliente no encontrado")
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{"mensaje": "eliminado"})
}

// GetDevoluciones devuelve el historial de devoluciones de un cliente puntual.
// Conecta el modulo de Clientes con el de Devoluciones.
func (h *ClienteHandler) GetDevoluciones(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if _, ok := h.almacen.BuscarClientePorID(id); !ok {
		writeError(w, http.StatusNotFound, "cliente no encontrado")
		return
	}
	writeJSON(w, http.StatusOK, h.almacen.ListarDevolucionesPorCliente(id))
}
