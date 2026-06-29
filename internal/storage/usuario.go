package storage

import (
	"time"

	"github.com/Mildreth-SC/Sistema_almacen_equipos/internal/models"
	"gorm.io/gorm"
)

type UsuarioRepository interface {
	CrearUsuario(u models.Usuario) (models.Usuario, error)
	BuscarUsuarioPorEmail(email string) (models.Usuario, bool)
}

type UsuarioGORM struct {
	db *gorm.DB
}

func NuevoUsuarioGORM(db *gorm.DB) *UsuarioGORM {
	return &UsuarioGORM{db: db}
}

func (r *UsuarioGORM) CrearUsuario(u models.Usuario) (models.Usuario, error) {
	u.CreadoEn = time.Now()
	if err := r.db.Create(&u).Error; err != nil {
		return models.Usuario{}, err
	}
	return u, nil
}

func (r *UsuarioGORM) BuscarUsuarioPorEmail(email string) (models.Usuario, bool) {
	var u models.Usuario
	if err := r.db.Where("email = ?", email).First(&u).Error; err != nil {
		return models.Usuario{}, false
	}
	return u, true
}

var _ UsuarioRepository = (*UsuarioGORM)(nil)
