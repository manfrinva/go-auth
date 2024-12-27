package main

import (
    "log"
    "net/http"

    "github.com/gorilla/mux"
    "go-auth/handlers"
    "go-auth/middleware"
)

func main() {
    r := mux.NewRouter()

    // Rota pública
    r.HandleFunc("/login", handlers.Login).Methods("POST")

    // Rota protegida
    r.Handle("/protected", middleware.AuthMiddleware(http.HandlerFunc(handlers.ProtectedEndpoint))).Methods("GET")

    log.Println("Servidor rodando na porta 8080")
    log.Fatal(http.ListenAndServe(":8080", r))
}