package service

import (
	"errors"
	"strings"

	"github.com/Mildreth-SC/Sistema_almacen_equipos/internal/models"
	"github.com/Mildreth-SC/Sistema_almacen_equipos/internal/storage"
)

type ClienteService struct {
	repo           storage.ClienteRepository
	devoluciones   storage.DevolucionRepository
	mantenimientos storage.MantenimientoRepository
}

func NewClienteService(
	repo storage.ClienteRepository,
	devoluciones storage.DevolucionRepository,
	mantenimientos storage.MantenimientoRepository,
) *ClienteService {
	return &ClienteService{repo: repo, devoluciones: devoluciones, mantenimientos: mantenimientos}
}

func (s *ClienteService) Listar() []models.Cliente {
	return s.repo.ListarClientes()
}

func (s *ClienteService) Obtener(id string) (models.Cliente, error) {
	c, ok := s.repo.BuscarClientePorID(id)
	if !ok {
		return models.Cliente{}, ErrNoEncontrado
	}
	return c, nil
}

func (s *ClienteService) Crear(c models.Cliente) (models.Cliente, error) {
	if err := validarCliente(c); err != nil {
		return models.Cliente{}, err
	}
	creado, err := s.repo.CrearCliente(c)
	if err != nil {
		if errors.Is(err, storage.ErrDuplicado) {
			return models.Cliente{}, ErrCedulaDuplicada
		}
		return models.Cliente{}, err
	}
	return creado, nil
}

func (s *ClienteService) Actualizar(id string, c models.Cliente) (models.Cliente, error) {
	if err := validarCliente(c); err != nil {
		return models.Cliente{}, err
	}
	actualizado, ok := s.repo.ActualizarCliente(id, c)
	if !ok {
		return models.Cliente{}, ErrNoEncontrado
	}
	return actualizado, nil
}

func (s *ClienteService) Borrar(id string) error {
	if _, ok := s.repo.BuscarClientePorID(id); !ok {
		return ErrNoEncontrado
	}
	for _, d := range s.devoluciones.ListarDevoluciones() {
		if d.ClienteID == id {
			return ErrClienteEnUso
		}
	}
	for _, m := range s.mantenimientos.ListarMantenimientos() {
		if m.ClienteID == id {
			return ErrClienteEnUso
		}
	}
	if !s.repo.BorrarCliente(id) {
		return ErrNoEncontrado
	}
	return nil
}

func validarCliente(c models.Cliente) error {
	if strings.TrimSpace(c.Nombre) == "" {
		return ErrNombreClienteVacio
	}
	if strings.TrimSpace(c.Cedula) == "" {
		return ErrCedulaVacia
	}
	return nil
}
