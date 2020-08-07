package model

import (
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

type Motor struct {
	Brand       string `json:"brand"`
	Color       string `json:"color"`
	Seats       string `json:"seats"`
	PlateNumber string `json:"plate"`
	Membership  string `json:"membership"`
}

type AuthResponse struct {
	Success    bool
	Message    string
	StatusCode int
	AuthToken  *AuthToken `json:"authToken"`
	Motor      *Motor     `json:"motor"`
}

type AuthToken struct {
	AccessToken string    `json:"accessToken"`
	ExpiredAt   time.Time `json:"expiredAt"`
}

func (m *Motor) GenToken() (*AuthToken, error) {
	expireAt := time.Now().Add(time.Hour * 24 * 7) // one week
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"membership": m.Membership,
	})

	accessToken, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return nil, err
	}

	return &AuthToken{
		AccessToken: accessToken,
		ExpiredAt:   expireAt,
	}, nil
}
