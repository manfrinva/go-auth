package middleware

import (
    "net/http"

    "go-auth/utils"
)

func AuthMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        token := r.Header.Get("Authorization")
        if token == "" {
            http.Error(w, "Token não fornecido", http.StatusUnauthorized)
            return
        }

        _, err := utils.ValidateJWT(token)
        if err != nil {
            http.Error(w, "Token inválido", http.StatusUnauthorized)
            return
        }

        next.ServeHTTP(w, r)
    })
}