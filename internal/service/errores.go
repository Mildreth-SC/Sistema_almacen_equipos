package service

import "errors"

var (
	ErrNoEncontrado = errors.New("no encontrado")

	// Piezas
	ErrNombreVacio         = errors.New("el nombre es requerido")
	ErrNumeroSerialVacio   = errors.New("el numero_serial es requerido")
	ErrCodigoBarrasVacio   = errors.New("el codigo_barras es requerido")
	ErrStockNegativo       = errors.New("el stock no puede ser negativo")
	ErrStockMinimoNegativo = errors.New("el stock minimo no puede ser negativo")
	ErrPrecioNegativo      = errors.New("los precios no pueden ser negativos")
	ErrGarantiaNegativa    = errors.New("la garantia no puede ser negativa")
	ErrStockInsuficiente   = errors.New("stock insuficiente")
	ErrRegistroDuplicado   = errors.New("numero_serial o codigo_barras ya existe")

	// Clientes
	ErrNombreClienteVacio = errors.New("el nombre del cliente es requerido")
	ErrCedulaVacia        = errors.New("la cedula es requerida")
	ErrCedulaDuplicada    = errors.New("ya existe un cliente con esa cedula")
	ErrClienteEnUso       = errors.New("el cliente tiene devoluciones o mantenimientos asociados")

	// Devoluciones — Ivanna Zamora
	ErrClienteIDVacio            = errors.New("el cliente_id es requerido")
	ErrPiezaIDVacio              = errors.New("el pieza_id es requerido")
	ErrNumeroFacturaVacio        = errors.New("el numero_factura es requerido")
	ErrMotivoVacio               = errors.New("el motivo es requerido")
	ErrMotivoInvalido            = errors.New("motivo invalido: use DEFECTUOSO, EQUIVOCADO o GARANTIA")
	ErrEstadoInvalido            = errors.New("estado invalido")
	ErrEstadoDevolucionPendiente = errors.New("use APROBADA o RECHAZADA para resolver")
	ErrDevolucionYaResuelta      = errors.New("la devolucion ya fue resuelta")
	ErrResolucionVacia           = errors.New("la resolucion es requerida al aprobar o rechazar")
	ErrAtendidoPorVacio          = errors.New("atendido_por es requerido al aprobar o rechazar")

	// Mantenimientos — José Mieles
	ErrEquipoDescripcionVacia  = errors.New("la equipo_descripcion es requerida")
	ErrFallaReportadaVacia     = errors.New("la falla_reportada es requerida")
	ErrTecnicoVacio            = errors.New("el tecnico es requerido")
	ErrTipoVacio               = errors.New("el tipo es requerido")
	ErrTipoInvalido            = errors.New("tipo invalido: use PREVENTIVO o CORRECTIVO")
	ErrCostoNegativo           = errors.New("el costo no puede ser negativo")
	ErrAnticipoInvalido        = errors.New("el anticipo no puede ser negativo ni mayor al costo")
	ErrTransicionEstadoInvalida = errors.New("transicion de estado invalida")

	ErrErrorInterno = errors.New("error interno del servidor")

	// Auth
	ErrEmailOContrasenaVacios = errors.New("email y contraseña son obligatorios")
	ErrCredencialesInvalidas  = errors.New("credenciales inválidas")
	ErrUsuarioYaExiste        = errors.New("usuario ya existe")
	ErrUsuarioNoEncontrado    = errors.New("usuario no encontrado")
)
