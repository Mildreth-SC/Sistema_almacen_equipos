// Punto de entrada — los 3 módulos del equipo corren en este mismo binario.
// Mildreth: inventario-piezas | Ivanna: devoluciones | José: mantenimientos

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Mildreth-SC/Sistema_almacen_equipos/internal/handlers"
	"github.com/Mildreth-SC/Sistema_almacen_equipos/internal/middleware"
	"github.com/Mildreth-SC/Sistema_almacen_equipos/internal/models"
	"github.com/Mildreth-SC/Sistema_almacen_equipos/internal/service"
	"github.com/Mildreth-SC/Sistema_almacen_equipos/internal/storage"
	"github.com/glebarez/sqlite"
	"github.com/go-chi/chi/v5"
	chimiddleware "github.com/go-chi/chi/v5/middleware"
	"gorm.io/gorm"
)

const dbPath = "db/almacen.db"

func main() {
	// 1. GORM es dueño del esquema: abre la DB, migra y siembra.
	if err := os.MkdirAll("db", 0o755); err != nil {
		log.Fatal("no se pudo crear carpeta db:", err)
	}
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		log.Fatal("no se pudo conectar a SQLite:", err)
	}
	if err := db.AutoMigrate(
		&models.Pieza{},
		&models.Cliente{},
		&models.Devolucion{},
		&models.RegistroMantenimiento{},
		&models.Usuario{},
	); err != nil {
		log.Fatal("no se pudo migrar tablas:", err)
	}

	almacen := storage.NewAlmacenSQLite(db)
	almacen.Sembrarvacio()
	usuarioRepo := storage.NuevoUsuarioGORM(db)

	// 2. Capa de servicios + server con inyección de dependencias.
	piezaSvc := service.NewPiezaService(almacen)
	clienteSvc := service.NewClienteService(almacen, almacen, almacen)
	devolucionSvc := service.NewDevolucionService(almacen, almacen, almacen)
	mantenimientoSvc := service.NewMantenimientoService(almacen, almacen, almacen)
	authSvc := service.NewAuthService(usuarioRepo)
	srv := handlers.NewServer(piezaSvc, clienteSvc, devolucionSvc, mantenimientoSvc, authSvc)

	// 3. Router + middleware.
	r := chi.NewRouter()
	r.Use(chimiddleware.Logger)
	r.Use(chimiddleware.Recoverer)
	r.Use(middleware.CORS)

	srv.RegisterRoutes(r)

	r.Get("/", func(w http.ResponseWriter, req *http.Request) {
		http.ServeFile(w, req, "frontend/index.html")
	})

	fmt.Println("Servidor en http://localhost:8080")
	fmt.Println("Frontend:  http://localhost:8080/")
	fmt.Println("Base de datos:", dbPath)
	log.Fatal(http.ListenAndServe(":8080", r))
}
