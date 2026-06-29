package storage

import (
	"sync"

	"github.com/Mildreth-SC/Sistema_almacen_equipos/internal/models"
)

// UsuarioMemoria es un fake en memoria para tests de handlers (no es la base real).
type UsuarioMemoria struct {
	mu       sync.RWMutex
	usuarios map[string]models.Usuario
	nextID   int
}

func NewUsuarioMemoria() *UsuarioMemoria {
	return &UsuarioMemoria{usuarios: make(map[string]models.Usuario)}
}

func (r *UsuarioMemoria) CrearUsuario(u models.Usuario) (models.Usuario, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.nextID++
	u.ID = r.nextID
	r.usuarios[u.Email] = u
	return u, nil
}

func (r *UsuarioMemoria) BuscarUsuarioPorEmail(email string) (models.Usuario, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	u, ok := r.usuarios[email]
	return u, ok
}

var _ UsuarioRepository = (*UsuarioMemoria)(nil)
