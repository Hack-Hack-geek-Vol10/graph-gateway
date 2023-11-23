package middleware

import (
	"time"

	"github.com/golang-jwt/jwt"
)

type CustomClaims struct {
	Name          string    `json:"name"`
	Picture       string    `json:"picture"`
	Iss           string    `json:"iss"`
	Aud           string    `json:"aud"`
	AuthTime      int64     `json:"auth_time"`
	UserId        string    `json:"user_id"`
	Sub           string    `json:"sub"`
	Iat           int64     `json:"iat"`
	Exp           int64     `json:"exp"`
	Email         string    `json:"email"`
	ScreenName    string    `json:"screenName"`
	LastRefreshAt time.Time `json:"lastRefreshAt"`
	jwt.StandardClaims
}
