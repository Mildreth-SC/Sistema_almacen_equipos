// MODULO REALIZADO POR JOSÉ MIELES — Handler HTTP de mantenimientos

package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Mildreth-SC/Sistema_almacen_equipos/internal/models"
	"github.com/go-chi/chi/v5"
)

func (s *Server) ListarMantenimientos(w http.ResponseWriter, r *http.Request) {
	estado := r.URL.Query().Get("estado")
	RespondJSON(w, http.StatusOK, s.Mantenimientos.Listar(estado))
}

func (s *Server) ObtenerMantenimiento(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	m, err := s.Mantenimientos.Obtener(id)
	if err != nil {
		RespondError(w, statusDeError(err), err.Error())
		return
	}
	RespondJSON(w, http.StatusOK, m)
}

func (s *Server) CrearMantenimiento(w http.ResponseWriter, r *http.Request) {
	var m models.RegistroMantenimiento
	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		RespondError(w, http.StatusBadRequest, "json invalido")
		return
	}
	creado, err := s.Mantenimientos.Crear(m)
	if err != nil {
		RespondError(w, statusDeError(err), err.Error())
		return
	}
	RespondJSON(w, http.StatusCreated, creado)
}

func (s *Server) ActualizarMantenimiento(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var m models.RegistroMantenimiento
	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		RespondError(w, http.StatusBadRequest, "json invalido")
		return
	}
	actualizado, err := s.Mantenimientos.Actualizar(id, m)
	if err != nil {
		RespondError(w, statusDeError(err), err.Error())
		return
	}
	RespondJSON(w, http.StatusOK, actualizado)
}

func (s *Server) CambiarEstadoMantenimiento(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var body struct {
		Estado models.EstadoMantenimiento `json:"estado"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		RespondError(w, http.StatusBadRequest, "json invalido")
		return
	}
	actualizado, err := s.Mantenimientos.CambiarEstado(id, body.Estado)
	if err != nil {
		RespondError(w, statusDeError(err), err.Error())
		return
	}
	RespondJSON(w, http.StatusOK, actualizado)
}

func (s *Server) BorrarMantenimiento(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if err := s.Mantenimientos.Borrar(id); err != nil {
		RespondError(w, statusDeError(err), err.Error())
		return
	}
	RespondJSON(w, http.StatusOK, map[string]string{"mensaje": "eliminado"})
}
