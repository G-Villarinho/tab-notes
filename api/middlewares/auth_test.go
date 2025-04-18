package middlewares

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/g-villarinho/tab-notes-api/configs"
	"github.com/g-villarinho/tab-notes-api/mocks"
	"github.com/g-villarinho/tab-notes-api/models"
	"github.com/g-villarinho/tab-notes-api/pkgs"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var testPrivateKey *ecdsa.PrivateKey

func init() {
	var err error
	testPrivateKey, err = ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		log.Fatalf("failed to generate ECDSA key: %v", err)
	}
}

func generateValidJWT(t *testing.T, key *ecdsa.PrivateKey, sub string, sid string) string {
	iat := time.Now().UTC()
	exp := iat.Add(10 * time.Minute)

	claims := models.AuthTokenClaims{
		SessionID: sid,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   sub,
			IssuedAt:  jwt.NewNumericDate(iat),
			ExpiresAt: jwt.NewNumericDate(exp),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	str, err := token.SignedString(key)
	assert.NoError(t, err)
	return str
}

func TestAuthMiddleware_Authenticated(t *testing.T) {
	t.Run("should return 401 if token is missing", func(t *testing.T) {
		kp := new(mocks.EcdsaKeyPairMock)
		rc := pkgs.NewRequestContext()
		ss := new(mocks.SessionServiceMock)
		mw := NewAuthMiddleware(kp, rc, ss)

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rr := httptest.NewRecorder()

		handler := mw.Authenticated(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			t.Fail()
		}))

		handler.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusUnauthorized, rr.Code)
	})

	t.Run("should return 401 if token is invalid JWT", func(t *testing.T) {
		kp := new(mocks.EcdsaKeyPairMock)
		rc := pkgs.NewRequestContext()
		ss := new(mocks.SessionServiceMock)
		mw := NewAuthMiddleware(kp, rc, ss)

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.AddCookie(&http.Cookie{Name: "tabnews_id", Value: "invalid-token"})
		rr := httptest.NewRecorder()

		kp.On("ParseECDSAPublicKey", configs.Env.Key.PublicKey).
			Return(&testPrivateKey.PublicKey, nil)

		handler := mw.Authenticated(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			t.Fail()
		}))

		handler.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusUnauthorized, rr.Code)
	})

	t.Run("should return 401 if session is revoked", func(t *testing.T) {
		kp := new(mocks.EcdsaKeyPairMock)
		rc := pkgs.NewRequestContext()
		ss := new(mocks.SessionServiceMock)
		mw := NewAuthMiddleware(kp, rc, ss)

		tokenStr := generateValidJWT(t, testPrivateKey, "user-1", "sess-1")

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.AddCookie(&http.Cookie{Name: "tabnews_id", Value: tokenStr})
		rr := httptest.NewRecorder()

		kp.On("ParseECDSAPublicKey", configs.Env.Key.PublicKey).
			Return(&testPrivateKey.PublicKey, nil)

		ss.On("IsSessionRevoked", mock.Anything, "sess-1").
			Return(true, nil)

		handler := mw.Authenticated(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			t.Fail()
		}))

		handler.ServeHTTP(rr, req)
		assert.Equal(t, http.StatusUnauthorized, rr.Code)
	})

	t.Run("should call next if token is valid and session is active", func(t *testing.T) {
		kp := new(mocks.EcdsaKeyPairMock)
		rc := pkgs.NewRequestContext()
		ss := new(mocks.SessionServiceMock)
		mw := NewAuthMiddleware(kp, rc, ss)

		tokenStr := generateValidJWT(t, testPrivateKey, "user-1", "sess-1")

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.AddCookie(&http.Cookie{Name: "tabnews_id", Value: tokenStr})
		rr := httptest.NewRecorder()

		kp.On("ParseECDSAPublicKey", configs.Env.Key.PublicKey).
			Return(&testPrivateKey.PublicKey, nil)

		ss.On("IsSessionRevoked", mock.Anything, "sess-1").
			Return(false, nil)

		called := false
		handler := mw.Authenticated(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			called = true
			userID, ok := rc.GetUserID(r.Context())
			assert.True(t, ok)
			assert.Equal(t, "user-1", userID)
			w.WriteHeader(http.StatusOK)
		}))

		handler.ServeHTTP(rr, req)

		assert.True(t, called)
		assert.Equal(t, http.StatusOK, rr.Code)
	})
}
