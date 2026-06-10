package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/Mildreth-SC/Sistema_almacen_equipos/internal/models"
	"github.com/Mildreth-SC/Sistema_almacen_equipos/internal/storage"
	"github.com/go-chi/chi/v5"
)

type InventarioPiezasHandler struct {
	almacen storage.Almacen
}

func NewInventarioPiezasHandler(a storage.Almacen) *InventarioPiezasHandler {
	return &InventarioPiezasHandler{almacen: a}
}

func validarPieza(p models.Pieza) string {
	if p.Nombre == "" {
		return "el nombre es requerido"
	}
	if p.Stock < 0 {
		return "el stock no puede ser negativo"
	}
	if p.StockMinimo < 0 {
		return "el stock minimo no puede ser negativo"
	}
	if p.PrecioUnit < 0 {
		return "el precio unitario no puede ser negativo"
	}
	return ""
}

func (h *InventarioPiezasHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, h.almacen.ListarPiezas())
}

func (h *InventarioPiezasHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	p, ok := h.almacen.BuscarPiezaPorID(id)
	if !ok {
		writeError(w, http.StatusNotFound, "no encontrado")
		return
	}

	writeJSON(w, http.StatusOK, p)
}

func (h *InventarioPiezasHandler) Create(w http.ResponseWriter, r *http.Request) {
	var p models.Pieza
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		writeError(w, http.StatusBadRequest, "json invalido")
		return
	}

	if msg := validarPieza(p); msg != "" {
		writeError(w, http.StatusBadRequest, msg)
		return
	}

	creada := h.almacen.CrearPieza(p)
	writeJSON(w, http.StatusCreated, creada)
}

func (h *InventarioPiezasHandler) Update(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var p models.Pieza
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		writeError(w, http.StatusBadRequest, "json invalido")
		return
	}

	if msg := validarPieza(p); msg != "" {
		writeError(w, http.StatusBadRequest, msg)
		return
	}

	actualizada, ok := h.almacen.ActualizarPieza(id, p)
	if !ok {
		writeError(w, http.StatusNotFound, "no encontrado")
		return
	}

	writeJSON(w, http.StatusOK, actualizada)
}

func (h *InventarioPiezasHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if !h.almacen.BorrarPieza(id) {
		writeError(w, http.StatusNotFound, "no encontrado")
		return
	}

	writeJSON(w, http.StatusOK, map[string]string{"mensaje": "eliminado"})
}

func (h *InventarioPiezasHandler) AjustarStock(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var body struct {
		Delta int `json:"delta"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		writeError(w, http.StatusBadRequest, "json invalido")
		return
	}

	p, err := h.almacen.AjustarStockPieza(id, body.Delta)
	if err != nil {
		if errors.Is(err, storage.ErrPiezaNoEncontrada) {
			writeError(w, http.StatusNotFound, "no encontrado")
			return
		}
		if errors.Is(err, storage.ErrStockInsuficiente) {
			writeError(w, http.StatusBadRequest, "stock insuficiente")
			return
		}
		writeError(w, http.StatusInternalServerError, "error interno")
		return
	}

	writeJSON(w, http.StatusOK, p)
}
