package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/Mildreth-SC/Sistema_almacen_equipos/internal/service"
)

type claveContexto string

const claveUsuarioID claveContexto = "usuarioID"

func Auth(authSvc *service.AuthService) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			encabezado := r.Header.Get("Authorization")
			partes := strings.SplitN(encabezado, " ", 2)
			if len(partes) != 2 || partes[0] != "Bearer" {
				respondeNoAutorizado(w)
				return
			}
			usuarioID, err := authSvc.ValidarToken(partes[1])
			if err != nil {
				respondeNoAutorizado(w)
				return
			}
			ctx := context.WithValue(r.Context(), claveUsuarioID, usuarioID)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func respondeNoAutorizado(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusUnauthorized)
	_, _ = w.Write([]byte(`{"error":"token no autorizado"}`))
}
