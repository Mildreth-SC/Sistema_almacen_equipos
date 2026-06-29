package handlers

import (
	"github.com/Mildreth-SC/Sistema_almacen_equipos/internal/middleware"
	"github.com/Mildreth-SC/Sistema_almacen_equipos/internal/service"
	"github.com/go-chi/chi/v5"
)

// Server agrupa los services de los 3 módulos + auth.
type Server struct {
	Piezas         *service.PiezaService
	Devoluciones   *service.DevolucionService
	Mantenimientos *service.MantenimientoService
	Auth           *service.AuthService
}

func NewServer(
	piezas *service.PiezaService,
	devoluciones *service.DevolucionService,
	mantenimientos *service.MantenimientoService,
	auth *service.AuthService,
) *Server {
	return &Server{
		Piezas:         piezas,
		Devoluciones:   devoluciones,
		Mantenimientos: mantenimientos,
		Auth:           auth,
	}
}

func (s *Server) RegisterRoutes(r chi.Router) {
	r.Route("/api/v1", func(r chi.Router) {
		// Rutas públicas
		r.Post("/auth/registrar", s.Registrar)
		r.Post("/auth/login", s.Login)

		// Rutas protegidas — requieren Bearer token
		r.Group(func(r chi.Router) {
			r.Use(middleware.Auth(s.Auth))

			r.Route("/inventario-piezas", func(r chi.Router) {
				r.Get("/", s.ListarPiezas)
				r.Get("/{id}", s.ObtenerPieza)
				r.Post("/", s.CrearPieza)
				r.Put("/{id}", s.ActualizarPieza)
				r.Delete("/{id}", s.BorrarPieza)
				r.Patch("/{id}/stock", s.AjustarStockPieza)
			})

			// Módulo devoluciones — Ivanna Zamora
			r.Route("/devoluciones", func(r chi.Router) {
				r.Get("/", s.ListarDevoluciones)
				r.Get("/{id}", s.ObtenerDevolucion)
				r.Post("/", s.CrearDevolucion)
				r.Put("/{id}", s.ActualizarDevolucion)
				r.Patch("/{id}/estado", s.CambiarEstadoDevolucion)
				r.Delete("/{id}", s.BorrarDevolucion)
			})

			// Módulo mantenimientos — José Mieles
			r.Route("/mantenimientos", func(r chi.Router) {
				r.Get("/", s.ListarMantenimientos)
				r.Get("/{id}", s.ObtenerMantenimiento)
				r.Post("/", s.CrearMantenimiento)
				r.Put("/{id}", s.ActualizarMantenimiento)
				r.Patch("/{id}/estado", s.CambiarEstadoMantenimiento)
				r.Delete("/{id}", s.BorrarMantenimiento)
			})
		})
	})
}
