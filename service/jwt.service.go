package service

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type JWTService interface {
	GenerateToken(userID string) string
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtCustomClaim struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims
}

type jwtService struct {
	secretKey string
	issuer    string
}

// NewJWTService membuat instance baru dari JWTservice
func NewJWTService() JWTService {
	return &jwtService{
		issuer:    "roihan",
		secretKey: getSecrectKey(),
	}

}

func getSecrectKey() string {
	secretKey := os.Getenv("JWT_SECRET")
	if secretKey != "" {
		secretKey = "inikuncijwt"
	}

	return secretKey
}

func (j *jwtService) GenerateToken(UserID string) string {

	claims := &jwtCustomClaim{
		UserID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().AddDate(1, 0, 0).Unix(),
			Issuer:    j.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(j.secretKey))
	if err != nil {
		panic(err)
	}
	return t
}

func (j *jwtService) ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {

		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method %v", t.Header["alg"])
		}

		return []byte(j.secretKey), nil
	})

}

// func (j *jwtService) ValidateToken(token string, error) *jwt.Token {
// 	t, err := jwt.Parse(token, func(t_ *jwt.Token) (interface{}, error) {
// 		if _, ok := t_.Method.(*jwt.SigningMethodHMAC); !ok {
// 			return nil, fmt.Errorf("Unexpected signing method %v", t_.Header["alg"])
// 		}
// 		return []byte(j.secretKey), nil
// 	})

// 	if err != nil {
// 		return nil
// 	}

// 	return t

// }
