// MODULO REALIZADO POR JOSÉ MIELES — Lógica de negocio de mantenimientos

package service

import (
	"strings"
	"time"

	"github.com/Mildreth-SC/Sistema_almacen_equipos/internal/models"
	"github.com/Mildreth-SC/Sistema_almacen_equipos/internal/storage"
)

type MantenimientoService struct {
	repo        storage.MantenimientoRepository
	piezaRepo   storage.PiezaRepository
	clienteRepo storage.ClienteRepository
}

func NewMantenimientoService(repo storage.MantenimientoRepository, piezaRepo storage.PiezaRepository, clienteRepo storage.ClienteRepository) *MantenimientoService {
	return &MantenimientoService{repo: repo, piezaRepo: piezaRepo, clienteRepo: clienteRepo}
}

func (s *MantenimientoService) Listar(estado string) []models.RegistroMantenimiento {
	lista := s.repo.ListarMantenimientos()
	if estado == "" {
		return lista
	}
	filtrada := make([]models.RegistroMantenimiento, 0)
	for _, m := range lista {
		if string(m.Estado) == estado {
			filtrada = append(filtrada, m)
		}
	}
	return filtrada
}

func (s *MantenimientoService) Obtener(id string) (models.RegistroMantenimiento, error) {
	m, ok := s.repo.BuscarMantenimientoPorID(id)
	if !ok {
		return models.RegistroMantenimiento{}, ErrNoEncontrado
	}
	return m, nil
}

func (s *MantenimientoService) Crear(m models.RegistroMantenimiento) (models.RegistroMantenimiento, error) {
	if err := s.validarMantenimiento(m); err != nil {
		return models.RegistroMantenimiento{}, err
	}
	if m.Estado == "" {
		m.Estado = models.MantenimientoPendiente
	}
	m.Pieza = models.Pieza{}
	m.Cliente = models.Cliente{}
	creado := s.repo.CrearMantenimiento(m)
	if creado.ID == "" {
		return models.RegistroMantenimiento{}, ErrErrorInterno
	}
	return creado, nil
}

func (s *MantenimientoService) Actualizar(id string, m models.RegistroMantenimiento) (models.RegistroMantenimiento, error) {
	if err := s.validarMantenimiento(m); err != nil {
		return models.RegistroMantenimiento{}, err
	}
	m.Pieza = models.Pieza{}
	m.Cliente = models.Cliente{}
	actualizado, ok := s.repo.ActualizarMantenimiento(id, m)
	if !ok {
		return models.RegistroMantenimiento{}, ErrNoEncontrado
	}
	return actualizado, nil
}

// CambiarEstado avanza el flujo: PENDIENTE → EN_PROCESO → LISTO → ENTREGADO.
func (s *MantenimientoService) CambiarEstado(id string, nuevoEstado models.EstadoMantenimiento) (models.RegistroMantenimiento, error) {
	if !models.EsEstadoMantenimientoValido(nuevoEstado) {
		return models.RegistroMantenimiento{}, ErrEstadoInvalido
	}

	m, ok := s.repo.BuscarMantenimientoPorID(id)
	if !ok {
		return models.RegistroMantenimiento{}, ErrNoEncontrado
	}
	if m.Estado == nuevoEstado {
		return m, nil
	}
	if !models.TransicionMantenimientoValida(m.Estado, nuevoEstado) {
		return models.RegistroMantenimiento{}, ErrTransicionEstadoInvalida
	}

	m.Estado = nuevoEstado
	if nuevoEstado == models.MantenimientoEntregado {
		ahora := time.Now()
		m.FechaEntrega = &ahora
	}
	m.Pieza = models.Pieza{}
	m.Cliente = models.Cliente{}

	actualizado, ok := s.repo.ActualizarMantenimiento(id, m)
	if !ok {
		return models.RegistroMantenimiento{}, ErrNoEncontrado
	}
	return actualizado, nil
}

func (s *MantenimientoService) Borrar(id string) error {
	if !s.repo.BorrarMantenimiento(id) {
		return ErrNoEncontrado
	}
	return nil
}

func (s *MantenimientoService) validarMantenimiento(m models.RegistroMantenimiento) error {
	if strings.TrimSpace(m.ClienteID) == "" {
		return ErrClienteIDVacio
	}
	if _, ok := s.clienteRepo.BuscarClientePorID(m.ClienteID); !ok {
		return ErrNoEncontrado
	}
	if strings.TrimSpace(m.EquipoDescripcion) == "" {
		return ErrEquipoDescripcionVacia
	}
	if strings.TrimSpace(m.FallaReportada) == "" {
		return ErrFallaReportadaVacia
	}
	if strings.TrimSpace(m.Tecnico) == "" {
		return ErrTecnicoVacio
	}
	if m.Tipo == "" {
		return ErrTipoVacio
	}
	if !models.EsTipoMantenimientoValido(m.Tipo) {
		return ErrTipoInvalido
	}
	if m.Costo < 0 {
		return ErrCostoNegativo
	}
	if m.Anticipo < 0 || m.Anticipo > m.Costo {
		return ErrAnticipoInvalido
	}
	if m.Estado != "" && !models.EsEstadoMantenimientoValido(m.Estado) {
		return ErrEstadoInvalido
	}
	if m.PiezaID != "" {
		if _, ok := s.piezaRepo.BuscarPiezaPorID(m.PiezaID); !ok {
			return ErrNoEncontrado
		}
	}
	return nil
}
