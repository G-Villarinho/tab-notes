package services

import (
	"context"
	"time"

	"github.com/g-villarinho/tab-notes-api/configs"
	"github.com/g-villarinho/tab-notes-api/models"
	"github.com/g-villarinho/tab-notes-api/pkgs"
	"github.com/golang-jwt/jwt/v5"
)

type TokenService interface {
	GenerateAuthToken(ctx context.Context, userID string, sessionID string, iat, exp time.Time) (string, error)
	GenerateMagicLinkToken(ctx context.Context, email string, iat, exp time.Time) (string, error)
}

type tokenService struct {
	kp pkgs.EcdsaKeyPair
}

func NewTokenService(kp pkgs.EcdsaKeyPair) TokenService {
	return &tokenService{kp: kp}
}

func (t *tokenService) GenerateAuthToken(ctx context.Context, userID string, sessionID string, iat, exp time.Time) (string, error) {
	privateKey, err := t.kp.ParseECDSAPrivateKey(configs.Env.Key.PrivateKey)
	if err != nil {
		return "", err
	}

	claims := models.AuthTokenClaims{
		SessionID: sessionID,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   userID,
			IssuedAt:  jwt.NewNumericDate(iat),
			ExpiresAt: jwt.NewNumericDate(exp),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	return token.SignedString(privateKey)
}

func (t *tokenService) GenerateMagicLinkToken(ctx context.Context, email string, iat, exp time.Time) (string, error) {
	privateKey, err := t.kp.ParseECDSAPrivateKey(configs.Env.Key.PrivateKey)
	if err != nil {
		return "", err
	}

	claims := models.MagicLinkTokenClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   email,
			IssuedAt:  jwt.NewNumericDate(iat),
			ExpiresAt: jwt.NewNumericDate(exp),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	return token.SignedString(privateKey)
}
