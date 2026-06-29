// MODULO REALIZADO POR MILDRETH GUANOLUISA — Handler HTTP de inventario de piezas

package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Mildreth-SC/Sistema_almacen_equipos/internal/models"
	"github.com/go-chi/chi/v5"
)

func (s *Server) ListarPiezas(w http.ResponseWriter, _ *http.Request) {
	RespondJSON(w, http.StatusOK, s.Piezas.Listar())
}

func (s *Server) ObtenerPieza(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	p, err := s.Piezas.Obtener(id)
	if err != nil {
		RespondError(w, statusDeError(err), err.Error())
		return
	}
	RespondJSON(w, http.StatusOK, p)
}

func (s *Server) CrearPieza(w http.ResponseWriter, r *http.Request) {
	var p models.Pieza
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		RespondError(w, http.StatusBadRequest, "json invalido")
		return
	}
	creada, err := s.Piezas.Crear(p)
	if err != nil {
		RespondError(w, statusDeError(err), err.Error())
		return
	}
	RespondJSON(w, http.StatusCreated, creada)
}

func (s *Server) ActualizarPieza(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var p models.Pieza
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		RespondError(w, http.StatusBadRequest, "json invalido")
		return
	}
	actualizada, err := s.Piezas.Actualizar(id, p)
	if err != nil {
		RespondError(w, statusDeError(err), err.Error())
		return
	}
	RespondJSON(w, http.StatusOK, actualizada)
}

func (s *Server) BorrarPieza(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if err := s.Piezas.Borrar(id); err != nil {
		RespondError(w, statusDeError(err), err.Error())
		return
	}
	RespondJSON(w, http.StatusOK, map[string]string{"mensaje": "eliminado"})
}

func (s *Server) AjustarStockPieza(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var body struct {
		Delta int `json:"delta"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		RespondError(w, http.StatusBadRequest, "json invalido")
		return
	}
	p, err := s.Piezas.AjustarStock(id, body.Delta)
	if err != nil {
		RespondError(w, statusDeError(err), err.Error())
		return
	}
	RespondJSON(w, http.StatusOK, p)
}
