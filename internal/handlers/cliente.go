package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Mildreth-SC/Sistema_almacen_equipos/internal/models"
	"github.com/go-chi/chi/v5"
)

func (s *Server) ListarClientes(w http.ResponseWriter, _ *http.Request) {
	RespondJSON(w, http.StatusOK, s.Clientes.Listar())
}

func (s *Server) ObtenerCliente(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	c, err := s.Clientes.Obtener(id)
	if err != nil {
		RespondError(w, statusDeError(err), err.Error())
		return
	}
	RespondJSON(w, http.StatusOK, c)
}

func (s *Server) CrearCliente(w http.ResponseWriter, r *http.Request) {
	var c models.Cliente
	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		RespondError(w, http.StatusBadRequest, "json invalido")
		return
	}
	creado, err := s.Clientes.Crear(c)
	if err != nil {
		RespondError(w, statusDeError(err), err.Error())
		return
	}
	RespondJSON(w, http.StatusCreated, creado)
}

func (s *Server) ActualizarCliente(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var c models.Cliente
	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		RespondError(w, http.StatusBadRequest, "json invalido")
		return
	}
	actualizado, err := s.Clientes.Actualizar(id, c)
	if err != nil {
		RespondError(w, statusDeError(err), err.Error())
		return
	}
	RespondJSON(w, http.StatusOK, actualizado)
}

func (s *Server) BorrarCliente(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if err := s.Clientes.Borrar(id); err != nil {
		RespondError(w, statusDeError(err), err.Error())
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
