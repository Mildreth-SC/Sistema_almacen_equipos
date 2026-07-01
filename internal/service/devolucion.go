// MODULO REALIZADO POR IVANNA ZAMORA — Lógica de negocio de devoluciones

package service

import (
	"strings"
	"time"

	"github.com/Mildreth-SC/Sistema_almacen_equipos/internal/models"
	"github.com/Mildreth-SC/Sistema_almacen_equipos/internal/storage"
)

type DevolucionService struct {
	repo        storage.DevolucionRepository
	piezaRepo   storage.PiezaRepository
	clienteRepo storage.ClienteRepository
}

func NewDevolucionService(repo storage.DevolucionRepository, piezaRepo storage.PiezaRepository, clienteRepo storage.ClienteRepository) *DevolucionService {
	return &DevolucionService{repo: repo, piezaRepo: piezaRepo, clienteRepo: clienteRepo}
}

func (s *DevolucionService) Listar(estado string) []models.Devolucion {
	lista := s.repo.ListarDevoluciones()
	if estado == "" {
		return lista
	}
	filtrada := make([]models.Devolucion, 0)
	for _, d := range lista {
		if string(d.Estado) == estado {
			filtrada = append(filtrada, d)
		}
	}
	return filtrada
}

func (s *DevolucionService) Obtener(id string) (models.Devolucion, error) {
	d, ok := s.repo.BuscarDevolucionPorID(id)
	if !ok {
		return models.Devolucion{}, ErrNoEncontrado
	}
	return d, nil
}

func (s *DevolucionService) Crear(d models.Devolucion) (models.Devolucion, error) {
	if err := s.validarDevolucion(d); err != nil {
		return models.Devolucion{}, err
	}
	if d.Estado == "" {
		d.Estado = models.EstadoPendiente
	}
	d.Pieza = models.Pieza{}
	d.Cliente = models.Cliente{}
	creada := s.repo.CrearDevolucion(d)
	if creada.ID == "" {
		return models.Devolucion{}, ErrErrorInterno
	}
	return creada, nil
}

func (s *DevolucionService) Actualizar(id string, d models.Devolucion) (models.Devolucion, error) {
	if err := s.validarDevolucion(d); err != nil {
		return models.Devolucion{}, err
	}
	d.Pieza = models.Pieza{}
	d.Cliente = models.Cliente{}
	actualizada, ok := s.repo.ActualizarDevolucion(id, d)
	if !ok {
		return models.Devolucion{}, ErrNoEncontrado
	}
	return actualizada, nil
}

// CambiarEstado resuelve una devolución pendiente (aprobar o rechazar).
func (s *DevolucionService) CambiarEstado(id string, estado models.EstadoDevolucion, resolucion, atendidoPor string) (models.Devolucion, error) {
	if !models.EsEstadoDevolucionValido(estado) {
		return models.Devolucion{}, ErrEstadoInvalido
	}
	if estado == models.EstadoPendiente {
		return models.Devolucion{}, ErrEstadoDevolucionPendiente
	}

	d, ok := s.repo.BuscarDevolucionPorID(id)
	if !ok {
		return models.Devolucion{}, ErrNoEncontrado
	}
	if d.Estado != models.EstadoPendiente {
		return models.Devolucion{}, ErrDevolucionYaResuelta
	}

	if strings.TrimSpace(resolucion) == "" {
		return models.Devolucion{}, ErrResolucionVacia
	}
	if strings.TrimSpace(atendidoPor) == "" {
		return models.Devolucion{}, ErrAtendidoPorVacio
	}

	ahora := time.Now()
	d.Estado = estado
	d.Resolucion = strings.TrimSpace(resolucion)
	d.AtendidoPor = strings.TrimSpace(atendidoPor)
	d.FechaResolucion = &ahora
	d.Pieza = models.Pieza{}
	d.Cliente = models.Cliente{}

	actualizada, ok := s.repo.ActualizarDevolucion(id, d)
	if !ok {
		return models.Devolucion{}, ErrNoEncontrado
	}
	return actualizada, nil
}

func (s *DevolucionService) Borrar(id string) error {
	if !s.repo.BorrarDevolucion(id) {
		return ErrNoEncontrado
	}
	return nil
}

func (s *DevolucionService) validarDevolucion(d models.Devolucion) error {
	if strings.TrimSpace(d.ClienteID) == "" {
		return ErrClienteIDVacio
	}
	if _, ok := s.clienteRepo.BuscarClientePorID(d.ClienteID); !ok {
		return ErrNoEncontrado
	}
	if strings.TrimSpace(d.PiezaID) == "" {
		return ErrPiezaIDVacio
	}
	if _, ok := s.piezaRepo.BuscarPiezaPorID(d.PiezaID); !ok {
		return ErrNoEncontrado
	}
	if strings.TrimSpace(d.NumeroFactura) == "" {
		return ErrNumeroFacturaVacio
	}
	if d.Motivo == "" {
		return ErrMotivoVacio
	}
	if !models.EsMotivoDevolucionValido(d.Motivo) {
		return ErrMotivoInvalido
	}
	if d.Estado != "" && !models.EsEstadoDevolucionValido(d.Estado) {
		return ErrEstadoInvalido
	}
	return nil
}
