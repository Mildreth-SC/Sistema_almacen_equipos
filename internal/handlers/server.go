package handlers

import (
	"github.com/Mildreth-SC/Sistema_almacen_equipos/internal/storage"
	"github.com/go-chi/chi/v5"
)

type Server struct {
	inventarioPiezas *InventarioPiezasHandler
	devoluciones     *DevolucionHandler
	mantenimientos   *MantenimientoHandler
}

func NewServer(a storage.Almacen) *Server {
	return &Server{
		inventarioPiezas: NewInventarioPiezasHandler(a),
		devoluciones:     NewDevolucionHandler(a),
		mantenimientos:   NewMantenimientoHandler(a),
	}
}

func (s *Server) RegisterRoutes(r chi.Router) {
	r.Route("/api/v1/inventario-piezas", func(r chi.Router) {
		r.Get("/", s.inventarioPiezas.GetAll)
		r.Get("/{id}", s.inventarioPiezas.GetByID)
		r.Post("/", s.inventarioPiezas.Create)
		r.Put("/{id}", s.inventarioPiezas.Update)
		r.Delete("/{id}", s.inventarioPiezas.Delete)
		r.Patch("/{id}/stock", s.inventarioPiezas.AjustarStock)
	})

	r.Route("/api/v1/devoluciones", func(r chi.Router) {
		r.Get("/", s.devoluciones.GetAll)
		r.Get("/{id}", s.devoluciones.GetByID)
		r.Post("/", s.devoluciones.Create)
		r.Put("/{id}", s.devoluciones.Update)
		r.Delete("/{id}", s.devoluciones.Delete)
	})

	r.Route("/api/v1/mantenimientos", func(r chi.Router) {
		r.Get("/", s.mantenimientos.GetAll)
		r.Get("/{id}", s.mantenimientos.GetByID)
		r.Post("/", s.mantenimientos.Create)
		r.Put("/{id}", s.mantenimientos.Update)
		r.Delete("/{id}", s.mantenimientos.Delete)
	})
}
