package services

import "github.com/golang-jwt/jwt"

type JWTservice interface {
	GenerateToken(userId string, email string, userName string) string
	ValidateToken(token string) (*jwt.Token, error)
}
