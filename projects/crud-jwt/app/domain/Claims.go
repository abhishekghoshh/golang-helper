package domain

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type TokenClaim struct {
	TokenType string `json:"token_type,omitempty"`
	Username  string `json:"username"`
	Role      string `json:"role"`
	jwt.StandardClaims
}

func (c TokenClaim) IsUserRole() bool {
	return c.Role == "user"
}
func (c TokenClaim) IsAdminRole() bool {
	return c.Role == "admin"
}

func CreateTokenClaims(tokeType string, username string, role string, d time.Duration) *TokenClaim {
	return &TokenClaim{
		// TokenType: "refresh_token",
		TokenType: tokeType,
		Username:  username,
		Role:      role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(d).Unix(),
		},
	}
}
