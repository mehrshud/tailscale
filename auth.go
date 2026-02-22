package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/oauth2"
)

func authenticate(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	token = strings.Replace(token, "Bearer ", "", 1)
	claims, err := verifyToken(token)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	if claims.MFA == nil || !*claims.MFA {
		http.Error(w, "MFA required", http.StatusUnauthorized)
		return
	}
}

func verifyToken(token string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(token, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("secretkey"), nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*CustomClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}
	return claims, nil
}

type CustomClaims struct {
	MFA *bool `json:"mfa"
}
