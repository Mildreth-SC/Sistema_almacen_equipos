package service

import (
	"testing"

	"github.com/Mildreth-SC/Sistema_almacen_equipos/internal/models"
	"github.com/Mildreth-SC/Sistema_almacen_equipos/internal/storage"
)

func crearClienteTest(t *testing.T, repo storage.ClienteRepository) models.Cliente {
	t.Helper()
	c, err := repo.CrearCliente(models.Cliente{
		Nombre: "María López",
		Cedula: "0923456789",
	})
	if err != nil {
		t.Fatalf("crear cliente: %v", err)
	}
	return c
}
