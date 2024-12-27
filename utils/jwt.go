package utils

import (
    "errors"
    "time"

    "github.com/golang-jwt/jwt/v4"
)

var secretKey = []byte("minha_chave_secreta")

func GenerateJWT(email string) (string, error) {
    claims := &jwt.RegisteredClaims{
        Subject:   email,
        ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(secretKey)
}

func ValidateJWT(tokenString string) (*jwt.RegisteredClaims, error) {
    token, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
        return secretKey, nil
    })
    if err != nil {
        return nil, err
    }

    claims, ok := token.Claims.(*jwt.RegisteredClaims)
    if !ok || !token.Valid {
        return nil, errors.New("token inválido")
    }

    return claims, nil
}