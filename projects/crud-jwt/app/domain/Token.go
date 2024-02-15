package domain

import "github.com/dgrijalva/jwt-go"

type Token struct {
	token *jwt.Token
}

func NewAuthToken(claims TokenClaim) Token {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return Token{
		token: token,
	}
}
