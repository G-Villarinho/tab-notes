package services

import (
	"context"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"errors"
	"testing"
	"time"

	"github.com/g-villarinho/tab-notes-api/configs"
	"github.com/g-villarinho/tab-notes-api/mocks"
	"github.com/g-villarinho/tab-notes-api/models"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
)

func TestGenerateAuthToken(t *testing.T) {
	ctx := context.Background()

	t.Run("should return error if private key is invalid", func(t *testing.T) {
		kp := new(mocks.EcdsaKeyPairMock)
		ts := NewTokenService(kp)

		kp.
			On("ParseECDSAPrivateKey", configs.Env.Key.PrivateKey).
			Return(nil, errors.New("invalid key"))

		token, err := ts.GenerateAuthToken(ctx, "user-1", "session-1", time.Now(), time.Now().Add(time.Hour))

		assert.Empty(t, token)
		assert.ErrorContains(t, err, "invalid key")
		kp.AssertExpectations(t)
	})

	t.Run("should generate signed JWT with correct claims", func(t *testing.T) {
		kp := new(mocks.EcdsaKeyPairMock)
		ts := NewTokenService(kp)

		privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
		assert.NoError(t, err)

		kp.
			On("ParseECDSAPrivateKey", configs.Env.Key.PrivateKey).
			Return(privateKey, nil)

		iat := time.Now().UTC()
		exp := iat.Add(10 * time.Minute)

		tokenStr, err := ts.GenerateAuthToken(ctx, "user-123", "session-abc", iat, exp)

		assert.NoError(t, err)
		assert.NotEmpty(t, tokenStr)

		parsedToken, err := jwt.ParseWithClaims(tokenStr, &models.AuthTokenClaims{}, func(token *jwt.Token) (any, error) {
			return &privateKey.PublicKey, nil
		})
		assert.NoError(t, err)
		assert.True(t, parsedToken.Valid)

		claims, ok := parsedToken.Claims.(*models.AuthTokenClaims)
		assert.True(t, ok)
		assert.Equal(t, "user-123", claims.Subject)
		assert.Equal(t, "session-abc", claims.SessionID)
	})
}

func TestGenerateMagicLinkToken(t *testing.T) {
	ctx := context.Background()

	t.Run("should return error if private key is invalid", func(t *testing.T) {
		kp := new(mocks.EcdsaKeyPairMock)
		ts := NewTokenService(kp)

		kp.
			On("ParseECDSAPrivateKey", configs.Env.Key.PrivateKey).
			Return(nil, errors.New("invalid key"))

		token, err := ts.GenerateMagicLinkToken(ctx, "joao@example.com", time.Now(), time.Now().Add(time.Minute))

		assert.Empty(t, token)
		assert.ErrorContains(t, err, "invalid key")
		kp.AssertExpectations(t)
	})

	t.Run("should generate signed JWT with correct claims", func(t *testing.T) {
		kp := new(mocks.EcdsaKeyPairMock)
		ts := NewTokenService(kp)

		privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
		assert.NoError(t, err)

		kp.
			On("ParseECDSAPrivateKey", configs.Env.Key.PrivateKey).
			Return(privateKey, nil)

		iat := time.Now().UTC()
		exp := iat.Add(15 * time.Minute)

		tokenStr, err := ts.GenerateMagicLinkToken(ctx, "joao@example.com", iat, exp)

		assert.NoError(t, err)
		assert.NotEmpty(t, tokenStr)

		parsedToken, err := jwt.ParseWithClaims(tokenStr, &models.MagicLinkTokenClaims{}, func(token *jwt.Token) (any, error) {
			return &privateKey.PublicKey, nil
		})
		assert.NoError(t, err)
		assert.True(t, parsedToken.Valid)

		claims, ok := parsedToken.Claims.(*models.MagicLinkTokenClaims)
		assert.True(t, ok)
		assert.Equal(t, "joao@example.com", claims.Subject)
		assert.WithinDuration(t, iat, claims.IssuedAt.Time, time.Second)
		assert.WithinDuration(t, exp, claims.ExpiresAt.Time, time.Second)
	})
}
