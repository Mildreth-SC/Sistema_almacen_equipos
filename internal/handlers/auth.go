package handlers

import (
	"encoding/json"
	"net/http"
)

type credenciales struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (s *Server) Registrar(w http.ResponseWriter, r *http.Request) {
	var creds credenciales
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		RespondError(w, http.StatusBadRequest, "json invalido: "+err.Error())
		return
	}
	if _, err := s.Auth.Registrar(creds.Email, creds.Password); err != nil {
		RespondError(w, statusDeError(err), err.Error())
		return
	}
	token, err := s.Auth.Login(creds.Email, creds.Password)
	if err != nil {
		RespondError(w, statusDeError(err), err.Error())
		return
	}
	RespondJSON(w, http.StatusOK, map[string]string{"token": token})
}

func (s *Server) Login(w http.ResponseWriter, r *http.Request) {
	var creds credenciales
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		RespondError(w, http.StatusBadRequest, "json invalido: "+err.Error())
		return
	}
	token, err := s.Auth.Login(creds.Email, creds.Password)
	if err != nil {
		RespondError(w, statusDeError(err), err.Error())
		return
	}
	RespondJSON(w, http.StatusOK, map[string]string{"token": token})
}
