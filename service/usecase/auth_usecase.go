package usecase

import (
	"errors"
	"github.com/golang-jwt/jwt"
)

type AuthService interface {
	AccessToken(phoneNumber string) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}

type authService struct {
}

var SECRET_KEY = []byte("MNC_TEST")

func NewAuthService() AuthService {
	return &authService{}
}

func (s *authService) AccessToken(phoneNumber string) (string, error) {
	claim := jwt.MapClaims{}
	claim["phone_number"] = phoneNumber

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	signedToken, err := token.SignedString(SECRET_KEY)
	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}
func (s *authService) ValidateToken(encodedToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("Invalid token")
		}

		return []byte(SECRET_KEY), nil
	})

	if err != nil {
		return token, err
	}

	return token, nil
}
