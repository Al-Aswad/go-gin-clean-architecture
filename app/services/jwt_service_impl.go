package services

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

type JWTServiceImpl struct {
	UserId string `json:"user_id"`
	jwt.StandardClaims
}

type jwtCustomClaim struct {
	UserId string `json:"user_id"`
	jwt.StandardClaims
}

type jwtService struct {
	secretKey string
	issuer    string
}

func CreateJwtService() JWTservice {
	return &jwtService{
		secretKey: getSecretKey(),
		issuer:    "aswad",
	}
}

func getSecretKey() string {
	secretKey := os.Getenv("SECRET_KEY")

	if secretKey == "" {
		secretKey = "secret"
	}

	return secretKey
}

func (j *jwtService) GenerateToken(userID string) string {
	claims := jwtCustomClaim{
		userID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().AddDate(1, 0, 0).Unix(),
			Issuer:    j.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(j.secretKey))
	if err != nil {
		return "err"
	}

	return t
}

func (j *jwtService) ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t_ *jwt.Token) (interface{}, error) {
		if _, ok := t_.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(j.secretKey), nil
	})
}
