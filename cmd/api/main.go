package main

import (
	"fmt"
	"net/http"

	"github.com/Mildreth-SC/Sistema_almacen_equipos/internal/handlers"
	"github.com/Mildreth-SC/Sistema_almacen_equipos/internal/storage"
	"github.com/go-chi/chi/v5"
)

func main() {
	fmt.Println("Iniciando servidor...")

	r := chi.NewRouter()

	invStorage := storage.NewInventarioStorage()
	invHandler := handlers.NewInventarioHandler(invStorage)
	r.Route("/api/v1/inventario", func(r chi.Router) {
		r.Get("/", invHandler.GetAll)
		r.Get("/{id}", invHandler.GetByID)
		r.Post("/", invHandler.Create)
		r.Put("/{id}", invHandler.Update)
		r.Delete("/{id}", invHandler.Delete)
		r.Patch("/{id}/stock", invHandler.AjustarStock)
	})

	http.ListenAndServe(":8080", r)
}
