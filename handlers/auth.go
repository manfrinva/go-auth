package handlers

import (
    "encoding/json"
    "net/http"

    "go-auth/models"
    "go-auth/utils"
)

var users = map[string]string{
    "user@example.com": "password123", // Simulação de base de dados
}

func Login(w http.ResponseWriter, r *http.Request) {
    var creds models.User
    if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
        http.Error(w, "Dados inválidos", http.StatusBadRequest)
        return
    }

    expectedPassword, ok := users[creds.Email]
    if !ok || expectedPassword != creds.Password {
        http.Error(w, "Credenciais inválidas", http.StatusUnauthorized)
        return
    }

    token, err := utils.GenerateJWT(creds.Email)
    if err != nil {
        http.Error(w, "Erro ao gerar token", http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(map[string]string{"token": token})
}

func ProtectedEndpoint(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Bem-vindo à rota protegida!"))
}