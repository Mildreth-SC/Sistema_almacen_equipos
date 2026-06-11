// MODULO REALIZO POR MILDRETH GUANOLUISA
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Mildreth-SC/Sistema_almacen_equipos/internal/handlers"
	"github.com/Mildreth-SC/Sistema_almacen_equipos/internal/middleware"
	"github.com/Mildreth-SC/Sistema_almacen_equipos/internal/models"
	"github.com/Mildreth-SC/Sistema_almacen_equipos/internal/storage"
	"github.com/glebarez/sqlite"
	"github.com/go-chi/chi/v5"
	chimiddleware "github.com/go-chi/chi/v5/middleware"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("cmd/api/almacen.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("no se pudo conectar a SQLite:", err)
	}

	if err := db.AutoMigrate(
		&models.Pieza{},
		&models.Devolucion{},
		&models.RegistroMantenimiento{},
	); err != nil {
		log.Fatal("no se pudo migrar tablas:", err)
	}

	almacen := storage.NewAlmacenSQLite(db)
	almacen.Sembrarvacio()

	srv := handlers.NewServer(almacen)

	r := chi.NewRouter()
	r.Use(chimiddleware.Logger)
	r.Use(chimiddleware.Recoverer)
	r.Use(middleware.CORS)

	srv.RegisterRoutes(r)

	fmt.Println("Servidor en http://localhost:8080")
	fmt.Println("Base de datos: cmd/api/almacen.db")
	log.Fatal(http.ListenAndServe(":8080", r))
}
