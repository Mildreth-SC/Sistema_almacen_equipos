// MODULO REALIZADO POR IVANNA ZAMORA — Handler HTTP de devoluciones

package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Mildreth-SC/Sistema_almacen_equipos/internal/models"
	"github.com/go-chi/chi/v5"
)

func (s *Server) ListarDevoluciones(w http.ResponseWriter, r *http.Request) {
	estado := r.URL.Query().Get("estado")
	RespondJSON(w, http.StatusOK, s.Devoluciones.Listar(estado))
}

func (s *Server) ObtenerDevolucion(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	d, err := s.Devoluciones.Obtener(id)
	if err != nil {
		RespondError(w, statusDeError(err), err.Error())
		return
	}
	RespondJSON(w, http.StatusOK, d)
}

func (s *Server) CrearDevolucion(w http.ResponseWriter, r *http.Request) {
	var d models.Devolucion
	if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
		RespondError(w, http.StatusBadRequest, "json invalido")
		return
	}
	creada, err := s.Devoluciones.Crear(d)
	if err != nil {
		RespondError(w, statusDeError(err), err.Error())
		return
	}
	RespondJSON(w, http.StatusCreated, creada)
}

func (s *Server) ActualizarDevolucion(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var d models.Devolucion
	if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
		RespondError(w, http.StatusBadRequest, "json invalido")
		return
	}
	actualizada, err := s.Devoluciones.Actualizar(id, d)
	if err != nil {
		RespondError(w, statusDeError(err), err.Error())
		return
	}
	RespondJSON(w, http.StatusOK, actualizada)
}

func (s *Server) CambiarEstadoDevolucion(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var body struct {
		Estado      models.EstadoDevolucion `json:"estado"`
		Resolucion  string                  `json:"resolucion"`
		AtendidoPor string                  `json:"atendido_por"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		RespondError(w, http.StatusBadRequest, "json invalido")
		return
	}
	actualizada, err := s.Devoluciones.CambiarEstado(id, body.Estado, body.Resolucion, body.AtendidoPor)
	if err != nil {
		RespondError(w, statusDeError(err), err.Error())
		return
	}
	RespondJSON(w, http.StatusOK, actualizada)
}

func (s *Server) BorrarDevolucion(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if err := s.Devoluciones.Borrar(id); err != nil {
		RespondError(w, statusDeError(err), err.Error())
		return
	}
	RespondJSON(w, http.StatusOK, map[string]string{"mensaje": "eliminado"})
}
