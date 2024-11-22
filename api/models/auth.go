package models

import "github.com/golang-jwt/jwt"

type CustomTokenClaims struct {
	Email string `json:"email"`
	Type  string `json:"type"`
	jwt.StandardClaims
}

type TokenResponse struct {
	TokenType    string `json:"token_type"`
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int64  `json:"expires_in"`
}
