package models

import "github.com/golang-jwt/jwt/v5"

type AuthTokenClaims struct {
	SessionID string `json:"sid"`
	jwt.RegisteredClaims
}

type MagicLinkTokenClaims struct {
	jwt.RegisteredClaims
}
