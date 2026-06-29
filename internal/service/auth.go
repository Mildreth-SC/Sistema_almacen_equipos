package service

import (
	"os"
	"strings"
	"time"

	"github.com/Mildreth-SC/Sistema_almacen_equipos/internal/models"
	"github.com/Mildreth-SC/Sistema_almacen_equipos/internal/storage"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func jwtSecret() []byte {
	if s := os.Getenv("JWT_SECRET"); s != "" {
		return []byte(s)
	}
	return []byte("secreto-jwt-almacen-dev")
}
var duracionToken = time.Hour * 24

type claims struct {
	UsuarioID int `json:"uid"`
	jwt.RegisteredClaims
}

type AuthService struct {
	repo storage.UsuarioRepository
}

func NewAuthService(repo storage.UsuarioRepository) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) Registrar(email, password string) (models.Usuario, error) {
	email = strings.TrimSpace(strings.ToLower(email))
	if email == "" || strings.TrimSpace(password) == "" {
		return models.Usuario{}, ErrEmailOContrasenaVacios
	}
	if _, ok := s.repo.BuscarUsuarioPorEmail(email); ok {
		return models.Usuario{}, ErrUsuarioYaExiste
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return models.Usuario{}, err
	}
	return s.repo.CrearUsuario(models.Usuario{
		Email:        email,
		PasswordHash: string(hash),
	})
}

func (s *AuthService) Login(email, password string) (string, error) {
	email = strings.TrimSpace(strings.ToLower(email))
	u, existe := s.repo.BuscarUsuarioPorEmail(email)
	if !existe {
		return "", ErrUsuarioNoEncontrado
	}
	if err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password)); err != nil {
		return "", ErrCredencialesInvalidas
	}
	return s.generarToken(u)
}

func (s *AuthService) generarToken(u models.Usuario) (string, error) {
	c := &claims{
		UsuarioID: u.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duracionToken)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString(jwtSecret())
}

func (s *AuthService) ValidarToken(tokenStr string) (int, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &claims{}, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrCredencialesInvalidas
		}
		return jwtSecret(), nil
	})
	if err != nil || !token.Valid {
		return 0, ErrCredencialesInvalidas
	}
	c, ok := token.Claims.(*claims)
	if !ok {
		return 0, ErrCredencialesInvalidas
	}
	return c.UsuarioID, nil
}
